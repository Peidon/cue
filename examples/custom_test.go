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
	"testing"
)

var(
    CTX = cuecontext.New()

	uniqueList = `
 
import "list"

#elem: {
	key: string
	val: _
}

elems: [...#elem] & [
	{key: "a", val: 1},
	{key: "b", val: 1},
	{key: "a", val: 2},
	{key: "b", val: 1},
]

// we compare the current element and add it if it does not appear in the remainder of the list
// in doing so, we add the last unique occurance to the result
uniq: [ for i, x in elems if !list.Contains(list.Drop(elems, i+1), x) {x}]`
)

func TestValue(t *testing.T) {
	fmt.Println("unique list by val")

	// compile our input
	value := CTX.CompileString(uniqueList, cue.Filename("input.cue"))
	if value.Err() != nil {
		fmt.Println("Error:", value.Err())
		return
	}

	b, e := value.MarshalJSON()
	if e != nil {
		fmt.Println("Error:", e)
		return
	}

	fmt.Println(string(b))

}