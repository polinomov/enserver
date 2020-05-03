//
//  go build  -buildmode=plugin -o /$GOPATH/bin/libgame.so pingpong.go
//  go build  -buildmode=plugin -o ../libgame.so pingpong.go
//

package main

import "fmt"

var V int
var  Msg string
var TheCallBack func( cmd string)

func F1() { fmt.Printf("Hello, number %d\n", V) }
func F2( ) { 
	fmt.Printf("Hello  %s\n", Msg) 
	TheCallBack("----so callback----")
}

func main(){
}