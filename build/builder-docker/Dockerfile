# Copyright 2017 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:1.10.1-alpine3.7 as builder

RUN apk add --no-cache git curl openssl

RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get github.com/chai2010/protorpc/protoc-gen-stdrpc
RUN go get golang.org/x/tools/cmd/goimports

# go-swagger
# RUN go get github.com/go-swagger/go-swagger/cmd/swagger
RUN mkdir -p /swagger && cd /swagger \
    && wget https://github.com/go-swagger/go-swagger/releases/download/0.14.0/swagger_linux_amd64 \
    && chmod +x swagger_linux_amd64 && mv swagger_linux_amd64 /go/bin/swagger

# the protoc can't run on alpine,
# we only need the protobuf's stdarnd library in the `/protoc/include`.
RUN mkdir -p /protoc && cd /protoc \
    && wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip \
    && unzip protoc-3.5.1-linux-x86_64.zip

RUN cd /go/src/github.com/golang/protobuf/protoc-gen-go \
    && echo "github.com/golang/protobuf/protoc-gen-go:" >> /VERSION \
    && git log -n 1 --pretty=format:%cI >> /VERSION \
    && cd /go/src/github.com/grpc-ecosystem/grpc-gateway \
    && echo >> /VERSION \
    && echo "github.com/grpc-ecosystem/grpc-gateway:" >> /VERSION \
    && git log -n 1 --pretty=format:%cI >> /VERSION \
    && cd /go/src/github.com/chai2010/protorpc \
    && echo >> /VERSION \
    && echo "github.com/chai2010/protorpc:" >> /VERSION \
    && git log -n 1 --pretty=format:%cI >> /VERSION \
    && cd /go/src/golang.org/x/tools \
    && echo >> /VERSION \
    && echo "golang.org/x/tools:" >> /VERSION \
    && git log -n 1 --pretty=format:%cI >> /VERSION \
    && echo >> /VERSION

FROM golang:1.10.1-alpine3.7

RUN apk add --no-cache git protobuf make curl openssl jq rsync

COPY --from=builder /protoc/include /usr/local/include
COPY --from=builder /go/bin /go/bin
COPY --from=builder /VERSION /VERSION
