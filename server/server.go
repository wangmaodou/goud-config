package server

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	LOG  = "[CONFIG_WEB_SERVICE]"
	PORT = ":7339"
)

var (
	VALUE = map[string]string{
		"name":   "maodou",
		"age":    "23",
		"gender": "m",
	}
	param = make(map[string]string)
)

func Start() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/update", postHandler)
	log.Println("Local config server server is starting...")
	log.Fatal(http.ListenAndServe(PORT, nil))
}

// show configs in server page.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	param = getConfigFromCenter()
	view, err := template.New("index").Parse(INDEX)
	checkError(err)
	view.Execute(w, param)
}

// update config immediately.
func postHandler(w http.ResponseWriter, r *http.Request) {
	for k, _ := range param {
		param[k] = r.FormValue(k)
	}
	updateConfigCenter()
	log.Println(LOG, param)
	io.WriteString(w, "ok")
}

func getConfigFromCenter() map[string]string {
	return VALUE
}

func updateConfigCenter() {

}

func check(err error) {
	if err != nil {
		log.Println(LOG, err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(LOG, err)
		os.Exit(1)
	}
}
