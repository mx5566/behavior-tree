// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  behavior_factory
 * @Version: 1.0.0
 * @Date: 2020/4/11 22:48
 */

package entity

import (
	"fmt"
	"sync"

	"github.com/mx5566/behavior-tree/example/action"
	"github.com/mx5566/behavior-tree/impl/decorate"

	btree "github.com/mx5566/behavior-tree"
)

var factoty *BehaviorFactory
var onceF sync.Once

// 行为树工厂
type BehaviorFactory struct {
	bTrees map[string]*btree.BehaviorTree
}

func GetInstanceBehaviorFactory() *BehaviorFactory {
	onceF.Do(func() {
		factoty = &BehaviorFactory{
			bTrees: map[string]*btree.BehaviorTree{}}
	})

	return factoty
}

func (this *BehaviorFactory) GetBTree(name string) *btree.BehaviorTree {
	tree, ok := this.bTrees[name]
	if !ok {
		return nil
	}
	return tree
}

func (this *BehaviorFactory) Load(file string) {
	// 加载配置文件
	// 没做这个
}

func (this *BehaviorFactory) CreateTestBT() {
	builder := btree.NewBehaviorTreeBuilder("test1")

	//tt := composite.NewSelectorImpl("select1")

	tree := builder. //AddBehavior(tt).
				AddBehavior(decorate.NewDecorateRepate("repeate", 3)).
				AddBehavior(action.NewActionAttack("attack")).
				End()
		//AddBehavior(composite.NewSequenceImpl("sequence")).
		//AddBehavior(condition.NewConditionIsSeeEnemy(true, "isseeenemy")).
		//Back().
		//AddBehavior(composite.NewSelectorImpl("select")).
		//AddBehavior(composite.NewSequenceImpl("sequence")).
		//AddBehavior(condition.NewConditionIsHeathLow(true, "isheathloww")).
		//Back().
		//AddBehavior(action.NewActionRunaway("actionrunaway")).
		//Back().
		//Back().
		//AddBehavior(composite.NewParallelImpl(common.RequireAll, common.RequireOne, "paralle")).
		//AddBehavior(condition.NewConditionIsEnemyDead(true, "isenemydead")).
		//Back().
		//AddBehavior(decorate.NewDecorateRepate("repeate", 3)).
		//AddBehavior(action.NewActionAttack("attack")).
		//Back().
		//Back().
		//Back().
		//Back().
		//Back().
		//AddBehavior(action.NewActionPatrol("patrol")).
		//End()

	fmt.Println(tree.GetRoot())
	this.bTrees[tree.GetName()] = tree
}
