// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  condition_mode
 * @Version: 1.0.0
 * @Date: 2020/4/9 16:25
 */

package common

type ConditionMode int

const (
	IsSeeEnemy ConditionMode = iota
	IsHealthLow
	IsEnemyDead
)
