package constant

type Status int

const UID = "UID"

const (
	StatusInactive Status = 0
	StatusActive   Status = 1
	StatusRemoved  Status = 2
)

const BaseCacheKey = "REVICE_COMMERCE"

const CacheSeparator = "::"

const JwksCacheKey = BaseCacheKey + CacheSeparator + "JWKS"
