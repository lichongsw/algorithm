package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var config map[string]string
var lock sync.Mutex
var flag atomic.Bool
var count atomic.Int32

func main() {
	count, _ := strconv.Atoi(os.Args[1])
	for x := 0; x < count; x++ {
		go getConfig3_2()
	}
	<-time.After(time.Second)

	fmt.Println("init config")
}

func getConfig1_1() map[string]string {
	lock.Lock()
	defer lock.Unlock()
	if config == nil {
		fmt.Println("init config")
		config = map[string]string{}
		return config
	}
	return config
}

func getConfig1_2() map[string]string {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("init config")
		config = map[string]string{}
		return config
	}
	return config
}

func getConfig2_1() map[string]string {
	if !flag.Load() {
		fmt.Println("init config")
		config = map[string]string{}
		flag.Store(true)
		return config
	}
	return config
}

func getConfig2_2() map[string]string {
	if !flag.Load() {
		if config == nil {
			fmt.Println("init config")
			config = map[string]string{}
			flag.Store(true)
			return config
		}
	}
	return config
}

func getConfig3_1() map[string]string {
	if !flag.Load() {
		if config == nil {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("init config")
			config = map[string]string{}
			flag.Store(true)
			return config
		}
	}
	return config
}

func getConfig3_2() map[string]string {
	if !flag.Load() {
		lock.Lock()
		defer lock.Unlock()
		count.Add(1)
		config = map[string]string{}
		flag.Store(true)
		return config
	}
	return config
}
