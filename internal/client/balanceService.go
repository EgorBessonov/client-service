package client

import (
	"context"
	"fmt"
	balanceService "github.com/EgorBessonov/balance-service/protocol"
)

func (client *Client) TopUp(ctx context.Context, userId string, shift float32) error {
	_, err := client.BalanceService.TopUp(ctx, &balanceService.TopUpRequest{
		UserId: userId,
		Shift:  shift,
	})
	if err != nil {
		return fmt.Errorf("client: can't top up balance")
	}
	return nil
}

func (client *Client) Withdraw(ctx context.Context, userId string, shift float32) error {
	_, err := client.BalanceService.Withdraw(ctx, &balanceService.WithdrawRequest{
		UserId: userId,
		Shift:  shift,
	})
	if err != nil {
		return fmt.Errorf("client: can't withdraw balance")
	}
	return nil
}
