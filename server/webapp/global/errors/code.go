package errors

// 错误码
type Enum int

const (
	Unknow               Enum = 1999
	OK                   Enum = 2000
	ParamsNil            Enum = 2001
	ParamsError          Enum = 2002
	PwdError             Enum = 2003
	NotLogin             Enum = 2004
	NoEmptyRoom          Enum = 2005
	NotFound             Enum = 2006
	NotEnoughChips       Enum = 2007
	NotEnoughDiamond     Enum = 2008
	NoSeat               Enum = 2009
	InvalidSessionToken  Enum = 2010
	ExpiredSessionToken  Enum = 2011
	UserNotExist         Enum = 2012
	CreateUserFailed     Enum = 2013
	CreateCfgShopFailed  Enum = 2014
	CreateTokenFailed    Enum = 2015
	UserNameAlreadyExist Enum = 2016
	EmailAlreadyExist    Enum = 2017
)

func (x Enum) String() string {
	return ""
}
