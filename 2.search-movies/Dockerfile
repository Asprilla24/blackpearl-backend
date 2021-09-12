FROM golang:1.16 AS builder
WORKDIR /go/src/search-movies/
COPY . ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o ./service ./cmd/...
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o ./client ./client/...

FROM alpine:latest AS release
WORKDIR /root/
COPY --from=builder /go/src/search-movies/service ./
COPY --from=builder /go/src/search-movies/client ./
RUN find $HOME -type d -exec 'chmod' '555' '{}' ';' && \
    find $HOME -type f -exec 'chmod' '444' '{}' ';' && \
    find $HOME -type f -exec 'chown' 'root:root' '{}' ';' && \
    chmod 555 ./service && chmod 555 ./client
USER nobody
ENTRYPOINT [ "./service" ]