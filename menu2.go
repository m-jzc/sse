package main

import (
	"fmt"

	"github.com/m-jzc/sse.git/linklist"
	//"linklist"
)

const (
	CMD_MAX_LEN = 128
	DESC_LEN    = 1024
	CMD_NUM     = 10
)

var head *linklist.DataNode = &linklist.DataNode{
	Cmd:     "help",
	Desc:    "this is help cmd!",
	Handler: nil,
	Next: &linklist.DataNode{
		Cmd:     "version",
		Desc:    "menu program v1.0",
		Handler: nil,
		Next:    nil,
	},
}
var Help = func() int {
	linklist.ShowAllCmd(head)
	return 0
}

func main() {
	for {
		cmd := make([]byte, CMD_MAX_LEN)
		fmt.Println("Input a cmd number: ")
		fmt.Scanln(&cmd)
		p := linklist.FindCmd(head, string(cmd))
		if p == nil {
			fmt.Println("This is a wrong cmd!")
			continue
		}
		fmt.Printf("%s - %s\n", p.Cmd, p.Desc)
		if p.Cmd == "help" {
			Help()
		}

		if p.Handler != nil {
			p.Handler()
		}
	}
}

