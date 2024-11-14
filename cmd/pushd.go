package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/apioo/typehub-sdk-go/sdk"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(pushdCmd)
}

var pushdCmd = &cobra.Command{
	Use:   "pushd [directory]",
	Short: "Push every file inside the provided directory to TypeHub",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var client = sdkClient.GetClient()

		var directory = args[0]

		user, err := client.Authorization().GetWhoami()
		if err != nil {
			log.Fatal(err)
		}

		// import directory
		files, err := os.ReadDir(directory)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			var schemaFile = file.Name()
			var fileExtension = filepath.Ext(schemaFile)

			if fileExtension != ".json" {
				continue
			}

			var documentName = schemaFile[0 : len(schemaFile)-5]

			spec, err := readFile(filepath.Join(directory, schemaFile))
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
				fmt.Println("* " + documentName + ": " + message.Message)
			} else {
				fmt.Println("* " + documentName + ": " + message.Message)
			}
		}
	},
}
