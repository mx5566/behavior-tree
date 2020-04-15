// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  condition
 * @Version: 1.0.0
 * @Date: 2020/4/9 20:23
 */

package base_behavior

import (
	"math/rand"
	"time"

	"github.com/mx5566/behavior-tree/common"
)

type BaseCondition struct {
	BaseBehavior
	BehaviorWork

	negation bool
}

func (this *BaseCondition) Init(neg bool, name common.BehaviorName) {
	this.SetNegation(neg)
	this.BaseBehavior.Init(name)
}

func (this *BaseCondition) IsNegation() bool {
	return this.negation
}

func (this *BaseCondition) SetNegation(neg bool) {
	this.negation = neg
}

func (this *BaseCondition) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	//    int i = random.intValue();
	return r
}
