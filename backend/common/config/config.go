package config

type KqPusherConf struct {
	Brokers []string
	Topic   string
}

type SparkModelConf struct {
	Url         string
	APIPassword string
	MaxTokens   int64
	TopK        int64
	Temperature float64
	Model       string
}
