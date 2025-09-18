package utils

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Global bundle and localizer getter
var bundle *i18n.Bundle

func InitTranslation() {
	// Initialize i18n bundle
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", i18n.UnmarshalFunc(func(data []byte, v interface{}) error {
		return json.Unmarshal(data, v)
	}))

	// Load translations
	bundle.MustLoadMessageFile("locales/en-US.json")
	bundle.MustLoadMessageFile("locales/id-ID.json")

}

func Translate(messageID string, data map[string]interface{}) string {
	lang := GetConfig().DefaultLanguage
	localizer := i18n.NewLocalizer(bundle, lang)

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
	})
	if err != nil {
		fmt.Println(err)
		return messageID // fallback
	}
	return msg
}
