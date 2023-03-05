package january

import "os"

// CreateDirIfNotExist TODO: create dirs if not exist

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
