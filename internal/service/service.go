package service

import (
	"context"
	"fmt"
	"githb.com/EgorBessonov/client-service/internal/client"
	"githb.com/EgorBessonov/client-service/internal/model"
)

type Service struct {
	Client    *client.Client
	UserID    string
	Positions map[string]int32
}

func NewService(client *client.Client) *Service {
	return &Service{
		Client:    client,
		UserID:    "",
		Positions: make(map[string]int32),
	}
}

func (service *Service) Authentication(ctx context.Context, email, password string) error {
	userID, err := service.Client.Authentication(ctx, email, password)
	if err != nil {
		return fmt.Errorf("client service: can't get user - %e", err)
	}
	service.UserID = userID
	return nil
}

func (service *Service) OpenPosition(ctx context.Context, shareType, shareCount int32) (string, error) {
	positionID, err := service.Client.OpenPosition(ctx, &model.OpenRequest{
		UserID:     service.UserID,
		ShareType:  shareType,
		ShareCount: shareCount,
	})
	if err != nil {
		return "", fmt.Errorf("client service: can't open position - %e", err)
	}
	service.Positions[positionID] = shareType
	return positionID, nil
}

func (service *Service) ClosePosition(ctx context.Context, positionID string) error {
	_, err := service.Client.ClosePosition(ctx, &model.CloseRequest{
		PositionID: positionID,
		ShareType:  service.Positions[positionID],
		UserID:     service.UserID,
	})
	if err != nil {
		return fmt.Errorf("client service: can't close position")
	}
	return nil
}

func (service *Service) TopUp(ctx context.Context, shift float32) error {
	err := service.Client.TopUp(ctx, service.UserID, shift)
	if err != nil {
		return fmt.Errorf("client service: can't top up balance - %e", err)
	}
	return nil
}
