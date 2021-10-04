package main

import (
	"fmt"
	"net/http"
)

// Contains Fan-In  Pattern
func main()  {
	urls := []string{"www.google.com","www.rafgfnsdsa.com","www.gmail.com","slack.k8s.io","x.yz.com","https://courses.calhoun.io/"}
	status := make(chan string)
	for _,u := range urls{
		go getResponse(u,status)
	}

	var statuses []string
	for range urls {
		select {
		case s:=<-status :
			statuses=append(statuses, s)
		}
	}

	fmt.Println(statuses)


}


func getResponse(url string, s chan string) {

	_,err := http.Get(url)

	if err != nil {
		s <- url + " ok"
		return
	}

	s <- url + " nok"
}
