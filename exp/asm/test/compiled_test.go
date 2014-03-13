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

package test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	protoparser "code.google.com/p/gogoprotobuf/parser"
	"code.google.com/p/gogoprotobuf/proto"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"github.com/awalterschulze/katydid/exp/asm/ast"
	"github.com/awalterschulze/katydid/exp/asm/compiler"
	katyexec "github.com/awalterschulze/katydid/exp/asm/exec"
	"github.com/awalterschulze/katydid/exp/asm/lexer"
	"github.com/awalterschulze/katydid/exp/asm/parser"
)

var (
	fileDescriptorSet *descriptor.FileDescriptorSet
)

func newJsonFile(name string, m proto.Message) {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	os.Remove(name + ".json")
	jsonFile, err := os.Create(name + ".json")
	if err != nil {
		panic(err)
	}
	if _, err := jsonFile.Write(data); err != nil {
		panic(err)
	}
	if err := jsonFile.Close(); err != nil {
		panic(err)
	}
}

func init() {
	var err error
	fileDescriptorSet, err = protoparser.ParseFile("test.proto", ".", "../../../../../../")
	if err != nil {
		panic(err)
	}
	/*newJsonFile("david", david)
	newJsonFile("robert", robert)
	newJsonFile("ioutil", ioUtil)
	newJsonFile("path", path)
	newJsonFile("runtime", runtime)
	newJsonFile("syscall", syscall)
	newJsonFile("mover", mover)
	newJsonFile("shaker", shaker)
	newJsonFile("routine", routine)
	newJsonFile("noname", noname)
	newJsonFile("john", john)
	newJsonFile("smith", smith)*/
}

