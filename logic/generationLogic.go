package logic

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/osteele/liquid"
	"github.com/razaibi/gmx/models"
)

func GenerateFile(item models.Item) {
	data := ReadJSON(
		filepath.Join(
			"_gmx",
			"data",
			item.DataFile,
		),
	)
	templateContent := ReadFile(
		filepath.Join(
			"_gmx",
			"templates",
			item.TemplateFile,
		),
	)

	// Parse the Liquid template
	engine := liquid.NewEngine()
	engine.RegisterFilter("pluralize", Pluralize)
	engine.RegisterFilter("kebabcase", ConvertToKebabCase)
	engine.RegisterFilter("camelcase", ConvertToCamelCase)
	engine.RegisterFilter("snakecase", ConvertToSnakeCase)
	engine.RegisterFilter("pascalecase", ConvertToPascaleCase)
	engine.RegisterFilter("uuid", GenerateUUID)
	engine.RegisterFilter("secret", Generate16bitSecret)
	engine.RegisterFilter("secret_complex", Generate64BitSecret)
	engine.RegisterFilter("env", ReadEnvValue)
	engine.RegisterFilter("path", JoinPath)
	engine.RegisterFilter("lower_first", LowerFirst)

	output, err := engine.ParseAndRenderString(templateContent, data)
	if err != nil {
		log.Fatalf("Failed to render template: %v", err)
	}

	// Write the output to the specified file
	fileErr := WriteFileWithCustomSeparator(item.OutputFile, []byte(output), 0644)
	if fileErr != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("Output generated successfully for %s!\n", item.OutputFile)

}
