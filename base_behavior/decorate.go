// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  decorate
 * @Version: 1.0.0
 * @Date: 2020/4/9 20:33
 */

package base_behavior

import (
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

// 装饰器
type BaseDecorate struct {
	BaseBehavior
	BehaviorWork

	child interf.IBehaviour //装饰器只会有一个子节点
}

func (this *BaseDecorate) Init(name common.BehaviorName) {
	this.BaseBehavior.Init(name)
}

func (this *BaseDecorate) AddChild(child interf.IBehaviour) {
	this.child = child
}

func (this *BaseDecorate) GetChild() interf.IBehaviour {
	return this.child
}
