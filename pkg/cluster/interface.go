package cluster

import (
	"net/http"
	"time"

	"github.com/DKeshavarz/armis/internal/logger"
	"github.com/google/uuid"
)

type Cluster interface {
	ACK() map[string]*node
	JoinReply() map[string]*node
}

type cluster struct {
	self           *node
	network        map[string]*node
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
		network:        make(map[string]*node),
		fanOut:         config.FanOut,
		gossipInterval: time.Duration(config.GossipInterval) * time.Second,
		logger:         logger.New("cluster-package"),
		client:         http.Client{},
	}

	for _, adr := range config.Network {
		cluster.network[adr] = &node{Address: adr}
	}

	go cluster.gossip()
	return cluster
}


func (c *cluster) ACK() map[string]*node {
	return c.selectNodes()
}

func (c *cluster) JoinReply() map[string]*node {
	return c.network
}






