// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  parallel_impl
 * @Version: 1.0.0
 * @Date: 2020/4/9 21:23
 */

package composite

import (
	"fmt"

	bbehavior "github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

type ParallelImpl struct {
	// 组合节点基类
	bbehavior.BaseComposite
	successPolicy common.Policy
	failPolocy    common.Policy
}

func NewParallelImpl(successPolicy, failPolocy common.Policy, name string) interf.IBehaviour {
	par := &ParallelImpl{
		successPolicy: successPolicy,
		failPolocy:    failPolocy,
	}
	par.Init(common.BehaviorName(name))

	return par
}

func (this *ParallelImpl) Update(inter interface{}) common.StatusType {
	successCount, failureCount := 0, 0

	childrenSize := len(this.GetChildren())
	for _, iBehaviour := range this.GetChildren() {
		// 如果行为已经终止则不再执行该行为
		fmt.Println("ParallelImpl Update ", iBehaviour.GetStatus())

		if !((iBehaviour.GetStatus() == common.Status_Success) || (iBehaviour.GetStatus() == common.Status_Failure)) {
			iBehaviour.Tick(inter)
		}

		fmt.Println("ParallelImpl Update ", iBehaviour.GetStatus())
		// 执行成功 把成功的个数加一
		// 根据策略 把节点初始化状态
		if iBehaviour.GetStatus() == common.Status_Success {
			successCount++
			if this.successPolicy == common.RequireOne {
				iBehaviour.Reset()
				return common.Status_Success
			}
		}

		// 执行成功 把失败的次数加一
		// 根据策略 把节点初始化状态
		if iBehaviour.GetStatus() == common.Status_Failure {
			failureCount++
			if this.failPolocy == common.RequireOne {
				iBehaviour.Reset()
				return common.Status_Failure
			}
		}
	}

	// 失败策略全部
	// 满足要求 所有孩子节点重置为初始状态
	if (this.failPolocy == common.RequireAll) && (failureCount == childrenSize) {
		for _, iBehaviour := range this.GetChildren() {
			iBehaviour.Reset()
		}

		fmt.Println("this.failPolocy == common.RequireAll")
		return common.Status_Failure
	}

	// 成功策略全部
	// 满足要求 所有孩子节点重置为初始状态
	if (this.successPolicy == common.RequireAll) && (successCount == childrenSize) {
		for _, iBehaviour := range this.GetChildren() {
			iBehaviour.Reset()
		}

		fmt.Println("this.successPolicy == common.RequireAll")

		return common.Status_Success
	}

	return common.Status_Running
}

func (this *ParallelImpl) OnTerminate(InStatus common.StatusType) {
	for _, iBeheavior := range this.GetChildren() {
		if iBeheavior.GetStatus() == common.Status_Running {
			iBeheavior.Abort()
		}

		iBeheavior.Reset()
	}
}

func (this *ParallelImpl) Abort() {
	this.OnTerminate(common.Status_Aborted)
	this.BaseComposite.Abort()
}
