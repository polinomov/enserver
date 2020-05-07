//
//  go build -o libgame.so -buildmode=c-shared pingpong.go
//

package main

import( 
// extern int goCallbackHandler(int, int);
// static int aa;
// static int doAdd(int a, int b) {
//     return goCallbackHandler(a, b);
// }
	"C"
	"fmt"
	"log"
	"unsafe"
	"github.com/golang/protobuf/proto"
	 pb "github.com/polinomov/enserver/enbuffer/cmd"
)

//export goCallbackHandler
func goCallbackHandler(a, b C.int) C.int {
    return a + b
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

func main(){
}