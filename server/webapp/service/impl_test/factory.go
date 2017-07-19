package impl_test

import (
	"server/webapp/service/interfaces"
)

// Factory 工厂类
type Factory struct {
}

// NewFactory ..
func NewFactory() *Factory {
	return &Factory{}
}

// NewAccountSrv ..
func (f *Factory) NewAccountSrv() interfaces.IAccountSrv {
	return newAccountSrv()
}
