package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/mailru/easyjson"

	"github.com/kzmake/freeza"
)

func main() {
	a := os.Getenv("ACCESS_KEY_ID")
	s := os.Getenv("SECRET_ACCESS_KEY")

	if !terminal.IsTerminal(0) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			r := freeza.Request{}
			easyjson.Unmarshal([]byte(scanner.Text()), &r)

			req, err := http.NewRequest(r.Method, r.URL, strings.NewReader(r.Body))
			if err != nil {
				panic(err)
			}

			for k, v := range r.Headers {
				req.Header.Set(k, v)
			}

			_, err = freeza.AWSSignV4(req, a, s)
			if err != nil {
				panic(err)
			}
			if r.Headers == nil {
				r.Headers = map[string]string{}
			}

			r.Headers["Authorization"] = req.Header.Get("Authorization")
			r.Headers["x-amz-content-sha256"] = req.Header.Get("x-amz-content-sha256")
			r.Headers["x-amz-date"] = req.Header.Get("x-amz-date")

			b, err := easyjson.Marshal(r)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(b))
		}
	}
}
