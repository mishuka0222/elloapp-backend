package kafka

// KafkaProducerConf kafka producer settings.
type KafkaProducerConf struct {
	Brokers []string
	Topic   string
}

// KafkaConsumerConf kafka client settings.
type KafkaConsumerConf struct {
	Brokers []string
	Topics  []string
	Group   string
}
