package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// 这里的命令与 Windows 的不符
// Windows 用 Git Bash 可以正常
func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	// grep 命令用于查找文件里符合条件的字符串
	// grep 目标字符 查找范围
	// 打印出含目标字符串的行

	// 明确获取输入/输出管道
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep\nhello again")) // 写入数据
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 在生成命令时，我们需要提供一个明确描述命令和参数的数组，而不能只传递 一个 命令行字符串。 如果你想使用 "一个字符串生" 成一个完整的命令，那么你可以使用 bash 命令的 -c 选项：
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
