// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  b_data
 * @Version: 1.0.0
 * @Date: 2020/4/13 18:19
 */

package base

import "github.com/mx5566/behavior-tree/base_behavior"

type IEntity interface {
}

type AIEntityWorkingData struct {
	base_behavior.BTWorkingData
	Entity    IEntity
	GameTime  float32
	DeltaTime float32
	Position  Position
}

func (this *AIEntityWorkingData) No() {

}

type Point struct {
	X int // 格子点
	Y int
}

type Position struct {
	XPos float32 // 坐标
	YPos float32
}
