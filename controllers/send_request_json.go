package controllers
import (
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
)

func sendJson(url string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil{
			json.NewEncoder(b).Encode(datajson)
	}
	client := &http.Client{}
	req, err := http.NewRequest(trequest, url, b)
	r, err:= client.Do(req)
  //r, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}


func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
