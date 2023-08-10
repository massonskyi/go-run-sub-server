package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
}

func run_sub_server() bool {
	cmd := exec.Command("cmd.exe", "/C", "C:/Users/user37/Desktop/server/launch.bat")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func MakeRequest() Response {

	resp, err := http.Get("http://localhost:8001/rand")
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	num, ok := result["number"].(float64)
	if !ok {
		log.Fatalln(err)
	}
	response := Response{
		Message: strconv.FormatFloat(num, 'f', -1, 64),
	}
	return response
}
func OutServer() Response {

	resp, err := http.Get("http://localhost:8001/exc")
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	txt, ok := result["number"].(float64)
	if !ok {
		log.Fatalln(err)
	}
	response := Response{
		Message: strconv.FormatFloat(txt, 'f', -1, 64),
	}
	return response
}
func main() {
	http.HandleFunc("/rand", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MakeRequest())

	})
	http.HandleFunc("/out", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(OutServer())

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if run_sub_server() {
			response := Response{
				Message: "sub server is runnig on port:" + "8001",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			response := Response{
				Message: "sub server crashed with error",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}

	})

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
