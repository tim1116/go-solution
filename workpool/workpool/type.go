package workpool

type Task func(param ...interface{})

type RecFun func(interface{})
