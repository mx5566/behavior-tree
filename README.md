# behavior-tree



下载通过命令行
```batch
go get github.com/mx5566/behavior-tree
```
```
目前行为树创建是通过硬代码的形式通过程序去创建行为树
下面会加入通过json文件导入的来创建行为树

```

代码组织分布 
***

1、***base文件夹是一些基本的数据结构  ***  
2、base-behavior里面是行为树节点的基础类  
3、&ensp;其中behavior和behavior_work是所有各种节点的基础，所有的都要继承这两个类  
4、black_board可以理解成是全局数据，节点之间共享  
5、common里面是公共的变量，常量  
6、impl是一些通用节点的实现  
7、interf是接口类型

***


例子在example文件夹
