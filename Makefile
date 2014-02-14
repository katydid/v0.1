# to attach to a docker for debugging
# lxc-attach -n <full docker id>
# this is not recommended

docker-build:
	docker build -t arborist - < arborist.docker

docker-run:
	docker run -d -p 8080:8080 arborist