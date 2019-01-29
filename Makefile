# https://www.gnu.org/software/make/manual/make.html
BASE=alpine_base
# TOP_DIR=$(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
TOP_DIR=$(shell dirname $(abspath $(firstword $(MAKEFILE_LIST))))
PROJECT=$(notdir $(TOP_DIR))
REPO=$(BASE)/$(PROJECT)
# Dockerfile | SDK.Dockerfile | SDKMAN.Dockerfile
DOCKER_FILE=Dockerfile
# docker-compose.yml | docker-compose.yaml
DOCKER_COMPOSE_FILE=docker-compose.yml
# alpine_network | <docker-compose-project>_alpine_network
DOCKER_NETWORK=alpine_network
OFFICIAL_PROJECT=$(shell echo $(PROJECT) | cut -d'_' -f2)

ifneq ("$(wildcard VERSION.txt)", "")
	TAG=$(shell grep -i version VERSION.txt | cut -d '=' -f 2  | tr -d '[:space:]')
	OFFICIAL_TAG=latest
else ifdef LATEST
	TAG=latest
	OFFICIAL_TAG=latest
else
	TAG=latest
	OFFICIAL_TAG=latest
endif

# COMMAND=docker run --rm -ti -v $(TOP_DIR)/$(SRC_FOLDER):$(PROJECT_FOLDER) -h $(PROJECT) --name $(PROJECT) $(REPO):$(TAG)
COMMAND=

ifneq ("$(wildcard ${HOME}/.env)", "")
	include ${HOME}/.env
else ifneq ("$(wildcard env)", "")
	include env
else
	GIT_USER=cclhsu
	GIT_USER_PASSWORD=
	GIT_USER_EMAIL=cclhsu@yahoo.com
	DOCKER_ID_USER=cclhsu
	DOCKER_PASSWORD=
endif

# project | HelloWorld
PROJECT_FOLDER=/HelloWorld
# . | src | example
SRC_FOLDER=.

##### HOWTO #####

# make package; make build_docker; make run_docker
# make package; make run_package
# make package; make run

##### ##### #####

.DEFAULT_GOAL := help

.PHONY: help
help:
	@make -rpn | sed -n -e '/^$$/ { n ; /^[^ .#][^ ]*:/ { s/:.*$$// ; p ; } ; }' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@#make -rpn | sed -n -e '/^$/ { n ; /^[^ ]*:/p ; }' | egrep --color '^[^ .]*:' | sort
	@#grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: compile

.PHONY: project_init
project_init:
	go mod init ${PROJECT} && \
	mkdir ${TOP_DIR}/bin && \
	mkdir ${TOP_DIR}/cmd && \
	mkdir ${TOP_DIR}/pb && \
	mkdir ${TOP_DIR}/proto

.PHONY: install_dependency_packages
install_dependency_packages:
	export GOPATH=${HOME}/go && \
	export PATH="${PATH}:$(shell go env GOPATH)/bin" && \
	go get -u <URL>


# .PHONY: build
# build:
# 	# export GOPATH=${TOP_DIR}
# 	export GOBIN=${TOP_DIR}/bin && \
# 	go env && \
# 	$(COMMAND) go build -o ${TOP_DIR}/bin/hello ${TOP_DIR}/src/hello.go

# .PHONY: build
# build:
# 	# export GOPATH=${TOP_DIR}
# 	export GOBIN=${TOP_DIR}/bin && \
# 	go env && \
# 	$(COMMAND) go build -o ${TOP_DIR}/bin/app app

.PHONY: compile_linux
compile_linux:
	# export GOPATH=${TOP_DIR}
	export GOBIN=${TOP_DIR}/bin && \
	export GOOS=linux && \
	export GOARCH=amd64 && \
	go env && \
	$(COMMAND) go install ${TOP_DIR}/src/hello.go

.PHONY: compile
compile:
	# export GOPATH=${TOP_DIR}
	export GOBIN=${TOP_DIR}/bin && \
	go env && \
	$(COMMAND) go install ${TOP_DIR}/src/hello.go

.PHONY: install
install:
	# export GOPATH=${TOP_DIR}
	export GOBIN=${TOP_DIR}/bin && \
	go env && \
	$(COMMAND) go install app

.PHONY: run
run:
	$(COMMAND) ./bin/hello || :
	$(COMMAND) ./bin/app || :

.PHONY: script
script:
	@# ./hello.sh
	$(COMMAND) go run ${TOP_DIR}/src/hello.go

.PHONY: package
package:
	make clean
	make compile
	@# make compile_linux

.PHONY: run_package
run_package:
	$(COMMAND) go ${TOP_DIR}/bin/hello

