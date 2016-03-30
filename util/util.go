package util

import (
	"encoding/json"
	"path/filepath"
	"os"
	"log"
	"io"
	"runtime/debug"
	"reflect"
)


func CheckErr(err error) {
	var rootpath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		var logfile, logfileerr = os.OpenFile(rootpath + "/client.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if logfileerr != nil {
			log.Fatalf("error opening file: %v", logfileerr)
		}
		mWriter := io.MultiWriter(os.Stdout, logfile)
		log.SetOutput(mWriter)

		log.Println(err)
		log.Println(string(debug.Stack()))
		logfile.Close();
		log.SetOutput(os.Stdout)
	}
}

func Jsonp(obj interface{}, callback string) string {
	b, _ := json.Marshal(obj)
	return callback + "(" + string(b) + ")"
}

func IsZero(val interface{}) bool {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && IsZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && IsZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}