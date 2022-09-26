//=============================================================================
/*
Copyright © 2022 Andrea Carboni andrea.carboni71@gmail.com

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

package exchange

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//=============================================================================

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all exchanges",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		listExchanges()
	},
}

//=============================================================================

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//=============================================================================

type Exchange struct {
	Code     string  `json:"code,omitempty"`
	Name     string  `json:"name"`
	MakerFee float64 `json:"maker_fee"`
	TakerFee float64 `json:"taker_fee"`
	ApiKey   string  `json:"api_key"`
	Secret   string  `json:"secret"`
	Test     bool    `json:"test"`
}

//=============================================================================

func listExchanges() {
	cert, err := ioutil.ReadFile("config/ca.crt")
	if err != nil {
		log.Fatalf("Could not open certificate file: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)

	certificate, err := tls.LoadX509KeyPair("config/client.crt", "config/client.key")
	if err != nil {
		log.Fatalf("Could not load certificate: %v", err)
	}

	client := &http.Client{
		Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certificate},
			},
		},
	}

	url := "https://bitfever-server:8443/api/exchange/v1/exchanges"

	res, err := client.Get(url)
	if err != nil {
		log.Fatalf("Error making get request: %v", err)
	}

	// Read the response body
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response: %v", err)
	}

	// Print the response body to stdout
	//fmt.Printf("%s\n", body)
	data_obj := []Exchange{}
	jsonErr := json.Unmarshal(body, &data_obj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for _, exc := range data_obj {
		fmt.Println("Code:", exc.Code, ", name:", exc.Name)
	}
}

//=============================================================================
