FROM golang:buster

EXPOSE 5555
EXPOSE 1234

RUN apt-get update
RUN apt-get install sudo
RUN apt-get install -y protobuf-compiler
RUN apt-get install -y libczmq-dev
RUN go get github.com/golang/protobuf/proto
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get gopkg.in/zeromq/goczmq.v4

#RUN useradd -m docker && echo "docker:docker" | chpasswd && adduser docker sudo
#USER docker

ADD app $GOPATH/src/enserver
WORKDIR $GOPATH/src/enserver 
RUN ./build.sh
CMD ["enserver"]
