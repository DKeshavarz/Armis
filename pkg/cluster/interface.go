package cluster

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

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
	url := "http://localhost:8000/"
	target := "cluster/ping"
	ticker := time.NewTicker(time.Duration(c.gossipInterval) * time.Second)
	time.Sleep(30 * time.Second)
	log.Println("start gossip")
	for range 5 {
		<-ticker.C
		ping(url + target)
		
	}

}

type response struct {
	Status string  `json:"status"`
	Info   []*node `json:"info"`
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
	var r response
	json.Unmarshal(body, &r)
	log.Println("get data:" , r.Status , " \n and \n", r.Info)
	defer resp.Body.Close()
	return r.Info
}

// *********** helpers ******************

func (c *cluster) selectNodes() []*node {
	return c.network
}
