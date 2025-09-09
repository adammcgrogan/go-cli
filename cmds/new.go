package cmds

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new projects or files from templates",
	Long:  `The 'new' command helps you bootstrap new projects by generating boilerplate code and directory structures.`,
}

var projectCmd = &cobra.Command{
	Use:   "project [name]",
	Short: "Create a new Go project in your Documents folder",
	Long: `Creates a new project in ~/Documents/<name> with a basic 'Hello, World!' main.go file,
initializes a Go module, and sets up a new Git repository.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding your home directory: %v\n", err)
			return
		}

		projectDir := filepath.Join(homeDir, "Documents", projectName)

		fmt.Printf("-> Creating project directory: %s\n", projectDir)
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}

		mainGoPath := filepath.Join(projectDir, "main.go")
		mainGoContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, New Go Project!")
}
`
		fmt.Println("-> Creating a starter main.go file...")
		if err := os.WriteFile(mainGoPath, []byte(mainGoContent), 0644); err != nil {
			fmt.Printf("Error writing to main.go: %v\n", err)
			return
		}

		fmt.Printf("-> Initializing Go module (%s)...\n", projectName)
		goModCmd := exec.Command("go", "mod", "init", projectName)
		goModCmd.Dir = projectDir

		output, err := goModCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error initializing Go module. Is Go installed correctly?\n")
			fmt.Printf("--- Go command output ---\n%s\n-------------------------\n", string(output))
			return
		}

		// --- 4. Execute 'git init' ---
		fmt.Println("-> Initializing Git repository...")
		gitInitCmd := exec.Command("git", "init")
		gitInitCmd.Dir = projectDir
		if err := gitInitCmd.Run(); err != nil {
			fmt.Printf("Warning: Could not initialize Git repository. Is Git installed? Error: %v\n", err)
		}

		fmt.Printf("\n✅ Successfully created Go project in '%s'.\n", projectDir)
		fmt.Printf("   You can now `cd %s` and start coding!\n", projectDir)
	},
}

func init() {
	newCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(newCmd)
}
