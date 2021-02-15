/*
@Time : 2021/2/15 7:30 下午
@Author : MuYiMing
@File : descartes_test
@Software: GoLand
*/
package descartes

import (
	"fmt"
	"testing"
)

func TestDesc(t *testing.T) {
	sets := Product([]interface{}{"a", "b", "c"}, []interface{}{1, 2, 3, 4}, []interface{}{"x", "y", "z"})
	for _, set := range sets {
		fmt.Println(set)
	}
}
