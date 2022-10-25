package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

// up memcached container using : docker run --name my-memcache -p 11211:11211 -d memcached memcached -m 64
func main() {
	mc := memcache.New("127.0.0.1:11211")

	err := mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	if err != nil {
		panic(err)
		return
	}

	it, err := mc.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println("get: ", string(it.Value))

	err = mc.Delete("foo")
	if err != nil {
		panic(err)
	}
}
