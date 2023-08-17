package january

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"reflect"
	"testing"
)

func TestSession_InitSession(t *testing.T) {
	c := &Session{
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "january",
		CookieDomain:   "127.0.0.1",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager
	ses := c.InitSession()

	var sessKind reflect.Kind
	var sessType reflect.Type

	rv := reflect.ValueOf(ses)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("loop: ", rv.Kind(), rv.Type(), rv)
		sessType = rv.Type()
		sessKind = rv.Kind()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("invalid type or kind; kind:", rv.Kind(), " type: ", rv.Type())
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong kind returned testing cookie session. expected: ", reflect.ValueOf(sm).Kind(), "and got ", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("wrong type returned testing cookie session. expected: ", reflect.ValueOf(sm).Type(), "and got ", sessType)
	}
}
