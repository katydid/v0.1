.PHONY: nuke regenerate gofmt build test

all: nuke regenerate gofmt build test

test:
	go test -v ./...

build:
	go build ./...

regenerate:
	(cd exp/asm && gocc asm.bnf)
	(cd exp/asm/ast && protoc --gogo_out=. -I=.:../../../../../../ asm.proto)
	(cd exp/asm/test && protoc --gogo_out=. -I=.:../../../../../../ test.proto)
	make gofmt

clean:
	(cd exp/asm && rm *.txt || true)

nuke: clean
	rm -rf exp/asm/errors
	rm -rf exp/asm/lexer
	rm -rf exp/asm/parser
	rm -rf exp/asm/token
	rm -rf exp/asm/util

gofmt:
	gofmt -l -s -w .