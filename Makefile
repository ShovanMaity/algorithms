test: 	dep vet fmt
	./test.sh
dep:
	go get github.com/stretchr/testify/assert
vet:
	go list ./... | xargs go vet
fmt:
	find . -type f -name "*.go" | xargs gofmt -s -w -l
.PHONY: test
