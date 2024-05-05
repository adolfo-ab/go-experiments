package main

func main() {
	listener1 := DataListener{
		Name: "L1",
	}
	listener2 := DataListener{
		Name: "L2",
	}

	subj := &DataSubject{}
	subj.registerObserver(listener1)
	subj.registerObserver(listener2)
	subj.ChangeItem("hi mum!")
}
