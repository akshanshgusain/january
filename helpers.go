package january

import (
	"crypto/rand"
	"os"
)

// CreateDirIfNotExist TODO: create dirs if not exist

const randomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

func (j *January) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomString)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

func (j *January) CreateDirIfNotExist(path string) error {
	// default permission : The folders will be readable and executed by others, but writable by the user only
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, mode); err != nil {
			return err
		}
	}
	return nil
}

func (j *January) CreateFileIfNotExist(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}
