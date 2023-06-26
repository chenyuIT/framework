package facades

import "github.com/chenyuIT/framework/contracts/mail"

func Mail() mail.Mail {
	return App().MakeMail()
}
