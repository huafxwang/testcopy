package main

import (
	"fmt"
	"io"
	"os"

)

func main(){
	file,err:=os.Open("./wenjian.go")
	if err != nil {
		fmt.Println("open file is failed",err)
		return 
	}
	defer file.Close()

//使用Read方法读取数据
var tmp []byte
tmp=make([]byte, 256)

	n,err:= file.Read(tmp)
	if err==io.EOF{
		fmt.Println("读取完毕")
	
	
	}
	if err != nil{
		fmt.Println("读取错误",err)
		return
	}
	fmt.Printf("已读取%d字节数据\n",n)
	fmt.Println(string(tmp[:n]))

}
