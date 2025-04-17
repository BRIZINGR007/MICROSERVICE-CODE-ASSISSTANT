package domain

import (
	"errors"
	"strings"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
)

type RepoRequest struct {
	GitHubURL    string `json:"github_url"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	Branch       string `json:"branch"`
	CodeBaseName string `json:"codebase_name" bson:"codebase_name"`
	FolderPath   string `json:"folder_path"`
}

type FileContent struct {
	FilePath string `json:"filePath"`
	Code     string `json:"code"`
	Chunks   []string
}

func ValidateCodeBaseData(req *models.CodeBaseData) error {
	if strings.TrimSpace(req.CodeBaseId) == "" {
		return errors.New("codeBaseId is  required ")
	}
	if strings.TrimSpace(req.GitHubURL) == "" {
		return errors.New("gitHubURL is required")
	}
	if strings.TrimSpace(req.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(req.Token) == "" {
		return errors.New("token is required")
	}
	if strings.TrimSpace(req.Branch) == "" {
		return errors.New("branch is required")
	}
	if strings.TrimSpace(req.FolderPath) == "" {
		return errors.New("folderPath is required")
	}
	return nil
}
