package observer

//抽象被观察者
type ISubject interface {
	AddObservers(observers ...IObserver) //添加观察者
	NotifyObservers() //通知观察者
}

//具体被观察者
type Subject struct {
	observers []IObserver
}

func (s *Subject) AddObservers(observer ...IObserver) {
	s.observers = append(s.observers, observer...)
}

func (s *Subject) NotifyObservers() {

	for k := range s.observers {
		s.observers[k].Notify() //触发观察者
	}
}