package main

type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener
	field     string
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataSubject) unregisterObserver(o DataListener) {
	var newObservers []DataListener
	for _, obs := range ds.observers {
		if o.Name != obs.Name {
			newObservers = append(newObservers, obs)
		}
	}
	ds.observers = newObservers
}

func (ds *DataSubject) notifyAll() {
	for _, obs := range ds.observers {
		obs.onUpdate(ds.field)
	}
}
