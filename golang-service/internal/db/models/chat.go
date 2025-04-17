package models

type Chat struct {
	ChatId       string `json:"chat_id" bson:"chat_id"`
	UserID       string `json:"user_id" bson:"user_id"`
	CodeBaseId   string `json:"code_base_id" bson:"code_base_id"`
	AIAnswer     string `json:"ai_answer" bson:"ai_answer"`
	UserQuestion string `json:"user_question" bson:"user_question"`
}
