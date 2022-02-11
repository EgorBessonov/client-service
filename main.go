package main

import (
	"context"
	balanceService "github.com/EgorBessonov/balance-service/protocol"
	"github.com/EgorBessonov/client-service/internal/client"
	"github.com/EgorBessonov/client-service/internal/config"
	"github.com/EgorBessonov/client-service/internal/service"
	priceService "github.com/EgorBessonov/price-service/protocol"
	tradeService "github.com/EgorBessonov/trade-service/protocol"
	userService "github.com/EgorBessonov/user-service/protocol"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("client service: can't parse config")
	}
	priceConn, err := grpc.Dial(cfg.PriceServerPort, grpc.WithInsecure())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("client service: can't connect to price service")
	}
	balanceConn, err := grpc.Dial(cfg.BalanceServerPort, grpc.WithInsecure())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("client service: can't connect to balance service")
	}
	userConn, err := grpc.Dial(cfg.UserServerPort, grpc.WithInsecure())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("client service: can't connect to user service")
	}
	tradeConn, err := grpc.Dial(cfg.TradeServerPort, grpc.WithInsecure())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("client service: can't connect to trade service")
	}
	defer func() {
		if err := balanceConn.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("client service: error while closing connection to balance service")
		}
		if err := userConn.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("client service: error while closing connection to user service")
		}
		if err := tradeConn.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("client service: error while closing connection to trade service")
		}
		if err := priceConn.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Error("client service: error while closing connection to price service")
		}
	}()
	balanceCli := balanceService.NewBalanceClient(balanceConn)
	userCli := userService.NewUserClient(userConn)
	tradeCli := tradeService.NewTraderClient(tradeConn)
	priceCli := priceService.NewPriceClient(priceConn)
	cli := client.NewClient(balanceCli, tradeCli, userCli, priceCli)
	clientService := service.NewService(cli)
	err = clientService.Authentication(context.Background(), "egor@gmail.com", "1111")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed")
	}
	err = clientService.TopUp(context.Background(), 10000)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed")
	}
	positionID, err := clientService.OpenPosition(context.Background(), 3, 5)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed")
	}
	err = clientService.ClosePosition(context.Background(), positionID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed")
	}
}
