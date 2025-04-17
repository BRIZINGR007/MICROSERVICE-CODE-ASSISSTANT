package utils

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/domain"
)

func ExtractCodeService(payload *models.CodeBaseData) ([]domain.FileContent, error) {
	tmpDir, err := os.MkdirTemp("", "repo")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	cloneURL := payload.GitHubURL
	if payload.Token != "" && payload.Username != "" {
		parts := strings.Split(payload.GitHubURL, "https://")
		if len(parts) > 1 {
			cloneURL = fmt.Sprintf("https://%s:%s@%s", payload.Username, payload.Token, parts[1])
		}
	}
	var args []string
	if payload.Branch != "" {
		args = []string{"clone", "-b", payload.Branch, "--single-branch", cloneURL, tmpDir}
	} else {
		args = []string{"clone", cloneURL, tmpDir}
	}

	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("git clone failed: %w, output: %s", err, output)
	}

	basePath := tmpDir
	if payload.FolderPath != "" {
		basePath = filepath.Join(tmpDir, payload.FolderPath)
		if _, err := os.Stat(basePath); os.IsNotExist(err) {
			return nil, fmt.Errorf("specified folder path does not exist: %s", payload.FolderPath)
		}
	}

	var files []domain.FileContent

	err = filepath.Walk(basePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(basePath, path)
			if err != nil {
				fmt.Printf("Error getting relative path for %s: %v\n", path, err)
				return nil
			}
			subString := "_init_.py"
			if strings.Contains(relPath, subString) {
				return nil
			}

			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", path, err)
				return nil
			}
			files = append(files, domain.FileContent{
				FilePath: relPath,
				Code:     string(content),
			})
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}
	log.Println("FILES  ....",files)
	return files, nil
}
