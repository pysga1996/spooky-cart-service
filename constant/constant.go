package constant

type Status int

const UID = "UID"

const (
	STATUS_INACTIVE Status = 0
	STATUS_ACTIVE   Status = 1
	STATUS_REMOVED  Status = 2
)
