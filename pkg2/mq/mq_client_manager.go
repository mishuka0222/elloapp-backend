package kafka

import (
	"github.com/zeromicro/go-zero/core/syncx"
	"io"
)

var (
	clientManager = syncx.NewResourceManager()
)

type client struct {
	*Producer
}

//func (c *client) Close() error {
//	return nil
//}

func GetCachedMQClient(c *KafkaProducerConf) *Producer {
	var (
		val io.Closer
		err error
	)
	val, err = clientManager.GetResource(c.Topic, func() (io.Closer, error) {
		cli := MustKafkaProducer(c)
		return cli, nil
	})
	if err != nil {
		panic(err)
	}

	return val.(*Producer)
}
