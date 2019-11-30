-include .$(env).sh

dependencies:
	go get github.com/golang/dep/cmd/dep
	dep ensure -v

	go get github.com/alecthomas/gometalinter
	gometalinter --install

spec:
	go test -v -tags=spec -race -cover ./...

feature:
	go test -v -tags=feature -race -cover -coverprofile=feature_coverage.out  ./...

coverage:
	go tool cover -func=feature_coverage.out

analysis:
	gometalinter --vendor ./... --disable-all --enable=gofmt --enable=golint --enable=vet --enable=errcheck
