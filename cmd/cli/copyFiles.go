package main

import (
	"embed"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(templatePath string, targetFile string) error {
	if fileExists(targetFile) {
		return errors.New(targetFile + " already exist")
	}

	data, err := templateFS.ReadFile(templatePath)

	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func appendFileFromTemplate(templatePath string, targetFile string) error {
	// Read the template file
	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	// Open the target file in append mode
	f, err := os.OpenFile(targetFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Append the template contents
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

// New function to copy all files from a directory
func copyFilesFromDir(sourceDir string, targetDir string) error {
	entries, err := templateFS.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("error reading source directory from templateFS: %w", err)
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(sourceDir, entry.Name())
		targetPath := filepath.Join(targetDir, entry.Name())

		if entry.IsDir() {
			// Create the target directory if it doesn't exist
			err = os.MkdirAll(targetPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("error creating directory %s: %w", targetPath, err)
			}

			// Recursively call copyFilesFromDir for subdirectories
			err = copyFilesFromDir(sourcePath, targetPath)
			if err != nil {
				return err
			}
		} else {
			data, err := templateFS.ReadFile(sourcePath)
			if err != nil {
				return fmt.Errorf("error reading file from templateFS: %w", err)
			}

			err = copyDataToFile(data, targetPath)
			if err != nil {
				return fmt.Errorf("error copying data to file: %w", err)
			}
		}
	}

	return nil
}
