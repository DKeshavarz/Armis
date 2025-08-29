package cluster

import (
	"net/http"
	"time"

	"github.com/DKeshavarz/armis/internal/logger"
	"github.com/google/uuid"
)

type Cluster interface {
	ACK() []*node
	JoinReply() []*node
}

type cluster struct {
	self           *node
	network        []*node
	fanOut         int
	gossipInterval time.Duration
	logger         logger.Logger
	client         http.Client
}

type node struct {
	Id          string ///binding:"required"
	Address     string
	State       State
	Incarnation int
}

func New(config Congig) Cluster {
	self := &node{
		Id:          uuid.NewString(),
		Address:     config.Self,
		State:       Alive,
		Incarnation: 0,
	}

	cluster := &cluster{
		self:           self,
		network:        make([]*node, 0),
		fanOut:         config.FanOut,
		gossipInterval: time.Duration(config.GossipInterval) * time.Second,
		logger:         logger.New("cluster-package"),
		client:         http.Client{},
	}

	for _, adr := range config.Network {
		cluster.network = append(cluster.network, &node{Address: adr})
	}

	go cluster.gossip()
	return cluster
}


func (c *cluster) ACK() []*node {
	return c.selectNodes()
}

func (c *cluster) JoinReply() []*node {
	return c.network
}






