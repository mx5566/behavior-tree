// Copyright (c) 2020 by meng.  All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/**
 * @Author: meng
 * @Description:
 * @File:  behavior
 * @Version: 1.0.0
 * @Date: 2020/4/9 16:49
 */

package base_behavior

import (
	"fmt"
	"reflect"

	"github.com/mx5566/behavior-tree/common"
	"github.com/mx5566/behavior-tree/interf"
)

/**
 * Behavior接口是所有行为树节点的核心，且我规定所有节点的构造和析构方法都必须是protected，以防止在栈上创建对象，
 * 所有的节点对象通过Create()静态方法在堆上创建，通过Release()方法销毁
 * 由于Behavior是个抽象接口，故没有提供Create()方法，本接口满足如下契约
 * 在Update方法被首次调用前，调用一次OnInitialize函数，负责初始化等操作
 * Update（）方法在行为树每次更新时调用且仅调用一次。
 * 当行为不再处于运行状态时，调用一次OnTerminate（），并根据返回状态不同执行不同的逻辑
 */

var UniqueID common.UniqueID = 0

// 节点额外的数据
// 继承了  IBTBehaviorContext
type BehaviorContext struct {
	Object interface{}
}

// 获取内部数据
// 这里的参数就类似c++模板类型T 通过反射获取类型
// 这样就不需要传递具体的类型，做到通用函数
func (this *BehaviorContext) GetObject(t interface{}) interface{} {
	if this.Object == nil {
		c := reflect.TypeOf(t)
		n := reflect.New(c).Interface()

		// 创建新的空的初始化0值对象
		this.Object = n

		return n
	}
	// 已经存在的对象
	return this.Object
}

type BaseBehavior struct {
	interf.IBehaviorWork
	status common.StatusType
	name   common.BehaviorName //节点名字
	ID     common.UniqueID     // 唯一ID值
}

func (this *BaseBehavior) Init(name common.BehaviorName) {
	this.status = common.Status_Invalid
	this.name = name
	this.ID = this.GenUniqueID()
}

func (this *BaseBehavior) Tick(inter interface{}) common.StatusType {

	//fmt.Println(this)
	//update方法被首次调用前执行OnInitlize方法，每次行为树更新时调用一次update方法
	//当刚刚更新的行为不再运行时调用OnTerminate方法
	if this.status != common.Status_Running {
		this.OnInitialize(inter)
	}

	//	fmt.Println("BaseBehavior Tick ", this.GetName(), " Status ", this.status, " Address ", &this)
	status := this.Update(inter)
	this.SetStatus(status)

	if status != common.Status_Running {
		this.OnTerminate(status)
	}

	return status
}

func (this *BaseBehavior) AddChild(child interf.IBehaviour) {
	return
}

func (this *BaseBehavior) Release() {

}

func (this *BaseBehavior) SetStatus(statusType common.StatusType) {
	this.status = statusType
}

func (this *BaseBehavior) GetStatus() common.StatusType {
	return this.status
}

// reset this behavior
func (this *BaseBehavior) Reset() {
	this.SetStatus(common.Status_Invalid)
}

// 终止节点行为
func (this *BaseBehavior) Abort() {
	this.OnTerminate(common.Status_Aborted)
	this.SetStatus(common.Status_Aborted)
}

// 设置名字
func (this *BaseBehavior) SetName(name string) {
	this.name = common.BehaviorName(name)
}

// 获取节点名字
func (this *BaseBehavior) GetName() common.BehaviorName {
	return this.name
}

// 获取唯一key
func (this *BaseBehavior) GetUniqueID() common.UniqueID {
	return this.ID
}

// generate id
func (this *BaseBehavior) GenUniqueID() common.UniqueID {
	UniqueID++

	return UniqueID
}

func (this *BaseBehavior) GetContextEx(data *BTWorkingData, t interface{}) interface{} {
	if data == nil || t == nil {
		return nil
	}

	id := this.GetUniqueID()

	//fmt.Printf("it is nil data.WData %v ", data.WData == nil)
	if context, ok := data.WData[id]; !ok {
		// 对应的节点数据没有被初始化
		// 做初始化处理
		c := reflect.TypeOf(t)
		n := reflect.New(c).Interface()

		data.WData[id] = n

		return n
	} else {
		return context
	}
}

func (this *BaseBehavior) GetContext(inter interface{}, t interface{}) interface{} {
	if inter == nil || t == nil {
		return nil
	}

	id := this.GetUniqueID()

	bt := reflect.ValueOf(inter).Elem().FieldByName("BTWorkingData")
	fmt.Printf("%v-------------", bt)

	data := inter.(*BTWorkingData)
	if context, ok := data.WData[id]; !ok {
		// 对应的节点数据没有被初始化
		// 做初始化处理
		c := reflect.TypeOf(t)
		n := reflect.New(c)

		data.WData[id] = n
		return n
	} else {
		return context
	}
}

func (this *BaseBehavior) SetBehaviorWork(work interf.IBehaviorWork) {
	this.IBehaviorWork = work
}

//
func (this *BaseBehavior) GetBehaviorWorker() interf.IBehaviorWork {
	return this.IBehaviorWork
}
