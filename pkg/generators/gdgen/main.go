package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func generate(outputDir string) error {
	var godotPath string
	switch runtime.GOOS {
	case "darwin":
		godotPath = "/Applications/Godot.app/Contents/MacOS/Godot"
	case "windows":
		godotPath = "godot.exe"
	default:
		godotPath = "godot"
	}
	godotPathEnv := os.Getenv("GODOT_PATH")
	if godotPathEnv != "" {
		godotPath = godotPathEnv
	}
	out, err := exec.Command(godotPath, "--headless", "--dump-gdextension-interface", "--dump-extension-api").Output()
	if err != nil {
		if err == exec.ErrNotFound {
			return fmt.Errorf("could not find Godot at %q, use the `GODOT_PATH` env var to specify the path to the godot binary", godotPath)
		}
		return fmt.Errorf("failed to run Godot (%q): %w", string(out), err)
	}
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fail("Usage: gdgen <output-dir>")
	}
	outputDir := args[0]
	err := generate(outputDir)
	if err != nil {
		fail(fmt.Sprintf("error: %s", err))
	}
}
