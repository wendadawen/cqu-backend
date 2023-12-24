package tool

import "log"

// TODO 整个都不会

type Handler func(input interface{}) (bool, interface{})
type ChainHandler struct {
	handlerChain []Handler
	// 默认的 handler，如果设置了 defaultHandler，将会在 chain 都没命中的时候进行处理
	defaultHandler Handler
	done           bool
}

func NewChainHandler(chain []Handler) *ChainHandler {
	if chain == nil {
		log.Fatal("Handler chain 为空")
	}

	return &ChainHandler{
		handlerChain:   chain,
		defaultHandler: nil,
		done:           false,
	}
}
func (c *ChainHandler) SetDefault(handler Handler) *ChainHandler {
	c.defaultHandler = handler
	return c
}

// Handle 方法只链检查并可能对 input 进行修改（比如 input 是 error 类型，可能会被 handler 修改）
func (c *ChainHandler) Handle(input interface{}) interface{} {
	var result interface{}
	for _, handler := range c.handlerChain {
		if c.done == false {
			c.done, result = handler(input)
		}
	}
	if c.defaultHandler != nil && c.done == false {
		_, result = c.defaultHandler(input)
	}
	return result
}
