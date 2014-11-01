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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"code.google.com/p/gogoprotobuf/proto"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/asm/compiler"
	"github.com/katydid/katydid/asm/lexer"
	"github.com/katydid/katydid/asm/parser"
	"github.com/katydid/katydid/serialize/proto/scanner"
	"github.com/katydid/katydid/serialize/proto/tokens"

	"text/template"
)

func Katydid(w http.ResponseWriter, req *http.Request) {
	var err error
	// x is the base name for .go, .6, executable files
	path := req.FormValue("path")

	if !exists(filepath.Join(path, "two.box")) {
		err = fmt.Errorf("Please remember to execute the previous textbox!")
		writeError(w, err, nil)
		return
	}

	src := filepath.Join(path, "three.box")

	// write body to x.proto
	body := new(bytes.Buffer)
	if _, err = body.ReadFrom(req.Body); err != nil {
		writeError(w, err, nil)
		return
	}
	bodyBytes := body.Bytes()
	if err = ioutil.WriteFile(src, bodyBytes, 0666); err != nil {
		writeError(w, err, nil)
		return
	}

	p := parser.NewParser()
	r, err := p.Parse(lexer.NewLexer(bodyBytes))
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

	bbb := bench{
		Msg:   fmt.Sprintf("%#v", data),
		Query: fmt.Sprintf("`%v`", string(bodyBytes)),
		Desc:  fmt.Sprintf("%#v", fileDescData),
	}

	protoTokens, err := tokens.NewZipped(rules, fileDescriptorSet)
	if err != nil {
		writeError(w, err, nil)
		return
	}

	e, rootToken, err := compiler.Compile(rules, protoTokens)
	if err != nil {
		writeError(w, err, nil)
		return
	}

	s := scanner.NewProtoScanner(protoTokens, rootToken)
	err = s.Init(data)
	if err != nil {
		writeError(w, err, nil)
		return
	}

	recog, err := e.Eval(s)
	if err != nil {
		writeError(w, err, nil)
		return
	}

	out := []byte(fmt.Sprintf("%v", recog))

	data, err = ioutil.ReadFile(dotPicFilename)
	if err != nil {
		writeError(w, err, nil)
		return
	}

	dot64String := base64.StdEncoding.EncodeToString(data)
	w.Write([]byte(`<table><tr><td valign="top">`))
	if strings.Contains(string(out), "true") {
		w.Write([]byte(`The populated protocol buffer was recognized`))
	} else if strings.Contains(string(out), "false") {
		w.Write([]byte(`The populated protocol buffer was NOT recognized`))
	} else {
		writeError(w, errors.New("Something unexpected happened"), output)
		return
	}
	w.Write([]byte(`</tr><tr><td valign="top">`))
	w.Write([]byte(`<img src="data:image/png;base64,` + dot64String + `" alt="error displaying image"/>`))
	w.Write([]byte(`</td></tr><tr><td valign="top">`))

	benchFile, err := os.Create(filepath.Join(path, "bench_test.go"))
	if err != nil {
		writeError(w, err, output)
		return
	}
	err2 := benchTemplate.Execute(benchFile, bbb)
	if err := benchFile.Close(); err != nil {
		writeError(w, err, output)
		return
	}
	if err2 != nil {
		writeError(w, err2, nil)
		return
	}

	cmdTestCompile := exec.Command("go", "test", "-c")
	cmdTestCompile.Dir = path
	output, err = cmdTestCompile.CombinedOutput()
	if err != nil {
		writeError(w, err, output)
		return
	}
	if len(output) > 0 {
		w.Write([]byte("<p>"))
		w.Write(output)
		w.Write([]byte("</p>"))
	}

	out, errBuf, err := run(path, filepath.Join(path, filepath.Base(path)+".test"), "-test.run=XXX", "-test.bench=.")
	if err != nil {
		out = append(errBuf, out...)
		writeError(w, err, out)
		return
	}

	w.Write([]byte(strings.Replace(string(out), "PASS", "", -1)))

	w.Write([]byte("</td></tr></table>"))

	return
}

type bench struct {
	Msg   string
	Query string
	Desc  string
}

var (
	benchTemplate *template.Template
)

func init() {
	benchTemplate = template.Must(template.New("a").Parse(benchStr))
}

var benchStr string = `
	package main

	import (
		"testing"
		"github.com/katydid/katydid/asm/ast"
		"github.com/katydid/katydid/asm/compiler"
		"github.com/katydid/katydid/asm/lexer"
		"github.com/katydid/katydid/asm/parser"
		"github.com/katydid/katydid/serialize/proto/scanner"
		"github.com/katydid/katydid/serialize/proto/tokens"
		"code.google.com/p/gogoprotobuf/proto"
		descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	)

	func BenchmarkQuery(b *testing.B) {
		buf := {{.Msg}}
		p := parser.NewParser()
		r, err := p.Parse(lexer.NewLexer([]byte({{.Query}})))
		if err != nil {
			panic(err)
		}
		rules := r.(*ast.Rules)
		desc := &descriptor.FileDescriptorSet{}
		err = proto.Unmarshal({{.Desc}}, desc)
		if err != nil {
			panic(err)
		}
		protoTokens, err := tokens.NewZipped(rules, desc)
		if err != nil {
			panic(err)
		}
		exec, rootToken, err := compiler.Compile(rules, protoTokens)
		if err != nil {
			panic(err)
		}
		s := scanner.NewProtoScanner(protoTokens, rootToken)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := s.Init(buf); err != nil {
				panic(err)
			}
			if _, err := exec.Eval(s); err != nil {
				panic(err)
			}
		}
	}
	
`

func init() {
	register(&Box{
		Title: "Write a Query",
		Name:  "three",
		Content: `root = main.Hello
main.Hello = start
start world = accept
start _ = start
accept _ = accept

if contains($string(main.Hello.World), "World") then world else noworld`,
		Func:  Katydid,
		Order: 3,
		Help: `Katydid is currently in an experimental phase. <br/>
		Here you can describe a bottom up hedge automaton which will either recognize or not recognize the populated protocol buffer.
		See <a href="https://github.com/katydid/katydid">https://github.com/katydid/katydid</a>`,
	})
}
