// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  patrol
 * @Version: 1.0.0
 * @Date: 2020/4/9 21:10
 */

package action

import (
	"fmt"

	bbehavior "github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

// 巡逻
type ActionPatrol struct {
	bbehavior.BaseAction
}

func NewActionPatrol(name string) interf.IBehaviour {
	aa := &ActionPatrol{}
	aa.Init(common.BehaviorName(name))

	return aa
}

func (this *ActionPatrol) Update(inter interface{}) common.StatusType {
	fmt.Println("Patrol Success")

	return common.Status_Success
}
