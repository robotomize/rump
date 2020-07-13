APP?=rump
NAME?=rump
RELEASE?=0.0.1
GOOS?=linux
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)

test:
	go test -v -cover -covermode=atomic ./...

.PHONY: build
build: clean
# build server
	CGO_ENABLED=0 GOOS=${GOOS} go build ./cmd/syncrcvpos \
		-ldflags "-X main.version=${RELEASE}  -X main.buildTime=${BUILD_TIME} -X main.name=${NAME}" \
		-o build/${APP}
# build client tool
	CGO_ENABLED=0 GOOS=${GOOS} go build ./tools/rcvposcli \
		-ldflags "-X main.version=${RELEASE}  -X main.buildTime=${BUILD_TIME} -X main.name=${NAME}" \
		-o build/${APP}

	CGO_ENABLED=0 GOOS=${GOOS} go build ./tools/syncposcli \
		-ldflags "-X main.version=${RELEASE}  -X main.buildTime=${BUILD_TIME} -X main.name=${NAME}" \
		-o build/${APP}

.PHONY: clean
clean:
	@rm -f build/${APP}