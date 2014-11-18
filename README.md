
# session

[![Build Status](https://travis-ci.org/zemirco/session.svg)](https://travis-ci.org/zemirco/session)
[![GoDoc](https://godoc.org/github.com/zemirco/session?status.svg)](https://godoc.org/github.com/zemirco/session)

Small wrapper for [gorilla/sessions](http://www.gorillatoolkit.org/pkg/sessions).

## Example

```go
package main

import "github.com/zemirco/session"

// login
func LoginHandlerFunc(w http.ResponseWriter, r *http.Request) {
  err := session.Set(w, r, "name", "john")
  if err != nil {
    panic(err)
  }
  http.Redirect(w, r, "/", http.StatusSeeOther)
}

// logout
func LogoutHandlerFunc(w http.ResponseWriter, r *http.Request) {
  err := session.Destroy(w, r)
  if err != nil {
    panic(err)
  }
  http.Redirect(w, r, "/", http.StatusSeeOther)
}
```

## Test

`go test`

## License

MIT
