package sms

type Sms interface {
	Send(mobile string) error
}
