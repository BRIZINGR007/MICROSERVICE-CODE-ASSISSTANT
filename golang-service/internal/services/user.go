package services

import (
	"context"
	"fmt"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/models"
	"github.com/BRIZINGR007/app-002-code-assistant/internal/db/repositories"
)

func SaveUser(user *models.User) error {
	repo := repositories.GetUserRepository()
	err := repo.AddUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
func GetUserByEmail(email string) (*models.User, error) {
	repo := repositories.GetUserRepository()
	user, err := repo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddCodeBaseDataForUser(email string, codeBaseData *models.CodeBaseData) error {
	repo := repositories.GetUserRepository()

	err := repo.AddCodeBaseData(context.Background(), email, codeBaseData)
	if err != nil {
		return fmt.Errorf("cannot add CodeBaseData to  the  mongo document")
	}
	return nil
}

func GetCodeBaseData(email string) ([]models.CodeBaseData, error) {
	repo := repositories.GetUserRepository()
	cd, err := repo.GetCodeBaseData(context.Background(), email)
	if err != nil {
		return nil, fmt.Errorf("error  in retrieving  code-base data")
	}
	return cd, nil

}

func DeleteCodeBaseContext(userId string, codeBaseId string) error {
	user_repo := repositories.GetUserRepository()
	chat_repo := repositories.GetChatRepository()
	codeassist_repo := repositories.GetCodeAssistRepository()
	var err error
	err = user_repo.DeleteCodeBaseIdFromUserMetaData(context.Background(), userId, codeBaseId)
	if err != nil {
		return fmt.Errorf("error in deleting the codedb context from user collection.")
	}
	err = chat_repo.DeleteChatsByCodeBaseId(context.Background(), codeBaseId)
	if err != nil {
		return fmt.Errorf("error in deleting chats .")
	}
	err = codeassist_repo.DeleteCodeContextByCodeBaseId(context.Background(), codeBaseId)
	if err != nil {
		return fmt.Errorf("error in deleting Code Context .")
	}
	return nil
}
