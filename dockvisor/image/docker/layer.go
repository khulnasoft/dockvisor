package docker

import (
	"strings"

	"github.com/khulnasoft/dockvisor/dockvisor/filetree"
	"github.com/khulnasoft/dockvisor/dockvisor/image"
)

// Layer represents a Docker image layer and metadata
type layer struct {
	history historyEntry
	index   int
	tree    *filetree.FileTree
}

// String represents a layer in a columnar format.
func (l *layer) ToLayer() *image.Layer {
	id := strings.Split(l.tree.Name, "/")[0]
	return &image.Layer{
		Id:      id,
		Index:   l.index,
		Command: strings.TrimPrefix(l.history.CreatedBy, "/bin/sh -c "),
		Size:    l.history.Size,
		Tree:    l.tree,
		// todo: query docker api for tags
		Names:  []string{"(unavailable)"},
		Digest: l.history.ID,
	}
}
