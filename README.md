# File Explorer Using Go
--
    import "."


## Usage

#### type CommandLineApp

```go
type CommandLineApp struct {
	CurrentDir string
}
```

CommandLineApp represents the state of the command-line application.

#### func  NewCommandLineApp

```go
func NewCommandLineApp() *CommandLineApp
```
NewCommandLineApp creates a new CommandLineApp instance and initializes it with
the current working directory.

#### func (*CommandLineApp) Run

```go
func (app *CommandLineApp) Run()
```
Run is the main application loop that handles user input and executes commands.
