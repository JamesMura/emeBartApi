package main

import (
	"encoding/xml"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/parnurzeal/gorequest"
	"github.com/unrolled/render"
)

type RouteInfo struct {
	Uri      string  `xml:"uri" json:"uri"`
	SchedNum int     `xml:"sched_num" json:"sched_num"`
	Routes   []Route `xml:"routes>route" json:"routes"`
	Message  string  `xml:"message" json:"message,omitempty"`
}

type Route struct {
	Name         string   `xml:"name" json:"name"`
	Abbr         string   `xml:"abbr" json:"abbr"`
	RouteID      string   `xml:"routeID" json:"routeID"`
	Number       int      `xml:"number" json:"number"`
	Origin       string   `xml:"origin" json:"origin"`
	Destination  string   `xml:"destination" json:"destination"`
	Direction    string   `xml:"direction" json:"direction"`
	Color        string   `xml:"color" json:"color"`
	Holidays     int      `xml:"holidays" json:"holidays"`
	StationCount int      `xml:"num_stns" json:"num_stns"`
	Stations     []string `xml:"config>station" json:"stations"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := render.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/route.aspx", func(w http.ResponseWriter, req *http.Request) {
		_, body, _ := gorequest.New().Get("http://api.bart.gov/api/route.aspx?cmd=routes&key=MW9S-E7SL-26DU-VV8V").EndBytes()
		routeInfo := RouteInfo{}
		err := xml.Unmarshal(body, &routeInfo)
		checkError(w, err)
		r.JSON(w, http.StatusOK, routeInfo)
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":" + port)
}

func checkError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
