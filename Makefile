BUILD_PARAMS="-tags=release"

all: build

build:
	go build "${BUILD_PARAMS}" .

clean:
	rm -f ./fly-exporter

fmt:
	find . -name \*.go -not -path bin/ -exec goimports -w {} \;

test:
	go test "${TEST_PARAM}" ./...
