package root

import (
	"github.com/cb-kubecd/jrfoo141/pkg"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "jrfoo141",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(pkg.NewCmdHelloWorld())
}
