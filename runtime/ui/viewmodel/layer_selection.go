package viewmodel

import (
	"github.com/khulnasoft/dockvisor/dockvisor/image"
)

type LayerSelection struct {
	Layer                                                      *image.Layer
	BottomTreeStart, BottomTreeStop, TopTreeStart, TopTreeStop int
}
