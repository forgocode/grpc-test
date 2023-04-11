package util

import (
	"serverMonitor/pkg/typed"
	"strconv"
)

func GenerateGrpcClientStr(endpoint typed.Endpoint) string {
	return endpoint.IP + ":" + strconv.Itoa(int(endpoint.Port))
}
