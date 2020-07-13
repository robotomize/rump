FROM golang:1.14 AS builder

RUN apt-get -qq update && apt-get -yqq install upx

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /src
COPY . .

RUN go build \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -installsuffix cgo \
  -tags netgo \
  -o /bin/service \
  ./cmd/syncrcvpos

RUN strip /bin/service
RUN upx -q -9 /bin/service


FROM scratch
COPY --from=builder /bin/service /bin/service

ENV SYNC_ADDR :5555
ENV RCV_ADDR :5577

ENTRYPOINT ["/bin/service"]
