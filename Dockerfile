FROM golang:1.14.2 as builder

RUN mkdir -p /go/src/github.com/boringdao/bridge
WORKDIR /go/src/github.com/boringdao/bridge

# Cache dependencies
COPY go.mod .
COPY go.sum .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download -x

# Build real binaries
COPY . .

RUN go get -u github.com/gobuffalo/packr/packr

# Build bitxhub node
RUN make install

# Final image
FROM frolvlad/alpine-glibc
RUN mkdir -p /root/.bridge
WORKDIR /

# Copy over binaries from the builder
COPY --from=builder /go/bin/bridge /usr/local/bin
COPY --from=builder /go/bin/packr /usr/local/bin


EXPOSE 4001

ENTRYPOINT ["bridge", "start"]


