package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file,err:=os.OpenFile("xx.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC,0666)
	if err!=nil{
		fmt.Println("open file faild err:",err)
		return
	}
	defer file.Close()
	// str:="hello 文件写入222："
	// file.Write([]byte(str))
	// file.WriteString("你好，文件第二种写入2222")
	writer:=bufio.NewWriter(file)
	// for i:=1;i<10;i++{
		writer.WriteString("hello,world\n")
	// }
	writer.Flush()

}