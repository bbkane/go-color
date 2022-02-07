# Deprecated

See [bbkane/gocolor](https://github.com/bbkane/gocolor)

# go-color

Forked from [TwiN/go-color](https://github.com/TwiN/go-color). The main change is an explicit `Init()` function that "activates" `Add()`. This allows calling code to determine whether or not colors are used. To keep the API surface small, I've also removed the `Colorize()` function in favor of `Add()` Many thanks to @TwiN!!

An extremely lightweight cross-platform package to colorize text in terminals.

This is not meant for maximal compatibility, nor is it meant to handle a plethora of scenarios.
It will simply wrap a message with the necessary characters, if the OS handles it.

There are many cases in which this would not work, such as the output being redirected to something other
than a terminal (such as a file, i.e. `executable >> file.txt`)

## Usage

### Function

You can use the `color.Add(color, message)` function
in conjunction with a variable like so:
```go
package main

import "github.com/TwiN/go-color"

func main() {
    Init()
    println(color.Add(color.Red, "This is red"))
}
```

### Variables only

Unlike using the aforementioned approach, directly using the color variables will required you to manually
prepend `color.Reset` at the end of your string.

You can either directly use the variables like so:

```go
package main

import "github.com/TwiN/go-color"

func main() {
    println(color.Red + "This is red" + color.Reset)
    println(color.Green + "This is green" + color.Reset)
    println(color.Blue + "This is blue" + color.Reset)
    println(color.Purple + "This is purple" + color.Reset)
    println(color.Yellow + "This is yellow" + color.Reset)
    println(color.Cyan + "This is cyan" + color.Reset)
    println(color.Gray + "This is gray" + color.Reset)
    println(color.White + "This is white" + color.Reset)
}
```

**NOTE**: If you're going to use this approach, don't forget to append your string with `color.Reset`,
otherwise everything else in your terminal will be that color until the color is reset or overridden.


## Installation

```
go get github.com/bbkane/go-color
```
