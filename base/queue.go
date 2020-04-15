// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  queue.go
 * @Version: 1.0.0
 * @Date: 2020/4/10 16:18
 */
package base

import "sync"

// IQueueValue
type IQueueValue interface {
}

// 节点值
type NodeValue interface {
}

// 链表节点
type Node struct {
	// 节点值
	value NodeValue
	// 前向节点
	prev *Node
	// 后一个节点
	next *Node
}

// 普通队列
type LinkQueue struct {
	tail *Node // 队列的尾
	head *Node // 队列的头

	sync.Mutex
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		tail: nil,
		head: nil,
	}
}

// 入队列
func (this *LinkQueue) Push(node *Node) {
	if node == nil {
		return
	}

	if this.head == nil && this.tail == nil {
		this.head = node
		this.tail = node
		return
	}

	// 设置尾节点
	this.tail.next = node
	this.tail.next.prev = this.tail

	this.tail = node
}

// 出队列
func (this *LinkQueue) Pop() *Node {
	// 空队列 返回空
	if this.head == nil {
		return nil
	}

	ret := this.head
	temp := this.head.next

	this.head.next.prev = nil
	this.head.next = nil

	this.head = temp
	return ret
}

// 队列是否为空
func (this *LinkQueue) IsEmpty() bool {
	if this.head == nil {
		return true
	}
	return false
}

// 循环数组队列
type CircleQueue struct {
	front      int           // 队列头
	back       int           // 队列尾
	nSize      int           // 队列容量
	array      []IQueueValue // 队列数组
	sync.Mutex               // 互斥锁
}

func NewCircleQueue(nSize int) *CircleQueue {
	return &CircleQueue{
		front: 0,
		back:  0,
		nSize: nSize,
		array: make([]IQueueValue, nSize),
	}
}

func (this *CircleQueue) Push(value IQueueValue) int {
	// 满了
	if this.IsFull() {
		return -1
	}

	this.Lock()
	defer this.Unlock()
	// 尾部插入
	this.array[this.back] = value
	this.back = (this.back + 1) % this.nSize
	return 1
}

func (this *CircleQueue) Pop() IQueueValue {
	if this.IsEmpty() {
		return nil
	}

	this.Lock()
	defer this.Unlock()
	// 头部出队列
	temp := this.array[this.front]
	this.front = (this.front + 1) % this.nSize
	return temp
}

func (this *CircleQueue) IsEmpty() bool {
	this.Lock()
	defer this.Unlock()

	if this.front == this.back {
		return true
	}

	return false
}

func (this *CircleQueue) IsFull() bool {
	this.Lock()
	defer this.Unlock()

	if (this.back+1)%this.nSize == this.front {
		return true
	}

	return false
}
