package xredis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/xmopen/golib/pkg/xlogging"
)

type subscribeCallBack func(msg *redis.Message)

// MultiSubScribe listen more redis channel,and run f()
func MultiSubScribe(client *redis.Client, channel []string, f subscribeCallBack) {
	xlog := xlogging.Tag("multi.subscribe")
	xlog.Infof("subScribe redis channel:[%+v]", channel)
	for msg := range client.Subscribe(context.TODO(), channel...).Channel() {
		f(msg)
		xlog.Infof("redis channel:[%+v] value:[%+v]", msg.Channel, msg.Payload)
	}
}
