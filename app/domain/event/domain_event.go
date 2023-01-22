package event

type DomainEvent interface {
	EventType() string
	Serialize() (string, error)
	Deserialize(data string) error
}
