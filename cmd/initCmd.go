package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

// initCmd represents the process command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize sample files",
	Long:  `Sets up workflow, data and template files.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := createFolderIfNotExists("_gmx")
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		err = createFolderIfNotExists(filepath.Join("_gmx", "data"))
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		err = createFolderIfNotExists(filepath.Join("_gmx", "templates"))
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		err = createFolderIfNotExists(filepath.Join("_gmx", "workflows"))
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}

		sampleDataPath := filepath.Join("_gmx", "data", "data.json")
		content := `{
  "title": "Hello, World!",
  "description": "This is a sample description."
}`
		// Call the function to create the file with content
		err = createFileWithContent(sampleDataPath, content)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		sampleWorkflowPath := filepath.Join("_gmx", "workflows", "workflow.yaml")
		content = `items:
  - action: generate
  	dataFile: data.json
    templateFile: template.liquid
    outputFile: output1.txt

  - action: create-file
	content: "Sample content."
    outputFile: sample-file-created.txt

  - action: exec
	cmd: "ls"

  - action: download
	source: "https://raw.githubusercontent.com/snippington/snp-go-basic/refs/heads/main/.gitignore"
	`
		// Call the function to create the file with content
		err = createFileWithContent(sampleWorkflowPath, content)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		sampleTemplatePath := filepath.Join("_gmx", "templates", "template.liquid")
		content = `<h1>{{ title }}</h1>
<p>{{ description }}</p>`
		// Call the function to create the file with content
		err = createFileWithContent(sampleTemplatePath, content)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		fmt.Println("gmx Initialized.")
	},
}

func createFileWithContent(fileName string, content string) error {
	// Create or open the file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func createFolderIfNotExists(path string) error {
	// Check if the folder already exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Folder does not exist, create it
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %w", err)
		}
	}
	return nil
}
