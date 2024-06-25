package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Switch struct {
	State bool
}

var switchState = Switch{State: false}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/switch", switchHandler)
	http.HandleFunc("/on", onHandler)
	http.HandleFunc("/off", offHandler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// indexHandler handles the HTTP request for the index page.
//
// Parameters:
//   w: http.ResponseWriter - Response writer to write the response back to the client.
//   r: *http.Request - HTTP request received from the client.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("index").Parse(`
		<html>
			<body>
				<h1>Switch {{if .State}}is on{{else}}is off{{end}}</h1>
				<form action="/on" method="post">
					<input type="submit" value="Turn On">
				</form>
				<form action="/off" method="post">
					<input type="submit" value="Turn Off">
				</form>
			</body>
		</html>
	`)
	t.Execute(w, switchState)
}

func switchHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func onHandler(w http.ResponseWriter, r *http.Request) {
	switchState.State = true
	http.Redirect(w, r, "/", http.StatusFound)
}

func offHandler(w http.ResponseWriter, r *http.Request) {
	switchState.State = false
	http.Redirect(w, r, "/", http.StatusFound)
}

// run with go run A3_webserver2.go
// compile to webserver2 with 
// go build -o webserver2 A3_webserver2.go