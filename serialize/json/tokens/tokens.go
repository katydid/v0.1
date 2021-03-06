//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package tokens

import (
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/asm/ast"
	prototokens "github.com/katydid/katydid/serialize/proto/tokens"
)

type errUnknown struct {
	Token string
}

func (this *errUnknown) Error() string {
	return "Could not find " + this.Token
}

type JsonTokens interface {
	Len() int
	GetTokenId(tokenString string) (int, error)
	Lookup(parentToken int, name []byte) int
}

type Visitor interface {
	GetToken(pkg, msg, field string) int
	AddTrans(srcToken int, field *descriptor.FieldDescriptorProto, dstToken int)
	StringTokens() []string
}

type visitor struct {
	*jsonTokens
	stringTokens []string
}

func newVisitor(stringTokens []string) *visitor {
	visitor := &visitor{
		jsonTokens: &jsonTokens{
			trans:       make(map[int]map[string]int),
			nameToToken: make(map[string]int),
			tokenToName: make(map[int]string),
		},
		stringTokens: stringTokens,
	}
	return visitor
}

func stateName(pkg, msg, field string) string {
	key := pkg + "." + msg
	if len(field) > 0 {
		key += "." + field
	}
	return key
}

func (this *visitor) GetToken(pkg, msg, field string) int {
	key := stateName(pkg, msg, field)
	return this.getToken(key)
}

func (this *jsonTokens) getToken(tokenStr string) int {
	s, ok := this.nameToToken[tokenStr]
	if ok {
		return s
	}
	s = len(this.nameToToken)
	this.nameToToken[tokenStr] = s
	this.tokenToName[s] = tokenStr
	return s
}

func (this *visitor) AddTrans(srcToken int, field *descriptor.FieldDescriptorProto, dstToken int) {
	input := field.GetName()
	if _, ok := this.trans[srcToken]; !ok {
		this.trans[srcToken] = make(map[string]int)
	}
	if _, ok := this.trans[srcToken][input]; ok {
		panic("transition already exists")
	}
	this.trans[srcToken][input] = dstToken
}

func (this *visitor) StringTokens() []string {
	return this.stringTokens
}

type jsonTokens struct {
	tokenToName map[int]string
	nameToToken map[string]int
	trans       map[int]map[string]int
}

func NewZipped(rules *ast.Rules, desc *descriptor.FileDescriptorSet) (JsonTokens, error) {
	stringTokens, err := ast.GetStringTokens(rules)
	if err != nil {
		return nil, err
	}
	tokens, err := new(rules.GetRoot().GetPackage(), rules.GetRoot().GetMessage(), desc, stringTokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func New(rules *ast.Rules, desc *descriptor.FileDescriptorSet) (JsonTokens, error) {
	srcPackage, srcMessage := rules.GetRoot().GetPackage(), rules.GetRoot().GetMessage()
	return new(srcPackage, srcMessage, desc, nil)
}

func new(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet, stringTokens []string) (JsonTokens, error) {
	visitor := newVisitor(stringTokens)
	if err := prototokens.Walk(srcPackage, srcMessage, desc, visitor); err != nil {
		return nil, err
	}
	return visitor.jsonTokens, nil
}

func (this *jsonTokens) GetTokenId(tokenString string) (int, error) {
	s, ok := this.nameToToken[tokenString]
	if !ok {
		return 0, &errUnknown{tokenString}
	}
	return s, nil
}

func (this *jsonTokens) Len() int {
	return len(this.tokenToName)
}

func (this *jsonTokens) Lookup(parentToken int, name []byte) int {
	strName := string(name)
	token, ok := this.trans[parentToken][strName]
	if ok {
		return token
	}
	parentName, ok := this.tokenToName[parentToken]
	if !ok {
		panic("no parent")
	}
	fullName := parentName + "." + strName
	newToken := this.getToken(fullName)
	this.trans[parentToken][strName] = newToken
	return newToken
}
