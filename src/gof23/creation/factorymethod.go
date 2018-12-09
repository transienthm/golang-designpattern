package creation

type FruitPicker interface {
	MakeFruit() Fruit
}
type Fruit interface {
}
type Banana struct {
	desc string
}

type BananaPicker struct {
}

func (this BananaPicker) MakeFruit(fruitName string) Fruit {
	switch fruitName {
	case "banana":
		return Banana{"我是一根香蕉"}
	default:
		return Apple{"可能我是一个苹果吧"}
	}
}
