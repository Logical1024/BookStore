package model

//Book 书
type Book struct {
	ID      int     //书号
	Title   string  //书名
	Author  string  //作者
	Price   float64 //价格
	Sales   int     //销量
	Stock   int     //库存
	ImgPath string  //图片路径
}
