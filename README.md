------

## 在原作者的版本上修改

$ go get github.com/gin-gonic/gin


$ go get github.com/go-sql-driver/mysql

$ go get github.com/go-xorm/xorm

mysql 创建数据库 test
建表
create table person(
id INT UNSIGNED AUTO_INCREMENT,
first_name VARCHAR(40) NOT NULL,
last_name VARCHAR(40) NOT NULL,
PRIMARY KEY(id)
);



vscode中打开  
0. 配置 mysql.go 中的数据库信息
1. 执行 env.bat (将当前工程目录路径加入gopath)
2. 工程目录内执行 
   go build main   go install main
3. 要用 vscode debug 的话，修改 .vscode 中的 settings.json 的go.gopath 为本机的对应地址

比如:
{
    "go.gopath": "C:\\Users\\allen\\go;D:\\gin-learn"
}


依次用 postman 进行以下测试:

get 127.0.0.1:8000/

post 127.0.0.1:8000/person/add?first_name=du&last_name=lili
  
get 127.0.0.1:8000/person/list   

get 127.0.0.1:8000/person/get/2

put 127.0.0.1:8000/person/update/1?first_name=kkk&last_name=ffff

delete 127.0.0.1:8000/person/delete/1

## 原作者的参考教程
https://www.jianshu.com/p/a3f63b5da74c
