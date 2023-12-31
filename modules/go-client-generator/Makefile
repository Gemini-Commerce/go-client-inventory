SHELL := /bin/bash

CUR_UID := $(shell id -u)
CUR_GID := $(shell id -g)

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))
openapi_generator = docker run --rm --user ${CUR_GID}:${CUR_UID} -v ${current_dir}/:/go-client-generator -v ${PWD}/:/go-client-${lower_service} openapitools/openapi-generator-cli:v7.1.0
lower_service = $(shell echo $(service) | sed -e 's/\(.*\)/\L\1/')
ucf_service = $(shell echo $(service) | sed -e "s/\b\(.\)/\u\1/g")
artifact_version = $(shell if [ ! -z "${version}" ]; then echo "artifactVersion=${version},"; fi)

generate-go-client:
	[ ! -z "${service}" ] || (printf "\e[31mError: this command is meant to be execute from specific service client repo\e[0m\n"; exit 1) && \
	${openapi_generator} \
	generate -i go-client-generator/modules/proto/${lower_service}/${lower_service}.swagger.json \
	-g go -t go-client-generator/templates/go -o go-client-${lower_service} \
	--git-host=github.com \
	--git-user-id=gemini-commerce \
	--git-repo-id=go-client-${lower_service} \
	--additional-properties ${artifact_version}\
	structPrefix=true,\
	packageName="${lower_service}"
