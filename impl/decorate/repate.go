// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  repate
 * @Version: 1.0.0
 * @Date: 2020/4/10 14:04
 */

package decorate

import (
	"fmt"

	bbehavior "github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

type DecorateRepate struct {
	bbehavior.BaseDecorate

	limited int // 上限次数
	count   int // 迭代的次数
}

func NewDecorateRepate(name string, limit int) interf.IBehaviour {
	dr := new(DecorateRepate)

	dr.limited = 3
	dr.count = 0
	// 初始化节点名字
	dr.Init(common.BehaviorName(name))

	return dr
}

func (this *DecorateRepate) Update(inter interface{}) common.StatusType {
	for true {
		child := this.GetChild()

		child.Tick(inter)
		switch child.GetStatus() {
		case common.Status_Running:
			return common.Status_Success
		case common.Status_Failure:
			return common.Status_Failure
		default:
			break
		}

		this.count++
		if this.count >= this.limited {
			fmt.Println("OK One+++++++")
			return common.Status_Success

		}
		child.Reset()
	}

	return common.Status_Invalid
}
