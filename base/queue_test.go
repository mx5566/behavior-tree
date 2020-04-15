// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  queue_test.go
 * @Version: 1.0.0
 * @Date: 2020/4/10 16:18
 */
package base

import (
	"fmt"
	"testing"
)

func TestNewLinkQueue(t *testing.T) {
	linkQueue := NewLinkQueue()

	node := new(Node)
	node.next = nil
	node.prev = nil
	node.value = 1

	linkQueue.Push(node)

	node = new(Node)
	node.value = 2
	linkQueue.Push(node)

	fmt.Println(linkQueue.IsEmpty())

	nodeRet := linkQueue.Pop()

	fmt.Println(nodeRet.value)
}

func TestNewCircleQueue(t *testing.T) {
	circleQueue := NewCircleQueue(3)
	circleQueue.Push(1)
	circleQueue.Push(2)
	circleQueue.Push(3)
	circleQueue.Push(4)

	fmt.Println(circleQueue.IsFull())

	fmt.Println(circleQueue.Pop())
	fmt.Println(circleQueue.Pop())
	fmt.Println(circleQueue.Pop())
	fmt.Println(circleQueue.Pop())
	fmt.Println(circleQueue.Pop())

	fmt.Println(circleQueue.IsEmpty())
}

func BenchmarkNewCircleQueue(b *testing.B) {
	circleQueue := NewCircleQueue(3)
	circleQueue.Push(1)
	circleQueue.Push(2)
	circleQueue.Push(3)
	circleQueue.Push(4)

	fmt.Println(circleQueue.IsFull())

	fmt.Println(circleQueue.Pop())
	fmt.Println(circleQueue.Pop())
	fmt.Println(circleQueue.Pop())
	fmt.Println(circleQueue.Pop())
	fmt.Println(circleQueue.Pop())

	fmt.Println(circleQueue.IsEmpty())
}
