package cmd

import (
	"log"
	"path/filepath"

	"github.com/razaibi/gmx/logic"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

// processCmd represents the process command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Process data and template files",
	Long:  `Reads a configuration YAML file, processes the specified data and template files, and generates output files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Usage: myapp process <config.yaml>")
		}
		workflowFile := args[0]

		workflow := logic.ReadConfig(
			filepath.Join(
				"_gmx",
				"workflows",
				workflowFile,
			),
		)

		for _, item := range workflow.Items {
			switch item.Action {
			case "generate":
				logic.GenerateFile(item)
			case "create-file":
				logic.WriteFileWithCustomSeparator(
					item.OutputFile,
					[]byte(item.Content),
					0644,
				)
			case "exec":
				logic.RunCommand(item)
			case "download":
				logic.DownloadFile(item)
			default:
				logic.GenerateFile(item)
			}

		}
	},
}
