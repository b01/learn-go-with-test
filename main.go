package main

import (
	dependencyinjection "learn-go-with-test/dependency-injection"
	"net/http"
)

func main()  {
	//dependencyinjection.Greet(os.Stdout, "Khalifah")
	http.ListenAndServe(":5000", http.HandlerFunc(dependencyinjection.MyGreeterHandler))
}
