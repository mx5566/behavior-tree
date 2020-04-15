// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  idecorate
 * @Version: 1.0.0
 * @Date: 2020/4/9 15:46
 */

package interf

/************************************************************************/
/* 装饰结点
/* 装饰：
*********************************************************************** */
type IDecorator interface {
	IBehaviour
	SetChild(child IBehaviour)
	GetChild() IBehaviour
}
