package cmd

import (
	"github.com/eldius/nft-pocs/internal/configs"
	"os"

	cfg "github.com/eldius/initial-config-go/configs"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nft-pocs",
	Short: "A POC to learn how to interact with Ethereum network",
	Long:  `A POC to learn how to interact with Ethereum network.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return cfg.InitSetup(
			configs.AppName,
			cfg.WithConfigFileToBeUsed(cfgFile),
			cfg.WithDefaultValues(map[string]any{
				cfg.LogFormatKey:         cfg.LogFormatJSON,
				cfg.LogLevelKey:          cfg.LogLevelDEBUG,
				cfg.LogOutputFileKey:     "./execution.log",
				cfg.LogOutputToStdoutKey: true,
			}),
		)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	cfgFile string
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nft-pocs.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
