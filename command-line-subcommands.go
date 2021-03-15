package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 声明一个子命令，定义一个专用的 flag
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// 子命令作为程序的 第一个参数传入
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// 检查哪一个子命令被调用
	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:]) // 解析
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'boo' or 'bar' subcommands")
		os.Exit(1)

	}
}

// 调用 foo 字命令
// command-line-subcommands foo -enable -name=joe a1 a2

// 调用 bar 子命令
// command-line-subcommands bar -level 8 a1

// bar 不接受 foo 的 flag
// command-line-subcommands bar -enable a1
