// a WAF proxy that forwards requests and blocks malicious requests

package main

import (
	// "errors"
	"io/ioutil"
	"fmt"
	"io"
	"net/http"
	"log"
	"strings"
	"os"
)

var forwardUrl string

// implement more blocking functionality here
func block_request(r *http.Request) bool {
	// return true or false based on blocking
	log.Println(r.URL.Path)
	if strings.Contains(r.URL.Path, "EICAR") { // EICAR test string
		return true; // block this request
	}

	return false;
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	if (block_request(r)) {
		w.WriteHeader(http.StatusNotAcceptable) // block with 406
		io.WriteString(w, "Blocked by WAF!\n")
		return
	} else {

		requestUrl := fmt.Sprintf("%s%s", forwardUrl, r.URL.Path)
		req, err := http.NewRequest(r.Method, requestUrl, r.Body)
		if err != nil {
			fmt.Printf("client: could not create request: %s\n", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
		}

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
			
		}

		// pass thru all headers as well
		for k, v := range res.Header {
	        w.Header().Set(k, v[0])
	    }

		w.WriteHeader(res.StatusCode)
		w.Write(resBody)

	}
	
	
}

func main() {
	forwardUrl = os.Args[1]
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":80", nil)
	if (err != nil) {
		log.Fatalf("Err starting http server: %s", err)
	}
}


