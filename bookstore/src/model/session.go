package model

//Session 结构
type Session struct {
	SessionID string   //Session编号
	UserName  string   //用户名
	UserID    int      //用户编号
	Cart      *Cart    //购物车
	OrderID   string   //订单号
	Orders    []*Order //订单
}
