package dao

import (
	"model"
	"utils"
	"net/http"
)

//AddSession 向数据库中添加Session
func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的Session
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSession 根据session的Id值从数据库中查询Session
func GetSession(sessID string) (*model.Session, error) {
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}

	row := inStmt.QueryRow(sessID)
	sess := &model.Session{}
	//扫描数据库中的字段值为Session的字段赋值
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}

//IsLogin 判断用户是否已经登录 false 没有登录 true 已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根据Cookie的name获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取Cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的Session
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			//已经登录
			return true, session
		}
	}
	//没有登录
	return false, nil
}

//GetSameSession 删除同样的Session
func GetSameSession(UserID int) {
	sqlStr := "select session_id from sessions where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, UserID)
	SessionID := ""
	row.Scan(&SessionID)

	if SessionID != "" {
		DeleteSession(SessionID)
	}

	return
}