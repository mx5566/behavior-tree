// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  entitymgr
 * @Version: 1.0.0
 * @Date: 2020/4/11 22:05
 */

package entity

import "sync"

var entityMgr *EntityMgr
var once sync.Once

// 实体管理器
type EntityMgr struct {
	entitys map[uint]*Entity
}

func NewEntityMgr() *EntityMgr {
	once.Do(func() {
		entityMgr = &EntityMgr{entitys: map[uint]*Entity{}}
	})
	return entityMgr
}

func (this *EntityMgr) AddEntity(entity *Entity) bool {
	if nil == entity {
		return false
	}
	_, ok := this.entitys[entity.GetID()]
	if ok {
		return true
	}

	this.entitys[entity.GetID()] = entity

	return true
}

func (this *EntityMgr) DeleteEntity(entity *Entity) bool {
	if entity == nil {
		return false
	}

	delete(this.entitys, entity.GetID())
	return true
}

func (this *EntityMgr) Update() {
	//for key, entity := range this.entitys {
	//
	//	}
}
