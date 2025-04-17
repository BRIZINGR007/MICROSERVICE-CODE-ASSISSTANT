package models

type CodeBaseData struct {
	CodeBaseId   string `json:"codebase_id" bson:"codebase_id"`
	CodeBaseName string `json:"codebase_name" bson:"codebase_name"`
	GitHubURL    string `json:"github_url" bson:"github_url"`
	Username     string `json:"username" bson:"username"`
	Token        string `json:"token" bson:"token"`
	Branch       string `json:"branch" bson:"branch"`
	FolderPath   string `json:"folder_path" bson:"folder_path"`
}
type User struct {
	UserID       string         `json:"userid" bson:"userid"`
	Name         string         `json:"name" bson:"name"`
	Email        string         `json:"email" bson:"email"`
	Password     string         `json:"password" bson:"password" `
	CodebaseData []CodeBaseData `json:"codebasedata" bson:"codebasedata"`
}
