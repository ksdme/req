package atoms

// Request is the underlying request configuration. It might not contain
// all the necessary information related to the request, you can always
// propagate upwards to request more data.
type Request struct {
	// The method of the request.
	// TODO: This needs to have a default value.
	Method Evalable

	// The host with the protocol of the request.
	// TODO: This should be called prefix?
	Host Evalable

	// The actual endpoint of the request.
	Endpoint Evalable

	// Search query parameters for the request.
	Query map[string]Evalable

	// Headers for the request.
	Headers map[string]Evalable
}
