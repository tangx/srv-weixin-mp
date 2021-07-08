package confredis

import (
	"fmt"
	"testing"
)

func Test_REdis(t *testing.T) {
	r := Redis{
		Host:     "127.0.0.1",
		Password: "password123",
		DB:       11,
	}

	r.Init()

	_ = r.Set(1, 2000)

	ret, _ := r.GetInt(1)
	fmt.Printf("%v", ret)

	ret2, _ := r.GetString(1)
	fmt.Printf("%v", ret2)
	println(ret2)

}
