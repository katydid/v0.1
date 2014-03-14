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
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func Go(w http.ResponseWriter, req *http.Request) {
	var err error

	path := req.FormValue("path")

	if !exists(filepath.Join(path, "one.box")) {
		err = fmt.Errorf("Please remember to execute the previous textbox!")
		writeError(w, err, nil)
		return
	}

	src := filepath.Join(path, "two.box")

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
	if err = ioutil.WriteFile(src+".go", body.Bytes(), 0666); err != nil {
		writeError(w, err, nil)
		return
	}

	mainSrc := `package main

	import "code.google.com/p/gogoprotobuf/proto"
	import "os"
	import "fmt"

	func main() {
		pop, err := Populate()
		if err != nil {
			panic(err)
		}
		data, err := proto.Marshal(pop)
		if err != nil {
			panic(err)
		}
		_, err = os.Stdout.Write(data)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, proto.MarshalTextString(pop)+"\n")
	}
	`

	err = ioutil.WriteFile(filepath.Join(path, "main.go"), []byte(mainSrc), 0666)
	if err != nil {
		writeError(w, err, nil)
		return
	}

	out, errBuf, err := run(path, "go", "run", "main.go", "one.box.pb.go", "two.box.go")
	if err != nil {
		out = append(errBuf, out...)
		writeError(w, err, out)
		return
	}

	if err = ioutil.WriteFile(filepath.Join(path, "two.pop"), out, 0666); err != nil {
		writeError(w, err, nil)
		return
	}

	w.Write([]byte("<pre>"))
	w.Write(errBuf)
	w.Write([]byte("</pre>"))

	return
}

func init() {
	register(&Box{
		Title: "Populate a Protocol Buffer using Go",
		Name:  "two",
		Content: `package main

import "code.google.com/p/gogoprotobuf/proto"

func Populate() (*Hello, error) {
	h := &Hello{}
	h.World = proto.String("World")
	return h, nil
}
`,
		Func:  Go,
		Order: 2,
		Help: `Go is a programming language. For more information see <a href="http://golang.org/">http://golang.org/</a>. 
		If you have compiled the protocol buffer defined above then the Go code has already been generated and is ready to use.
		Write a Populate function that returns the appropriate structure and an error.`,
	})
}
