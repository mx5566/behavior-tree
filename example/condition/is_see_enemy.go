// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  is_see_enemy
 * @Version: 1.0.0
 * @Date: 2020/4/10 14:03
 */

package condition

import (
	"fmt"

	bbehavior "github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

type ConditionIsSeeEnemy struct {
	bbehavior.BaseCondition
}

func NewConditionIsSeeEnemy(neg bool, name string) interf.IBehaviour {
	con := &ConditionIsSeeEnemy{}
	con.Init(neg, common.BehaviorName(name))

	return con
}

func (this *ConditionIsSeeEnemy) Update(inter interface{}) common.StatusType {

	// 需要根据玩家
	random := this.GetRandom()
	if random < 50 {
		fmt.Println("SeeEnemy")
		if !this.IsNegation() {
			return common.Status_Success
		}

		return common.Status_Failure
	} else {
		fmt.Println("Not SeeEnemy")
		if !this.IsNegation() {
			return common.Status_Failure
		}

		return common.Status_Success
	}
}
