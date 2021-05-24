package serverutils

type ContextKey string

var (
	ContextKeySystemId   = ContextKey("systemId")
	ContextKeySystemType = ContextKey("systemType")
)
