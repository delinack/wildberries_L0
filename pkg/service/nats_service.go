package service

import (
	"L0/pkg/domain"
	"L0/pkg/repository"
	"encoding/json"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type NatsService struct {
	OrderService Order

	NatsServer stan.Conn
	ClusterID  string
}

func NewNatsService(repo *repository.Repository, clusterID string, clientID string) *NatsService {
	return &NatsService{
		OrderService: NewOrderService(repo),
		NatsServer:   newNatsServer(clusterID, clientID),
		ClusterID:    clusterID,
	}
}

func (s *NatsService) CreateOrder(message *stan.Msg) {
	var order *domain.Order

	err := json.Unmarshal(message.Data, &order)
	if err != nil {
		logrus.Errorf("failed to unmarshall order: %v", err)
		return
	}

	o, err := s.OrderService.CreateOrder(order)
	if err != nil {
		logrus.Errorf("failed to create order: %v", err)
		return
	}
	logrus.Infof("created order with id: %s", o.OrderUid)
}

func newNatsServer(clusterID, clientID string) stan.Conn {
	natsServer, err := stan.Connect(clusterID, clientID)
	if err != nil {
		logrus.Fatalf("cannot init natsServer: %v", err)
	}

	return natsServer
}
