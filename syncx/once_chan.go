package syncx

import "sync"

// OnceChan 一次性通道
type OnceChan struct {
	done chan struct{}
	once sync.Once
}

func NewOnceChan() *OnceChan {
	return &OnceChan{
		done: make(chan struct{}),
	}
}

// IsClose 判断是否已经关闭
func (c *OnceChan) IsClose() bool {
	select {
	case <-c.Done():
		return true
	default:
		return false
	}
}

// Close 关闭 done channel
func (c *OnceChan) Close() {
	c.once.Do(func() {
		close(c.done)
	})
}

// Done 返回关闭信号
func (c *OnceChan) Done() <-chan struct{} {
	return c.done
}
