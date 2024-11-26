package cmd

import (
	"context"
	"github.com/eldius/nft-pocs/internal/client/eth"
	"github.com/eldius/nft-pocs/internal/configs"

	"github.com/spf13/cobra"
)

// ethNftCmd represents the nft command
var ethNftCmd = &cobra.Command{
	Use:   "nft",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if err := eth.Mint(ctx, configs.GetNetworkEndpoint(), ethNftOpts.contactPath, ""); err != nil {
			panic(err)
		}
	},
}

var (
	ethNftOpts struct {
		contactPath string
	}
)

func init() {
	ethCmd.AddCommand(ethNftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ethNftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ethNftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	ethNftCmd.Flags().StringVar(&ethNftOpts.contactPath, "contract-path", "", "Path to compiled contract file")
}
