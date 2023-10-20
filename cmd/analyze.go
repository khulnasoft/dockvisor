package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/khulnasoft/dockvisor/dockvisor"
	"github.com/khulnasoft/dockvisor/runtime"
)

// doAnalyzeCmd takes a docker image tag, digest, or id and displays the
// image analysis to the screen
func doAnalyzeCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		printVersionFlag, err := cmd.PersistentFlags().GetBool("version")
		if err == nil && printVersionFlag {
			printVersion(cmd, args)
			return
		}

		fmt.Println("No image argument given")
		os.Exit(1)
	}

	userImage := args[0]
	if userImage == "" {
		fmt.Println("No image argument given")
		os.Exit(1)
	}

	initLogging()

	isCi, ciConfig, err := configureCi()

	if err != nil {
		fmt.Printf("ci configuration error: %v\n", err)
		os.Exit(1)
	}

	var sourceType dockvisor.ImageSource
	var imageStr string

	sourceType, imageStr = dockvisor.DeriveImageSource(userImage)

	if sourceType == dockvisor.SourceUnknown {
		sourceStr := viper.GetString("source")
		sourceType = dockvisor.ParseImageSource(sourceStr)
		if sourceType == dockvisor.SourceUnknown {
			fmt.Printf("unable to determine image source: %v\n", sourceStr)
			os.Exit(1)
		}

		imageStr = userImage
	}

	ignoreErrors, err := cmd.PersistentFlags().GetBool("ignore-errors")
	if err != nil {
		logrus.Error("unable to get 'ignore-errors' option:", err)
	}

	runtime.Run(runtime.Options{
		Ci:           isCi,
		Source:       sourceType,
		Image:        imageStr,
		ExportFile:   exportFile,
		CiConfig:     ciConfig,
		IgnoreErrors: viper.GetBool("ignore-errors") || ignoreErrors,
	})
}
