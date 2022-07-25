package callback

type MessageCallback struct {
	mc func([]byte)
}

func (m *MessageCallback) SetCallback(f func([]byte)) {
	m.mc = f
}

func (m *MessageCallback) Callback(data []byte) {
	if m.mc != nil {
		m.mc(data)
	}
}

type MCallback interface {
	Callback([]byte)
}
