package atoms

// Context that is used for evaluation of the request at various stages.
type Context struct {
	Variables map[string]interface{}
}
