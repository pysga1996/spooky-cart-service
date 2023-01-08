package constant

type Status int

const UID = "UID"

const (
	StatusInactive Status = 0
	StatusActive   Status = 1
	StatusRemoved  Status = 2
)
