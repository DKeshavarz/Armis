package cluster

type State string

const (
	Alive   State = "alive"
	Suspect State = "suspect"
	Failed  State = "deaD"
)