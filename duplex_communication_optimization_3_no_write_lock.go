package main

import (
	"encoding/binary"
	"fmt"
	_ "io"
	"net"
	"sync"
	"time"
)

var zRecvCount = uint32(0) // 张大爷听到了多少句话
var lRecvCount = uint32(0) // 李大爷听到了多少句话
var total = uint32(100000) // 总共需要遇见多少次

var z0 = "吃了没，您吶?"
var z3 = "嗨！吃饱了溜溜弯儿。"
var z5 = "回头去给老太太请安！"
var l1 = "刚吃。"
var l2 = "您这，嘛去？"
var l4 = "有空家里坐坐啊。"

type RequestResponse struct {
	Serial  uint32 // 序号
	Payload string // 内容
}

// 发送缓冲
func writeLoop(conn *net.TCPConn, writeChanelQueue chan []byte) {
	for b := range writeChanelQueue {
		_, err := conn.Write(b)
		// fmt.Println("write to socket data", string(b[8:]))
		if err != nil {
			fmt.Println("write err:", err)
			return
		}
	}
}

// 序列化RequestResponse，并发送
// 序列化后的结构如下：
//   长度  4字节
//   Serial 4字节
//   PayLoad 变长
func writeTo(r *RequestResponse, conn *net.TCPConn, writeChanelQueue chan []byte) {
	payloadBytes := []byte(r.Payload)
	length := uint32(len(payloadBytes) + 4)
	packageBytes := make([]byte, length+4)

	binary.BigEndian.PutUint32(packageBytes, length)
	binary.BigEndian.PutUint32(packageBytes[4:8], r.Serial)
	copy(packageBytes[8:], payloadBytes)

	writeChanelQueue <- packageBytes
}

// 接收数据，反序列化成RequestResponse
func readFrom(conn *net.TCPConn, recvBuf []byte, recvIndex int) ([]*RequestResponse, int, error) {
	retResponses := make([]*RequestResponse, 0)
	// fmt.Println("recv buf length", recvIndex, "recv buf cap", cap(recvBuf))
	if recvIndex != 0 {
		// fmt.Println("leagecy data before read", recvBuf[:recvIndex])
	}
	n, err := conn.Read(recvBuf[recvIndex:cap(recvBuf)])
	// fmt.Println("read length", n)
	if err != nil {
		return nil, n + recvIndex, fmt.Errorf("读数据故障：%s", err.Error())
	}

	index := 0
	for {
		ret := &RequestResponse{}
		if index+8 <= n+recvIndex {
			length := int(binary.BigEndian.Uint32(recvBuf[index : index+4]))
			// fmt.Println("parse length", length)
			ret.Serial = binary.BigEndian.Uint32(recvBuf[index+4 : index+8])
			// fmt.Println("parse serial", ret.Serial)

			if index+8+length-4 <= n+recvIndex {
				ret.Payload = string(recvBuf[index+8 : index+8+length-4])
				// fmt.Println("parse payload", ret.Payload)
				index += length + 4
				retResponses = append(retResponses, ret)
			} else {
				// fmt.Println("uncomplete payload", string(recvBuf[index+8:n+recvIndex]))
				break
			}
		} else {
			break
		}
	}

	if n+recvIndex-index > 0 {
		// fmt.Println("parse uncomplete package", n+recvIndex-index)
		copy(recvBuf[:n+recvIndex-index], recvBuf[index:n+recvIndex])
		// fmt.Println("uncomplete leagecy date", recvBuf[index:n+recvIndex])
		// fmt.Println("move leagecy date", recvBuf[:n+recvIndex-index])
	}

	// fmt.Println("recvbuf next index", n+recvIndex-index)
	return retResponses, n + recvIndex - index, nil
}

// 张大爷的耳朵
func zhangDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup, writeChanelQueue chan []byte) {
	defer wg.Done()
	recvBuf := make([]byte, 0, 1024)
	recvIndex := 0
	for zRecvCount < total*3 {
		responses, nextRecvIndex, err := readFrom(conn, recvBuf, recvIndex)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// fmt.Println("zhangDaYe listen and read response count: ", len(responses))

		recvIndex = nextRecvIndex

		for _, r := range responses {
			// fmt.Println("张大爷收到：" + r.Payload)
			if r.Payload == l2 { // 如果收到：您这，嘛去？
				go writeTo(&RequestResponse{r.Serial, z3}, conn, writeChanelQueue) // 回复：嗨！吃饱了溜溜弯儿。
			} else if r.Payload == l4 { // 如果收到：有空家里坐坐啊。
				go writeTo(&RequestResponse{r.Serial, z5}, conn, writeChanelQueue) // 回复：回头去给老太太请安！
			} else if r.Payload == l1 { // 如果收到：刚吃。
				// 不用回复
			} else {
				fmt.Println("张大爷听不懂：" + r.Payload)
				break
			}
			zRecvCount++
		}
	}
}

// 张大爷的嘴
func zhangDaYeSay(conn *net.TCPConn, writeChanelQueue chan []byte) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeTo(&RequestResponse{nextSerial, z0}, conn, writeChanelQueue)
		nextSerial++
	}
}

// 李大爷的耳朵，实现是和张大爷类似的
func liDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup, writeChanelQueue chan []byte) {
	defer wg.Done()
	recvBuf := make([]byte, 0, 1024)
	recvIndex := 0
	for lRecvCount < total*3 {
		// fmt.Println("liDaYe listen and read")
		responses, nextRecvIndex, err := readFrom(conn, recvBuf, recvIndex)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		recvIndex = nextRecvIndex
		// fmt.Println("liDaYeListen: ", len(responses))
		for _, r := range responses {
			// fmt.Println("李大爷收到：" + r.Payload)
			if r.Payload == z0 { // 如果收到：吃了没，您吶?
				writeTo(&RequestResponse{r.Serial, l1}, conn, writeChanelQueue) // 回复：刚吃。
			} else if r.Payload == z3 {
				// do nothing
			} else if r.Payload == z5 {
				// do nothing
			} else {
				fmt.Println("李大爷听不懂：" + r.Payload)
				break
			}
			lRecvCount++
		}
	}
}

// 李大爷的嘴
func liDaYeSay(conn *net.TCPConn, writeChanelQueue chan []byte) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeTo(&RequestResponse{nextSerial, l2}, conn, writeChanelQueue)
		nextSerial++
		writeTo(&RequestResponse{nextSerial, l4}, conn, writeChanelQueue)
		nextSerial++
	}
}

func startServer(wg *sync.WaitGroup) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("张大爷在胡同口等着 ...")
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("碰见一个李大爷:" + conn.RemoteAddr().String())

		zhangdayeWriteChanel := make(chan []byte, 100)
		go writeLoop(conn, zhangdayeWriteChanel)

		go zhangDaYeListen(conn, wg, zhangdayeWriteChanel)
		go zhangDaYeSay(conn, zhangdayeWriteChanel)
	}

}

func startClient(wg *sync.WaitGroup) *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	lidayeWriteChanel := make(chan []byte, 100)
	go writeLoop(conn, lidayeWriteChanel)

	go liDaYeListen(conn, wg, lidayeWriteChanel)
	go liDaYeSay(conn, lidayeWriteChanel)
	return conn
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go startServer(&wg)
	time.Sleep(time.Second)
	conn := startClient(&wg)
	t1 := time.Now()
	wg.Wait()
	elapsed := time.Since(t1)
	conn.Close()
	fmt.Println("耗时: ", elapsed)
}

