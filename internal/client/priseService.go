package client

import (
	"context"
	"github.com/EgorBessonov/client-service/internal/model"
	priceService "github.com/EgorBessonov/price-service/protocol"
	"github.com/sirupsen/logrus"
)

func (client *Client) SubscribePriceService(ctx context.Context) {
	streamOpts := []int32{1, 2, 3, 4, 5}
	stream, err := client.PriseService.Get(ctx, &priceService.GetRequest{Name: streamOpts})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("client: can't send request to price server")
	}
	for {
		select {
		case <-stream.Context().Done():
			return
		default:
			shares, err := stream.Recv()
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("client: error while reading gRPC stream")
			}
			share := &model.Share{
				Name:      shares.Share.Name,
				Bid:       shares.Share.Bid,
				Ask:       shares.Share.Ask,
				UpdatedAt: shares.Share.Time,
			}
			client.SaveOrUpdate(share)
		}
	}
}
