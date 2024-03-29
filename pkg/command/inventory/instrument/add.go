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

package instrument

import (
	"fmt"
	"github.com/bit-fever/shell/pkg/model"
	"github.com/bit-fever/shell/pkg/tool"
	"github.com/spf13/cobra"
)

//=============================================================================

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new inventory",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		addInstrument()
	},
}

//=============================================================================

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//=============================================================================

func addInstrument() {
	url := "https://bitfever-server:8443/api/inventory/v1/instruments"

	data := model.Instrument{
		Symbol:        "@TS",
		Name:          "Test",
		MarketType:    "GA",
		SecurityType:  "AA",
		ExchangeId:    1,
		DatasourceId:  1,
		CurrencyId:    1,
		PriceScale:    1,
		MinMovement:   2,
		BigPointValue: 3,
	}

	tool.DoPut(url, &data, &data)

	fmt.Println("Instrument added with id=", data.Id)
}

//=============================================================================
