package main

import (
    "io"
    "fmt"
    "encoding/json"
    "html/template"
    "net/http"
    "log"
    "time"
    "bytes"
    "math/rand"
)

// BackendDNS - DNS to connect to backend
var BackendDNS=getEnv("BACKEND_DNS", "localhost")
// BackendPort - Port to connect to backend
var BackendPort=getEnv("BACKEND_PORT", "9000")

type fortune struct {
	ID      string `json:"id" redis:"id"`
	Message string `json:"message" redis:"message"`
}

type newFortune struct {
    Message string `json:"message"`
}

// use a custom client, because we don't do blocking operations wihout timeouts
var myClient = &http.Client{Timeout: 10 * time.Second}

// HealthzHandler - Handler for requests to /healthz
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, "healthy")
}

func main() {

    rand.Seed(time.Now().UnixNano())

    http.HandleFunc("/healthz", HealthzHandler)

    http.HandleFunc("/api/random", func (w http.ResponseWriter, r *http.Request) {
        resp, err := myClient.Get(fmt.Sprintf("http://%s:%s/fortunes/random", BackendDNS, BackendPort))
        if err != nil {
            log.Fatalln(err)
            fmt.Fprint(w, err)
            return
        }

        f := new(fortune)
        json.NewDecoder(resp.Body).Decode(f)

        fmt.Fprint(w, f.Message)
        return
    })

    http.HandleFunc("/api/all", func (w http.ResponseWriter, r *http.Request) {
        resp, err := myClient.Get(fmt.Sprintf("http://%s:%s/fortunes", BackendDNS, BackendPort))
        if err != nil {
            log.Fatalln(err)
            fmt.Fprint(w, err)
            return
        }

        fortunes := new([]fortune)
        json.NewDecoder(resp.Body).Decode(fortunes)

        tmpl, err := template.ParseFiles("./templates/fortunes.html")

        if err != nil {
            log.Fatalln(err)
            fmt.Fprint(w, err)
            return
        }

        tmpl.Execute(w, fortunes)
        return
    })

    http.HandleFunc("/api/add", func (w http.ResponseWriter, r *http.Request) {

        if r.Method != "POST" {
            http.Error(w, "Use POST", http.StatusMethodNotAllowed)
            return
        }

        f := new(newFortune)
        json.NewDecoder(r.Body).Decode(f)

        var postURL = fmt.Sprintf("http://%s:%s/fortunes", BackendDNS, BackendPort)
        var jsonStr = []byte(fmt.Sprintf(`{"id": "%d", "message": "%s"}`, rand.Intn(10000), f.Message))

        _, err := myClient.Post(postURL, "application/json", bytes.NewBuffer(jsonStr))
        if err != nil {
            log.Fatalln(err)
            fmt.Fprint(w, err)
            return
        }

        fmt.Fprint(w, "Cookie added!")

        return
    })

    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.ListenAndServe(":8080", nil)
}
