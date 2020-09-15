FROM golang:1.14 AS builder
ENV CGO_ENABLED 0
WORKDIR /go/src/app
ADD . .
RUN go build -mod vendor -o /patchwork

FROM alpine:3.12
COPY --from=builder /patchwork /patchwork
CMD ["/patchwork"]