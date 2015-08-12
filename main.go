package main

import (
	"github.com/go-martini/martini"
	"flag"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"github.com/larspensjo/config"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"sharego/auth"
	util "sharego/util"
)

func main() {
	m := getMartini()
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("goodshare_session", store))
	m.Use(render.Renderer())
	//配置文件
	configFile := flag.String("configfile", "config.ini", "配置文件")
	inicfg, _ := config.ReadDefault(*configFile)
	m.Map(inicfg)
	//数据库
	db := util.GetDB(inicfg)
	m.Map(db)
	//缓存
	cache := make(map[string]interface{})
	m.Map(cache)
	m.Any("/login", auth.Login)
	m.Get("/logout", auth.Logout)
	m.Get("/", index)
	//静态内容
//	m.Use(martini.Static("static"))
	//需要权限的内容
//	m.Group("/admin", admin.Router, auth.Auth)
	m.Run();
//	m.RunOnAddr(":3333")
}

func index(db *mgo.Database, r render.Render, req *http.Request, inicfg *config.Config) {
	ret := make(map[string]interface{})
	r.HTML(200, "index", ret)
}

func getMartini() *martini.ClassicMartini {
	//与martini.Classic相比,去掉了martini.Logger() handler  去掉讨厌的http请求日志
	base := martini.New()
	router := martini.NewRouter()
	base.Use(martini.Recovery())
	base.Use(martini.Static("static"))
	base.MapTo(router, (*martini.Routes)(nil))
	base.Action(router.Handle)
//	return &martini.ClassicMartini{base, router}
	return martini.Classic()
}
