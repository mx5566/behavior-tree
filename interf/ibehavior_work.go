// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  ibehavior_work
 * @Version: 1.0.0
 * @Date: 2020/4/15 01:04
 */

package interf

import "github.com/mx5566/behavior-tree/common"

type IBehaviorWork interface {
	OnInitialize(inter interface{}) //当节点调用前
	//
	Update(inter interface{}) common.StatusType //节点操作的具体实现
	//
	OnTerminate(statusType common.StatusType) //节点调用后执行
}
