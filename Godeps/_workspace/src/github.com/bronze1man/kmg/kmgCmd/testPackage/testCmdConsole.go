package testPackage

import (
	"fmt"
	"github.com/bronze1man/kmg/kmgConsole"
	"time"
)

func TestCmdConsole() {
	fmt.Println("children init finish")
	kmgConsole.WaitForExit()

	fmt.Println("children after WaitForExit")
	time.Sleep(time.Second)
	fmt.Println("children after sleep")
}
