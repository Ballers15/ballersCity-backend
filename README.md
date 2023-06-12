# Ballers City Backend

## Description
   Game backend is built in golang and postgresql as DB and being  run in dockerized container 

## Prerequisites
Make sure you have installed all of the following prerequisites on your development machine: Git - Download & Install Git. OSX and Linux machines typically have this already installed.Install Golang  go1.18 - Download & Install GO.

## Dependencies

```
go install github.com/spf13/cobra-cli@latest
go install github.com/luckybet100/protodeps@v1.0.3
```
## Project structure
```
-> internal // code not meant to be imported from outside
-> -> cmd // direct run code with parameters
-> api // protobuf+grpc files
-> pb // implementation of proto on go (does not commit to git)
-> pkg //code that can be imported from outside
-> cmd // scripts for running tasks (reading from flags and environment variables)
```
## Updated Branch

master

## Working Branch

master

## Get Started
Get started developing...

```shell
# clone backend in your local system
git clone https://github.com/Ballers15/ballersCity-backend.git
git checkout master
go run main.go
``

## Development server

Run `go run main.go` for a dev server.

## Staging server

[Staging Server](https://staging.ballers.fun/).

