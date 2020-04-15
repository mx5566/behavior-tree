// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  behavior_tree
 * @Version: 1.0.0
 * @Date: 2020/4/10 15:28
 */

package behavior_tree

import (
	"github.com/mx5566/behavior-tree/interf"
)

type BehaviorTree struct {
	root interf.IBehaviour
	name string
}

func NewBehaviorTree(root interf.IBehaviour, name string) *BehaviorTree {
	return &BehaviorTree{root: root, name: name}
}

func (this *BehaviorTree) GetName() string {
	return this.name
}

func (this *BehaviorTree) Tick(inter interface{}) {
	this.root.Tick(inter)
}

func (this *BehaviorTree) HaveRoot() bool {
	if this.root != nil {
		return true
	}

	return false
}

func (this *BehaviorTree) SetRoot(inNode interf.IBehaviour) {
	this.root = inNode
}

func (this *BehaviorTree) Release() {
	this.root.Release()
}

func (this *BehaviorTree) GetRoot() interf.IBehaviour {
	return this.root
}
