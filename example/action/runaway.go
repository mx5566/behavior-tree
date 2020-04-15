// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  runaway
 * @Version: 1.0.0
 * @Date: 2020/4/9 21:13
 */

package action

import (
	"fmt"

	bbehavior "github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

// 逃跑
type ActionRunaway struct {
	bbehavior.BaseAction
}

func NewActionRunaway(name string) interf.IBehaviour {
	aa := &ActionRunaway{}
	aa.Init(common.BehaviorName(name))

	return aa
}

func (this *ActionRunaway) Update(inter interface{}) common.StatusType {
	fmt.Println("Runaway Success")

	return common.Status_Success
}
