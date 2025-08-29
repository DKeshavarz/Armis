package cluster

import (
	"fmt"
)

func (c *cluster) gossip() {
	//TODO: join an create network
	c.join()
	// TODO: send message
	// TODO:grasfully shutdoin

}


func (c *cluster) join() {
	var resp JoinResponse
	newNetwork := make([]*node, 0, len(c.network))
	for _, ip := range c.network {
		if ip.Address == c.self.Address {
			continue
		}
		
		url := fmt.Sprintf("%s/%s/%s", PROTOCOL, ip.Address, JOIN)
		err := c.Post(1, url, JoinRequest{Self: c.self}, &resp)

		if err == nil { // other node okay
			newNetwork = append(newNetwork, resp.Info...)
			break
		}
	}
	newNetwork = append(newNetwork, c.self)
	c.network = newNetwork
}

func (c *cluster) selectNodes() []*node {
	return c.network
}

// func ping(url string) []*node {

// 	resp, err := http.Get(url)
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