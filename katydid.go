//  Copyright 2014 Walter Schulze
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

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"

	"code.google.com/p/gogoprotobuf/proto"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"github.com/awalterschulze/katydid/exp/readable"
	"github.com/awalterschulze/katydid/exp/readable/ast"
	"github.com/awalterschulze/katydid/exp/readable/lexer"
	"github.com/awalterschulze/katydid/exp/readable/parser"
)

func Katydid(w http.ResponseWriter, req *http.Request) {
	var err error
	// x is the base name for .go, .6, executable files
	path := req.FormValue("path")
	src := filepath.Join(path, "three.box")

	// write body to x.proto
	body := new(bytes.Buffer)
	if _, err = body.ReadFrom(req.Body); err != nil {
		writeError(w, err, nil)
		return
	}
	if err = ioutil.WriteFile(src, body.Bytes(), 0666); err != nil {
		writeError(w, err, nil)
		return
	}

	p := parser.NewParser()
	r, err := p.Parse(lexer.NewLexer(body.Bytes()))
	if err != nil {
		writeError(w, err, nil)
		return
	}
	rules := r.(*ast.Rules)
	dotString := rules.Dot()

	dotFilename := filepath.Join(path, "three.dot")
	if err := ioutil.WriteFile(dotFilename, []byte(dotString), 0666); err != nil {
		writeError(w, err, []byte(dotString))
		return
	}
	dotPicFilename := filepath.Join(path, "three.png")
	cmd := exec.Command("dot", dotFilename, "-Tpng", "-o", dotPicFilename)
	output, err := cmd.CombinedOutput()
	if err != nil {
		writeError(w, err, output)
		return
	}

	data, err := ioutil.ReadFile(filepath.Join(path, "two.pop"))
	if err != nil {
		writeError(w, err, nil)
		return
	}

	fileDescData, err := ioutil.ReadFile(filepath.Join(path, "one.desc"))
	if err != nil {
		writeError(w, err, nil)
		return
	}

	fileDescriptorSet := &descriptor.FileDescriptorSet{}
	if err := proto.Unmarshal(fileDescData, fileDescriptorSet); err != nil {
		writeError(w, err, nil)
		return
	}

	matcher, err := readable.NewInterpreter(fileDescriptorSet, rules)
	if err != nil {
		writeError(w, err, nil)
		return
	}
	match, err := matcher.Match(data)
	if err != nil {
		writeError(w, err, nil)
		return
	}

	out := []byte(fmt.Sprintf("%v", match))

	data, err = ioutil.ReadFile(dotPicFilename)
	if err != nil {
		writeError(w, err, nil)
		return
	}

	dot64String := base64.StdEncoding.EncodeToString(data)
	w.Write([]byte(`<table><tr><td valign="top">`))
	w.Write([]byte(`<img src="data:image/png;base64,` + dot64String + `" alt="error displaying image"/>`))
	w.Write([]byte("</td>"))
	w.Write([]byte(`<td valign="top">`))
	w.Write(out)
	w.Write([]byte("</td></tr></table>"))

	return
}

func init() {
	register(&Box{
		Title: "Write a Query",
		Name:  "three",
		Content: `root = main.Hello
main.Hello = start
start world = accept
start _ = start
accept _ = accept

if contains(decString(main.Hello.World.value), "World") then world else noworld`,
		Func:  Katydid,
		Order: 3,
	})
}
