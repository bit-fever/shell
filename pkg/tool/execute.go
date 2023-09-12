//=============================================================================
/*
Copyright © 2023 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package tool

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//=============================================================================

func DoGet(url string, output any) {
	client := createClient()
	res, err := client.Get(url)
	buildResponse(res, err, &output)
}

//=============================================================================

func DoPut(url string, params any, output any) {
	client := createClient()

	body, err := json.Marshal(&params)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(body)
	res, err := client.Post(url, "Application/json", reader)
	buildResponse(res, err, &output)
}

//=============================================================================

func createClient() *http.Client {
	cert, err := os.ReadFile("config/ca.crt")
	if err != nil {
		log.Fatalf("Could not open certificate file: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)

	certificate, err := tls.LoadX509KeyPair("config/client.crt", "config/client.key")
	if err != nil {
		log.Fatalf("Could not load certificate: %v", err)
	}

	return &http.Client{
		Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certificate},
			},
		},
	}
}

//=============================================================================

func buildResponse(res *http.Response, err error, output any) {
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if res.StatusCode >= 400 {
		log.Fatalf("Error from the server: %v", res.Status)
	}

	// Read the response body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		log.Fatalf("Bad JSON response from server:\n%v", err)
	}
}

//=============================================================================
