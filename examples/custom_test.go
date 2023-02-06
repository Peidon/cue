/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/2/2 2:18 下午
 * @Version: cue
 */

package main

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"
	"cuelang.org/go/encoding/protobuf"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var (
	CTX = cuecontext.New()

	cueFile   = "api_proxy.cue"
	protoFile = "protobuf/item.proto"
)

func TestValue(t *testing.T) {
	pp, err := protobuf.Extract(protoFile, nil, &protobuf.Config{
		Paths: []string{},
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	schema, _ := format.Node(pp, format.Simplify())

	fmt.Println(string(schema))

	fp := filepath.Join("cue_script", cueFile)
	input, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// compile our input
	data := CTX.CompileBytes(input, cue.Filename("input.cue"))
	if data.Err() != nil {
		fmt.Println("Error:", data.Err())
		return
	}

	// mock api return
	dataPath := cue.ParsePath("data")
	filledData := data.FillPath(dataPath, []Item{
		{
			ID:   1,
			Name: "item",
			ItemPrice: Price{
				Value: 100.0,
			},
		},
	})

	cp := cue.ParsePath("flat_response")
	data = filledData.LookupPath(cp)

	b, e := data.MarshalJSON()
	if e != nil {
		fmt.Println("Error:", e)
		return
	}

	fmt.Println(string(b))

}

type Item struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ItemPrice Price  `json:"price"`
}

type Price struct {
	Value float64 `json:"val"`
}
