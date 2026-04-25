package config

import (
	"dakhl/app/facades"
)

func init() {
	config := facades.Config()
	config.Add("sms", map[string]any{
		"driver": "default",
	})
}
