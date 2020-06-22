package main

import (
	"log"
	"gopkg.in/zeromq/goczmq.v4"
  "github.com/golang/protobuf/proto"
  pb "github.com/polinomov/enserver/enbuffer/cmd"
)

type DebugSocket struct{
  socket* goczmq.Sock
  port string
  running bool
  mainThread *ServerData
};

func ( c *DebugSocket) Init(server *ServerData,port string) bool {
  log.Println("DebugSocket Init called with port ",port)
  repsocket, err := goczmq.NewRep("tcp://*:" + port)
  if err != nil{
    log.Println(err);
    return false;
  }

  repsocket.Bind("tcp://*:" + port)
  c.socket = repsocket
  c.port = port
  c.mainThread = server
  c.running = true;
  go c.run()
  return true;
}

func (c *DebugSocket) Destroy(){
  c.running = false
}

func (c *DebugSocket) run(){
  log.Println("DebugSocket started")
  for c.running {
    log.Println("Reading Frame")
	recdata,_,rerr := c.socket.RecvFrame()
	log.Println("Reading Frame HELLO-----------------")
    if rerr != nil {
      log.Println(rerr)
      continue
    }

    command := &pb.DebugCommand{}
    if err := proto.Unmarshal(recdata, command);
    err != nil {
      log.Println("DebugSocket Error: Failed to Unmarshal Command", err)
      continue
	}
	
	log.Printf("Queing---------------- %p\n", c.mainThread)
    c.mainThread.QueEvent(func (server *ServerData) bool {
      log.Println("Hello From MainThread")
      return true
    })
	log.Println("Qued------------------")

	log.Println("Recieved Debug Command") 
    reply := []byte("Hello")
    c.socket.SendFrame(reply,goczmq.FlagNone)
  }
  log.Println("DebugSocket closed")
}
