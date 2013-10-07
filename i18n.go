package i18n

import (
  "fmt"
  "errors"
)

type LocaleMissing struct{
  Code string
}

func (m LocaleMissing) Error() string {
  return fmt.Sprintf("Missing `%s` locale", m.Code)
}

var (
  locales       map[string]*Locale
  currentLocale *Locale
)

func init() {
  resetLocales()
}

func resetLocales() {
  locales = make(map[string]*Locale)
}

func AddLocale(locale *Locale) {
  locales[locale.Code] = locale
}

func GetLocale(code string) *Locale {
  return locales[code]
}

func SetLocale(code string) error {
  currentLocale = GetLocale(code)

  if currentLocale == nil {
    return &LocaleMissing{ code }
  }

  return nil
}

func CurrentLocale() *Locale {
  return currentLocale
}

func Translate(key string, args ...interface{}) (string, error) {
  if currentLocale != nil {
    return currentLocale.Translate(key, args...)
  }

  return "", errors.New("Current locale is not set")
}
