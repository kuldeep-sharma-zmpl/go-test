
type redis interface {
    Set(key string, value inteface{}) int
    Get(param string) int
}

func MyFunc(r redis) {
	r.Set(somevalue)
}

import "goredis"
type redis struct {
    
}

