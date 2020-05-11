
/*
 go install google.golang.org/protobuf/cmd/protoc-gen-go

 To generate proto go file call:
   protoc --go_out ./ protodef/enginecmd.proto
*/

package main

import (
    /*
     #cgo LDFLAGS: libgame.so
     #include <libgame.h>
    // callback_function_type
    extern int goCallbackHandler123(int, int);
    
   static int doAdd123(int a, int b) {
         goCallbackHandler123(a, b);
    }
    static callback_function_type getCallBackPtr(){
        return &doAdd123;
    }

    static int beginGameLoop(){
        printf(" HELLO CALLBACK\n");
        StartGameLoop(doAdd123);
        return 0;
    }
   
     */
    "C"
    "fmt"
	"log"
	"gopkg.in/zeromq/goczmq.v4"
    "time"
   // "plugin"
    "os"
   // "unsafe"
   // "io/ioutil"
    //"github.com/golang/protobuf/proto"
     pb "github.com/polinomov/enserver/enbuffer/cmd"
 )

//export goCallbackHandler123
func goCallbackHandler123(a, b C.int) C.int {
    fmt.Printf("############################# 123 ######################## %d %d \n",a,b)
    return a + b
}
 

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

 /*
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
*/

func callC( a[] uint8){
   // C.test_func( (*C.uchar)(&a[0]))
}

func testMe(){
    var cc = addOneCommand(123);
    var buf = marshalCommand(cc)
    C.ProcessCmd(C.CString("oninit"), (*C.uchar)(&buf[0]), (C.int)(len(buf)) )
}

func testMe123( theCall C.callback_fcn ){
   // C.cb_wrapper(theCall)
}

func main(){
    log.Println("MAIN PUBSUB1")
    _, err := os.Stat("libgame.so")
    if err != nil {
        fmt.Printf("Can not find file %s\n", "libgame.so")
    } else {
        fmt.Printf("found %s\n", "libgame.so")
    }
   // var cb = C.getCallBackPtr()
    //C.cb(C.int(1), C.int(2))
   // C.doAdd123(C.int(1), C.int(2));
    //MyAdd(1, 2);
    //testMe123(C.callback_fcn(C.i_am_callback))
    //C.StartGameLoop(cb)
    C.beginGameLoop()
    testMe()
 
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