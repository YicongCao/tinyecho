img_ver=v1
img_name=tinyecho
img_repo=hub.docker.com/sample/

build:
	docker build -t $(img_repo)$(img_name):$(img_ver) .
	docker image ls | grep $(img_name)

export:
	docker save -o ../$(img_name)$(img_ver).tar $(img_repo)$(img_name):$(img_ver)

debug:
	docker run -it --rm --name "hello_instance" -p 8081:8080 $(img_repo)$(img_name):$(img_ver)

push:
	docker push $(img_repo)$(img_name):$(img_ver)