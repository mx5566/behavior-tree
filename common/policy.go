// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  policy
 * @Version: 1.0.0
 * @Date: 2020/4/9 16:26
 */

package common

type Policy int

//Parallel节点成功与失败的要求，是全部成功/失败，还是一个成功/失败
const (
	RequireOne Policy = iota
	RequireAll
)
