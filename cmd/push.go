package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/apioo/typehub-cli/sdk"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a local document to TypeHub",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var client = sdkClient.GetClient()

		var documentName = args[0]
		var schemaFile = args[1]

		user, err := client.Authorization().GetWhoami()
		if err != nil {
			log.Fatal(err)
		}

		// import document
		spec, err := readFile(schemaFile)
		if err != nil {
			log.Fatal(err)
		}

		schema := sdk.Passthru{}
		err = json.Unmarshal(spec, &schema)
		if err != nil {
			log.Fatal(err)
		}

		message, err := client.Document().Import(user.Name, documentName, schema)
		if err != nil {
			log.Fatal(err)
		}

		if message.Success == true {
			fmt.Println(message.Message)
		} else {
			log.Fatal(message.Message)
		}
	},
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
