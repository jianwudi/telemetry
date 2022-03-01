package collect

import (
	"encoding/json"
	"fmt"
)

const (
	RedisPubChannel = "sensorpath.notifications"
)

type Notification struct {
	SensorPath string `json:"sensorpath"`
	Interval   uint32 `json:"interval"`
}

func Notify(sensorpath string, internal uint32) {
	r := &RedisCluster{IsSwitch: true}
	message, _ := json.Marshal(Notification{sensorpath, internal})
	client := r.singleton()
	if err := client.Publish(RedisPubChannel, string(message)).Err(); err != nil {
		//		log.Errorf("Error trying to set value: %s", err.Error())
		fmt.Printf("Error trying to set value: %s", err.Error())
	}
}
