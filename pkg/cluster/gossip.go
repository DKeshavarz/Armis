package cluster

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/DKeshavarz/armis/internal/logger"
)

func (c *cluster) gossip() {
	c.join()
	// TODO: send message
	ticker := time.NewTicker(c.gossipInterval)

	for {
		select {
		case <-ticker.C:
			c.ping()
		case <-c.shutdownCh:
			c.logger.Debug("hello i am shutdown")
			return
		}
	}
	// TODO:grasfully shutdoin

}

func (c *cluster) join() {
	c.mu.Lock()
	defer c.mu.Unlock()

	var resp JoinResponse
	tmpMap := make(map[string]*node)
	for _, ip := range c.network {
		if ip.Address == c.self.Address {
			continue
		}

		url := fmt.Sprintf("%s/%s/%s", PROTOCOL, ip.Address, JOIN)
		err := c.client.Post(1, url, JoinRequest{Self: map[string]*node{c.self.Address: c.self}}, &resp)

		if err == nil && resp.Info != nil {
			tmpMap = resp.Info
			break
		}
	}
	c.logger.Debug("check map", logger.Field{Key: "map", Value: tmpMap})
	c.network = tmpMap
	c.network[c.self.Address] = c.self
}

func (c *cluster) GetUpdate(nodes map[string]*node) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//TODO: add chanel
	//TODO: save for multy thread
	for adr, node := range nodes {
		if _, ok := c.network[adr]; !ok {
			c.network[adr] = node
			continue
		}

		if !node.isValid() || *c.network[adr] == *node || c.network[adr].Incarnation > node.Incarnation { //WTF ???? works ok but WTF
			continue
		}

		if c.network[adr].Incarnation < node.Incarnation {
			c.network[adr] = node
		} else if c.network[adr].State < node.State { // equal Incarnation
			c.network[adr] = node
		}
	}
}

// ****************** helpers ************************
func (c *cluster) selectNodes(nodeCnt int) map[string]*node {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if nodeCnt > len(c.network) {
		return c.network
	}

	selected := make(map[string]*node)
	keys := make([]string, 0, len(c.network))
	for k := range c.network {
		keys = append(keys, k)
	}
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	for i := 0; i < nodeCnt; i++ {
		key := keys[i]
		selected[key] = c.network[key]
	}

	return selected
}

func (n *node) isValid() bool {
	return n.Id != "" && n.Address != ""
}

func (c *cluster) ping() {
	nodes := c.selectNodes(c.fanOut)
	for adr := range nodes {
		if adr == c.self.Address {
			continue
		}
		url := fmt.Sprintf("%s/%s/%s", PROTOCOL, adr, PING)

		go func(url, adr string) {
			var resp PingResponse
			err := c.client.Get(1, url, &resp)
			if err != nil {
				c.logger.Error("catch error in calling api", logger.Field{
					Key:   "error",
					Value: err,
				})
				c.findSuspect(adr)
				return
			}

			c.GetUpdate(resp.Info)
		}(url, adr)

	}
}

func (c *cluster) findSuspect(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	n, ok := c.network[key]
	if !ok {
		return
	}

	if n.State == Alive {
		n.State = Suspect
		n.SuspectTime = time.Now()
	}
}
