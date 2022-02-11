package client

import (
	balanceService "github.com/EgorBessonov/balance-service/protocol"
	"github.com/EgorBessonov/client-service/internal/model"
	priceService "github.com/EgorBessonov/price-service/protocol"
	tradeService "github.com/EgorBessonov/trade-service/protocol"
	userService "github.com/EgorBessonov/user-service/protocol"
	"sync"
	"time"
)

type Client struct {
	BalanceService balanceService.BalanceClient
	TradeService   tradeService.TraderClient
	UserService    userService.UserClient
	PriseService   priceService.PriceClient
	ShareList      []*model.Share
	mutex          sync.Mutex
}

func NewClient(balanceService balanceService.BalanceClient, tradeService tradeService.TraderClient, userService userService.UserClient, priceService priceService.PriceClient) *Client {
	return &Client{
		BalanceService: balanceService,
		TradeService:   tradeService,
		UserService:    userService,
		PriseService:   priceService,
		mutex:          sync.Mutex{},
		ShareList: []*model.Share{
			{
				Name:      1,
				Bid:       2874.16,
				Ask:       2878.31,
				UpdatedAt: time.Now().Format(time.RFC3339Nano),
			},
			{
				Name:      2,
				Bid:       170.02,
				Ask:       171.71,
				UpdatedAt: time.Now().Format(time.RFC3339Nano),
			},
			{
				Name:      3,
				Bid:       307.90,
				Ask:       308.54,
				UpdatedAt: time.Now().Format(time.RFC3339Nano),
			},
			{
				Name:      4,
				Bid:       382.95,
				Ask:       384.11,
				UpdatedAt: time.Now().Format(time.RFC3339Nano),
			},
			{
				Name:      5,
				Bid:       54.27,
				Ask:       55.16,
				UpdatedAt: time.Now().Format(time.RFC3339),
			},
		},
	}
}

func (client *Client) SaveOrUpdate(share *model.Share) {
	client.mutex.Lock()
	client.ShareList[share.Name-1] = share
	client.mutex.Unlock()
}

func (client *Client) Get(shareType int) *model.Share {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	return client.ShareList[shareType]
}
