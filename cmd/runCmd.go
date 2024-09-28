package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/razaibi/gmx/logic"

	"github.com/osteele/liquid"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Config holds the filenames for data, template, and output
type Config struct {
	Items []Item `yaml:"items"`
}

// Item holds individual data, template, and output file paths
type Item struct {
	DataFile     string `yaml:"dataFile"`
	TemplateFile string `yaml:"templateFile"`
	OutputFile   string `yaml:"outputFile"`
}

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

		workflow := readConfig(
			filepath.Join(
				"_gmx",
				"workflows",
				workflowFile,
			),
		)

		for _, item := range workflow.Items {
			data := readJSON(
				filepath.Join(
					"_gmx",
					"data",
					item.DataFile,
				),
			)
			templateContent := readFile(
				filepath.Join(
					"_gmx",
					"templates",
					item.TemplateFile,
				),
			)

			// Parse the Liquid template
			engine := liquid.NewEngine()
			engine.RegisterFilter("pluralize", logic.Pluralize)
			engine.RegisterFilter("kebabcase", logic.ConvertToKebabCase)
			engine.RegisterFilter("camelcase", logic.ConvertToCamelCase)
			engine.RegisterFilter("snakecase", logic.ConvertToSnakeCase)
			engine.RegisterFilter("pascalecase", logic.ConvertToPascaleCase)
			engine.RegisterFilter("uuid", logic.GenerateUUID)
			engine.RegisterFilter("secret", logic.Generate16bitSecret)
			engine.RegisterFilter("secret_complex", logic.Generate64BitSecret)
			engine.RegisterFilter("env", logic.ReadEnvValue)
			engine.RegisterFilter("path", logic.JoinPath)
			engine.RegisterFilter("lower_first", logic.LowerFirst)

			output, err := engine.ParseAndRenderString(templateContent, data)
			if err != nil {
				log.Fatalf("Failed to render template: %v", err)
			}

			// Write the output to the specified file
			fileErr := writeFileWithCustomSeparator(item.OutputFile, []byte(output), 0644)
			if fileErr != nil {
				log.Fatalf("Failed to write output file: %v", err)
			}

			fmt.Printf("Output generated successfully for %s!\n", item.OutputFile)
		}
	},
}

func writeFileWithCustomSeparator(filePath string, data []byte, perm os.FileMode) error {
	// Replace custom path separator with OS-specific path separator
	normalizedPath := strings.ReplaceAll(filePath, ">", string(os.PathSeparator))

	// Get the directory path
	dirPath := filepath.Dir(normalizedPath)

	// Create directories if they don't exist
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Write the file
	if err := os.WriteFile(normalizedPath, data, perm); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// readConfig reads and parses the YAML configuration file
func readConfig(filename string) Config {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	return config
}

// readJSON reads and parses the JSON data file
func readJSON(filename string) map[string]interface{} {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read data file: %v", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatalf("Failed to parse data file: %v", err)
	}

	return data
}

// readFile reads the content of a file
func readFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read template file: %v", err)
	}

	return string(content)
}
