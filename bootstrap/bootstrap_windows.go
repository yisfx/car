package bootstrap

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadBody(body io.Reader) []byte {
	b, _ := ioutil.ReadAll(body)
	return b
}
func MustJSONDecode(b []byte, i interface{}) {
	err := json.Unmarshal(b, i)
	if err != nil {
		panic(err)
	}
}

type ActionRequest struct {
	Action string
	Type   string
}
type ActionResponse struct {
	Result string `json"result"`
	Gear   int    `json"Gear"`
}

func Run() {

	fmt.Println("Starting")
	http.HandleFunc("/", index)
	http.HandleFunc("/bootstrap.min.css", css)
	http.HandleFunc("/jquery-3.5.1.min.js", js)
	http.HandleFunc("/action", action)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ending")
}

var gear int = 0

func action(res http.ResponseWriter, req *http.Request) {

	request := ActionRequest{Action: "", Type: ""}
	MustJSONDecode([]byte(string(ReadBody(req.Body))), &request)

	fmt.Println(request.Type, request.Action)

	gear = gear + 1
	res.Header().Set("content-type", "application/json; charset=utf-8")

	b, err := json.Marshal(ActionResponse{Result: "ok", Gear: gear})
	if err != nil {
		panic(err)
	}
	res.Write(b)
}

func js(res http.ResponseWriter, _ *http.Request) {
	content, err := ioutil.ReadFile("./jquery-3.5.1.min.js")
	if err != nil {
		res.Write([]byte("err"))
		return
	}
	res.Header().Set("content-type", "application/javascript")
	res.Write(content)
}

func css(res http.ResponseWriter, _ *http.Request) {
	content, err := ioutil.ReadFile("./bootstrap.min.css")
	if err != nil {
		res.Write([]byte("err"))
		return
	}
	res.Header().Set("content-type", "text/css")
	res.Write(content)
}
func index(res http.ResponseWriter, _ *http.Request) {
	content, err := ioutil.ReadFile("./index.html")
	if err != nil {
		res.Write([]byte("err"))
		return
	}
	res.Write(content)

}
