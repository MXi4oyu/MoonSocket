package main

import (
	"fmt"
	"os"
	"net"
	"log"
	"github.com/mxi4oyu/MoonSocket/protocol"
	"time"
)


//定义CheckError方法，避免写太多到 if err!=nil
func CheckError(err error)  {

	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())

		os.Exit(1)
	}

}

//自定义log
func Log(v... interface{})  {

	log.Println(v...)
}

func main()  {

    server_listener,err:=net.Listen("tcp","localhost:8848")

	CheckError(err)

	defer server_listener.Close()

	Log("Waiting for clients connect")



	for{
		new_conn,err:=server_listener.Accept()

		CheckError(err)

		go ServerMsgHandler(new_conn)
	}

	
}

//服务端消息处理
func ServerMsgHandler(conn net.Conn)  {

	//存储被截断的数据
	tmpbuf:=make([] byte,0)
	buf:=make([] byte,1024)

	defer conn.Close()

	//接收解包
	readchan:=make(chan [] byte,16)
	go ReadChan(readchan)

	for{
		//读取客户端发来的消息
		n,err:=conn.Read(buf)

		if err!=nil{

			fmt.Println("connection close")
			return
		}

		//解包
		tmpbuf = protocol.Depack(append(tmpbuf,buf[:n]...))
		fmt.Println("client say:",string(tmpbuf))

		Msg:=tmpbuf

		//向客户端发送消息
		go WriteMsgToClient(conn)

		beatch :=make(chan byte)
		//心跳计时，默认30秒
		go HeartBeat(conn,beatch,30)
		//检测每次Client是否有数据传来
		go HeartChanHandler(Msg,beatch)

	}

}

//处理心跳,根据HeartChanHandler判断Client是否在设定时间内发来信息
func HeartBeat(conn net.Conn,heartChan chan byte,timeout int)  {
	select {
	case hc:=<-heartChan:
		Log("<-heartChan:",string(hc))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout)*time.Second))
		break
	case <-time.After(time.Second*30):
		Log("timeout")
		conn.Close()
	}
}

//服务端向客户端发送消息
func WriteMsgToClient(conn net.Conn)  {

	talk:="wordpress"
	//将信息封包
	smsg:=protocol.Enpack([]byte(talk))
	conn.Write(smsg)
}

//处理心跳channel
func HeartChanHandler( n [] byte,beatch chan byte)  {
	for _,v:=range n{
		beatch<-v
	}
	close(beatch)
}

//从channell中读取数据
func ReadChan(readchan chan [] byte)  {

	for{
		select {
		case data:=<-readchan:
			Log(string(data))
		}
	}
}