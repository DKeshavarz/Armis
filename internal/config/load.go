package config

func New() (*Config, error) {
	config := Config{}

	config.Cluster.Self = GetEnv("CLUSTER_SELF", "localhost") + ":" + GetEnv("PORT", "8080")
	config.Cluster.Network = GetEnvAsSlice("CLUSTER_NETWORK", nil, ",")
	config.Cluster.FanOut = GetEnvAsInt("CLUSTER_FANOUT", 3)
	config.Cluster.GossipInterval = GetEnvAsInt("CLUSTER_GOSSIP_INTERVAL", 2)
	return &config, nil
}