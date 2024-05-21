package lang

import (
	"fmt"
	"os"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var (
	bundle *i18n.Bundle
)

func init() {
	bundle = i18n.NewBundle(language.AmericanEnglish)
	const format = "yaml"
	bundle.RegisterUnmarshalFunc(format, yaml.Unmarshal)
	loadFiles("lang/", format)
}

func loadFiles(path string, format string) {
	items, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Read %d files: %v\n", len(items), err)
		return
	}
	for _, item := range items {
		if !strings.HasSuffix(item.Name(), "."+format) {
			return
		}
		_, err = bundle.LoadMessageFile(strings.TrimRight(path, "/\\") + string(os.PathSeparator) + item.Name())
		if err != nil {
			fmt.Printf("Failed to load language file '%s': %v\n", item.Name(), err)
		}
	}
}

func Localize(key string, lang language.Tag) (string, error) {
	localized, err := i18n.
		NewLocalizer(bundle, lang.String()).
		Localize(&i18n.LocalizeConfig{MessageID: key})
	if _, ok := err.(*i18n.MessageNotFoundErr); ok {
		localized = key
		err = nil
	}
	return localized, err
}

func MustLocalize(key string, lang language.Tag) string {
	localized, err := Localize(key, lang)
	if err != nil {
		panic(err)
	}
	return localized
}
