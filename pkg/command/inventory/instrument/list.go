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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all instruments",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		listInstruments()
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

func listInstruments() {
	url := "https://bitfever-server:8443/api/inventory/v1/instruments"

	// Print the response body to stdout
	//fmt.Printf("%s\n", body)
	data_obj := []model.Exchange{}
	tool.DoGet(url, &data_obj)

	for _, exc := range data_obj {
		fmt.Println("Code:", exc.Code, ", name:", exc.Name)
	}
}

//=============================================================================