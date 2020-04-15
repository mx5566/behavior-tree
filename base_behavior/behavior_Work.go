// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  behavior_Work
 * @Version: 1.0.0
 * @Date: 2020/4/15 01:05
 */

package base_behavior

import (
	"fmt"

	"github.com/mx5566/behavior-tree/common"
)

type BehaviorWork struct {
}

//当节点调用前
func (this *BehaviorWork) OnInitialize(inter interface{}) {

}

//节点操作的具体实现
func (this *BehaviorWork) Update(inter interface{}) common.StatusType {
	fmt.Println("BaseBehavior Update Address ")
	return common.Status_Invalid
}

//节点调用后执行
func (this *BehaviorWork) OnTerminate(statusType common.StatusType) {

}
