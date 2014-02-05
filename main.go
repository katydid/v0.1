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
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var (
	httpListen = flag.String("http", "127.0.0.1:4000", "host:port to listen on")
)

type Boxes []*Box

func (this Boxes) Less(i, j int) bool {
	return this[i].Order < this[j].Order
}

func (this Boxes) Len() int {
	return len(this)
}

func (this Boxes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

var boxes = Boxes([]*Box{})

func main() {
	flag.Parse()

	http.HandleFunc("/share/", Share)
	http.HandleFunc("/shared/", Shared)
	for _, box := range boxes {
		http.HandleFunc("/func/"+box.Name, box.Func)
	}
	http.HandleFunc("/", Main)
	log.Fatal(http.ListenAndServe(*httpListen, nil))
}

type Box struct {
	Title   string
	Name    string
	Content string
	Func    func(w http.ResponseWriter, req *http.Request)
	Path    string
	Help    string
	Order   int
}

func Share(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("share\n")
	path := req.URL.Path[1:]
	if len(path) <= 4 || !strings.HasPrefix(path, "share") {
		writeError(w, errors.New("path too short"), nil)
		return
	}
	if dir := filepath.Dir(filepath.Dir(path)); dir != "." {
		writeError(w, errors.New("path too long by: "+dir), nil)
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		writeError(w, err, nil)
		return
	}

	cwd := filepath.Join(wd, path)
	cwd = strings.Replace(cwd, "share", "tmp", -1)
	newShared := paths.NewShared()

	scwd := filepath.Join(wd, newShared)
	fmt.Printf("copying from %v to %v\n", cwd, scwd)
	if err = copyDir(cwd, scwd); err != nil {
		writeError(w, err, nil)
		return
	}
	w.Write([]byte(`<a href="../` + newShared + `">../` + newShared + `</a>`))
}

func Shared(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("shared\n")
	path := req.URL.Path[1:]
	if len(path) <= 4 || !strings.HasPrefix(path, "shared") {
		writeError(w, errors.New("path too short"), nil)
		return
	}
	if dir := filepath.Dir(filepath.Dir(path)); dir != "." {
		writeError(w, errors.New("path too long by: "+dir), nil)
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		writeError(w, err, nil)
		return
	}
	path = req.URL.Path[1:]
	fmt.Printf("path = %v\n", path)

	cwd := filepath.Join(wd, path)
	newTemp := paths.NewTemp()

	tcwd := filepath.Join(wd, newTemp)
	fmt.Printf("copying to %v\n", tcwd)
	if err = copyDir(cwd, tcwd); err != nil {
		writeError(w, err, nil)
		return
	}
	http.Redirect(w, req, "../"+newTemp, http.StatusFound)
}

func copyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}
	}
	return
}

func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

func copyDir(source string, dest string) error {
	if err := os.MkdirAll(dest, 0777); err != nil {
		return err
	}
	filenames, err := readDirNames(source)
	if err != nil {
		return err
	}
	for _, name := range filenames {
		_, f := filepath.Split(name)
		err := copyFile(filepath.Join(source, name), filepath.Join(dest, f))
		if err != nil {
			return err
		}
	}
	return nil
}

func Main(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	if len(path) == 0 {
		http.Redirect(w, req, paths.NewTemp(), http.StatusFound)
		return
	}
	if len(path) <= 4 || !strings.HasPrefix(path, "tmp") {
		writeError(w, errors.New("path too short"), nil)
		return
	}
	if dir := filepath.Dir(filepath.Dir(path)); dir != "." {
		writeError(w, errors.New("path too long by: "+dir), nil)
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cwd := filepath.Join(wd, path)
	these := []*Box{}
	for i, box := range boxes {
		this := &Box{}
		this.Path = cwd
		this.Name = box.Name
		this.Title = box.Title
		this.Func = box.Func
		this.Content = box.Content
		this.Help = box.Help
		data, err := ioutil.ReadFile(filepath.Join(cwd, boxes[i].Name+".box"))
		if err == nil {
			this.Content = string(data)
		}
		these = append(these, this)
	}
	head.Execute(w, nil)
	for _, box := range these {
		style.Execute(w, box)
	}
	headFuncs.Execute(w, nil)
	for _, box := range these {
		funcs.Execute(w, box)
	}
	tailFuncs.Execute(w, req.URL.Path[5:])
	for _, box := range these {
		table.Execute(w, box)
	}
	tail.Execute(w, nil)
}

func writeError(w http.ResponseWriter, err error, more []byte) {
	w.WriteHeader(404)
	w.Write([]byte("<pre>" + err.Error() + "\n" + string(more) + "</pre>"))
}

func run(dir string, args ...string) ([]byte, []byte, error) {
	var buf bytes.Buffer
	var errBuf bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdout = &buf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	return buf.Bytes(), errBuf.Bytes(), err
}

func register(box *Box) {
	boxes = append(boxes, box)
	sort.Sort(boxes)
}
