package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/mailru/easyjson"

	"github.com/kzmake/namec"
)

var (
	accessKeyID     string
	secretAccessKey string
	service         string
	region          string
)

func init() {
	flag.StringVar(&accessKeyID, "a", os.Getenv("ACCESS_KEY_ID"), "access key id")
	flag.StringVar(&secretAccessKey, "s", os.Getenv("SECRET_ACCESS_KEY"), "secret access key")
	flag.StringVar(&service, "service", "", "service id")
	flag.StringVar(&region, "region", "", "region name")
}

func main() {
	flag.Parse()

	if !terminal.IsTerminal(0) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			r := namec.Request{}
			easyjson.Unmarshal([]byte(scanner.Text()), &r)

			req, err := http.NewRequest(r.Method, r.URL, strings.NewReader(r.Body))
			if err != nil {
				panic(err)
			}

			for k, vs := range r.Header {
				for _, v := range vs {
					req.Header.Add(k, v)
				}
			}

			_, err = namec.AWSSignV4(req, accessKeyID, secretAccessKey, service, region)
			if err != nil {
				panic(err)
			}
			if r.Header == nil {
				r.Header = map[string][]string{}
			}

			for k, v := range req.Header {
				r.Header[k] = v
			}

			b, err := easyjson.Marshal(r)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(b))
		}
	}
}
