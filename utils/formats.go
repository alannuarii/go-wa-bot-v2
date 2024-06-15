package utils

import(
	"fmt"
	"strings"
	"go-wa-bot-v2/models"
)

func FormatUsers(users []models.User) string {
	var sb strings.Builder
	for _, user := range users {
		sb.WriteString(fmt.Sprintf("ID: %d, Name: %s, Email: %s, Role: %s\n", user.ID, user.Name, user.Email, user.Role))
	}
	return sb.String()
}