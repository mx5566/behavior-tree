// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  composite
 * @Version: 1.0.0
 * @Date: 2020/4/9 17:25
 */

package base_behavior

import (
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

// 已经继承了 IComposite
type BaseComposite struct {
	BaseBehavior
	BehaviorWork

	children []interf.IBehaviour
}

func (this *BaseComposite) Init(name common.BehaviorName) {
	this.BaseBehavior.Init(name)
}

func (this *BaseComposite) AddChild(child interf.IBehaviour) {
	//fmt.Println("BaseComposite AddChild ", child.GetName())
	this.children = append(this.children, child)
}

func (this *BaseComposite) RemoveChild(child interf.IBehaviour) {
	length := len(this.children)
	for i := 0; i < length; i++ {
		if this.children[i] != child {
			continue
		}
		this.children = append(this.children[0:i], this.children[i+1:]...)
		break
	}
}

func (this *BaseComposite) ClearChild() {
	this.children = this.children[0:0]
}

func (this *BaseComposite) GetChildren() []interf.IBehaviour {
	return this.children
}

func (this *BaseComposite) SetChildren(behaviours []interf.IBehaviour) {
	this.children = behaviours
}
