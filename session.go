// Package session is a wrapper for http://www.gorillatoolkit.org/pkg/sessions.
//
// Login:
//
//  func LoginHandlerFunc(w http.ResponseWriter, r *http.Request) {
//    err := session.Set(w, r, "name", "john")
//    if err != nil {
//      panic(err)
//    }
//    http.Redirect(w, r, "/", http.StatusSeeOther)
//  }
//
// Logout:
//
//  func LogoutHandlerFunc(w http.ResponseWriter, r *http.Request) {
//    err := session.Destroy(w, r)
//    if err != nil {
//      panic(err)
//    }
//    http.Redirect(w, r, "/", http.StatusSeeOther)
//  }
package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Make session secret, cookie store and session name configurable
var (
	Secret = "something-very-secret"
	Store  = sessions.NewCookieStore([]byte(Secret))
	Name   = "session"
)

// Set creates a new session with given key and value.
func Set(w http.ResponseWriter, r *http.Request, key, value interface{}) error {
	sess, err := Store.Get(r, Name)
	if err != nil {
		return err
	}
	sess.Values[key] = value
	return sess.Save(r, w)
}

// Get returns the session value for given key.
func Get(r *http.Request, key string) (interface{}, error) {
	sess, err := Store.Get(r, Name)
	if err != nil {
		return nil, err
	}
	return sess.Values[key], nil
}

// Destroy removes session cookie by setting MaxAge to a value in the past.
func Destroy(w http.ResponseWriter, r *http.Request) error {
	sess, err := Store.Get(r, Name)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	return sess.Save(r, w)
}
