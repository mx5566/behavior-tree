// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  action-mode
 * @Version: 1.0.0
 * @Date: 2020/4/9 16:22
 */

package common

type ActionMode int

const (
	Attack ActionMode = iota
	Patrol
	Runaway
)
