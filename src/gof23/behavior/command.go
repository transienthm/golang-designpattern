package behavior

import (
	"fmt"
)

/**
Command: 抽象命令类
ConcreteCommand: 具体命令类
Invoker: 调用者
Receiver: 接收者
 */

// receiver
type TV struct {
}

func (p TV) Open() {
	fmt.Println("play...")
}

func (p TV) Close() {
	fmt.Println("stop...")
}

//command
type Command interface {
	Press()
}

//ConcreteCommand
type OpenCommand struct {
	tv TV
}

func (p OpenCommand) Press() {
	p.tv.Open()
}

//ConcreteCommand
type CloseCommand struct {
	tv TV
}


func (p CloseCommand) Press() {
	p.tv.Close()
}

//invoker
type Invoke struct {
	cmd Command

}

func (p *Invoke) SetCommand(cmd Command) {
	p.cmd = cmd
}

func (p *Invoke) Do() {
	p.cmd.Press()
}

type OpenCloseCommand struct {
	index int
	cmds []Command
}


func NewOpenCLoseCommand() *OpenCloseCommand {
	openCLose := &OpenCloseCommand{}
	openCLose.cmds = make([]Command, 2)
	return openCLose
}

func (p *OpenCloseCommand) AddCommand(cmd Command) {
	p.cmds[p.index] = cmd
	p.index++
}

func (p *OpenCloseCommand) Press() {
	for _, item := range p.cmds {
		item.Press()
	}
}

