package main

import (
	"fmt"
	"github.com/bronze1man/kmg/kmgCmd"
	"os"
	"os/signal"
	"syscall"
	//"time"
)

func main() {
	//var exitSignalProcessor func(signal os.Signal) // 只杀掉一层子进程不管用,bash -c 不会传递信号.
	go func() {
		kmgCmd.MustRun("kmg go install github.com/bronze1man/kmg/kmgCmd/testPackage/testCmdChildren")
		cmd := kmgCmd.CmdBash("./bin/testCmdChildren cmd | tee -i /tmp/1.log")
		//exitSignalProcessor = func(signal os.Signal){
		//	err := cmd.GetExecCmd().Process.Signal(signal)
		//	if err!=nil{
		//		panic(err)
		//	}
		//}
		cmd.MustRun()
		fmt.Println("parent Must Run return")
	}()
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	thisSignal := <-ch
	fmt.Println("parent", thisSignal)
	//if exitSignalProcessor!=nil {
	//	exitSignalProcessor(thisSignal)
	//}
}
