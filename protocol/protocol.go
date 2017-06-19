package protocol

import (

	"bytes"
	"encoding/binary"
)

const (
	ConstHeader = "MoonSocket"
	ConstHeaderLength = 10
	ConstDataLength = 4
)


//封包
func Enpack(msg [] byte) [] byte  {

	//使用ConstHeader+msg长度+msg来封装一条数据包
	return append(append([]byte(ConstHeader),IntToBytes(len(msg))...),msg...)
}

//解包

func Depack(buffer [] byte) [] byte  {

	length:=len(buffer)
	data:=make([] byte,32)
	var i int

	for i=0;i<length;i++{

		if length<i+ConstHeaderLength+ConstDataLength{
			//解包完毕
			break
		}

		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader{

			msgLength := BytesToInt(buffer[i+ConstHeaderLength:i+ConstHeaderLength+ConstDataLength])

			if length<i+ConstHeaderLength+ConstDataLength+msgLength{
				break
			}

			data=buffer[i+ConstHeaderLength+ConstDataLength:i+ConstHeaderLength+ConstDataLength+msgLength]
		}
	}

	if i==length{

		return make([] byte,0)
	}

	return  data
}

//整型转换为字节
func IntToBytes(n int) [] byte  {
	data:=int32(n)

	bytesBuffer:=bytes.NewBuffer([] byte{})
	//将data参数里面包含的数据写入到bytesBuffer中
	//
	binary.Write(bytesBuffer,binary.BigEndian,data)

	return bytesBuffer.Bytes()
}

//字节转换成整型
func BytesToInt(b [] byte) int  {

	bytesBuffer:=bytes.NewBuffer(b)

	var data int32
	binary.Read(bytesBuffer,binary.BigEndian,&data)

	return int(data)
}