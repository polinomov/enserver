package main

import (
    "fmt"
	"log"
    //"math/rand"
	"gopkg.in/zeromq/goczmq.v4"
    "time"
 )


 type CmdStruct struct {
	n int
	data string
}

func fromClient(cmdBuff chan CmdStruct)  {
    fmt.Printf("Start client channel type : %T \n",cmdBuff) 
    pullsocket,err := goczmq.NewPull("tcp://0.0.0.0:1234") 
    if err != nil {
        log.Fatal(err)
        return;
    }
   // defer pullsocket.Destroy()
   // pullsocket.Bind("tcp://127.0.0.1:1234")
    var i = 0
    //var t_now = time.Now().UnixNano()
    var t_old = time.Now().UnixNano()
    for {
        //log.Printf("About to read")
        recdata, rerr :=pullsocket.RecvMessage()
        if rerr != nil {
           log.Fatal(rerr)
        }
        if (i%1000)== 0 {
            var t_now = time.Now().UnixNano()
            if  i<0 {
                log.Printf("received '%s' ---", string(recdata[0]))
            }
            fmt.Printf("%d cap=%d\n",(t_now-t_old)/1e6, cap(recdata[0]))
            t_old = t_now
        }
        //fmt.Printf(" ---- len = %d  cap = %d\n", len(cmdBuff), cap(cmdBuff))
        cmdBuff <- CmdStruct{i, "blah"}
       // time.Sleep(time.Millisecond * 100)
        i = i + 1
    }
 }

 func toClient(cmdBuff chan CmdStruct){
    for {
        rec := <- cmdBuff
        rec.n++
        //fmt.Printf("%d %s\n",rec.n,rec.data)
       // time.Sleep(time.Millisecond*1000)
    }
 }

func main(){
  log.Println("MAIN PUBSUB1")
  cmdBuff := make(chan CmdStruct, 32)
  go fromClient(cmdBuff)
  go toClient(cmdBuff)
  for{
    time.Sleep(time.Millisecond*1000) 
  }
  log.Println("MAIN DONE")
  
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