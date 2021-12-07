package main

import (
	"fmt"
	"github.com/Xuanwo/go-locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
	"log"
)

// go install  github.com/nicksnyder/go-i18n/v2/goi18n

func main() {
	// Step 1: Create bundle
	bundle := i18n.NewBundle(language.English)
	// Step 2: Create localizer for that bundle using one or more language tags
	loc := i18n.NewLocalizer(bundle, language.English.String())
	// Step 3: Define messages
	messages := &i18n.Message{
		ID:          "Emails",
		Description: "The number of unread emails a user has",
		One:         "{{.Name}} has {{.Count}} email.",
		Other:       "{{.Name}} has {{.Count}} emails.",
	}
	// Step 3: Localize Messages
	messagesCount := 2
	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: messages,
		TemplateData: map[string]interface{}{
			"Name":  "Theo",
			"Count": messagesCount,
		},
		PluralCount: messagesCount,
	})

	fmt.Println(translation)

	// Define different delimiters
	messages = &i18n.Message{
		ID:          "Notifications",
		Description: "The number of unread notifications a user has",
		One:         "<<.Name>> has 一(<<.Count>>) notification.",
		Other:       "<<.Name>> has <<.Count>> notifications.",
		LeftDelim:   "<<",
		RightDelim:  ">>",
	}

	notificationsCount := 1
	translation = loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: messages,
		TemplateData: map[string]interface{}{
			"Name":  "Theo",
			"Count": notificationsCount,
		},
		PluralCount: notificationsCount,
	})

	fmt.Println(translation)

	// Unmarshaling from files
	//bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	//bundle.MustLoadMessageFile("en.json")
	//bundle.MustLoadMessageFile("zh.json")
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.MustLoadMessageFile("./locales/en-US/en.yaml")
	bundle.MustLoadMessageFile("./locales/zh-CN/zh.yaml")

	tag, err := locale.Detect()
	if err != nil {
		log.Fatal(err)
	}
	// Have fun with language.Tag!
	fmt.Println(tag)

	loc = i18n.NewLocalizer(bundle, tag.String())

	translation = loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "everything",
	})
	fmt.Println(translation)

	messagesCount = 1
	translation = loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "tdp.hello_world",
		TemplateData: map[string]interface{}{
			"Name":  "Daniel",
			"Count": messagesCount,
		},
		PluralCount: messagesCount,
	})

	fmt.Println(translation)

}

// https://phrase.com/blog/posts/internationalisation-in-go-with-go-i18n/
