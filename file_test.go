package i18n

import (
  "os"
  "path"
  "testing"
  "path/filepath"
  assert "github.com/pilu/miniassert"
)

func withTestData(files map[string]string, callback func(string)) {
  folderName := "test_locales"
  err := os.Mkdir(folderName, 0755)
  if err != nil {
    panic(err)
  }
  defer func() {
    os.RemoveAll(folderName)
  }()

  for fileName, content := range files {
    filePath := path.Join(folderName, fileName)
    file, err := os.Create(filePath)
    if err != nil {
      panic(err)
    }

    file.WriteString(content)
  }

  callback(folderName)
}

func TestLoadFile(t *testing.T) {
  defer resetLocales()

  files := map[string]string{
    "en.conf": `
    greeting: hello

    [en_GB]
    greeting: hello GB

    [en_US]
    greeting: hello US

    [en_GB]
    greeting2: hello 2 GB
    `,
  }

  withTestData(files, func(folderPath string) {
    err := LoadFile(filepath.Join(folderPath, "en.conf"))
    assert.Nil(t, err)
    assert.Equal(t, 3, len(locales))

    tests := [][]string {
      []string{"en",    "greeting",   "hello"},
      []string{"en_GB", "greeting",   "hello GB"},
      []string{"en_GB", "greeting2",  "hello 2 GB"},
      []string{"en_US", "greeting",   "hello US"},
    }

    for _, items := range tests {
      code   := items[0]
      key    := items[1]
      value  := items[2]
      locale := GetLocale(code)

      assert.NotNil(t, locale)
      assert.Type(t, "*i18n.Locale", locale)
      assert.Equal(t, code, locale.Code)

      tr := locale.translations[key]
      assert.Equal(t, value, tr.Value())
    }
  })
}

func TestLoadFiles(t *testing.T) {
  defer resetLocales()

  files := map[string]string{
    "en.conf": `
    greeting: hello
    `,
    "it.conf": `
    greeting: ciao
    `,
    "other_en.conf": `
    [en] // add more translations for locale "en"
    greeting2: hello2
    `,
    "other_it.conf": `
    [it] // add more translations for locale "it"
    greeting2: ciao2
    `,
    "es.conf": `// blank file`,
  }

  withTestData(files, func(folderPath string) {
    err := LoadFiles(filepath.Join(folderPath, "*.conf"))
    assert.Nil(t, err)
    assert.Equal(t, 2, len(locales))

    tests := [][]string {
      []string{"en", "greeting",  "hello"},
      []string{"en", "greeting2", "hello2"},
      []string{"it", "greeting",  "ciao"},
      []string{"it", "greeting2", "ciao2"},
    }

    for _, items := range tests {
      code   := items[0]
      key    := items[1]
      value  := items[2]
      locale := GetLocale(code)

      assert.NotNil(t, locale)
      assert.Type(t, "*i18n.Locale", locale)
      assert.Equal(t, code, locale.Code)

      tr := locale.translations[key]
      assert.Equal(t, value, tr.Value())
    }
  })
}
