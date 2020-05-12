
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
    extern int goGameCallBack(int, char*, float);
    
    static int call_proxy(int objId, char *pAttr, float val) {
        goGameCallBack(objId, pAttr,val);
    }
  
    static int beginGameLoop(){
        printf(" HELLO CALLBACK\n");
        StartGameLoop(call_proxy);
        return 0;
    }
     */
    "C"
    "fmt"
	"log"
	"gopkg.in/zeromq/goczmq.v4"
    "time"
    "os"
    //"runtime/debug"
   // "unsafe"
   // "io/ioutil"
    "github.com/golang/protobuf/proto"
     pb "github.com/polinomov/enserver/enbuffer/cmd"
 )

 type GRecord struct{
    id int32
    attr string 
    val float32
}

type Context struct{
     //frame []GRecord
     pubsocket* goczmq.Sock
     protodata* pb.CommandList
}

var TheContext = &Context{}

func (c *Context) initSocket()  {
    fmt.Printf("Context init socket\n")
    //var opt = goczmq.SockSetConflate(1)
   // pubsocket, err := goczmq.NewPub("tcp://*:5555",opt)
    pubsocket, err := goczmq.NewPub("tcp://*:5555")
    fmt.Printf(" pubsocket TYPE IS %T\n", pubsocket)
    if err != nil {
        log.Fatal(err)
    }
    pubsocket.Bind("tcp://*:5555")
    c.pubsocket = pubsocket
}

func ( c *Context) sendData(){
    /*
    cmdList := &pb.CommandList{}
    cmd := &pb.Command { Name : "command-name", Opa : 1,  Opb : 2,Opc : 3,}
    cmdList.Cmd = append(cmdList.Cmd, cmd)
    out, err := proto.Marshal(cmdList)
    */
   // fmt.Printf("sendData\n");
    out, err := proto.Marshal(c.protodata)
   	if err != nil {
		log.Fatalln("Failed to Marshall", err)
	}
    c.pubsocket.SendFrame(out, goczmq.FlagNone)
}

func ( c *Context) frameBegin(){
    //fmt.Printf("frameBegin\n");
    c.protodata  = &pb.CommandList{}
}

func ( c *Context) saveRecord( objId int32, attrName string, val float32){
    //cmdList := &pb.CommandList{}
    //fmt.Printf(" type : %T \n",cmdList) 
    var ival = (int32)(val*10000.0);
    cmd := &pb.Command { Name : attrName, Opa : objId,  Opb : ival,Opc : 3,}
    c.protodata.Cmd = append(c.protodata.Cmd, cmd)
   // fmt.Printf("#saveRecord %d %s %f \n", objId, attrName, val)
}

func ( c *Context) frameEnd(){
    //fmt.Printf("frameEnd\n");
    c.sendData()
    c.protodata  = nil;
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

 
//export goGameCallBack
func goGameCallBack( objId C.int, attrName *C.char, attrValue float32) C.int {
    var idd = int32(objId)
    if( idd == -1){
        TheContext.frameBegin();
        return 0;
    }
    if( idd == -2){
        TheContext.frameEnd();
        return 0;
    }

    TheContext.saveRecord(int32(objId), C.GoString(attrName),  attrValue)
    return 0
}

 func toClient(cmdBuff chan pb.Command){
   // var opt = goczmq.SockSetSndbuf(1)
    pubsocket, err := goczmq.NewPub("tcp://*:5555")
    fmt.Printf(" pubsocket TYPE IS %T\n", pubsocket)
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
   // debug.SetGCPercent(-1)
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
    TheContext.initSocket()
    C.beginGameLoop()
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