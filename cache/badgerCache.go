package cache

import "github.com/dgraph-io/badger/v3"

// TODO: implement embedded-cache, (Optional)

type BadgerCache struct {
	Conn   *badger.DB
	Prefix string
}
