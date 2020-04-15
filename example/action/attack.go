// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  attack
 * @Version: 1.0.0
 * @Date: 2020/4/9 20:58
 */

package action

import (
	"fmt"

	"github.com/mx5566/behavior-tree/example/base"

	bbehavior "github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

// 攻击行为
type ActionAttack struct {
	bbehavior.BaseAction
}

type UserContextData struct {
	AttackTime float32 // 攻击持续时间
}

func NewActionAttack(name string) interf.IBehaviour {
	aa := new(ActionAttack)
	aa.Init(common.BehaviorName(name))
	return aa
}

func (this *ActionAttack) OnInitialize(inter interface{}) {
	thisData := inter.(*base.AIEntityWorkingData)
	userData := this.GetContextEx(&thisData.BTWorkingData, bbehavior.BehaviorContext{}).(*bbehavior.BehaviorContext).
		GetObject(UserContextData{AttackTime: 0}).(*UserContextData)

	userData.AttackTime = 5

	fmt.Println("Enter OnInitialize")
	//fmt.Println("-----------------------------")
	//fmt.Println(thisData, " ", userData, " ", inter)
	//fmt.Println("-----------------------------")

}

func (this *ActionAttack) Update(inter interface{}) common.StatusType {
	fmt.Println("Attack Success")

	thisData := inter.(*base.AIEntityWorkingData)
	userData := this.GetContextEx(&thisData.BTWorkingData, bbehavior.BehaviorContext{}).(*bbehavior.BehaviorContext).
		GetObject(UserContextData{AttackTime: 0}).(*UserContextData)

	if userData.AttackTime > 0 {
		userData.AttackTime -= thisData.DeltaTime

		//fmt.Println("-----------------------------")
		//fmt.Println(inter, " ", userData)
		//fmt.Println("-----------------------------")

		if userData.AttackTime <= 0 {
			return common.Status_Success
		}
	}
	// 执行攻击
	// 攻击时间
	return common.Status_Running
}
