package client

import (
	"context"
	"fmt"
	"github.com/EgorBessonov/client-service/internal/model"
	tradeService "github.com/EgorBessonov/trade-service/protocol"
)

func (client *Client) OpenPosition(ctx context.Context, posRequest *model.OpenRequest) (string, error) {
	result, err := client.TradeService.OpenPosition(ctx, &tradeService.OpenPositionRequest{
		UserId:    posRequest.UserID,
		ShareType: posRequest.ShareType,
		Count:     posRequest.ShareCount,
		Price:     client.ShareList[posRequest.ShareType].Bid,
	})
	if err != nil {
		return "", fmt.Errorf("client: can't open position - %e", err)
	}
	return result.PositionID, nil
}

func (client *Client) ClosePosition(ctx context.Context, closeRequest *model.CloseRequest) (string, error) {
	result, err := client.TradeService.ClosePosition(ctx, &tradeService.ClosePositionRequest{
		PositionId: closeRequest.PositionID,
		Price:      client.ShareList[closeRequest.ShareType].Ask,
		UserId:     closeRequest.UserID,
	})
	if err != nil {
		return "", fmt.Errorf("client: can't close position - %e", err)
	}
	return result.Result, nil
}
