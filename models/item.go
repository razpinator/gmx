package models

// Item holds individual data, template, and output file paths
type Item struct {
	Action       string `yaml:"action"`
	DataFile     string `yaml:"dataFile"`
	TemplateFile string `yaml:"templateFile"`
	OutputFile   string `yaml:"outputFile"`
	Content      string `yaml:"content"`
	Cmd          string `yaml:"cmd"`
	Source       string `yaml:"source"`
}
