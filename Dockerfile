# golang image
FROM golang:latest

# golang image 的 GOPATH 为：/go
WORKDIR /go/src/github.com/xiaosongfu/link-wxmp

# 复制当前目录下的文件到 WORKDIR
COPY . .

# install dep and then run it
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure

# clean and build
RUN go clean && go install

CMD ["link-wxmp"]