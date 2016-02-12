package kmgCmd

import "os/exec"

/*
运行命令,并且把命令输入输出和当前的输入输出接起来
会回显输入的命令
会显示输出的结果
应该可以从命令行输入输入,但是没有使用过.

传入一个字符串是命令, 不允许参数中包含空格,如果命令有复杂参数请使用 RunOsStdioCmd
不支持bash语法, 不能在里面使用bash的各种连接符之类的.
只能写一条命令

不支持下列用法:
	ls 'a b' //不允许参数中包含空格.
	echo 1 > /dev/null //不支持bash语法,会输出 1 > /dev/null
*/
func Run(cmd string) (err error) {
	return CmdString(cmd).Run()
}

//相比Run 不回显命令 并且使用slice作为输入方式
func StdioSliceRun(args []string) (err error) {
	return CmdSlice(args).StdioRun()
}

func MustRun(cmd string) {
	CmdString(cmd).MustRun()
}

//代理运行命令,同run,出现错误时,会退出进程
// 请仅在命令行程序中使用
func ProxyRun(cmd string) {
	CmdString(cmd).ProxyRun()
}

//相比MustRun 如果进程返回值不是0,不报错
func MustRunNotExistStatusCheck(cmd string) {
	CmdString(cmd).MustRunAndNotExitStatusCheck()
}

//相比MustRun 也返回输出结果
func MustRunAndReturnOutput(cmd string) []byte {
	return CmdString(cmd).MustRunAndReturnOutput()
}

//相比MustRun 返回输出结果，但是不会输出到 stdout
func MustCombinedOutput(cmd string) []byte {
	return CmdString(cmd).MustCombinedOutput()
}

func MustCombinedOutputWithErrorPrintln(cmd string) []byte {
	return CmdString(cmd).MustCombinedOutputWithErrorPrintln()
}

func MustCombinedOutputAndNotExitStatusCheck(cmd string) []byte {
	return CmdString(cmd).MustCombinedOutputAndNotExitStatusCheck()
}

//相比MustRun 输入的命令会被放到bash中执行,cmd的语法和bash一致.
func MustRunInBash(cmd string) {
	CmdBash(cmd).MustRun()
}

func IsErrorExitStatus(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*exec.ExitError)
	return ok
}
