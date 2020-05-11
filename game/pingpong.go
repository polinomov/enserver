//
//  go build -o libgame.so -buildmode=c-shared pingpong.go
//

package main

import( 
/*	
 #include <stdio.h>
 typedef int (*callback_fcn)(int);
 typedef int (*callback_function_type)(int,int);
 static void wrap_xa(callback_function_type cbb, int a, int b)
 {
	 cbb( a,b);
 }
 */
	"C"
	"fmt"
	"log"
	"time"
	"unsafe"
	"math/rand"
	"github.com/golang/protobuf/proto"
	 pb "github.com/polinomov/enserver/enbuffer/cmd"
)

//export goCallbackHandler
func goCallbackHandler(a C.int) C.int {
    return a 
}
func MyCallback(x C.int) {
    fmt.Println("callback with", x)
}

//export ProcessCmd
func ProcessCmd(info *C.char, protodata *C.uchar, datlen C.int) *C.char {
	var cmd = C.GoString(info)
	fmt.Printf("ProcessCmd %s len=%d\n",cmd,int(datlen))
	var pr =C.GoBytes(unsafe.Pointer(protodata),datlen)
	unmarshalCommand (pr)
	return nil
}



// TYPE IS *cmd.CommandList
func  unmarshalCommand (dat []uint8) (  *pb.CommandList ) {
	cmdList := &pb.CommandList{}
	if err := proto.Unmarshal(dat, cmdList); err != nil {
		log.Fatalln("Failed to unmarshall", err)
	}
	fmt.Printf( "len=%d cap=%d Opa=%d \n",len(cmdList.Cmd),cap(cmdList.Cmd),cmdList.Cmd[0].Opa)
	for i := range cmdList.Cmd {
		fmt.Printf("name= %s Opa= %d\n",cmdList.Cmd[i].Name, cmdList.Cmd[i].Opa)
	}
	fmt.Printf("TYPE IS %T\n", cmdList)
	return cmdList
 }

 /*------------------------------------------------------------------------------*/
 //export SetUpdateCallBack
func SetUpdateCallBack( theCall C.callback_fcn ){

}

//export StartGameLoop
func StartGameLoop( cbfunc C.callback_function_type){
	fmt.Printf("START GAME LOOP\n");
    C.wrap_xa(cbfunc,C.int(555),C.int(666))
	go gameLoop(cbfunc)
}

type GObject struct{
	 id uint32
	 rad float32
     px,py,pz float32
	 vx,vy,vz float32
}


func gameLoop(cbfunc C.callback_function_type){
	var gObjMap = map[uint32]*GObject{}
	var rf = float32(0.01)
	for i := 0; i < 1; i++ {
		var xf = rand.Float32() * ( 1.0 - 2.0 * rf) + rf
		var yf = rand.Float32() * ( 1.0 - 2.0 * rf) + rf
		var velx =  float32(0.02)
		var vely =  float32(0.01)
		gObjMap[uint32(i)] = &GObject{id: 0, rad:rf, px:xf , py:yf , pz:0, vx:velx, vy:vely, vz:0}
	}

	for{
		//fmt.Printf("loop\n")
		time.Sleep(time.Millisecond*1000)
		for k := range gObjMap {
			var isInside = true;
			var pxn = gObjMap[k].px + gObjMap[k].vx;
			var pyn = gObjMap[k].py + gObjMap[k].vy;
			if(( pxn < rf) || (pxn > 1.0- rf))  {
				gObjMap[k].vx  = -gObjMap[k].vx 
				isInside = false;
			}
			if(( pyn < rf) || (pyn > 1.0- rf))  {
				gObjMap[k].vy  = -gObjMap[k].vy
				isInside = false;
			}
			if ( isInside==true){
				gObjMap[k].px = pxn;
				gObjMap[k].py = pyn;
			}
			fmt.Printf( "id=%d ( %f %f)\n",k, gObjMap[k].px,gObjMap[k].py)
		} 
		fmt.Printf("---------------\n")
		C.wrap_xa(cbfunc,C.int(555),C.int(666))
		//AuxFunc(theCall)
		//C.theCall(555)
   }
}

func main(){
}