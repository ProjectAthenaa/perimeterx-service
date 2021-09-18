package generator

import (
	"time"
)

func GenerateTimestamp() int{
	return int(time.Now().UnixNano()/1000000)
}