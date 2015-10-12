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
//	"strings"
	"reflect"
	"io"
	"encoding/json"
)

func Router(m martini.Router){
	m.Any("/upload", videoupload)
	m.Any("/list/:id", clientlist)
	m.Any("/list", videolist)
	m.Any("/version/:id", videoversion)
	m.Any("/uploadpage", uploadpage)
	m.Any("/upload", videouploadpage)
	m.Any("/del/:id", del)
	m.Any("/changename/:id/:name", changename)
	m.Any("/index", videoindex)
	m.Any("/clients", clients)
	m.Any("/client/add/:id", client_add)
	m.Any("/client/videoadd/:id/:videoid", client_video_add)
	m.Any("/client/videochange/:id/:idx/:videoid", client_video_change)
	m.Any("/client/videodel/:id/:idx", client_video_del)
	m.Any("/client/del/:id", client_del)
}

func clients(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	result := []bson.M{}
	db.C("video_client").Find(bson.M{}).Sort("_id").All(&result)
	for _, value := range result {
		list := bson.M{}
		db.C("video_list").Find(bson.M{"_id":value["id"]}).One(&list)
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

func clientlist(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	result := bson.M{}
	db.C("video_client").Find(bson.M{"_id": params["id"]}).One(&result)
	tmp := result["videolist"];
	fmt.Println(reflect.TypeOf(result));
	fmt.Println(reflect.TypeOf(tmp));
	fmt.Println([]string(result["videolist"]));
//	bson.Mash
//	fmt.Println(len(tmp));
//	ary := *(*string)(unsafe.Pointer(&tmp))
//	for  i ,v:= range *tmp{
//		fmt.Println(i,v)
//		//arry[i] = "/uploadvideo/"+arry[i]
//	}
	r.JSON(200, tmp)
}

func videolist(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	result := []bson.M{}
	db.C("video_list").Find(bson.M{}).All(&result)
	r.JSON(200,result)
}
func videoversion(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	result := bson.M{}
	db.C("video_client").Find(bson.M{"_id": params["id"]}).One(&result)
	//ret := map[string]interface{}{"version":1,"files":[]string{"/uploadjs/tasklib.js"}}
	r.JSON(200, result)
}

func videoversionpage(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	r.HTML(200, "videoversio", nil)
}

func videoversionupdate(r render.Render, db *mgo.Database, params martini.Params, req *http.Request, w http.ResponseWriter) {
	result := bson.M{}
	db.C("video_client").Find(bson.M{"_id": params["id"]}).One(&result)
	//ret := map[string]interface{}{"version":1,"files":[]string{"/uploadjs/tasklib.js"}}
	r.JSON(200, result)
}


func uploadpage(r render.Render, w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.HTML(200, "upload", nil)
}

func videoupload(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter,db *mgo.Database,) {
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
		newname := uuid.NewV4().String()
		newfilename := newname + ext
		if _, err := os.Stat(newfilename); err == nil {
			newfilename = uuid.NewV4().String()+ext
		}
		fmt.Println(dir + "/static/uploadvideo/" + newfilename)
		dst, _ := os.Create(dir + "/static/uploadvideo/" + newfilename)
		filenames = append(filenames, newfilename)
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		db.C("video_list").Insert(bson.M{"_id": newfilename,"name":newfilename,"src":"/uploadvideo/"+newfilename})
	}
	r.JSON(200, filenames)
}

func del(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter,db *mgo.Database,) {
	db.C("video_list").Remove(bson.M{"_id": params["id"]});
	r.JSON(200, nil)
}

func changename(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter,db *mgo.Database) {
	db.C("video_list").Update(bson.M{"_id": params["id"]},bson.M{"$set":bson.M{"name":params["name"]}})
}

func client_add(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter,db *mgo.Database) {
	db.C("video_client").Insert(bson.M{"_id": params["id"]})
}
func client_del(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter,db *mgo.Database) {
	db.C("video_client").Remove(bson.M{"_id": params["id"]})
}

func client_video_add(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter,db *mgo.Database) {
	result := bson.M{}
	db.C("video_list").Find(bson.M{"_id": params["videoid"]}).One(&result);
	db.C("video_client").Update(bson.M{"_id": params["id"]},bson.M{"$push":bson.M{"videolist":bson.M{"_id":params["videoid"],"src":result["src"]}}})
}
func client_video_del(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter,db *mgo.Database) {
//	db.C("video_client").Remove(bson.M{"_id": params["id"] ,"videolist":params["idx"]});
	db.C("video_client").Update(bson.M{"_id": params["id"]} , bson.M{"$unset" : bson.M{"videolist."+params["idx"] : 1 }});
	db.C("video_client").Update(bson.M{"_id": params["id"]} , bson.M{"$pull" : bson.M{"videolist" : nil}});
}

func client_video_change(r render.Render, params martini.Params, req *http.Request, w http.ResponseWriter,db *mgo.Database) {
	result := bson.M{}
	db.C("video_list").Find(bson.M{"_id": params["videoid"]}).One(&result);
	db.C("video_client").Update(bson.M{"_id": params["id"]},bson.M{"$set":bson.M{"videolist."+params["idx"]:bson.M{"videolist":bson.M{"_id":params["videoid"],"src":result["src"]}}})
}

