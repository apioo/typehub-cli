package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "typehub",
	Short: "TypeHub is a next generation API and data design platform.",
	Long:  `TypeHub is a next generation API and data design platform. Complete documentation is available at https://typehub.cloud/`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var sdkClient = SdkClient{
	ClientId:     "",
	ClientSecret: "",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&sdkClient.ClientId, "client-id", "", "The client id is either your username or an app key which you can create at our typehub.cloud backend.")
	rootCmd.PersistentFlags().StringVar(&sdkClient.ClientSecret, "client-secret", "", "This client secret is either your password or an app secret which you can create at our typehub.cloud backend.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
