FROM golang
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get google.golang.org/grpc
RUN go get github.com/mattn/go-sqlite3
ADD . /go/src/raftAlgo.com/service/

