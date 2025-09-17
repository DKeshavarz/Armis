package cluster

import (
	"sync"
	"time"

	"github.com/DKeshavarz/armis/internal/logger"
	"github.com/DKeshavarz/armis/pkg/client"
	"github.com/google/uuid"
)

type Cluster interface {
	ACK() map[string]*node
	JoinReply() map[string]*node
	GetUpdate(nodes map[string]*node)
	Shutdown() error
}

type cluster struct {
	self           *node
	network        map[string]*node
	fanOut         int
	gossipInterval time.Duration
	logger         logger.Logger
	client         client.Client
	shutdownCh     chan any
	mu             sync.RWMutex
}

type node struct {
	Id          string ///binding:"required"
	Address     string
	State       State
	Incarnation int
	SuspectTime time.Time
	DeadTime    time.Time
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
		client:         client.New(),
		shutdownCh:     make(chan any),
	}

	for _, adr := range config.Network {
		cluster.network[adr] = &node{Address: adr}
	}

	go cluster.gossip()

	return cluster
}

func (c *cluster) ACK() map[string]*node {
	return c.selectNodes(c.fanOut)
}

func (c *cluster) JoinReply() map[string]*node {
	return c.network
}

func (c *cluster) Shutdown() error {
	close(c.shutdownCh)
	c.logger.Trace("grasfully shutdown cluster")
	return nil
}
