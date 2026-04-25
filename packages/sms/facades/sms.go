package facades

import (
	"log"

	"dakhl/packages/sms"
	"dakhl/packages/sms/contracts"
)

func Sms() contracts.Sms {
	instance, err := sms.App.Make(sms.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Sms)
}
