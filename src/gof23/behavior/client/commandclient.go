package main

import (
	"fmt"
	"gof23/behavior"
)

func main() {
	//单一命令
	tv := behavior.TV{}
	openCommand := behavior.OpenCommand{tv}
	invoker := behavior.Invoke{openCommand}
	invoker.Do()

	closeCommand := behavior.CloseCommand{tv}
	invoker.SetCommand(closeCommand)
	invoker.Do()

	//复合命令
	fmt.Println("############复合命令###############")
	openClose := behavior.NewOpenCLoseCommand()
	openClose.AddCommand(openCommand)
	openClose.AddCommand(closeCommand)

	invoker.SetCommand(openClose)
	invoker.Do()
}
