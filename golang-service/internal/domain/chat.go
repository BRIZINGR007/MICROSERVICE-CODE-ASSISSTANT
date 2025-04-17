package domain

type ChatReq struct {
	CodeBaseId   string `json:"code_base_id"`
	AIAnswer     string `json:"ai_answer"`
	UserQuestion string `json:"user_question"`
}
