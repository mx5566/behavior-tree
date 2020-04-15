// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  test_builder
 * @Version: 1.0.0
 * @Date: 2020/4/11 12:32
 */

package main

import (
	"fmt"
	"time"

	"github.com/mx5566/behavior-tree/example/action"
	"github.com/mx5566/behavior-tree/example/condition"
	"github.com/mx5566/behavior-tree/example/decorate"

	btree "github.com/mx5566/behavior-tree"
	"github.com/mx5566/behavior-tree/common"

	"github.com/mx5566/behavior-tree/impl/composite"
)

func main() {
	testBT()
}

//测试行为树
func testBT() {
	builder := btree.NewBehaviorTreeBuilder("testBT")

	tt := composite.NewSelectorImpl("select1")

	tree := builder.AddBehavior(tt).
		AddBehavior(composite.NewSequenceImpl("sequence")).
		AddBehavior(condition.NewConditionIsSeeEnemy(false, "isseeenemy")).
		Back().
		AddBehavior(composite.NewSelectorImpl("select")).
		AddBehavior(composite.NewSequenceImpl("sequence")).
		AddBehavior(condition.NewConditionIsHeathLow(false, "isheathloww")).
		Back().
		AddBehavior(action.NewActionRunaway("actionrunaway")).
		Back().
		Back().
		AddBehavior(composite.NewParallelImpl(common.RequireAll, common.RequireOne, "paralle")).
		AddBehavior(condition.NewConditionIsEnemyDead(true, "isenemydead")).
		Back().
		AddBehavior(decorate.NewDecorateRepate("repeate", 3)).
		AddBehavior(action.NewActionAttack("attack")).
		Back().
		Back().
		Back().
		Back().
		Back().
		AddBehavior(action.NewActionPatrol("patrol")).
		End()

	//模拟执行行为树
	if tree != nil {
		fmt.Println("tree is not nil")
	}
	for i := 0; i < 1; i++ {
		tree.Tick()
	}

	time.Sleep(time.Millisecond * time.Duration(10000))
}
