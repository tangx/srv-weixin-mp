package confredis

import (
	"testing"
)

func Test_REdis(t *testing.T) {
	r := Redis{
		Host:     "127.0.0.1",
		Password: "password123",
		DB:       11,
		prefix:   "hahahah",
	}

	r.Init()

}
