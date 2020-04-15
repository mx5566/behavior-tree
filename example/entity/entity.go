// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  entity
 * @Version: 1.0.0
 * @Date: 2020/4/11 01:02
 */

package entity

import (
	"math"

	"github.com/mx5566/behavior-tree/example/base"

	behavior_tree "github.com/mx5566/behavior-tree"
)

var ID uint = 0

type ObjType int

const (
	Obj_Player ObjType = iota + 1
	Obj_Npc
)

const (
	XGrid  = 10
	YGrid  = 10
	Width  = 100
	Length = 100
)

type Entity struct {
	ID            uint
	Hp            int
	Type          ObjType
	Tree          *behavior_tree.BehaviorTree // 行为树
	Position      base.Position
	aiWorkingData *base.AIEntityWorkingData // 运行时数据 包括各个节点的数据
}

func NewEntity() *Entity {
	ID++
	instance := GetInstanceBehaviorFactory()
	instance.CreateTestBT()
	return &Entity{
		ID:   ID,
		Hp:   100,
		Type: Obj_Player,
		Tree: instance.GetBTree("test1"),
		Position: base.Position{
			XPos: 0,
			YPos: 0,
		},
		aiWorkingData: NewAIEntityWorkingData(),
	}
}

func (this *Entity) Init() {
	this.aiWorkingData.Entity = this
	this.aiWorkingData.Position = this.Position

}

func (this *Entity) GetPoint() base.Point {
	return base.Point{
		X: int(this.Position.XPos / XGrid),
		Y: int(this.Position.YPos / YGrid),
	}
}

func (this *Entity) GetID() uint {
	return this.ID
}

func (this *Entity) GetHp() int {
	return this.Hp
}

func (this *Entity) SetHp(hp int) {
	this.Hp = hp
}

func (this *Entity) AddHp(hp int) {
	if hp == 0 {
		return
	}

	this.Hp += hp

	this.Hp = int(math.Max(0, float64(this.Hp)))
}

func (this *Entity) UpdateBehavior(gameTime, deltaTime float32) {
	this.Tree.Tick(this.aiWorkingData)
}
