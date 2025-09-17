package cluster

const (
	PROTOCOL = "http:/"
	JOIN     = "cluster/join/"
	PING     = "cluster/ping/"
)

type JoinResponse struct {
	Msg  string           `json:"message"`
	Info map[string]*node `json:"info"`
}

type JoinRequest struct {
	Self map[string]*node `json:"self"` //binding:"required"
}

type PingResponse struct {
	Msg  string           `json:"message"`
	Info map[string]*node `json:"info"`
}
