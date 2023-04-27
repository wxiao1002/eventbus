# eventbus
简单实现 eventbus: 事件消息通知机制

## 是什么
为了解决iot edge 南向采集数据并转发到北向过程中消息传递问题，使用eventbus 也就是生产者消费者模式

## 如何使用
- 创建eventbus 
```
// 系统默认提供一个eventbus
// 可以使 PubLish ,SubScribe DeSubScribe操作

eventbus.NewEventBus()
```

- 发布消息
```
eventbus.Pub("xxx",data any)
```

- 订阅消息
```
f  := func(data any) {
		fmt.Print(data)
	}
subId := eventbus.SubScribe("xxxx", f)
```

- 取消订阅
```
eventbus.DeSub("xxxx", subId)
```

