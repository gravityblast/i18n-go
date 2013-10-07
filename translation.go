package i18n

import (
  "fmt"
)

type Translation [2]string

func (t *Translation) Key() string {
  return t[0]
}

func (t *Translation) Value() string {
  return t[1]
}

func (t *Translation) Format(args ...interface{}) string {
  return fmt.Sprintf(t.Value(), args...)
}

func NewTranslation(key, value string) *Translation {
  return &Translation{key, value}
}

