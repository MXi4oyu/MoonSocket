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

	msg:="测试自定义协议"
	SendMsg(conn,msg)

}

func GetSession() string{
	gs1:=time.Now().Unix()
	gs2:=strconv.FormatInt(gs1,10)
	return gs2
}

func SendMsg(conn net.Conn,msg string)  {

	for i:=0;i<100;i++{
		session:=GetSession()

		words := "{\"ID\":"+ strconv.Itoa(i) +"\",\"Session\":"+session +",\"Meta\":\"golang\",\"Message\":\""+msg+"\"}"
		conn.Write([] byte(words))
		protocol.Enpack([]byte(words))
		conn.Write(protocol.Enpack([]byte(words)))
	}

	fmt.Println("send over")

	defer conn.Close()
}