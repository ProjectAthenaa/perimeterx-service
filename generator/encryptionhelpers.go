package generator

import (
	"github.com/satori/go.uuid"
	"time"
)

func GenerateTimestamp() int{
	return int(time.Now().UnixNano()/1000000)
}

func GenerateUUID() string{
	return uuid.NewV4().String()
}