package tilt

type Compose struct {
	File      string
	Resources []DCResource
}
type DCResource struct {
	Name         string
	ResourceDeps []string
	Labels       []string
	TriggerMode  TriggerMode
	Links        []string
}

type Local struct {
	Name   string
	Cmd    string
	Deps   []string
	Labels []string
}

type Data struct {
	Compose *Compose // nil if not using compose
	Locals  []Local
}
