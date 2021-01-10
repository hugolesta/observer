package observer

type Observable interface {
	AddObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}
