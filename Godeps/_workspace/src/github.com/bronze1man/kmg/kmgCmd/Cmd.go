package kmgCmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	//"os/signal"
	//"syscall"
)

//please use Cmd* function to new a Cmd,do not create one yourself.
type Cmd struct {
	cmd *exec.Cmd
}

//you need at least one args: the path of the command, or it will panic
func CmdSlice(args []string) *Cmd {
	if len(args) == 0 {
		panic("[CmdSlice] need the path of the command")
	}
	return &Cmd{
		cmd: exec.Command(args[0], args[1:]...),
	}
}

func CmdString(cmd string) *Cmd {
	if cmd == "" {
		panic("[CmdString] need the path of the command")
	}
	args := strings.Split(cmd, " ")
	return &Cmd{
		cmd: exec.Command(args[0], args[1:]...),
	}
}

func CmdBash(cmd string) *Cmd {
	return CmdSlice([]string{"bash", "-c", cmd})
}

func (c *Cmd) MustSetEnv(key string, value string) *Cmd {
	err := SetCmdEnv(c.cmd, key, value)
	if err != nil {
		panic(err)
	}
	return c
}

func (c *Cmd) SetDir(path string) *Cmd {
	c.cmd.Dir = path
	return c
}

func (c *Cmd) PrintCmdLine() {
	c.FprintCmdLine(os.Stdout)
}

func (c *Cmd) FprintCmdLine(w io.Writer) {
	fmt.Fprintln(w, ">", strings.Join(c.cmd.Args, " "))
}

//回显命令,并且运行,并且和标准输入输出接起来
func (c *Cmd) Run() error {
	c.PrintCmdLine()
	c.cmd.Stdin = os.Stdin
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	return c.cmd.Run()
}

//如果代理运行失败,当前进程会退出
func (c *Cmd) ProxyRun() {
	err := c.Run()
	if err != nil {
		//不使用 kmgConsole.ExitOnErr ,以避免依赖循环
		fmt.Println(err)
		os.Exit(2)
		return
	}
}

//get the os/exec.Cmd
func (c *Cmd) GetExecCmd() *exec.Cmd {
	return c.cmd
}

//回显命令,并且运行,返回运行的输出结果.并且把输出结果放在stdout中
func (c *Cmd) RunAndReturnOutput() (b []byte, err error) {
	c.PrintCmdLine()
	buf := &bytes.Buffer{}
	w := io.MultiWriter(buf, os.Stdout)
	c.cmd.Stdout = w
	c.cmd.Stderr = w
	err = c.cmd.Run()
	return buf.Bytes(), err
}

func (c *Cmd) CombinedOutput() (b []byte, err error) {
	return c.cmd.CombinedOutput()
}

// 不能传递signel
func (c *Cmd) RunAndTeeOutputToFile(path string) (err error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0777))
	if err != nil {
		return err
	}
	w := io.MultiWriter(f, os.Stdout)
	c.FprintCmdLine(w)
	c.cmd.Stdout = w
	c.cmd.Stderr = w
	c.cmd.Stdin = os.Stdin
	return c.cmd.Run()
}

//不回显命令,运行,并且返回运行的输出结果
func (c *Cmd) StdioRun() error {
	c.cmd.Stdin = os.Stdin
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	return c.cmd.Run()
}

//回显命令,并且运行,并且忽略输出结果
func (c *Cmd) RunAndNotExitStatusCheck() error {
	err := c.Run()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return nil
		}
		return err
	}
	return nil
}

func (c *Cmd) MustStdioRun() {
	err := c.StdioRun()
	if err != nil {
		panic(err)
	}
}

func (c *Cmd) MustRunAndReturnOutput() (b []byte) {
	b, err := c.RunAndReturnOutput()
	if err != nil {
		panic(err)
	}
	return b
}

func (c *Cmd) MustRunAndReturnOutputAndNotExitStatusCheck() (b []byte) {
	b, err := c.RunAndReturnOutput()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return b
		}
		panic(err)
	}
	return b
}

//允许命令,返回命令的内容,不回显任何东西
func (c *Cmd) MustCombinedOutput() (b []byte) {
	b, err := c.cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return b
}

//允许命令,返回命令的内容,不回显任何东西,
// 出现错误就回显输出结果,没有错误,什么也不显示.
func (c *Cmd) MustCombinedOutputWithErrorPrintln() (b []byte) {
	b, err := c.cmd.CombinedOutput()
	if err != nil {
		fmt.Println(">", strings.Join(c.cmd.Args, " "))
		os.Stdout.Write(b)
		panic(err)
	}
	return b
}

func (c *Cmd) MustCombinedOutputAndNotExitStatusCheck() (b []byte) {
	b, err := c.cmd.CombinedOutput()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return b
		}
		panic(err)
	}
	return b
}

func (c *Cmd) MustRunAndNotExitStatusCheck() {
	err := c.RunAndNotExitStatusCheck()
	if err != nil {
		panic(err)
	}
}

func (c *Cmd) MustRun() {
	err := c.Run()
	if err != nil {
		panic(err)
	}
}

func (c *Cmd) MustHiddenRunAndGetExitStatus() int {
	err := c.cmd.Run()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if !ok {
			panic(err)
		}
	}
	return GetExecCmdExitStatus(c.cmd)
}

func (c *Cmd) MustHiddenRunAndIsSuccess() bool {
	err := c.cmd.Run()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if !ok {
			panic(err)
		}
	}
	return c.cmd.ProcessState.Success()
}

func (c *Cmd) MustRunWithStdin(stdin []byte) {
	c.PrintCmdLine()
	c.cmd.Stdin = bytes.NewBuffer(stdin)
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	err := c.cmd.Run()
	if err != nil {
		panic(err)
	}
}

// TODO 实现和编程复杂度高,而且接口不能很容易的变成单独函数接口.
//func (c *Cmd) MustRunWithSingle(){
//	c.MustRun()
//}

type exitStatuser interface {
	ExitStatus() int
}

func GetExecCmdExitStatus(cmd *exec.Cmd) int {
	return cmd.ProcessState.Sys().(exitStatuser).ExitStatus()
}

func Exist(cmd string) bool {
	_, err := CmdBash("which " + cmd).cmd.CombinedOutput()
	if err == nil {
		return true
	} else {
		return false
	}
}