.PHONY: clean
clean:
	@# $(COMMAND) gradle clean
	rm -rf ${TOP_DIR}/bin/*
	rm -rf ${TOP_DIR}/api/pb
	rm -rf ${TOP_DIR}/api/swagger

.PHONY: bash
bash:
	$(COMMAND) /bin/bash

.PHONY: build_docker
build_docker:
	make clean_docker
	# docker build -f ${DOCKER_FILE} -t $(REPO):$(TAG) .
	docker build --rm -f ${DOCKER_FILE} -t $(REPO):$(TAG) .

.PHONY: clean_docker
clean_docker:
	docker rm -f $(PROJECT) 2>/dev/null; true
	docker rmi $(REPO):$(TAG) 2>/dev/null; true
	docker rm $(shell docker ps -a -q) 2>/dev/null; true
	docker rmi $(shell docker images -f "dangling=true" -q) 2>/dev/null; true
	docker volume prune -f 2>/dev/null; true
	rm -rf tmp

.PHONY: run_docker
run_docker:
	@# $(COMMAND)
	docker run --rm -ti -v $(TOP_DIR):$(PROJECT_FOLDER) -v $(TOP_DIR)/entrypoint.sh:/usr/bin/entrypoint.sh -h $(PROJECT) --name $(PROJECT) $(REPO):$(TAG)

.PHONY: run_docker_with_mount_volume
run_docker_with_mount_volume:
	@# docker run --rm -ti -v $(TOP_DIR)/$(SRC_FOLDER):$(PROJECT_FOLDER) -h $(PROJECT) --name $(PROJECT) $(REPO):$(TAG) <command>
	docker run --rm -ti -v $(TOP_DIR):$(PROJECT_FOLDER) -h $(PROJECT) --name $(PROJECT) $(REPO):$(TAG)

.PHONY: run_docker_with_mount_entrypoint
run_docker_with_mount_entrypoint:
	@# docker run --rm -ti -v $(TOP_DIR)/entrypoint.sh:/usr/bin/entrypoint.sh -h $(PROJECT) --name $(PROJECT) $(REPO):$(TAG) <command>
	docker run --rm -ti -v $(TOP_DIR)/entrypoint.sh:/usr/bin/entrypoint.sh -h $(PROJECT) --name $(PROJECT) $(REPO):$(TAG)

.PHONY: bash_docker
bash_docker:
	@# $(COMMAND) /bin/bash
	docker run --rm -ti -h $(PROJECT) --name $(PROJECT) $(REPO):$(TAG) /bin/bash

.PHONY: bash_container
bash_container:
	bash -c "clear && docker exec -it $(PROJECT) /bin/bash"

.PHONY: list_docker
list_docker:
	docker ps -a
	docker volume ls
	docker images

.PHONY: publish
publish:
	@# https://docs.docker.com/docker-cloud/builds/push-images/
	@# docker login -u $(DOCKER_ID_USER) -p $(DOCKER_PASSWORD)
	docker tag $(REPO):$(TAG) $(DOCKER_ID_USER)/$(PROJECT)
	docker push $(DOCKER_ID_USER)/$(PROJECT)

.PHONY: untag_publish
untag_publish:
	docker rmi $(DOCKER_ID_USER)/$(PROJECT):$(TAG)

.PHONY: dockerhub
dockerhub:
	open https://hub.docker.com/u/${DOCKER_ID_USER}/

.PHONY: publish_to_registry
publish_to_registry:
	@# https://docs.docker.com/docker-cloud/builds/push-images/
	@# docker login -u $(DOCKER_ID_USER) -p $(DOCKER_PASSWORD)
	# docker tag $(REPO):$(TAG) $(DOCKER_ID_USER)/$(PROJECT)
	# docker push $(DOCKER_ID_USER)/$(PROJECT)
	docker tag alpine:latest 0.0.0.0:5000/alpine/alpine:latest
	docker push 0.0.0.0:5000/alpine/alpine:latest
	curl -X GET http://0.0.0.0:5000/v2/_catalog
	curl -X GET http://0.0.0.0:5000/v2/alpine/alpine/tags/list

.PHONY: untag_publish_to_registry
untag_publish_to_registry:
	docker rmi $(DOCKER_ID_USER)/$(PROJECT):$(TAG)

.PHONY: git_init
git_init:
	git config --global user.name ${GIT_USER}
	git config --global user.email ${GIT_USER_EMAIL}
	curl -u ${GIT_USER}:${GIT_USER_PASSWORD} https://api.github.com/user/repos -d "{\"name\":\"${PROJECT}\"}"
	touch README.md
	touch .gitignore
	git init
	git checkout -b develop
	git add .gitignore
	git commit -m "first commit"
	git add *
	git rm -r --cached Makefile
	git commit -m "add project skeleton"
	git checkout -b master
	git merge develop
	git flow init -fd
	# git remote add origin git@github.com:${GIT_USER}/${PROJECT}.git
	git remote add origin https://github.com/${GIT_USER}/${PROJECT}.git
	echo ${GIT_USER_PASSWORD} | git push -u origin master
	git remote set-head origin master
	git remote set-head origin -a
	echo ${GIT_USER_PASSWORD} | git push --all

.PHONY: git_push
git_push:
	git push

.PHONY: git_delete_github_repository
git_delete_github_repository:
	curl -u ${GIT_USER}:${GIT_USER_PASSWORD} -X DELETE https://api.github.com/repos/${GIT_USER}/${PROJECT}