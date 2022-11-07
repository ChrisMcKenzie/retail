FROM golang:alpine AS builder

WORKDIR $GOPATH/src/ChrisMcKenzie/retail/
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -ldflags '-w -s' \
    -o /go/bin/retail .

FROM scratch

COPY --from=builder /go/bin/retail /go/bin/retail

ENTRYPOINT ["/go/bin/retail"]

