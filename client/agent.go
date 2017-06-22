package main

import (
	"fmt"
	"os"
	"net"
	"strconv"
	"time"
	"github.com/mxi4oyu/MoonSocket/protocol"
)

//定义CheckError方法，避免写太多到 if err!=nil
func CheckError(err error)  {

	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())

		os.Exit(1)
	}

}

func main()  {

	if len(os.Args) !=2 {

		fmt.Fprintf(os.Stderr,"Usage:%s IP:Port\n",os.Args[0])

		os.Exit(1)
	}

	//动态传入服务端IP和端口号
	service:=os.Args[1]

	tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)

	CheckError(err)

	conn,err:=net.DialTCP("tcp",nil,tcpAddr)

	CheckError(err)
	ch:=make(chan int,100)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for{
		select {
		case <-ticker.C:
			ch<-1
			go ClientMsgHandler(conn,ch)
		case <-time.After(time.Second*10):
			defer conn.Close()
			fmt.Println("timeout")
		}
	}

}

//客户端消息处理
func ClientMsgHandler(conn net.Conn,ch chan int)  {

	<-ch
	//获取当前时间
	msg:=time.Now().String()
	SendMsg(conn,msg)
}

func GetSession() string{
	gs1:=time.Now().Unix()
	gs2:=strconv.FormatInt(gs1,10)
	return gs2
}

func SendMsg(conn net.Conn,msg string)  {

session:=GetSession()

	words := "{\"Session\":"+session +",\"Meta\":\"Monitor\",\"Message\":\""+msg+"\"}"
	conn.Write([] byte(words))
	protocol.Enpack([]byte(words))
	conn.Write(protocol.Enpack([]byte(words)))

}