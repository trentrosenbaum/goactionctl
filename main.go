package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const repoName = "my-golang-action"

var templateRepoURL = fmt.Sprintf("https://github.com/trentrosenbaum/%s.git", repoName)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a name for the new GitHub Action project.")
		return
	}

	projectName := os.Args[1]
	fmt.Printf("Generating GitHub Action project '%s'...\n", projectName)

	err := cloneProject(projectName)
	if err != nil {
		log.Fatalf("Error generating the project: %s\n", err)
	}

	err = removeGitDirectory(projectName)
	if err != nil {
		log.Fatalf("Error removing the .git directory: %s\n", err)
	}

	fmt.Printf("Project '%s' generated successfully!\n", projectName)
}

func cloneProject(projectName string) error {
	cmd := exec.Command("git", "clone", templateRepoURL, projectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error cloning the template project: %w", err)
	}

	// Rename the GitHub Action name in the generated project
	err = renameAction(projectName)
	if err != nil {
		return fmt.Errorf("error renaming the GitHub Action in the generated project: %w", err)
	}

	return nil
}

func renameAction(projectName string) error {
	// replace the references to the cloned repo name with the new project name
	return replaceStringInFiles(projectName, repoName, projectName)
}

func replaceStringInFiles(rootDir, searchString, replacement string) error {
	filenames := []string{"Makefile", "go.mod"}
	extensions := []string{".go", ".md"}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check the file is not a dir and is in the filter lists, (filenames and extensions)
		if !info.IsDir() && contains(filenames, info.Name()) || contains(extensions, filepath.Ext(info.Name())) {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Replace the search string with the new value
			newContent := strings.Replace(string(content), searchString, replacement, -1)

			// Write the new content back to the file
			err = os.WriteFile(path, []byte(newContent), info.Mode())
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error while replacing string: %w", err)
	}

	return nil
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func removeGitDirectory(projectName string) error {
	gitDir := fmt.Sprintf("%s/.git", projectName)
	err := os.RemoveAll(gitDir)
	if err != nil {
		return fmt.Errorf("error removing .git directory: %w", err)
	}

	return nil
}
