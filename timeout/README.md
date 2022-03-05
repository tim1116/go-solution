# 超时控制

当一段时间如果没有完成某项任务,则舍弃该任务,执行其他业务逻辑。

在服务端编程中经常使用,例如RPC,HTTP请求超时控制,DB超时控制等

## 使用方法

##### 1:goroutine中利用select控制超时

[代码示例1](./timeout1.go)

##### 2:借助time.After改造

[代码示例2](./timeout2.go)

##### 3:context.WithCancel 实现

[代码示例3](./timeout3.go)

##### 4:context.WithTimeout 实现

[代码示例4](./timeout4.go)

## 注意事项

- 其中 示例3 和 示例4 使用到content包实现在goroutine之间传递上下文消息,功能更加强大,实际开发中使用较多

- 使用中需要避免goroutine泄露的发生