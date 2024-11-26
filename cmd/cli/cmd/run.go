package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eldius/nft-pocs/internal/client/ethereum"
	"github.com/eldius/nft-pocs/internal/configs"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if err := ethereum.Mint(ctx, configs.GetNetworkEndpoint(), configs.GetPrivateKey(), readImgFile(runOpts.inputFile)); err != nil {
			panic(err)
		}
	},
}

var (
	runOpts struct {
		inputFile string
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	runCmd.Flags().StringVarP(&runOpts.inputFile, "input", "i", "", "input file")
}

func readImgFile(f string) string {
	bytes, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)
	return base64Encoding
}
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
