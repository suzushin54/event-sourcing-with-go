package event

type DomainEvent interface {
	EventType() string
	Serialize() (string, error)
}
