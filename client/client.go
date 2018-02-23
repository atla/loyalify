package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"loyalify/common"
	"net/http"
	"unsafe"
)

var serverURL = "http://localhost:8000/api"

func registerClientForProgram(id string) {

	//resp, err := http.Get(serverURL + "/programs/" + id)
	resp, err := http.Get("http://localhost:8000/api/programs/1")

	if err != nil {
		log.Fatal("http get fatal error", err)
	}

	program := common.Program{}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	s := *(*string)(unsafe.Pointer(&body))

	err = json.Unmarshal([]byte(s), program)

	log.Println("Response:")
	program.Print()

}

// Run runs the client
func Run() {
	registerClientForProgram("1")
}
