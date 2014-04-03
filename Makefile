all: jconfig

jconfig: test

test: config_test.go
	@go version
	go test