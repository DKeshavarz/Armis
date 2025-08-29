package cluster

import (
	"fmt"

	"github.com/DKeshavarz/armis/internal/logger"
)

func (c *cluster) gossip() {
	//TODO: join an create network
	c.join()
	// TODO: send message
	// TODO:grasfully shutdoin

}


func (c *cluster) join() {
	var resp JoinResponse
	tmpMap := make(map[string]*node)
	for _, ip := range c.network {
		if ip.Address == c.self.Address {
			continue
		}
		
		url := fmt.Sprintf("%s/%s/%s", PROTOCOL, ip.Address, JOIN)
		err := c.Post(1, url, JoinRequest{Self: map[string]*node{c.self.Address:c.self}}, &resp)

		if err == nil { // other node okay
			tmpMap = resp.Info
			break
		}
	}
	c.logger.Debug("check map", logger.Field{Key:"map", Value: tmpMap})
	c.network = tmpMap
	c.network[c.self.Address] = c.self
}

func (c *cluster) GetUpdate(nodes map[string]*node){
	//TODO: add chanel
	//TODO: save for multy thread
	for adr, node := range nodes{
		c.logger.Debug("panic happend", 
		logger.Field{Key: "adr", Value: adr}, 
		logger.Field{Key: "node", Value: node}, 
		logger.Field{Key: "netwrot", Value: c.network})
		if _, ok := c.network[adr]; !ok {
			c.network[adr] = node
			continue
		}

		if !node.isValid() || *c.network[adr] == *node || c.network[adr].Incarnation > node.Incarnation{ //WTF ???? works ok but WTF
			continue
		}

		c.logger.Debug("do you seee hereeeeee ??")
		if c.network[adr].Incarnation < node.Incarnation {
			c.network[adr] = node
		}else if c.network[adr].State < node.State{ // equal Incarnation
			c.network[adr] = node
		}
	}
}
// ****************** helpers ************************
func (c *cluster) selectNodes() map[string]*node {
	return c.network
}

func (n *node) isValid()bool{
	return  n.Id != "" && n.Address != ""
}

// func ping(url string) []*node {

// 	resp, ervr := http.Get(url)
// 	if err != nil {
// 		log.Fatalf("Error making GET request: %v", err)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading response body: %v", err)
// 	}
// 	var r JoinResponse
// 	json.Unmarshal(body, &r)
// 	log.Println("get data:", r.Msg, " \n and \n", r.Info)
// 	defer resp.Body.Close()
// 	return r.Info
// }