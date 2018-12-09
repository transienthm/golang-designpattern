package behavior

import "fmt"

/**
Mediator: 抽象中介者
ConcreteMediator: 具体中介者
Colleague: 抽象同事类
ConcreteColleague: 具体同事类
 */

//mediator 及 ConcreteMediator
type UnitedNations interface {
	ForwardMessage(message string, country Country)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Iraq
}

func (unsc UnitedNationsSecurityCouncil) ForwardMessage(message string, country Country) {
	switch country.(type) {
	case USA:
		unsc.Iraq.GetMessage(message)
	case Iraq:
		unsc.USA.GetMessage(message)
	default:
		fmt.Printf("The country is not a member of UNSC")
	}
}

type Country interface {
	SendMessage(message string)
	GetMessage(message string)
}

//Colleague以及ConcreteColleague类
type USA struct {
	UnitedNations
}

func (usa USA) SendMessage(message string) {
	usa.UnitedNations.ForwardMessage(message, usa)
}

func (usa USA) GetMessage(message string) {
	fmt.Printf("美国收到对方消息: %s\n", message)
}

type Iraq struct {
	UnitedNations
}

func (iraq Iraq) SendMessage(message string) {
	iraq.UnitedNations.ForwardMessage(message, iraq)
}

func (iraq Iraq) GetMessage(message string) {
	fmt.Printf("伊拉克收到对方消息: %s\n", message)
}
