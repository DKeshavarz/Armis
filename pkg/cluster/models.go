package cluster

const (
	PROTOCOL = "http:/"
	JOIN     = "cluster/join/"
	PING     = "cluster/ping/"
)

type JoinResponse struct {
	Msg  string  `json:"status"`
	Info []*node `json:"info"`
}

type JoinRequest struct {
	Self node `json:"self"`
}

type PingResponse struct {
	Msg  string  `json:"status"`
	Info []*node `json:"info"`
}

