package lib 

type flagsInterface interface {
	SetHeaders(string)
	SetRequestNumber(int)
	SetConcurrentRequestNumber(int)
	SetRequestType(string)
	SetURL(string)
	GetHeaders() (string)
	GetRequestNumber() (int)
	GetConcurrentRequestNumber() (int)
	GetRequestType() (string)
	GetURL() (string)
}