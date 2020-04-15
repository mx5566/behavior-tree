// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  behavior_tree_builder
 * @Version: 1.0.0
 * @Date: 2020/4/10 15:47
 */

package behavior_tree

import (
	"fmt"

	"github.com/mx5566/behavior-tree/base"
	"github.com/mx5566/behavior-tree/interf"
)

type BehaviorTreeBuilder struct {
	stack    *base.Stack       //存储节点
	treeRoot interf.IBehaviour // 行为树跟节点
	name     string
}

func NewBehaviorTreeBuilder(name string) *BehaviorTreeBuilder {
	return &BehaviorTreeBuilder{
		stack:    new(base.Stack),
		treeRoot: nil,
		name:     name,
	}
}

func (this *BehaviorTreeBuilder) AddBehavior(behaviour interf.IBehaviour) *BehaviorTreeBuilder {
	//如果没有根节点设置新节点为根节点
	//否则设置新节点为堆栈顶部节点的子节点
	//fmt.Println("AddBehavior Name ", behaviour.GetName())

	// 作用是用来回调用户实现类--也是行为节点
	behaviour.SetBehaviorWork(behaviour.(interf.IBehaviorWork))
	//fmt.Println("Address ", behaviour, " Name ", behaviour.GetName())
	if this.treeRoot == nil {
		this.treeRoot = behaviour
	} else {
		node := this.stack.Peek()
		if node == nil {
			fmt.Println("Are You Ok")
			return this
		}

		node.GetV().(interf.IBehaviour).AddChild(behaviour)
	}

	//fmt.Println(behaviour)
	//将新节点压入堆栈
	this.stack.Push(behaviour)
	return this
}

func (this *BehaviorTreeBuilder) Back() *BehaviorTreeBuilder {
	_ = this.stack.Pop()

	//fmt.Println("Back Name", node.GetV().(interf.IBehaviour).GetName())
	return this
}

func (this *BehaviorTreeBuilder) End() *BehaviorTree {
	// 清空栈空间
	this.stack.Clear()
	tree := NewBehaviorTree(this.treeRoot, this.name)

	return tree
}
