package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ChildCount     = 3
	ConsumerPeriod = 3
)

func main() {
	channelNotify()
	fmt.Println("main process exit!")
}

// 不同粒度的运行单位之间的调度通常借用系统底层的信号。上下文切换导致粒度对效率的影响是巨大的，用户态的协程是成本较低的。
// 各种语言都可以有，没有的后期自行扩展都行，无非是看runtime（库）的实现程度。golang的协程从语言设计开始作为一类实体，用法以及推广都会比库友好一点。
// chanel作为调度来使用时是一种语法糖级别的封装，但对可读性和可维护性提升巨大
func channelNotify() {
	var ops uint64
	var wg sync.WaitGroup

	messages := make(chan int, ChildCount*ConsumerPeriod)
	// 带缓冲的通道，放满消费品
	for i := 0; i < ConsumerPeriod*ChildCount; i++ {
		messages <- i
	}
	defer close(messages)

	// 设置多个消费者
	chans := make([]chan bool, ChildCount)
	for i := 0; i < ChildCount; i++ {
		chans[i] = make(chan bool)
	}

	for i := 0; i < ChildCount; i++ {
		wg.Add(1)
		child := chans[i]
		go func() {
			ticker := time.NewTicker(1 * time.Microsecond)
			for _ = range ticker.C {
				select {
				case <-child:
					fmt.Println("child process interrupt...")
					wg.Done()
					return
				default:
					if len(messages) > 0 {
						fmt.Printf("fetch message: %d\n", <-messages)
						atomic.AddUint64(&ops, 1)
						fmt.Println("ops:", ops)
					}
				}
			}
		}()
	}

	// 多等1秒，消费完所有物品
	time.Sleep((1) * time.Second)
	for i := 0; i < ChildCount; i++ {
		close(chans[i])
	}
	wg.Wait()
	fmt.Println("final ops:", ops)
}
