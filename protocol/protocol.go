package protocol

import (

	"bytes"
	"encoding/binary"
)

const (
	ConstHeader = "Headers"
	ConstHeaderLength = 7
	ConstMLength = 4
)

//封包

func Enpack(msg [] byte) [] byte  {

	return append(append([]byte(ConstHeader),IntToBytes(len(msg))...),msg...)

}

//解包

func Depack(buffer [] byte) [] byte  {

	length:=len(buffer)
	data:=make([] byte,32)
	var i int

	for i=0;i<length;i++{

		if length<i+ConstHeaderLength+ConstMLength{

			break
		}

		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader{

			msgLength := BytesToInt(buffer[i+ConstHeaderLength:i+ConstHeaderLength+ConstMLength])

			if length<i+ConstHeaderLength+ConstMLength+msgLength{
				break
			}

			data=buffer[i+ConstHeaderLength+ConstMLength:i+ConstHeaderLength+ConstMLength+msgLength]
		}
	}

	if i==length{

		return make([] byte,0)
	}

	return  data
}

//整型转换为字节
func IntToBytes(n int) [] byte  {
	x:=int32(n)

	bytesBuffer:=bytes.NewBuffer([] byte{})

	binary.Write(bytesBuffer,binary.BigEndian,x)

	return bytesBuffer.Bytes()
}

//字节转换成整型
func BytesToInt(b [] byte) int  {

	bytesBuffer:=bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer,binary.BigEndian,&x)

	return int(x)
}