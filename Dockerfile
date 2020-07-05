FROM golang
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get google.golang.org/grpc
ADD . /go/src/raftAlgo.com/service/

