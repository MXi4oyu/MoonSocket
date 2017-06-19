package main

import (
	"fmt"
	"os"
	"net"
	"log"
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

		go MsgHandler(new_conn)
	}
	
}

//处理业务逻辑

func MsgHandler(conn net.Conn)  {

	buf:=make([] byte,1024)

	defer conn.Close()

	for{
		n,err:=conn.Read(buf)

		if err!=nil{

			fmt.Println("connection close")
			return
		}

		fmt.Println("client say:",string(buf[:n]))

		clientIp:=conn.RemoteAddr()

		Log(clientIp)

		conn.Write([] byte("hello:"+clientIp.String()+"\n"))


	}

}