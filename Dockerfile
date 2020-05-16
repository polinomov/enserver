FROM golang:buster

EXPOSE 5555
EXPOSE 1234

RUN apt-get update
RUN apt-get install -y protobuf-compiler
RUN apt-get install -y libczmq-dev
RUN go get github.com/golang/protobuf/proto
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get gopkg.in/zeromq/goczmq.v4

RUN mkdir /app 
ADD app/libEnserver  $GOPATH/src/libEnserver
ADD app/game $GOPATH/src/game

ADD app/protodef/ /home/protodef
ADD app/build.sh /home/build.sh

RUN /home/build.sh
CMD ["enserver"]
