package main

import (
	//"os"
	//"os/signal"
	//"syscall"
	"fmt"
	"github.com/bronze1man/kmg/kmgConsole"
	"time"
	//"net/http"
	//"github.com/bronze1man/kmg/kmgSys"
	//"github.com/bronze1man/kmg/kmgCmd"
	"github.com/bronze1man/kmg/kmgFile"
)

func main() {
	//go func(){
	//	for {
	//		time.Sleep(time.Second)
	//		fmt.Println("sleep")
	//	}
	//}()
	//go func(){
	//	err := http.ListenAndServe(":23456",nil)
	//	if err!=nil{
	//		panic(err)
	//	}
	//}()
	//go func(){
	//	tun,err:=kmgSys.NewTunNoName()
	//	if err!=nil{
	//		panic(err)
	//	}
	//	buf:=make([]byte,4096)
	//	_,err=tun.Read(buf)
	//	if err!=nil{
	//		panic(err)
	//	}
	//}()
	//kmgCmd.MustRun("ls")
	kmgConsole.AddCommandWithName("cmd", cmd)
	kmgConsole.Main()
}

func cmd() {
	fmt.Println("children init finish")
	kmgConsole.WaitForExit()

	kmgFile.MustWriteFile("/tmp/2.log", []byte(time.Now().String()))
	fmt.Println("children after WaitForExit")
	time.Sleep(time.Second)
	fmt.Println("children after sleep")
}
