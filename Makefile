###########################################################################
#                             Mike's Makefile                              
#                  Copyright (C) 2022 - Michael Jeuga
#
###########################################################################

repo=$(shell basename "`pwd`")

gopher:
	@git init
	@touch main.go
	@touch Dockerfile
	@touch docker-compose.yml
	@go get github.com/google/uuid
	@go get github.com/adamluzsi/testcase
	@go get github.com/gorilla/mux
	@go mod init github.com/mikejeuga/$(repo)
	@go mod tidy

run:
	@go run ./cmd/main.go

t: test
test:
	@go test -v ./...


ut: unit-test
unit-test:
	@go test -v -tags unit ./...

at: acceptance-test
acceptance-test:
	@go test -v -tags acceptance ./...

ic: init
init:
	@git add .
	@git commit -m "Initial commit"
	@git remote add origin git@github.com:mikejeuga/${repo}.git
	@git branch -M main
	@git push -u origin main

c: commit
commit:
	@git add .
	@git commit -m "$m"
	@git pull --rebase
	git push
