package workpool

type pool struct {
	max         int
	chanQueue   chan struct{}
	RecoverHand RecFun
}

func NewPool(max int) *pool {
	if max < 1 {
		panic("max参数异常")
	}
	p := &pool{
		max: max,
		RecoverHand: func(i interface{}) {
			panic(i)
		},
		chanQueue: make(chan struct{}, max),
	}
	for i := 0; i < max; i++ {
		p.chanQueue <- struct{}{}
	}
	return p
}

// 添加任务
func (p *pool) Do(t Task, param ...interface{}) {
	<-p.chanQueue
	go func() {
		// 处理panic
		defer func() {
			if pac := recover(); pac != nil {
				p.RecoverHand(pac)
			}
			p.chanQueue <- struct{}{}
		}()
		t(param...)
	}()

}

// 阻塞等待任务结束
func (p *pool) Wait() {
	for i := 0; i < p.max; i++ {
		<-p.chanQueue
	}
	close(p.chanQueue)
}
