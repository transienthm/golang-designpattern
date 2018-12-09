package creation

type FruitFactory interface {
	makeApple() Apple
	makeOrange() Orange
}

type Apple struct {
	desc string
}

type Orange struct {
	desc string
}

type Factory struct {
}

func (this Factory) MakeApple() Apple {
	return Apple{"我是一个苹果"}
}

func (this Factory) MakeOrange() Orange {
	return Orange{"我是一个橘子"}
}

