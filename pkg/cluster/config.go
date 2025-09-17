package cluster

type Congig struct {
	Self           string
	Network        []string
	FanOut         int
	GossipInterval int
}
