package strategy

//抽象策略角色
type cashStrategy interface {
	AcceptCash(float64) float64
}

//具体策略角色
type cashNormal struct {
}

func (normal *cashNormal) AcceptCash(money float64) float64 {
	return money
}

type cashRebate struct {
	moneyRebate float64
}

func (rebate *cashRebate) AcceptCash(money float64) float64 {
	return money * rebate.moneyRebate
}

type cashReturn struct {
	moneyCondition float64
	moneyReturn float64
}

func (returned *cashReturn) AcceptCash(money float64) float64 {
	if money >= returned.moneyCondition {
		return money - float64(int(money / returned.moneyCondition)) * returned.moneyReturn
	} else {
		return money
	}
}

// context角色
type CashContext struct {
	Stratege cashStrategy
}

func NewCashContext(cashType string) *CashContext {
	c := new(CashContext)

	switch cashType {
	case "打八折":
		c.Stratege = &cashRebate{0.8}
	case "满300返100":
		c.Stratege = &cashReturn{300, 100}
	case "正常收费":
		c.Stratege = &cashNormal{}
	default:
		c.Stratege = &cashNormal{}
	}
	return c
}