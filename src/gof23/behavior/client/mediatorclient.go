package main

import "gof23/behavior"

func main() {
	tMediator := behavior.UnitedNationsSecurityCouncil{}
	usa := behavior.USA{tMediator}

	iraq := behavior.Iraq{tMediator}

	tMediator.USA = usa
	tMediator.Iraq = iraq

	usa.SendMessage("停止大规模杀伤性武器的研发，否则发动战争")
	iraq.SendMessage("我们没有研发大规模杀伤性武器，也不怕战争")
}
