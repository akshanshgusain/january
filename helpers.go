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
