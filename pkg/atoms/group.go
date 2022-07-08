package atoms

// Group is a collection of requests, it helps with grouping and sharing
// resources between a group of requests.
type Group struct {
	// A group contains a number of keyed requests. The key is the name
	// of the request.
	Requests map[string]Request `yaml:"requests"`
}
