## gmx

A Golang based implementation of genmax. This is a code/file generation tool that feeds data into templates. This can be used for scaffolding projects.

#### Installation

**Quick Install (Recommended):**

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/razpinator/gmx/main/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/razpinator/gmx/main/install.ps1 | iex
```

**Go-based Installer:**

For users who prefer a native Go installer, you can use the Go-based installer:

```bash
# Download and run the Go installer
go install github.com/razpinator/gmx/installer@latest
gmx-installer
```

Or build from source:
```bash
git clone https://github.com/razpinator/gmx.git
cd gmx/installer
go build -o gmx-installer .
./gmx-installer
```

**Manual Installation:**

If you have Golang installed using below command.

```bash
go install github.com/razpinator/gmx@latest
```

*Note: After manual installation, you may need to add `$(go env GOPATH)/bin` to your PATH.*

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
