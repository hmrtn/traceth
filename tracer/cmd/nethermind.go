/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/hansmrtn/tracing-apis/tracer/utils"
	"github.com/spf13/cobra"
)

// nethermindCmd represents the nethermind command
var nethermindCmd = &cobra.Command{
	Use:   "nethermind",
	Short: "Trace Nethermind transactions.",
	Long:  `Requests debug_traceTransaction to trace a recent transaction.`,
	Run: func(cmd *cobra.Command, args []string) {
		tx, _ := cmd.Flags().GetString("tx")
		url, _ := cmd.Flags().GetString("url")
		filename, _ := cmd.Flags().GetString("filename")

		if tx != "" {
			var trace []byte
			if url != "" {
				trace = requestNethermindTrace(url, tx)
			} else {
				trace = requestNethermindTrace("https://api.archivenode.io/n21le9m5ogypuubc2hdh8n21le9vhv6n/nethermind", tx)
			}
			if filename != "" {
				utils.Save(filename, trace)
			} else {
				fmt.Print(string(trace))
			}
		} else {
			fmt.Println("A transaction hash is required to trace.")
		}

	},
}

func init() {
	rootCmd.AddCommand(nethermindCmd)
	nethermindCmd.PersistentFlags().String("tx", "", "Transaction Hash")
	nethermindCmd.PersistentFlags().String("url", "", "JSON RPC API url")
	nethermindCmd.PersistentFlags().String("filename", "", "Save Response to filename")
}

func requestNethermindTrace(url, tx string) []byte {
	fmt.Printf("Nethermind Tracing Tx: %s ...\n", tx)
	data := fmt.Sprintf(`{"id": 1, "jsonrpc":"2.0", "method": "debug_traceTransaction", "params": ["%s"]}`, tx)
	trace := utils.RequestTrace(url, data)
	return trace
}
