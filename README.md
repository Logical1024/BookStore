# BookStore使用说明 
 
一、软件安装 
 
1.浏览器（如chrome浏览器、QQ浏览器）； 
2.MySQL数据库（除了安装MySQL数据库，最好再安装MySQL数据库的可视化工具，如Navicat）； 
3.Golang编译环境（在Go语言中文网https://studygolang.com/ 点击"下载"再选择对应版本下载）； 
4.Golang编辑环境（如GoLand、VSCode）； 
 
二、环境配置 
 
1.打开\bookstore\sql\mysql.sql，根据先后顺序在MySQL数据库命令行窗口或者可视化工具执行里面的SQL语句； 
2.打开\bookstore\sql\books.sql，在MySQL数据库命令行窗口或者可视化工具执行里面的SQL语句； 
3.打开\bookstore\src\utils\db.go，根据电脑MySQL数据库的设置，更改init()函数里的MySQL数据库的连接参数并保存（一般来说，只需更改root账户的用户密码）； 
4.更改GOPATH，可以在\bookstore\src路径，通过命令行窗口执行 export GOPATH=`pwd` 命令；也可以更改环境变量（以win10为例），我的电脑（鼠标右键）-属性-高级系统设置-环境变量-更改GOPATH变量的值为\bookstore\src路径； 
5.编译和运行代码，在命令行窗口，在\bookstore\src路径，执行 go build main.go 命令编译生成新的main.exe，再执行 ./main.exe  命令运行代码（或者鼠标双击main.exe运行代码），也可以执行 go run main.go  命令编译并运行代码（不会生成新的main.exe）；如运行代码后出现网络防火墙警告，则选择允许访问；每次正式使用之前都要先运行代码； 
6.打开浏览器，访问主页http://localhost:8080/main ，开始正式使用； 
 
三、使用事项 
 
1.可通过网页进行用户注册；用户名为admin为管理员，登录后可进行图书商品管理，图书订单发货等操作；其他用户名的用户为会员，登录后可进行购买图书等操纵； 
2.管理员添加图书商品时，需要在\bookstore\src\views\static\img目录下添加以 图书名称.jpg 命名的图书商品的图片； 
 
四、其他事项 
 
1.项目地址https://github.com/Logical1024/BookStore ，转载和使用请注明出处； 
2.后续作者应该还会继续进行更新，请持续关注，有空麻烦给stars； 
3.作者WeChat:freedoms66，欢迎联系作者； 
4.编辑时间：2021.05.15； 
 
