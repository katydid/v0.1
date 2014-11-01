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

RUN apt-get install -y python-software-properties git wget mercurial build-essential protobuf-compiler graphviz net-tools

ENV GOVERSION 1.3
ENV GODOWNLOAD http://golang.org/dl/
ENV GOFILENAME go$GOVERSION.linux-amd64.tar.gz
RUN wget $GODOWNLOAD/$GOFILENAME && \
	tar -C / -xzf $GOFILENAME && \
	rm $GOFILENAME

ENV GOROOT /go
ENV GOPATH /gopath
ENV PATH $PATH:$GOPATH/bin:$GOROOT/bin

RUN mkdir $GOPATH
RUN mkdir -p $GOPATH/src/code.google.com/p
RUN mkdir -p $GOPATH/src/github.com/awalterschulze
RUN mkdir -p $GOPATH/src/github.com/katydid

ENV SHAREPATH shared
RUN mkdir $SHAREPATH && \
	chmod 777 $SHAREPATH

# RUN go get -v code.google.com/p/go.text/unicode/norm
ENV GOTEXTPATH $GOPATH/src/code.google.com/p/go.text
RUN hg clone https://code.google.com/p/go.text/ $GOTEXTPATH && \
	(cd $GOTEXTPATH && hg checkout f8db539672d0 ) && \
	(cd $GOTEXTPATH && go install ./... )

ENV GOCCPATH /gopath/src/code.google.com/p/gocc
RUN git clone https://code.google.com/p/gocc $GOCCPATH && \
	(cd $GOCCPATH && git checkout 87a3f29bc7f9 ) && \
	(cd $GOCCPATH && go install ./...)

ENV GOGOPATH $GOPATH/src/code.google.com/p/gogoprotobuf
RUN git clone https://code.google.com/p/gogoprotobuf $GOGOPATH && \
	(cd $GOGOPATH && git checkout 6c9802773308 ) && \
	(cd $GOGOPATH && make)

ENV KATYPATH $GOPATH/src/github.com/awalterschulze/katydid
ENV EXAMPLEPATH example
RUN git clone https://github.com/awalterschulze/katydid $KATYPATH && \
	(cd $KATYPATH && git checkout e69e084bd1 ) && \
	(cd $KATYPATH && make) && \
	mkdir $EXAMPLEPATH && \
	(cd $KATYPATH/asm/test && go test -c && ./test.test) && \
	(mv $KATYPATH/asm/test/example/* /$EXAMPLEPATH/)

ENV ARBOPATH $GOPATH/src/github.com/katydid/arborist
EXPOSE 8080
RUN mkdir $ARBOPATH

# run webserver
RUN git clone https://github.com/katydid/arborist $ARBOPATH && \
	go install .
CMD ["arborist"]
USER daemon

# development environment
# ENTRYPOINT bash

