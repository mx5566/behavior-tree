// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  selector_impl
 * @Version: 1.0.0
 * @Date: 2020/4/9 21:23
 */

package composite

import (
	"github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

type SelectorImpl struct {
	base_behavior.BaseComposite
	currChild interf.IBehaviour
}

func NewSelectorImpl(name string) interf.IBehaviour {
	var sel = &SelectorImpl{
		currChild: nil,
	}
	//sel.BaseComposite
	sel.Init(common.BehaviorName(name))

	return sel
}

func (this *SelectorImpl) OnInitialize(inter interface{}) {
	childs := this.GetChildren()
	if len(childs) > 0 {
		this.currChild = this.GetChildren()[0]
	}
}

func (this *SelectorImpl) Update(inter interface{}) common.StatusType {
	childs := this.GetChildren()
	_ = len(childs)
	for _, child := range childs {
		this.currChild = child
		status := child.Tick(inter)
		if status != common.Status_Failure {
			return status
		}
	}
	//循环全部失败 或者空
	return common.Status_Failure
}
