package main

import "fmt"

/*
中介者是一种行为设计模式， 让程序组件通过特殊的中介者对象进行间接沟通， 达到减少组件之间依赖关系的目的。
中介者能使得程序更易于修改和扩展， 而且能更方便地对独立的组件进行复用， 因为它们不再依赖于很多其他的类。
*/
func main() {
	stationManager := newStationManger()

	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	freightTrain := &freightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}

// 组件
type train interface {
	arrive()
	depart()
	permitArrival()
}

// 具体组件
type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *passengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.arrive()
}

// 具体组件
type freightTrain struct {
	mediator mediator
}

func (g *freightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *freightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.arrive()
}

// 中介者接口
type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

// 具体中介者
type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func newStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}

func (s *stationManager) canArrive(t train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
