package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

func (t Translation) Translate(key string) string {
	log.Printf("目前的语言为%s\n", t.Language)
	switch t.Language {
	case English:
		//创建一个翻译器
		bundle := i18n.NewBundle(language.English)
		// 加载翻译文件
		_, err := bundle.LoadMessageFile("locales/en.json")
		if err != nil {
			log.Fatalf("Failed to load en.json: %v", err)
		}
		localized := i18n.NewLocalizer(bundle, language.English.String())
		return localized.MustLocalize(&i18n.LocalizeConfig{MessageID: key})
	case Chinese:
		//创建一个翻译器
		bundle := i18n.NewBundle(language.Chinese)
		// 加载翻译文件
		_, err := bundle.LoadMessageFile("locales/cn.json")
		if err != nil {
			log.Fatalf("Failed to load cn.json: %v", err)
		}
		localized := i18n.NewLocalizer(bundle, language.Chinese.String())
		return localized.MustLocalize(&i18n.LocalizeConfig{MessageID: key})
	default:
		//创建一个翻译器
		bundle := i18n.NewBundle(language.Chinese)
		// 加载翻译文件
		_, err := bundle.LoadMessageFile("locales/cn.json")
		if err != nil {
			log.Fatalf("Failed to load cn.json: %v", err)
		}
		localized := i18n.NewLocalizer(bundle, language.Chinese.String())
		return localized.MustLocalize(&i18n.LocalizeConfig{MessageID: key})
	}
}
