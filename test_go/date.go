package main

import (
     "fmt"
     "time"
     //"strings"

)

func main(){
  var a, b time.Time
  a = time.Date(2016, 02, 1, 0, 0, 0, 0, time.UTC)
  b = time.Date(2016, 06, 16, 1, 1, 1, 1, time.UTC)
  y,m,d,_,_,_ := diff(a,b)
  var meses_contrato float64
  meses_contrato = (float64(y*12))+float64(m)+(float64(d)/30)
  fmt.Println(float64(d)/30) // Expected: 1 1 1 1 1 1
  fmt.Println(m)
    fmt.Println(y)
    fmt.Println(meses_contrato)
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



    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)


    // Normalize negative values

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
