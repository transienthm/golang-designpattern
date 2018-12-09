package structure

import "fmt"

type Noddles interface {
	Description() string
	Price() float32
}

//普通拉面
type Ramen struct {
	name string
	price float32
}

func (p Ramen) Description() string {
	return p.name
}

func (p Ramen) Price() float32 {
	return p.price
}

//想吃鸡蛋拉面怎么办？
type Egg struct {
	//多用组合，少用继承，何况golang没有继承
	noddles Noddles
	name string
	price float32
}

func (p Egg) SetNoddles(noddles Noddles)  {
	p.noddles = noddles
}

func (p Egg) Description() string {
	return p.noddles.Description() + " + " + p.name
}

func (p Egg) Price() float32 {
	return p.noddles.Price() + p.price
}

//加个香肠吧！
type Sausage struct {
	noddles Noddles
	name string
	price float32
}

func (p Sausage) SetNoddles(noddles Noddles) {
	p.noddles = noddles
}

func (p Sausage) Description() string {
	return p.noddles.Description() + " + " + p.name
}

func (p Sausage) Price() float32 {
	return p.noddles.Price() + p.price
}

func main() {
	ramen := Ramen{
		name:"ramen",
		price:8,
	}

	egg := Egg{
		noddles:ramen,
		name:"egg",
		price:2,
	}

	egg2 := Egg{
		noddles:egg,
		name:"egg",
		price:2,
	}

	sausage := Sausage{
		noddles:egg,
		name:"sausage",
		price:2,
	}

	fmt.Println("客官，您的普通拉面来了。。。")
	fmt.Println(ramen.Description())
	fmt.Println(ramen.Price())

	fmt.Println("客官，您的鸡蛋拉面来了。。。")
	fmt.Println(egg.Description())
	fmt.Println(egg.Price())

	fmt.Println("客官，您的双蛋拉面来了。。。")
	fmt.Println(egg2.Description())
	fmt.Println(egg2.Price())

	fmt.Println(sausage.Description())
	fmt.Println(sausage.Price())
	fmt.Println("客官，您的香肠拉面来了。。。")

}
