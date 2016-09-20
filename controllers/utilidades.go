package controllers
import (
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
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
func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
    if a.Location() != b.Location() {
        b = b.In(a.Location())
    }
    if a.After(b) {
        a, b = b, a
    }
    y1, M1, d1 := a.Date()
    y2, M2, d2 := b.Date()

    h1, m1, s1 := a.Clock()
    h2, m2, s2 := b.Clock()

    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)
    hour = int(h2 - h1)
    min = int(m2 - m1)
    sec = int(s2 - s1)

    // Normalize negative values
    if sec < 0 {
        sec += 60
        min--
    }
    if min < 0 {
        min += 60
        hour--
    }
    if hour < 0 {
        hour += 24
        day--
    }
    if day < 0 {
        // days in month:
        t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
        day += 32 - t.Day()
        month--
    }
    if month < 0 {
        month += 12
        year--
    }

    return
}
