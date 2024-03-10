package cmd

import (
	"fmt"
	"github.com/apioo/typehub-sdk-go/sdk"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(triggerCmd)
}

var triggerCmd = &cobra.Command{
	Use:   "trigger [document] [owner] [repo] [workflow_id] [access_token]",
	Short: "Configures a GitHub trigger for a specific document",
	Args:  cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		var client = sdkClient.GetClient()

		var documentName = args[0]
		var owner = args[1]
		var repo = args[2]
		var workflowId = args[3]
		var accessToken = args[4]

		user, err := client.Authorization().GetWhoami()
		if err != nil {
			log.Fatal(err)
		}

		// add trigger
		config := sdk.TriggerConfig{}
		config["owner"] = owner
		config["repo"] = repo
		config["workflow_id"] = workflowId
		config["access_token"] = accessToken

		trigger := sdk.TriggerCreate{
			Type:   "github",
			Config: &config,
		}

		message, err := client.Trigger().Create(user.Name, documentName, trigger)
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
