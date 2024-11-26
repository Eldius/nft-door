package cmd

import (
	"context"
	"github.com/eldius/nft-pocs/internal/client/eth"
	"github.com/eldius/nft-pocs/internal/configs"

	"github.com/spf13/cobra"
)

// ethRunCmd represents the run command
var ethRunCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if err := eth.Connect(ctx, configs.GetNetworkEndpoint()); err != nil {
			panic(err)
		}
	},
}

func init() {
	ethCmd.AddCommand(ethRunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethRunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethRunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
