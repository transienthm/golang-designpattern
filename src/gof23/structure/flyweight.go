package structure

import (
	"fmt"
	"reflect"
)

type Coordinate struct {
	x, y int
}

type ChessFlyWeight interface {
	getColor() string
	display(c Coordinate)
}

type ConcreteChess struct {
	Color string
}

func (chess ConcreteChess) display(c Coordinate) {
	fmt.Printf("棋子颜色:%s\n", chess.Color)
	fmt.Printf("棋子位置:%d----%d\n", c.x, c.y)
}

func (chess ConcreteChess) getColor() string {
	return chess.Color
}

type ChessFlyWeightFactory struct {
	pool map[string]ChessFlyWeight
}

func (factory ChessFlyWeightFactory) getChess(color string) ChessFlyWeight {
	chess := factory.pool[color];

	if chess == nil {
		chess = ConcreteChess{color}
		factory.pool[color] = chess
	}

	return chess
}

func main() {
	pool  := make(map[string]ChessFlyWeight)

	factory := ChessFlyWeightFactory{pool	}

	chess1 := factory.getChess("black")
	chess2 := factory.getChess("black")

	fmt.Println(reflect.DeepEqual(chess1, chess2))
	c1 := Coordinate{1,2}
	c2 := Coordinate{2, 4}
	chess1.display(c1)
	chess1.display(c2)
}

