package video

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	util "github.com/payallmoney/sharego/util"

	"fmt"
	"path/filepath"
	"os"
	"github.com/satori/go.uuid"
	"strings"
	"io"
	"encoding/json"
)

func Router(m martini.Router){
	m.Any("/upload", videoupload)
	m.Any("/list/:id", videolist)
	m.Any("/version/:id", videoversion)
	m.Any("/uploadpage", uploadpage)
	m.Any("/upload", videouploadpage)
	m.Any("/index", videoindex)
	m.Any("/clients", clients)
}

func clients(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	result := []bson.M{}
	db.C("video_client").Find(bson.M{}).All(&result)
	for _, value := range result {
		list := bson.M{}
		db.C("video_list").Find(bson.M{"id":value["id"]}).One(&list)
		value["list"] = list["videolist"];
	}
	r.JSON(200, result)
}
func videouploadpage(r render.Render, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.HTML(200, "videoupload", nil)
}

func videoindex(r render.Render, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.HTML(200, "video_index", nil)
}

func videolist(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	result := bson.M{}
	db.C("video_list").Find(bson.M{"id": params["id"]}).One(&result)
	r.JSON(200, result)
}
func videoversion(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	result := bson.M{}
	db.C("video_client").Find(bson.M{"id": params["id"]}).One(&result)
	//ret := map[string]interface{}{"version":1,"files":[]string{"/uploadjs/tasklib.js"}}
	r.JSON(200, result)
}

func videoversionpage(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	r.HTML(200, "videoversio", nil)
}

func videoversionupdate(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	result := bson.M{}
	db.C("video_client").Find(bson.M{"id": params["id"]}).One(&result)
	//ret := map[string]interface{}{"version":1,"files":[]string{"/uploadjs/tasklib.js"}}
	r.JSON(200, result)
}


func uploadpage(r render.Render, w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.HTML(200, "upload", nil)
}

func videoupload(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	util.CheckErr(err)
	err = req.ParseMultipartForm(100000)
	util.CheckErr(err)
	//get a ref to the parsed multipart form
	m := req.MultipartForm

	//get the *fileheaders
	fmt.Println(json.Marshal(m))
	files := m.File["file_data"]

	filenames := []string{}
	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, _ := files[i].Open()
		defer file.Close()
		//create destination file making sure the path is writeable.
		ext := filepath.Ext(files[i].Filename)
		newfilename := uuid.NewV4().String() + ext
		if _, err := os.Stat(newfilename); err == nil {
			newfilename = uuid.NewV4().String()+ext
		}
		fmt.Println(dir + "/static/uploadvideo/" + newfilename)
		dst, _ := os.Create(dir + "/static/uploadvideo/" + newfilename)
		filenames = append(filenames, "/uploadvideo/"+newfilename)
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	ret := map[string]string{"urls":strings.Join(filenames, ",")}
	r.JSON(200, ret)
}
