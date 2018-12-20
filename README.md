# hahiye  [![CircleCI Build Status](https://circleci.com/gh/mucyomiller/hahiye.svg?style=shield)](https://circleci.com/gh/mucyomiller/hahiye) [![GoDoc](https://godoc.org/github.com/mucyomiller/hahiye?status.svg)](https://godoc.org/github.com/mucyomiller/hahiye) [![Go Report Card](https://goreportcard.com/badge/mucyomiller/hahiye)](https://goreportcard.com/report/mucyomiller/hahiye) [![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/mucyomiller/hahiye/master/LICENSE)
something new(project description soon... )   

### _Run project:_
```go run main.go```   
### _Generate protobuffer from proto file:_   
```protoc -I proto/ proto/hahiye.proto  --go_out=plugins=grpc:hahiye/```   
## pre-requisite   
This project depends on dgraph database   
run dgraph DB by:   
```shell
$ dgraph zero
```
```shell
$ dgraph server --lru_mb=2048
```
```shell
$ dgraph-ratel
```

Or simply if you use Docker   
```shell
$ docker-compose up
```

