package tasks

import "sync"

type Method struct {
	method interface{}
	parameterNum int
}

type Factory struct {
	wait *sync.WaitGroup
	methods []Method
}




