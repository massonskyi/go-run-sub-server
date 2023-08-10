package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
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

/*
func get_number_from_sub_api(response string) int {

}
*/
func main() {
	destURL, err := url.Parse("http://localhost:8001")
	if err != nil {
		log.Fatal("Ошибка при разборе URL назначения:", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(destURL)

	http.HandleFunc("/rand", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Переадресация запроса:", r.URL.Path)
		proxy.ServeHTTP(w, r)
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