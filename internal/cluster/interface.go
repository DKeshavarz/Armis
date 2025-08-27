package cluster

import "github.com/google/uuid"

type Cluster interface {
	ACK() []*node
	JoinReply() []*node
}

type cluster struct {
	self           *node
	network        []*node
	fanOut         int
	gossipInterval int
}

type node struct {
	id          string
	address     string
	state       State
	Incarnation int
}

func New(address string, network []string, fanOut, gossipInterval int) Cluster {
	self := &node{
		id:          uuid.NewString(),
		address:     address,
		state:       Alive,
		Incarnation: 0,
	}

	cluster := &cluster{
		self:           self,
		network:        make([]*node, 0),
		fanOut:         fanOut,
		gossipInterval: gossipInterval,
	}

	return cluster
}

// ***********************************
func (c *cluster) ACK() []*node {
	return c.selectNodes()
}

func (c *cluster) JoinReply() []*node {
	return c.network
}

// ********************************

// *********** helpers ******************

func (c *cluster) selectNodes() []*node {
	return c.network
}
