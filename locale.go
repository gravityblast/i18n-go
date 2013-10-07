package i18n

import (
  "fmt"
)

type TranslationMissing struct {
  Key string
}

func (t TranslationMissing) Error() string {
  return fmt.Sprintf("Translation missing for key `%s`", t.Key)
}

type Locale struct {
  Code         string
  translations map[string]*Translation
}

func (locale *Locale) Add(key, value string) *Translation {
  translation := NewTranslation(key, value)
  locale.translations[translation.Key()] = translation

  return translation
}

func (locale *Locale) Translate(key string, args ...interface{}) (string, error) {
  var value string
  translation := locale.translations[key]

  if translation == nil {
    return value, &TranslationMissing{ key }
  }

  return translation.Format(args...), nil
}

func (locale *Locale) T(key string, args ...interface{}) (string, error) {
  return locale.Translate(key, args...)
}

func NewLocale(code string) *Locale {
  locale := &Locale{
    Code: code,
    translations: make(map[string]*Translation),
  }

  return locale
}