type tester interface {
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type test struct {
	t     tester
	str   string
	rules *ast.Rules
	exec  *katyexec.Exec
}

func newTest(name string, t tester, str string) test {
	p := parser.NewParser()
	r, err := p.Parse(lexer.NewLexer([]byte(str)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	rules := r.(*ast.Rules)
	e, err := compiler.Compile(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	dotString := rules.Dot()
	os.Remove(name + ".dot")
	os.Remove(name + ".pdf")
	os.Remove(name + ".txt")
	textFile, err := os.Create(name + ".txt")
	if err != nil {
		panic(err)
	}
	if _, err := textFile.Write([]byte(str)); err != nil {
		panic(err)
	}
	if err := textFile.Close(); err != nil {
		panic(err)
	}
	file, err := os.Create(name + ".dot")
	if err != nil {
		panic(err)
	}
	if _, err := file.Write([]byte(dotString)); err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}
	cmd := exec.Command("dot", name+".dot", "-Tpdf", "-o", name+".pdf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("dot error: " + err.Error() + ":" + string(output))
	}
	return test{
		t:     t,
		str:   str,
		rules: rules,
		exec:  e,
	}
}

func (this test) match(m proto.Message, positive bool) test {
	fmt.Printf("======== Testing %v\n", m)
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	if match, err := this.exec.Eval(data); err != nil {
		this.t.Errorf("Error: %v", err)
	} else if match != positive {
		this.t.Errorf("Expected a %v match from \n%v \non \n%v", positive, this.str, m)
	}
	return this
}

var david = &Person{
	Name: proto.String("David"),
	Addresses: []*Address{
		{
			Number: proto.Int64(123),
			Street: proto.String("TheStreet"),
		},
		{
			Number: proto.Int64(456),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var robert = &Person{
	Name: proto.String("Robert"),
	Addresses: []*Address{
		{
			Number: proto.Int64(456),
			Street: proto.String("TheStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var contextPerson = `root = test.Person
	test.Person = start
	start numberAndStreet = accept
	start _ = start

	test.Address = address
	address number = number
	address street = street
	address _ = address
	number street = numberAndStreet
	number _ = number
	street number = numberAndStreet
	street _ = street

	if (decInt64(test.Address.Number) == int64(456)) then number else noNumber

	if contains(nfkc(decString(test.Address.Street)), nfkc("TheStreet")) then street else noStreet
	`

func TestContextPerson(t *testing.T) {
	//Does this Person live at 456 The Street
	newTest("TestContextPerson", t, contextPerson).match(david, false).match(robert, true)
}

var ioUtil = &SrcTree{
	PackageName: proto.String("io/ioutil"),
	Imports: []*SrcTree{
		{
			PackageName: proto.String("io"),
			Imports: []*SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("sync"),
				},
			},
		},
		{
			PackageName: proto.String("os"),
			Imports: []*SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("io"),
				},
				{
					PackageName: proto.String("runtime"),
				},
			},
		},
	},
}

var path = &SrcTree{
	PackageName: proto.String("path"),
	Imports: []*SrcTree{
		{
			PackageName: proto.String("errors"),
		},
		{
			PackageName: proto.String("strings"),
			Imports: []*SrcTree{
				{
					PackageName: proto.String("errors"),
				},
				{
					PackageName: proto.String("io"),
				},
				{
					PackageName: proto.String("uncode"),
				},
				{
					PackageName: proto.String("uncode/utf8"),
				},
			},
		},
		{
			PackageName: proto.String("unicode/utf8"),
		},
	},
}

var runtime = &SrcTree{
	PackageName: proto.String("runtime"),
	Imports: []*SrcTree{
		{
			PackageName: proto.String("unsafe"),
		},
	},
}

var syscall = &SrcTree{
	PackageName: proto.String("syscall"),
	Imports: []*SrcTree{
		{
			PackageName: proto.String("errors"),
		},
		{
			PackageName: proto.String("runtime"),
		},
		{
			PackageName: proto.String("sync"),
			Imports: []*SrcTree{
				{
					PackageName: proto.String("sync/atomic"),
				},
				{
					PackageName: proto.String("unsafe"),
				},
			},
		},
		{
			PackageName: proto.String("unsafe"),
		},
	},
}

var recursiveSrcTree = `root = test.SrcTree
	test.SrcTree = start
	start accept = accept
	start _ = start
	accept _ = accept

	if (decString(test.SrcTree.PackageName) == "io") 
	  then accept 
	  else packageName

	`

func TestRecursiveSrcTree(t *testing.T) {
	//Does this SrcTree depend on io or is its packageName io
	newTest("TestRecursiveSrcTree", t, recursiveSrcTree).match(ioUtil, true).match(path, true).match(runtime, false).match(syscall, false)
}

var mover = &Person{
	Name: proto.String("Mover"),
	Addresses: []*Address{
		{
			Number: proto.Int64(123),
			Street: proto.String("TheStreet"),
		},
		{
			Number: proto.Int64(456),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var shaker = &Person{
	Name: proto.String("Shaker"),
	Addresses: []*Address{
		{
			Number: proto.Int64(55),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(2),
			Street: proto.String("SomeStreet"),
		},
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var routine = &Person{
	Name: proto.String("Routine"),
	Addresses: []*Address{
		{
			Number: proto.Int64(3),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0124444444"),
}

var listIndexAddress = `root = test.Person
	test.Person = start
	start numberTwo = topNumberTwo
	start _ = start
	topNumberTwo numberOne = accept
	topNumberTwo numberTwo = topNumberTwo
	topNumberTwo _ = start
	accept numberTwo = topNumberTwo
	accept _ = start

	test.Address = address
	address numberTwo = numberTwo
	address numberOne = numberOne
	address _ = address
	numberTwo numberTwo = numberTwo
	numberTwo numberOne = numberOne
	numberOne numberTwo = numberTwo
	numberOne numberOne = numberOne

	if (decInt64(test.Address.Number) == int64(1))
	  then numberOne
	  else {
	    if (decInt64(test.Address.Number) == int64(2))
	    then numberTwo
	    else noNumber
	  }
`

func TestListIndexAddress(t *testing.T) {
	//Is this Person's newest streetnumber 1 and second newest streetnumber 2.
	//Assume that addresses are appended to the list, so the last address is the newest address.
	newTest("TestListIndexAddress", t, listIndexAddress).match(mover, false).match(shaker, true).match(routine, false)
	// find test.Person where { test.Person { Addresses[-2].Number == 2 && Addresses[-1].Number == 1 } }
}

var noname = &Person{
	Addresses: []*Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var john = &Person{
	Name: proto.String("John"),
	Addresses: []*Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0123456789"),
}

var smith = &Person{
	Name: proto.String(""),
	Addresses: []*Address{
		{
			Number: proto.Int64(1),
			Street: proto.String("SomeStreet"),
		},
	},
	Telephone: proto.String("0127897897"),
}

var nilName = `root = test.Person
	test.Person = accept
	accept name = reject
	accept _ = accept

	if exists(test.Person.Name)
	  then name
	  else noname
	`

func TestNilName(t *testing.T) {
	//Is this Person's name missing
	newTest("TestNilName", t, nilName).match(noname, true).match(john, false).match(smith, false)
}

var lenName = `root = test.Person
	test.Person = start
	start name = reject
	start noname = accept
	start _ = start

	if (length(decString(test.Person.Name)) == int64(0))
	  then noname
	  else name`

func TestLenName(t *testing.T) {
	//Is this Person's name an empty string
	newTest("TestLenName", t, lenName).match(noname, false).match(john, false).match(smith, true)
}

var emptyOrNil = `root = test.Person
	test.Person = accept
	accept name = reject
	accept _ = accept

	if (length(decString(test.Person.Name)) == int64(0))
	  then noname
	  else name`

func TestEmptyOrNil(t *testing.T) {
	//Is this Person's name empty or an empty string
	newTest("TestEmptyOrNil", t, emptyOrNil).match(noname, true).match(john, false).match(smith, true)
}

var incorrentNotName = `root = test.Person
	test.Person = start
	start notname = accept
	start _ = start

	if not((decString(test.Person.Name) == "David")) 
	  then notname 
	  else name
		`

func TestIncorrectNotName(t *testing.T) {
	newTest("TestIncorrectNotName", t, incorrentNotName).match(noname, false).match(robert, true).match(smith, true).match(david, false)
}

var correctNotName = `root = test.Person
	test.Person = accept
	accept name = reject
	reject _ = reject
	accept _ = accept

	if (decString(test.Person.Name) == "David") 
	  then name 
	  else noname
		`

func TestCorrectNotName(t *testing.T) {
	newTest("TestCorrectNotName", t, correctNotName).match(noname, true).match(robert, true).match(smith, true).match(david, false)
}

var andNameTelephone = `root = test.Person
	test.Person = start
	start name = name
	start tel = tel
	start _ = start
	name tel = accept
	name _ = name
	tel name = accept
	tel _ = tel

	if (decString(test.Person.Name) == "David") 
	  then name 
	  else noname

	if (decString(test.Person.Telephone) == "0123456789") 
	  then tel 
	  else notel
		`

func TestAndNameTelephone(t *testing.T) {
	newTest("TestAndNameTelephone", t, andNameTelephone).match(david, true).match(john, false).match(mover, false).match(smith, false)
}

var orNameTelephone = `root = test.Person
	test.Person = start
	start name = accept
	start tel = accept
	start _ = start
	accept _ = accept

	if (decString(test.Person.Name) == "David") 
	  then name 
	  else noname

	if (decString(test.Person.Telephone) == "0123456789") 
	  then tel 
	  else notel
		`

func TestOrNameTelephone(t *testing.T) {
	newTest("TestOrNameTelephone", t, orNameTelephone).match(david, true).match(john, true).match(mover, true).match(smith, false)
}