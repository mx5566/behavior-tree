// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  status
 * @Version: 1.0.0
 * @Date: 2020/4/9 15:04
 */

package common

type StatusType int

const (
	Status_Invalid StatusType = iota // 初始状态
	Status_Success                   // 成功
	Status_Failure                   // 失败
	Status_Running                   // 运行
	Status_Aborted                   // 终止
)
