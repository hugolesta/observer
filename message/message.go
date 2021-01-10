package message

type Message struct {
	message string
}

func (m *Message) AddObserver(o Observer) {

}

func (m *Message) RemoveObserver(o Observer) {}

func (m *Message) NotifyObservers() {}
