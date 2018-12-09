package observer

import "fmt"

 //抽象观察者
type IObserver interface {
	Notify() //当被观察对象有理性的时候，触发观察者的Notify()方法
}

//具体观察者
type Observer struct {
}

func (o *Observer) Notify() {
	fmt.Println("已经触发了观察者")
}
