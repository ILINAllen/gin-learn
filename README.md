------

## 在原作者的版本上修改

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

  
   



## 原作者的参考教程
https://www.jianshu.com/p/a3f63b5da74c
