## gmx

A Golang based implementation of genmax. This is a code/file generation tool that feeds data into templates. This can be used for scaffolding projects.

#### Installation

If you have Golang installed using below command.

```bash
go install github.com/razpinator/gmx
```

Alternatively, you can visit the [Releases](https://github.com/razpinator/gmx/releases) page for platform specific files.

#### Commands

| Command | Description |
|:---|:---|
|gmx init | Create a new project.|
|gmx run \<workflow-name> | Run a workflow and generate your files.|

#### Extensions supported in templates

The following methods are supported in the template:

| Description | Usage |
|:---|:---|
|Pluralize| `{{ "dog" \| pluralize }}`|
|Kebab Case| `{{ "Hello World" \| kebabcase }}`|
|Camel case| `{{ "Hello World" \| camelcase }}`|
|Snake case| `{{ "Hello World" \| snakecase }}`|
|Pascale case| `{{ "hello world" \| pascalecase }}`|
|UUID Generation| `{{ "" \| uuid }}`|
|Generate secret in 16 bit - hexadecimal| `{{ "" \| secret }}`|
|Generate secret in 64 bit - hexadecimal| `{{ "" \| secret_complex }}`|
|Read value from env file| `{{ "MY_CONFIG_KEY" \| config: ".env" }}`|
|Join strings to make a file path.| `{{ "home" \| joinpath: ["documents", "file.txt"] }}`|
|Convert first character to lower case.| ` {{ "Hello World" \| lowerfirst }}`|
