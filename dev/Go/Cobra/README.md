—-

## Cobra Command

Additional commands can be defined and typically are each given their own file inside of the cmd/ directory.

If you wanted to create a version command you would create cmd/version.go and populate it with the following:

<details><summary>show</summary>
<p>

```go

package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}

```

</p>
</details>

—-

## Cobra- Persistent Flags

A flag can be 'persistent' meaning that this flag will be available to the command it's assigned to as well as every command under that command. For global flags, assign a flag as a persistent flag on the root.

<details><summary>show</summary>
<p>

``` 

rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

```

</p>
</details>

—-

## Cobra- Local Flags

A flag can also be assigned locally which will only apply to that specific command.

<details><summary>show</summary>
<p>

```go
rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")

```

</p>
</details>

—-

## Cobra- Local Flag on Parent Commands

By default Cobra only parses local flags on the target command, any local flags on parent commands are ignored. By enabling Command.TraverseChildren Cobra will parse local flags on each command before executing the target command.

<details><summary>show</summary>
<p>

```go
command := cobra.Command{
  Use: "print [OPTIONS] [COMMANDS]",
  TraverseChildren: true,
}

```

</p>
</details>

—-

## Cobra- Bind Flags with Config

You can also bind your flags with viper:

<details><summary>show</summary>
<p>

```
var author string

func init() {
  rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
  viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}

```

</p>
</details>

—-

## Cobra- Required flags

Flags are optional by default. If instead you wish your command to report an error when a flag has not been set, mark it as required:

<details><summary>show</summary>
<p>

```
rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
rootCmd.MarkFlagRequired("region")

```

</p>
</details>

—-

## Cobra- Positional and Custom Arguments

Validation of positional arguments can be specified using the Args field of Command.



<details><summary>show</summary>
<p>

```
The following validators are built in:

* NoArgs - the command will report an error if there are any positional args.
* ArbitraryArgs - the command will accept any args.
* OnlyValidArgs - the command will report an error if there are any positional args that are not in the ValidArgs field of Command.
* MinimumNArgs(int) - the command will report an error if there are not at least N positional args.
* MaximumNArgs(int) - the command will report an error if there are more than N positional args.
* ExactArgs(int) - the command will report an error if there are not exactly N positional args.
* RangeArgs(min, max) - the command will report an error if the number of args is not between the minimum and maximum number of expected args.

```

</p>
</details>

—-

## Cobra- Setting the custom validator:

<details><summary>show</summary>
<p>

```golang
var cmd = &cobra.Command{
  Short: "hello",
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("requires at least one arg")
    }
    if myapp.IsValidColor(args[0]) {
      return nil
    }
    return fmt.Errorf("invalid color specified: %s", args[0])
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello, World!")
  },
}

```

</p>
</details>

—-

## Cobra- Example

In the example below, we have defined three commands. Two are at the top level and one (cmdTimes) is a child of one of the top commands. In this case the root is not executable meaning that a subcommand is required. This is accomplished by not providing a 'Run' for the 'rootCmd'.

<details><summary>show</summary>
<p>

```
package main

import (
  "fmt"
  "strings"

  "github.com/spf13/cobra"
)

func main() {
  var echoTimes int

  var cmdPrint = &cobra.Command{
    Use:   "print [string to print]",
    Short: "Print anything to the screen",
    Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Print: " + strings.Join(args, " "))
    },
  }

  var cmdEcho = &cobra.Command{
    Use:   "echo [string to echo]",
    Short: "Echo anything to the screen",
    Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Print: " + strings.Join(args, " "))
    },
  }

  var cmdTimes = &cobra.Command{
    Use:   "times [# times] [string to echo]",
    Short: "Echo anything to the screen more times",
    Long: `echo things multiple times back to the user by providing
a count and a string.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      for i := 0; i < echoTimes; i++ {
        fmt.Println("Echo: " + strings.Join(args, " "))
      }
    },
  }

  cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

  var rootCmd = &cobra.Command{Use: "app"}
  rootCmd.AddCommand(cmdPrint, cmdEcho)
  cmdEcho.AddCommand(cmdTimes)
  rootCmd.Execute()
}

```

</p>
</details>

—-

