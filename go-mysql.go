package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

// go连接Mysql示例
var db *sql.DB // 是一个连接对象
func initDB() (err error) {
	//数据库信息
	//用户名：密码@tcp(ip:端口)/数据库的名字
	dsn := "root:root@tcp(127.0.0.1:3306)/goday10"
	//连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(100)
	return

}

//查询
func queryone() {

}

//查询多条
func queryMore(n int){
	//1.SQL语句
	sqlStr:=`select id,name,age from user where id >?;`
	//2.执行
	rows,err:=db.Query(sqlStr,n)
	if err!=nil{
		fmt.Printf("exec %s query failed,err:%v\n",sqlStr,err)
		return	
	}
	//3.一定要关闭rows
	defer rows.Close()
	//4.循环取值
	for rows.Next(){
		var u1 user
		err:=rows.Scan(&u1.id,&u1.name,&u1.age)
		if err!=nil{
			fmt.Printf("scan failed, err:%v\n",err)
		}
		fmt.Printf("u1.%#v\n",u1)
	}
}
//插入
func insert() {
	//1.写SQL语句
	sqlStr:=`insert into user(name,age) values("张一",20)`
	//2.exec
	ret,err:=db.Exec(sqlStr)
	if err!=nil{
		fmt.Printf("insert into failed,err:%v\n",err)
		return
	}
//如果是插入数据的操作，能够拿到数据的id
id,err:=ret.LastInsertId()
if err!=nil{
	fmt.Printf("get id failed,err:%v\n",err)
	return
}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功")
	var u1 user
	//1.写查询单条记录的sql语句
	sqlstr := `select id,name,age from user01 where id=2;`
	//2.执行
	rowObj := db.QueryRow(sqlstr)
	//3.拿到结果 
	rowObj.Scan(&u1.id, &u1.name, &u1.age) //必须对rowObj对象调用Scan方法，因为该方法会释放数据库链接
	//4.打印结果 
	fmt.Printf("u1:%#v\n",u1)

	queryMore(2)
	insert()
}
