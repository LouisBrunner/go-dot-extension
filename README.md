# go-dot-extension

**WARNING: This project is still in early development and the API might break frequently.**

This project allows you to create a Godot extension in Go.

The API is heavily inspired by https://github.com/ShadowApex/godot-go.

## Usage

In order to use Go in your Godot project, you need to compile the Go code into a shared library and load it in Godot.
This is done by using the Go `-buildmode=c-shared` flag on a `main` package with a `main` function.

### Boilerplate code

Create a go file containing the following:

```go
package main

func init() {
}

func main() {
}
```

We will be using the `init` function to register our custom Godot classes, implemented in Go.

## Create a custom Godot class

A basic Godot class is defined as a struct with an embbeded `gdapi` struct. For example, to inherit from `Node2D`:

```go
type MyNode2D struct {
  gdapi.Node2D
}
```

Any public method receiver defined on `MyNode2D` will be available as a method on the Godot class and any public field will be available as a property.

```go
type MyNode2D struct {
  gdapi.Node2D

  Speed int
}

func (n *MyNode2D) Move(vec gdapi.Vector2) {
  Node2D.SetPosition(vec.MulScalar(n.Speed))
}
```

## Register the custom Godot class

In the `init` function, register the custom Godot class:

```go
package main

import (
  gde "github.com/LouisBrunner/go-dot-extension"
)

func NewMyNode2D() gde.Class {
  return &MyNode2D{}
}

func init() {
  gde.Register(NewMyNode2D)
}

func main() {
}
```

### Compiling the shared library

Your package is now ready to be built as a shared library.

```bash
# Linux
go build -buildmode=c-shared -o libmyextension.so pkg_folder
# macOS
go build -buildmode=c-shared -o libmyextension.dylib pkg_folder
# Windows
go build -buildmode=c-shared -o libmyextension.dll pkg_folder
```

You can name the shared library whatever you want, the example above uses `libmyextension`.

### How to write your class

Here are some notes on how to write your class:

- Any exported method will be available as a method on the Godot class (apart from any function called `Destroy`)
- Any exported field will be available as a property on the Godot class (apart from any with type `gdapi.Signal` or `gdapi.SignalSubscribers`)
- Any method starting with `X_` will be assumed to be a virtual method (e.g. `X_Ready` for `_ready`)
- All methods and fields will be exported as `snake_case` (e.g. `MyMethod` will be `my_method`), this might lead to name conflicts, so make sure to use unique names
- All exported methods must have 1 or no return value, multiple return values are not supported
- Any exported property which is a `gdapi.Object` (basically anything from `gdapi`) will be automatically initialized for you, if it is a pointer however, you will need to initialize it yourself
- Any `gdapi.SignalSubscribers` field will be automatically initialized for you, even if its not exported, as stated above, it must be a field and not a pointer however

### Including the shared library in Godot

You will need to create a file with the extension `.gdextension` and place it in your project for Godot to be able to load the shared library.

```gdscript
[configuration]
entry_symbol = "go_dot_gdextension_entry"
compatibility_minimum = "4.1"

[libraries]
macos.debug.arm64 = "res://your_folder/libmyextension-darwin-arm64.dylib"
macos.release.arm64 = "res://your_folder/libmyextension-darwin-arm64.dylib"
macos.debug.amd64 = "res://your_folder/libmyextension-darwin-amd64.dylib"
macos.release.amd64 = "res://your_folder/libmyextension-darwin-amd64.dylib"
windows.debug.amd64 = "res://your_folder/libmyextension-windows-amd64.dll"
windows.release.amd64 = "res://your_folder/libmyextension-windows-amd64.dll"
linux.debug.amd64 = "res://your_folder/libmyextension-linux-amd64.so"
linux.release.amd64 = "res://your_folder/libmyextension-linux-amd64.so"
```

You will need to replace `your_folder` with the path to the folder containing the shared library and `libmyextension` with the name of the shared library.

### Examples

You can find examples in the `examples` folder.

- `simple`: a simple example of a custom Godot class, which basically reproduces the above usage
- `property-getset`: an example of how to use getter and setter for your properties
- `signals`: an example of how to use signals
- `gdscript`: an example of how to use the custom Godot class in GDScript

## Issues

- ptrcall methods return type conversion is unimplemented
- `Signal`s have no arguments in the Godot documentation/type hint system
- potential panic when using functions which return `Array`s (bad support for the cast from `Variant`)
- potential panic when using varargs functions (bad error handling)
- potential panic when using functions which deal with enums (bad pointer passing)

## Acknowledgements

- https://github.com/ShadowApex/godot-go: for their great work and straight-forward API which inspired this project

## See also

- https://github.com/godot-go/godot-go: more mature project with a similar goal
