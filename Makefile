docker-build:
	docker build -t arborist .

docker-run:
	docker rm arborist-container | true
	mkdir -p /tmp/arborist/shared
	chmod 777 /tmp/arborist/shared
	docker run -d -t -p 80:8080 -v /tmp/arborist/shared:/shared:rw --name arborist-container arborist
	docker ps

docker-run-mac:
	docker rm arborist-container | true
	mkdir -p /Users/walter/tmp/arborist/shared
	chmod 777 /Users/walter/tmp/arborist/shared
	docker run -d -t -p 80:8080 -v /Users/walter/tmp/arborist/shared:/shared:rw --name arborist-container arborist
	docker ps

docker-build-env-mac:
	docker run --rm=true -p 80:8080 -i -t -v /Users/walter/code/gopath/src/github.com/katydid/arborist:/gopath/src/github.com/katydid/arborist:rw -v /Users/walter/tmp/arborist/shared:/shared:rw --name arborist-container arborist 

docker-open-web-mac:
	open http://$(boot2docker ip 2>/dev/null)/

docker-stop:
	docker kill arborist-container
	docker ps

docker-restart:
	make docker-stop
	make docker-run

docker-rebuild:
	make docker-build
	make docker-restart

docker-attach:
	docker attach arborist-container

gofmt:
	gofmt -l -s -w .