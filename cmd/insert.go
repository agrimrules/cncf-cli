/*
Copyright © 2019 Agrim Asthana <dev@agrimasthana.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

// InsertSuccess is the default response from the API
type InsertSuccess struct {
	Response string
}

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Inserts a note",
	Long:  `Inserts a note via HTTP/GRPC to the backend`,
	Run: func(cmd *cobra.Command, args []string) {
		// HTTP Client
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetResult(&InsertSuccess{}).
			SetBody(`{ "user": "` + args[0] + `", "message": "` + args[1] + `" }`).
			Post("http://localhost:8080/post")
		if err != nil {
			fmt.Println("Error :", err)
			os.Exit(1)
		}
		fmt.Println("Successfully ", resp.Result().(*InsertSuccess).Response)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
