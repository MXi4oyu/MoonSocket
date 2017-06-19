package main

import (
	"fmt"
	"os"
	"net"
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

	conn.Write([] byte("hello server!"))

	buf:=make([] byte,1024)

	n,_:=conn.Read(buf)

	fmt.Print(string(buf[:n]))

}