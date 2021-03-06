package main

import (
	"fmt"
	"log"
	"github.com/golang/protobuf/proto"
	 pb "github.com/polinomov/enserver/enbuffer/cmd"
 )
// extern int goCallbackHandler456(int, int);
// typedef int (*function_type)(int,int);
// static int doAdd(int a, int b) {
//     return goCallbackHandler456(a, b);
// }
// static function_type getMyFunc(){
//		return &doAdd;
// }
import "C"

//export goCallbackHandler456
func goCallbackHandler456(a, b C.int) C.int {
	fmt.Printf("------------- goCallbackHandler----------\n")
    return a + b
}

func MyAdd(a, b int) int {
   return int( C.doAdd( C.int(a), C.int(b)) )
}


// []uint8
 func addOneCommand( v1 int32) ( *pb.Command) {
    c := pb.Command {
		Name : "command-name",
        Opa : v1,
        Opb : 2,
        Opc : 3,
	}
    return &c; 
 }

 func  marshalCommand( c *pb.Command) ([]uint8) {
	fmt.Printf("Do marshal %d \n",c.Opa)
	cmdList := &pb.CommandList{}
    cmdList.Cmd = append(cmdList.Cmd, c)
	out, err := proto.Marshal(cmdList)
   	if err != nil {
		log.Fatalln("Failed to Marshall", err)
	}
	return out
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

