// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  icomposite
 * @Version: 1.0.0
 * @Date: 2020/4/9 15:41
 */

package interf

type IComposite interface {
	IBehaviour
	//
	RemoveChild(child IBehaviour)
	//
	ClearChild()
	//
	GetChildren() []IBehaviour
	//
	SetChildren(behaviours []IBehaviour)
}
