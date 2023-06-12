# Ballers City Backend
  Game backend is built in golang and postgresql as DB and being  run in dockerized container 
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
