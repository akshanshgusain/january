package handlers

import (
	"context"
	"github.com/akshanshgusain/january"
	"net/http"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, tmpl string, variables, data interface{}) error {
	return h.App.TemplateEngine.Page(w, r, tmpl, variables, data)
}

func (h *Handlers) sessionPut(ctx context.Context, key string, val interface{}) {
	h.App.Session.Put(ctx, key, val)
}

func (h *Handlers) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) sessionGet(ctx context.Context, key string) interface{} {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) randomString(n int) string {
	return h.App.RandomString(n)
}

func (h *Handlers) encrypt(text string) (string, error) {
	e := january.Encryption{Key: []byte(h.App.EncryptionKey)}

	encrypted, err := e.Encrypt(text)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (h *Handlers) decrypt(cypher string) (string, error) {
	e := january.Encryption{Key: []byte(h.App.EncryptionKey)}

	encrypted, err := e.Decrypt(cypher)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}
