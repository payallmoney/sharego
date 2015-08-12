package auth

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	 util "sharego/util"
)

type JsonRet struct {
	Type string
	Status int
	Msg  string
	Data interface{}
}

func Login(session sessions.Session, db *mgo.Database, r render.Render, req *http.Request ,writer http.ResponseWriter) string {
	writer.Header().Set("Content-Type", "text/javascript")
	userid := req.FormValue("userid")
	callback := req.FormValue("callback")
	password := req.FormValue("password")
	fmt.Println("userid",userid)
	if userid == "" {
		return util.Jsonp(JsonRet{"login",401,"请登录", "abc"},callback)
	} else {
		result := bson.M{}
		err := db.C("auth_user").Find(bson.M{"id": userid}).One(&result)
		fmt.Println(password)
		fmt.Println(result)

		if err == nil {
			values := result
			if values["password"] == password {
				session.Set("userid", values["loginname"])
				session.Set("username", values["username"])
				fmt.Println("登录成功!")
				return util.Jsonp( JsonRet{"login",200,"登录成功", values["username"]},callback)

			}else {
				return util.Jsonp( JsonRet{"login",401,"登录失败!密码错误!", nil},callback)
			}
		}else {
			return util.Jsonp( JsonRet{"login",401,"登录失败!用户名错误!", nil},callback)
		}
	}
}

func Logout(session sessions.Session, r render.Render) {
	session.Delete("userid")
	r.HTML(200, "login", "登出成功")
}

func Auth(session sessions.Session, c martini.Context, r render.Render) {
	v := session.Get("userid")
	if v == nil {
		r.Redirect("/login")
	}else {
		c.Next();
	}
}

