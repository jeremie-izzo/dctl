package registry

import (
	"github.com/jeremie-izzo/dctl/pkg/contracts"
	"github.com/jeremie-izzo/dctl/pkg/presets"
	"github.com/jeremie-izzo/dctl/pkg/tilt"
)

type coreRegistry struct {
	engines map[string]contracts.TemplateEngine
	//providers   map[string]contracts.ArtifactProvider
	presets     map[string]presets.ServicePresets
	tiltPlugins map[string]tilt.Plugin
}

func New() *coreRegistry {
	return &coreRegistry{
		engines: make(map[string]contracts.TemplateEngine),
		//providers:   make(map[string]contracts.ArtifactProvider),
		tiltPlugins: make(map[string]tilt.Plugin),
		presets:     make(map[string]presets.ServicePresets),
	}
}

func (r *coreRegistry) RegisterEngine(e contracts.TemplateEngine) {
	r.engines[e.EngineName()] = e
}

func (r *coreRegistry) TemplateEngine(name string) (contracts.TemplateEngine, bool) {
	e, ok := r.engines[name]
	return e, ok
}

//
//func (r *coreRegistry) RegisterProvider(p contracts.ArtifactProvider) {
//	r.providers[p.Name()] = p
//
//}
//func (r *coreRegistry) ArtifactProvider(name string) (contracts.ArtifactProvider, bool) {
//	p, ok := r.providers[name]
//	return p, ok
//}

func (r *coreRegistry) RegisterTiltPlugin(p tilt.Plugin) {
	r.tiltPlugins[p.Name()] = p
}

func (r *coreRegistry) TiltPlugin(name string) (tilt.Plugin, bool) {
	p, ok := r.tiltPlugins[name]
	return p, ok
}

func (r *coreRegistry) RegisterPreset(p presets.ServicePresets) {
	r.presets[p.Name()] = p
}

func (r *coreRegistry) Preset(name string) (presets.ServicePresets, bool) {
	p, ok := r.presets[name]
	return p, ok
}

var Global = New()
