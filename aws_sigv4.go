package namec

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// AWSSignV4 ...
func AWSSignV4(req *http.Request, accessKey, secretKey, service, region string) (http.Header, error) {
	signer := v4.Signer{
		Credentials: aws.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			"",
		),
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return signer.Sign(context.Background(), req, bytes.NewReader(body), service, region, time.Now())
}
