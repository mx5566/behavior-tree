// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  data
 * @Version: 1.0.0
 * @Date: 2020/4/12 00:14
 */

package base_behavior

import "github.com/mx5566/behavior-tree/common"

// 数据
type IBTBehaviorContext interface {
}

type IBTWorkingData interface {
	No()
}

// 行为树数据
type BTWorkingData struct {
	WData map[common.UniqueID]IBTBehaviorContext

	// 看看是不是简单些逻辑
	//WData1 map[common.UniqueID]interface{}
}

func (this *BTWorkingData) No() {

}
