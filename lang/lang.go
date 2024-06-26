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

// Localizer is implemented by any value that has a Localize method, which, similarly to
// [fmt.Stringer], defines a readable form of that value. In addition to String, Localize takes an
// additional parameter langTag which is a [language.Tag] defining in which language to localize to.
//
// When implementing Localizer in a type, the String method should be just a call to Localize with a
// default language, e.g. [language.English], for the langTag parameter.
type Localizer interface {
	Localize(langTag language.Tag) string
}

func Localize(key string, langTag language.Tag) (string, error) {
	localized, err := i18n.
		NewLocalizer(bundle, langTag.String()).
		Localize(&i18n.LocalizeConfig{MessageID: key})
	if _, ok := err.(*i18n.MessageNotFoundErr); ok {
		localized = key
		err = nil
	}
	return localized, err
}

func MustLocalize(key string, langTag language.Tag) string {
	localized, err := Localize(key, langTag)
	if err != nil {
		panic(err)
	}
	return localized
}

// LocalizeInterface tries to localize v if it's type implements [Localizer]. Otherwise it returns v
// in the default formatting (using %v).
func LocalizeInterface(v any, langTag language.Tag) string {
	if localizer, ok := v.(Localizer); ok {
		return localizer.Localize(langTag)
	}
	return fmt.Sprintf("%v", v)
}

// AllLangs returns a slice of all loaded languages.
func AllLangs() []language.Tag {
	return bundle.LanguageTags()
}
