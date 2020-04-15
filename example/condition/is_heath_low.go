// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  is_heath_low
 * @Version: 1.0.0
 * @Date: 2020/4/10 14:02
 */

package condition

import (
	"fmt"

	bbehavior "github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

type ConditionIsHeathLow struct {
	bbehavior.BaseCondition
}

func NewConditionIsHeathLow(neg bool, name string) interf.IBehaviour {
	con := &ConditionIsHeathLow{}
	con.Init(neg, common.BehaviorName(name))

	return con
}

func (this *ConditionIsHeathLow) Update(inter interface{}) common.StatusType {
	random := this.GetRandom()
	if random < 60 {
		fmt.Println("Health is low")
		if !this.IsNegation() {
			return common.Status_Success
		}

		return common.Status_Failure
	} else {
		fmt.Println("Health is Not low")
		if !this.IsNegation() {
			return common.Status_Failure
		}

		return common.Status_Success
	}
}
