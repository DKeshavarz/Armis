package cluster

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	Id          string
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

	for _, val := range config.Network {
		cluster.network = append(cluster.network, &node{Address: val})
	}

	go cluster.gossip()
	return cluster
}

// ***********************************
func (c *cluster) ACK() []*node {
	return c.selectNodes()
}

func (c *cluster) JoinReply() []*node {
	return c.network
}

// **********************************
func (c *cluster) gossip() {
	//TODO: join an create network
	c.join()
	// TODO: send message
	// TODO:grasfully shutdoin

}


func (c *cluster) join() {
	var resp JoinResponse
	for _, ip := range c.network {
		if ip.Address == c.self.Address {
			continue
		}
		
		url := fmt.Sprintf("%s/%s/%s", PROTOCOL, ip.Address, JOIN)
		c.logger.Info(url)

		err := c.Post(1, url, JoinRequest{Self: c.self}, &resp)

		c.logger.Info("", logger.Field{Key:"err", Value:err}, 
		logger.Field{Key:"msg", Value:resp})
	}
}



func ping(url string) []*node {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	var r JoinResponse
	json.Unmarshal(body, &r)
	log.Println("get data:", r.Msg, " \n and \n", r.Info)
	defer resp.Body.Close()
	return r.Info
}

// *********** helpers ******************

func (c *cluster) selectNodes() []*node {
	return c.network
}


