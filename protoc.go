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
	"io/ioutil"
	"net/http"
	"path/filepath"

	"code.google.com/p/gogoprotobuf/proto"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
)

func Protoc(w http.ResponseWriter, req *http.Request) {
	var err error
	// x is the base name for .go, .6, executable files
	path := req.FormValue("path")
	src := filepath.Join(path, "one.box")

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

	dir, file := filepath.Split(src)

	out, errBuf, err := run(dir, "protoc", "--descriptor_set_out=/dev/stdout", file)
	if err != nil {
		out = append(errBuf, out...)
		writeError(w, err, out)
		return
	}

	if err = ioutil.WriteFile(filepath.Join(path, "one.desc"), out, 0666); err != nil {
		writeError(w, err, nil)
		return
	}

	desc := &descriptor.FileDescriptorSet{}
	err = proto.Unmarshal(out, desc)
	if err != nil {
		writeError(w, err, nil)
		return
	}
	out = []byte(proto.MarshalTextString(desc) + "\n")
	out = append(out, errBuf...)

	w.Write([]byte("<pre>"))
	w.Write(out)
	w.Write([]byte("</pre>"))

	out, errBuf, err = run(dir, "protoc", "--gogo_out="+dir, file)
	if err != nil {
		out = append(errBuf, out...)
		writeError(w, err, out)
		return
	}
	out = append(out, errBuf...)

	w.Write([]byte("<pre>"))
	w.Write(out)
	w.Write([]byte("</pre>"))

	out, err = ioutil.ReadFile(src + ".pb.go")
	if err != nil {
		writeError(w, err, nil)
		return
	}
	w.Write([]byte("<pre>"))
	w.Write(out)
	w.Write([]byte("</pre>"))

	return
}

func init() {
	register(&Box{
		Title: "Define a Protocol Buffer",
		Name:  "one",
		Content: `package main;

message Hello {
	optional string World = 1;
}`,
		Func:  Protoc,
		Order: 1,
		Help: `
Protocol Buffers are a way of encoding structured data in an efficient yet extensible format. 
For more information please see <a href="http://code.google.com/p/protobuf/">http://code.google.com/p/protobuf/</a>.
`,
	})
}
