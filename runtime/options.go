package runtime

import (
	"github.com/spf13/viper"

	"github.com/khulnasoft/dockvisor/dockvisor"
)

type Options struct {
	Ci           bool
	Image        string
	Source       dockvisor.ImageSource
	IgnoreErrors bool
	ExportFile   string
	CiConfig     *viper.Viper
	BuildArgs    []string
}
