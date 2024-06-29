package handler

import (
	"EmegencyNotificationSystem/profile_service/models"
	"errors"
	"github.com/graphql-go/graphql"
)

func (h *Handler) signUp(p graphql.ResolveParams) (interface{}, error) {
	input, ok := p.Args["input"].(map[string]interface{})
	if !ok {
		return nil, newErrorResponse("invalid input body", errors.New("invalid input body"))
	}
	name, nameOk := input["name"].(string)
	email, emailOk := input["email"].(string)
	password, passwordOk := input["password"].(string)
	privateKey, privateKeyOk := input["privateKey"].(string)

	if !nameOk || !emailOk || !passwordOk || !privateKeyOk ||
		name == "" || email == "" || password == "" || privateKey == "" {
		return nil, newErrorResponse("invalid input body", errors.New("invalid input body"))
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	id, err := h.services.Authorization.CreateUser(user, privateKey)
	if err != nil {
		return nil, newErrorResponse("service failure", err)
	}

	user.Id = id
	return user, nil
}
