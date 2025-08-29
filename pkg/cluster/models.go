package cluster

const (
	PROTOCOL = "http:/"
	JOIN     = "cluster/join/"
	PING     = "cluster/ping/"
)

type JoinResponse struct {
	Msg  string  `json:"message"` 
	Info []*node `json:"info"`
}

type JoinRequest struct {
	Self *node `json:"self"` //binding:"required"
}

type PingResponse struct {
	Msg  string  `json:"message"`
	Info []*node `json:"info"`
}

