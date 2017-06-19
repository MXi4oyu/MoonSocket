package protocol

import (

	"bytes"
	"encoding/binary"
)

const (
	MsgHeader = "MoonSocket"
	HeaderLength = 10
	DataLength = 4
)


//封包
//封包信息由 header + 信息长度 ＋ 信息内容组成
func Enpack(msg [] byte) [] byte  {

	return append(append([]byte(MsgHeader),IntToBytes(len(msg))...),msg...)
}

//解包

func Depack(buffer [] byte) [] byte  {

	length:=len(buffer)
	data:=make([] byte,64)
	var i int

	for i=0;i<length;i++{

		if length<i+HeaderLength+DataLength{
			//解包完毕
			break
		}

		//如果解析到头部，则解析包信息到data
		if string(buffer[i:i+HeaderLength]) == MsgHeader{

			//将msg的长度转换为int
			msgLength := BytesToInt(buffer[i+HeaderLength:i+HeaderLength+DataLength])

			if length<i+HeaderLength+DataLength+msgLength{
				//解包完毕
				break
			}

			//解析包信息到data
			data=buffer[i+HeaderLength+DataLength:i+HeaderLength+DataLength+msgLength]
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