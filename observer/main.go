package main

import "fmt"

// Subject, the instance which publishes an event when anything happens. aka Publiser
type ISubject interface {
	subscribe(IObserver)
	unsubscribe(IObserver)
	notify(string)
}

// Concrete Subject
func newMatch(name string) Match {
	return Match{
		name:         name,
		observerList: make(map[string]IObserver),
	}
}

type Match struct {
	observerList map[string]IObserver
	name         string
}

func (m *Match) subscribe(observer IObserver) {
	m.observerList[observer.getID()] = observer
}

func (m *Match) unsubscribe(observer IObserver) {
	delete(m.observerList, observer.getID())
}

func (m Match) notify(event string) {
	for _, observer := range m.observerList {
		observer.update(m.name, event)
	}
}

// Observer. aka Subscriptor
type IObserver interface {
	getID() string
	update(string, string)
}

// Concrete Observer
type Fan struct {
	id string
}

func (f Fan) getID() string {
	return f.id
}

func (f Fan) update(eventName, event string) {
	fmt.Printf("%s - New event notification from %s - %s \n", f.id, eventName, event)
}

// client
func main() {

	match1 := newMatch("Nacional vs Milan")

	fan1 := Fan{
		id: "001",
	}
	fan2 := Fan{
		id: "002",
	}
	fan3 := Fan{
		id: "003",
	}

	match1.subscribe(fan1)
	match1.subscribe(fan2)
	match1.subscribe(fan3)
	match1.notify("0' Start the match")
	match1.notify("7' Yellow card for Milan")
	match1.unsubscribe(fan2)
	match1.notify("21' Yellow card for Nacional")
	match1.subscribe(fan2)
	match1.notify("45' + 2' first time ends")

}
