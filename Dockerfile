#  Copyright 2014 Walter Schulze
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

FROM ubuntu

RUN sed 's/main$/main universe/' -i /etc/apt/sources.list
RUN apt-get update

RUN apt-get install -y python-software-properties
RUN apt-get install -y git
RUN apt-get install -y wget
RUN apt-get install -y mercurial
RUN apt-get install -y build-essential
RUN apt-get install -y protobuf-compiler
RUN apt-get install -y graphviz
RUN apt-get install -y net-tools

RUN wget 'https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz'
RUN tar -C / -xzf go1.2.linux-amd64.tar.gz
RUN rm go1.2.linux-amd64.tar.gz
ENV GOROOT /go

RUN mkdir gopath
ENV GOPATH /gopath
ENV PATH $PATH:$GOPATH/bin:$GOROOT/bin

RUN mkdir -p /gopath/src/code.google.com/p
RUN git clone https://code.google.com/p/gocc /gopath/src/code.google.com/p/gocc
RUN (cd /gopath/src/code.google.com/p/gocc && git checkout gocc2 && go install ./...)

RUN git clone https://code.google.com/p/gogoprotobuf /gopath/src/code.google.com/p/gogoprotobuf
RUN (cd /gopath/src/code.google.com/p/gogoprotobuf && make)

RUN go get -v code.google.com/p/go.text/unicode/norm

RUN mkdir shared
RUN chmod 777 shared

ADD ./README.md /forceUpdateFromThisPoint.md

RUN mkdir -p /gopath/src/github.com/awalterschulze
RUN git clone https://github.com/awalterschulze/katydid /gopath/src/github.com/awalterschulze/katydid
RUN (cd /gopath/src/github.com/awalterschulze/katydid/ && make)

RUN mkdir example
RUN (cd /gopath/src/github.com/awalterschulze/katydid/asm/test && go test -c && ./test.test)
RUN (mv /gopath/src/github.com/awalterschulze/katydid/asm/test/example/* /example/)

RUN git clone https://github.com/awalterschulze/arborist /gopath/src/github.com/awalterschulze/arborist
RUN (cd /gopath/src/github.com/awalterschulze/arborist && go install .)

CMD ["arborist"]
USER daemon
EXPOSE 8080
