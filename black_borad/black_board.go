// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  black_board
 * @Version: 1.0.0
 * @Date: 2020/4/12 12:04
 */

package black_borad

import "errors"

type BlackBoardData interface {
}

type BlackBoard struct {
	// 黑板数据
	bBData map[string]BlackBoardData
}

func (this *BlackBoard) SetD(key string, value BlackBoardData) {
	this.bBData[key] = value
}

func (this *BlackBoard) GetD(key string) (BlackBoardData, error) {
	value, ok := this.bBData[key]
	if !ok {
		return nil, errors.New("Not Find")
	}

	return value, nil
}
