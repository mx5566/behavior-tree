// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  data
 * @Version: 1.0.0
 * @Date: 2020/4/12 00:33
 */

package entity

import (
	"github.com/mx5566/behavior-tree/base_behavior"
	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/example/base"
)

//

func NewAIEntityWorkingData() *base.AIEntityWorkingData {
	return &base.AIEntityWorkingData{
		Entity:    nil,
		GameTime:  0.5,
		DeltaTime: 0.5,
		BTWorkingData: base_behavior.BTWorkingData{
			WData: make(map[common.UniqueID]base_behavior.IBTBehaviorContext),
		},
	}
}
