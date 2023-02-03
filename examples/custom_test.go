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
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var(
    CTX = cuecontext.New()

	filename = "api_proxy.cue"
)

func TestValue(t *testing.T) {
	fmt.Println("unique list by val")

	fp := filepath.Join("cue_script", filename)
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

	b, e := data.MarshalJSON()
	if e != nil {
		fmt.Println("Error:", e)
		return
	}

	fmt.Println(string(b))

}