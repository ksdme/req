package atoms

// Collection represents groups of request configurations which are at the
// core of this utility.
type Collection struct {
	// The version of the current collection.
	Version uint `yaml:"version"`

	// The request groups under the collection.
	Groups map[string]Group `yaml:"groups"`
}
