package rconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rgonzalezNetel/rConfig/internal/config"
)

func GenerateProjectStructure(configPath string) error {
	config, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Create folders
	for _, folder := range config.Folders.Names {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create folder '%s': %w", folder, err)
		}
	}

	// Create files with package declaration
	for _, file := range config.Files.Paths {
		packageName := getPackageName(file)
		fileContent := fmt.Sprintf("package %s\n", packageName)
		if err := createFile(file, fileContent); err != nil {
			return fmt.Errorf("failed to create file '%s': %w", file, err)
		}
	}

	return nil
}

func getPackageName(filePath string) string {
	return filepath.Base(filepath.Dir(filePath))
}

func createFile(filePath, content string) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", dir, err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", filePath, err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write content to file '%s': %w", filePath, err)
	}
	return nil
}
