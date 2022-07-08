package atoms

// Request is the core underlying request configuration.
type Request struct {
	// The method of the request.
	// TODO: This needs to have a default value.
	Method string `yaml:"method"`
}
