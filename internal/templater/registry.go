package templater

type Registry struct{ m map[Kind]Templater }

func NewRegistry() *Registry                     { return &Registry{m: map[Kind]Templater{}} }
func (r *Registry) Register(t Templater)         { r.m[Kind(t.Kind())] = t }
func (r *Registry) Get(k Kind) (Templater, bool) { t, ok := r.m[k]; return t, ok }
func (r *Registry) List() []Kind {
	out := make([]Kind, 0, len(r.m))
	for k := range r.m {
		out = append(out, k)
	}
	return out
}

// Global is an optional global registry to support side-effect registration.
var Global = NewRegistry()

// Register adds a templater to the global registry (for use in init()).
func Register(t Templater) {
	if t == nil {
		panic("nil templater")
	}
	Global.Register(t)
}
