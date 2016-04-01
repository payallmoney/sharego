package video

import (
	"github.com/go-martini/martini"
//	"strings"
//	"encoding/json"
)

func Router(m martini.Router) {
	m.Any("/upload", videoupload)
	m.Any("/list/:id", clientlist)
	m.Any("/list", videolist)
	m.Any("/reg/:id", reg)
	m.Any("/active/:id", active)
	m.Any("/version/:id", videoversion)
	m.Any("/uploadpage", uploadpage)
	//m.Any("/upload", videouploadpage)
	m.Any("/del/:id", del)
	m.Any("/changename/:id/:name", changename)
	m.Any("/index", videoindex)
	m.Any("/clients", clients)
	m.Any("/client/add/:id", client_add)
	m.Any("/client/videoadd/:id/:videoid", client_video_add)
	m.Any("/client/videochange/:id/:idx/:videoid", client_video_change)
	m.Any("/client/videodel/:id/:idx", client_video_del)
	m.Any("/client/del/:id", client_del)
	m.Any("/program/version", programVersion)
	m.Any("/program/upload", programUpload)
	m.Any("/program/delete/:version", programDel)
	m.Any("/program/list", programList)
	m.Any("/program/reset", reset)
	m.Any("/status/:id", client_status)
}


