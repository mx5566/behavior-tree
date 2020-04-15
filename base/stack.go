// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  stack
 * @Version: 1.0.0
 * @Date: 2020/4/10 16:24
 */

package base

import "sync"

type StackValue interface {
}

// 链表节点
type StackNode struct {
	// 节点值
	value StackValue
	// 前向节点
	prev *StackNode
	// 后一个节点
	next *StackNode
}

func (this *StackNode) GetV() StackValue {
	return this.value
}

type Stack struct {
	head       *StackNode // 队列的头
	sync.Mutex            // 锁
}

// 初始化栈
func NewStack() *Stack {
	return &Stack{
		head: nil,
	}
}

func (this *Stack) Push(value StackValue) {
	this.Lock()
	defer this.Unlock()

	node := new(StackNode)
	node.value = value
	node.next = nil
	node.prev = nil
	if this.head == nil {
		this.head = node
		return
	}

	this.head.next = node
	node.prev = this.head

	this.head = node
}

func (this *Stack) Pop() *StackNode {
	this.Lock()
	defer this.Unlock()

	// empty
	if this.head == nil {
		return nil
	}

	ret := this.head

	temp := this.head.prev

	this.head.prev.next = nil
	this.head.prev = nil
	this.head = temp

	return ret
}

// 返回栈顶元素 不删除
func (this *Stack) Peek() *StackNode {
	this.Lock()
	defer this.Unlock()

	return this.head
}

func (this *Stack) IsEmpty() bool {
	this.Lock()
	defer this.Unlock()

	if this.head == nil {
		return true
	}
	return false
}

func (this *Stack) Clear() {
	this.Lock()
	defer this.Unlock()

	this.head = nil
}
