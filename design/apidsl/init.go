package apidsl

import (
	"github.com/filewalkwithme/goa/design"
	"github.com/filewalkwithme/goa/dslengine"
)

// Setup API DSL roots.
func init() {
	design.Design = design.NewAPIDefinition()
	design.GeneratedMediaTypes = make(design.MediaTypeRoot)
	design.ProjectedMediaTypes = make(design.MediaTypeRoot)
	dslengine.Register(design.Design)
	dslengine.Register(design.GeneratedMediaTypes)
}
