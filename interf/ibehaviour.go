// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  ibehaviour
 * @Version: 1.0.0
 * @Date: 2020/4/9 14:31
 */

package interf

import (
	"github.com/mx5566/behavior-tree/common"
)

// 行为接口
type IBehaviour interface {
	Tick(inter interface{}) common.StatusType //设置调用顺序，onInitialize--update--onTerminate
	//
	Release() //释放对象所占资源
	//
	AddChild(child IBehaviour)
	//
	SetStatus(statusType common.StatusType)
	//
	GetStatus() common.StatusType
	//
	Reset()
	//
	Abort()
	//
	SetName(name string)
	//
	GetName() common.BehaviorName
	//
	GetUniqueID() common.UniqueID
	// Gen
	GenUniqueID() common.UniqueID
	// set behavior work
	SetBehaviorWork(work IBehaviorWork)
	//
	GetBehaviorWorker() IBehaviorWork
}
