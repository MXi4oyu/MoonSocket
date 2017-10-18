package main

import (
	"fmt"
	"os"
	"net"
	"strconv"
	"time"
	"github.com/MXi4oyu/MoonSocket/protocol"
	"errors"
)

//定义CheckError方法，避免写太多到 if err!=nil
func CheckError(err error)  {

	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())

		os.Exit(1)
	}

}

//解决断线重连问题
func doWork(conn net.Conn) error {
	ch:=make(chan int,100)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for{
		select {
		case stat:=<-ch:
			if stat==2{
				return errors.New("None Msg")
			}
		case <-ticker.C:
			ch<-1
			go ClientMsgHandler(conn,ch)

		case <-time.After(time.Second*10):
			defer conn.Close()
			fmt.Println("timeout")
		}

	}

	return nil
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

	for{

		conn,err:=net.DialTCP("tcp",nil,tcpAddr)

		if err!=nil{
			fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
		}else{
			defer conn.Close()
			doWork(conn)
		}

		time.Sleep(3 * time.Second)

	}

}

//客户端消息处理
func ClientMsgHandler(conn net.Conn,ch chan int)  {

	<-ch
	//获取当前时间
	msg:=time.Now().String()
	go SendMsg(conn,msg)
	go ReadMsg(conn,ch)

}

func GetSession() string{
	gs1:=time.Now().Unix()
	gs2:=strconv.FormatInt(gs1,10)
	return gs2
}

//接收服务端发来的消息
func ReadMsg(conn net.Conn,ch chan int)  {

	//存储被截断的数据
	tmpbuf:=make([] byte,0)
	buf:=make([] byte,1024)

	//将信息解包
	n,_:=conn.Read(buf)
	tmpbuf = protocol.Depack(append(tmpbuf,buf[:n]...))
	msg:=string(tmpbuf)
	fmt.Println("server say:",msg)
	if len(msg)==0{
		//服务端无返回信息
		ch<-2
	}
}

//向服务端发送消息
func SendMsg(conn net.Conn,msg string)  {

session:=GetSession()

	words := "{\"Session\":"+session +",\"Meta\":\"Monitor\",\"Message\":\""+msg+"\"}"
	//将信息封包
	smsg:=protocol.Enpack([]byte(words))
	conn.Write(smsg)

}