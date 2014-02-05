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
	"crypto/rand"
	"encoding/base64"
	"os"
	"path/filepath"
	"sync"
)

var paths = &Path{}

type Path struct {
	sync.Mutex
}

func NewPath() *Path {
	return &Path{}
}

func randURL(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (this *Path) new(prefix string) string {
	this.Lock()
	defer this.Unlock()
	i := 4
	p := filepath.Join(prefix, randURL(i))
	for exists(p) {
		i++
		p = filepath.Join(prefix, randURL(i))
	}
	err := os.MkdirAll(p, 0777)
	if err != nil {
		panic(err)
	}
	return p
}

func (this *Path) NewTemp() string {
	return this.new("./tmp/")
}

func (this *Path) NewShared() string {
	return this.new("./shared/")
}

func exists(filename string) bool {
	_, err := os.Lstat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}
