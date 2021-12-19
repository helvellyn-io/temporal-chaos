package redis

import "testing"

func TestRedisConn(t *testing.T) {

	_, err := RedisClient()

	if err != nil {
		t.Errorf("Connection to Redis failed")
	}

}
