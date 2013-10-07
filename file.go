package i18n

import (
  "path/filepath"
  "github.com/pilu/config"
)

func loadFile(filePath string) error {
  baseName        := filepath.Base(filePath)
  extension       := filepath.Ext(baseName)
  mainSectionName := baseName[:len(baseName) - len(extension)]

  sections, err := config.ParseFile(filePath, mainSectionName)
  if err != nil {
    return err
  }

  for section, options := range sections {
    if len(options) > 0 {
      locale := GetLocale(section)
      if locale == nil {
        locale = NewLocale(section)
      }
      for key, value := range options {
        locale.Add(key, value)
      }
      AddLocale(locale)
    }
  }

  return nil
}

func Load(globPattern string) error {
  paths, err := filepath.Glob(globPattern)
  if err != nil {
    return nil
  }

  for _, filePath := range paths {
    err := loadFile(filePath)
    if err != nil {
      return err
    }
  }

  return nil
}

