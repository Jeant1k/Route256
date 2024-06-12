package main

import "sync"

func main() {
	mx := sync.Mutex{}
	mx.Lock() // блокировка и ожидание пока кто-то его не разблокирует
	mx.Lock()

	mx.Unlock()
}
