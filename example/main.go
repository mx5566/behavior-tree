// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/4/9 20:57
 */

package main

import (
	"fmt"
	"reflect"
	"time"

	entity2 "github.com/mx5566/behavior-tree/example/entity"
)

type TT struct {
	A int
}

func Test(inter interface{}) {
	a := reflect.TypeOf(inter)
	b := reflect.New(a)

	fmt.Printf("%v", b)
}

func main() {

	entity := entity2.NewEntity()
	entity.Init()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		entity.UpdateBehavior(10, 10)
	}
}
