/*
Package i18n implements i18n with simple config file or custom backends.

Files are simple config files based on https://github.com/pilu/config.
You can load a file with i18n.Load:

    err := i18n.Load("en.conf")

You can also load multiple files passing a glob pattern:

    i18n.Load("/path/to/locals/*.conf")

In the first example, all the translations of "en.conf" are added to the locale "en",
which is the file name without extension.

You can also create a file with a different name and add all your translations under
the "en" section:

  [en]
  greeting: hello

Inside a file you can have multiple sections. Each section will be a different locale:

  // file en.conf

  greeting: hello

  [it]
  greeting: ciao

  [es]
  greeting: hola

All top level translations are added to the locale "en", which is the file name.

Check the example folder for a complete example https://github.com/pilu/i18n-go/tree/master/example
*/
package i18n

