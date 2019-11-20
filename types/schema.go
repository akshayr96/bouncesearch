package types

type Schema map[string]AttributeMeta

type AttributeMeta struct {
	DefaultValue string
	Weight       float32
	Optional     bool
}

type Record map[string]string
