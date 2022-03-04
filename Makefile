GOCMD?=go

build:
	${GOCMD} build -o bin/detect cmd/main.go
	[ -f ./bin/build ] || (cd bin && ln -s ./detect build)

test: build
	${GOCMD} test ./integration/*.go