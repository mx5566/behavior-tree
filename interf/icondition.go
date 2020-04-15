// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  icondition
 * @Version: 1.0.0
 * @Date: 2020/4/9 15:44
 */

package interf

/************************************************************************/
/* 条件基类
/* 具体执行动作                                   */

/************************************************************************/
type ICondition interface {
	IBehaviour

	IsNegation() bool

	SetNegation(negation bool)
}
