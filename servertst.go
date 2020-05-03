
/*
 go install google.golang.org/protobuf/cmd/protoc-gen-go

 To generate proto go file call:
   protoc --go_out ./ protodef/enginecmd.proto
*/

package main

import (
    "fmt"
	"log"
	"gopkg.in/zeromq/goczmq.v4"
    "time"
    "plugin"
    "os"
   // "io/ioutil"
    //"github.com/golang/protobuf/proto"
     pb "github.com/polinomov/enserver/enbuffer/cmd"
 )


func fromClient(cmdBuff chan pb.Command)  {
    fmt.Printf("Start client channel type : %T \n",cmdBuff) 
    pullsocket,err := goczmq.NewPull("tcp://0.0.0.0:1234") 
    if err != nil {
        log.Fatal(err)
        return;
    }
    defer pullsocket.Destroy()
    var i = 0
     for {
        //log.Printf("About to read")
        recdata, rerr :=pullsocket.RecvMessage()
        if rerr != nil {
           log.Fatal(rerr)
        }
        var ucmd = unmarshalCommand(recdata[0])
        fmt.Printf( "--- len---=%d \n",len(ucmd.Cmd))

        cmdBuff <- *ucmd.Cmd[0]
       // time.Sleep(time.Millisecond * 100)
        i = i + 1
    }
 }

 func toClient(cmdBuff chan pb.Command){
   // var opt = goczmq.SockSetSndbuf(1)
    pubsocket, err := goczmq.NewPub("tcp://*:5555")
    if err != nil {
        log.Fatal(err)
        return;
    }
    defer pubsocket.Destroy()
    pubsocket.Bind("tcp://*:5555")
    var rec = *addOneCommand(0)
    for {
        rec = <- cmdBuff
        var dat = marshalCommand(&rec)
        pubsocket.SendFrame(dat, goczmq.FlagNone)
        fmt.Printf(" SendFrame\n")
    }
 }

 func soCallBack( s string){
    fmt.Printf("soCallBack %s\n", s)  
 }

 func loadPlugIn( plgname string){
    _, err := os.Stat(plgname)
    if err != nil {
        fmt.Printf("Can not find file %s\n", plgname)
    } else {
        fmt.Printf("found %s\n", plgname)
    }
    
    p, err := plugin.Open(plgname)
    if err != nil {
        log.Fatal(err)
    } 
    v, err := p.Lookup("V")
    if err != nil {
	    log.Fatal(err)
    }
    f, err := p.Lookup("F1")
    if err != nil {
	    log.Fatal(err)
    }   
    *v.(*int) = 7
    f.(func())() 

    msg, err := p.Lookup("Msg")
    if err != nil { log.Fatal(err) }
    *msg.(*string) = "BLAH"

   
    
    cbFunc, err := p.Lookup("TheCallBack")
    if err != nil { log.Fatal(err) }
    var xerr = *cbFunc.(*func(string))
    *cbFunc.(*func(string))  =  soCallBack
    fmt.Printf("xerr type : %T \n",  xerr)

    procFunc, err := p.Lookup("F2")
    if err != nil { log.Fatal(err) }
    procFunc.(func())()
 }


func main(){
  log.Println("MAIN PUBSUB1")
  /*
  d1 := []byte("hello\ngo\n")
  err := ioutil.WriteFile("datxxxxxx", d1, 0644)
  if err != nil {
    panic(err)
  }  
  */
  loadPlugIn("libgame.so")

 // var cc = addOneCommand(123);
 // var dat = marshalCommand(cc)
 // unmarshalCommand(dat)

  cmdBuff := make(chan pb.Command, 32)
  go fromClient(cmdBuff)
  go toClient(cmdBuff)
  for{
    time.Sleep(time.Millisecond*1000) 
  }
  //log.Println("MAIN DONE")
  
  /*
  var opt = goczmq.SockSetSndbuf(1)
  subsocket, _ := goczmq.NewPub("tcp://*:5555",opt)
 
  defer subsocket.Destroy()
  subsocket.Bind("tcp://*:5555")
  rand.Seed(time.Now().UnixNano())
  
   var bb [1024*1024]byte

  
   var i = 0
   rand.Seed(time.Now().UnixNano())
    // loop for a while aparently
    for {
       //  zipcode := rand.Intn(100000)
       //  temperature := rand.Intn(215) - 80
        // relhumidity := rand.Intn(50) + 10
		var tc = time.Now().UnixNano()
        i = i+ 1
		//bb[0]  = 0;
		//bb[255] = 123
		
        msg := fmt.Sprintf("%d %l %s", i,tc, "this-is-message")
	    //time.Sleep(time.Second)
		subsocket.SendFrame([]byte(msg), goczmq.FlagNone)
		var s []byte= bb[0:1023*1024]
		s[0] = 123
		subsocket.SendFrame(s, goczmq.FlagNone)
	}
  */
}