package client

import (
	"context"
	"fmt"
	userService "github.com/EgorBessonov/user-service/protocol"
)

func (client *Client) Authentication(ctx context.Context, email, password string) (string, error) {
	user, err := client.UserService.Authentication(ctx, &userService.AuthenticationRequest{
		UserEmail:    email,
		UserPassword: password,
	})
	if err != nil {
		return "", fmt.Errorf("client: authentication failed - %e", err)
	}
	return user.UserId, nil
}

func (client *Client) Registration(ctx context.Context, email, name, password string) error {
	_, err := client.UserService.Registration(ctx, &userService.RegistrationRequest{
		UserEmail:    email,
		UserName:     name,
		UserPassword: password,
	})
	if err != nil {
		return fmt.Errorf("client: registration failed - %e", err)
	}
	return nil
}
