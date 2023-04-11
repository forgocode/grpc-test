package monitor

import (
)

func Start() {
	InitServiceRoot()
	go NewTimer()
}
