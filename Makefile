docker-build:
	sudo docker build -t arborist - < arborist.docker

docker-run:
	sudo docker rm arborist-container | true
	mkdir -p /tmp/arborist/shared
	chmod 777 /tmp/arborist/shared
	sudo docker run -d -t -p 80:8080 -v /tmp/arborist/shared:/shared:rw --name arborist-container arborist
	sudo docker ps

docker-stop:
	sudo docker kill arborist-container
	sudo docker ps

docker-restart:
	make docker-stop
	make docker-run

docker-rebuild:
	make docker-build
	make docker-restart

docker-attach:
	sudo docker attach arborist-container

gofmt:
	gofmt -l -s -w .