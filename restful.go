package restful_skel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"restful_skel/config"
	"restful_skel/log"
)

type SimpleResult struct {
	Result bool   `json:"result"`
	Desc   string `json:"description"`
}

func init() {
}

func RESTfulAPIServe() {
	log.Printf("Trying to open RESTful API port(%d).", config.ListeningPort)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", restfulIndex)
	router.HandleFunc("/api/hello/{cid}", restfulHello)
	log.Printf("APIs are ready.")
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(config.ListeningPort), router)
		log.Fatal(err)
		finish(config.EXITPORT)
	}()
}

func restfulIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("index")
}

func restfulHello(w http.ResponseWriter, r *http.Request) {
	var outBuffer bytes.Buffer
	var result = SimpleResult{Result: false}

	defer func() {
		outBuffer.WriteString(result.Desc)
		outBuffer.WriteString("\n")
		text := outBuffer.String()
		prettyText := strings.Replace(text, "\n", "\n\t", -1)
		log.Println(prettyText)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}()

	outBuffer.WriteString("Process API: hello\n")
	vars := mux.Vars(r)
	cid := vars["cid"]
	outBuffer.WriteString(fmt.Sprintf("sub symbol is %s\n", cid))
}
