package dao

import (
	"model"
	"utils"
)

//AddOrderItem 向数据库中插入订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	sql := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}

//GetOrderItemsByOrderID 根据订单号获取该订单的所有订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	sql := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id = ?"
	rows, err := utils.Db.Query(sql, orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
