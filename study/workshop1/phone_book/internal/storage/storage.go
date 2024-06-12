package storage

import (
	"os"
	"errors"
)

type Storage struct {
	fileName string
}

// NewStorage .. TODO сделать описание функции
func NewStorage(name string) Storage {
	return Storage{fileName: name}
}

func (S Storage) addContact(telephone models.Telephone) error {
	if _, err := os.Stat(s.fileName); errors.Is(err, os.ErrNotExist) {
		// создаем файл
		if errCreateFile := s.createFile(); errCreateFile != nil {
			return errCreateFile
		}
	}

	// прочитать
	b, err := os.ReadFile(s.fileName)
	if err != nil {
		return err
	}

	var records []telephoneRecord
	if errUnmarshal := json.Unmarshal(b, &records); err != nil {
		return errUnmarshal
	}
	records = append(records, transform())

	bWrite, errMarshal := json.Indent(records)
	if errMarshal != nil {
		return errMarshal
	}

	return os.WriteFile(s.fileName, bWrite, 0666)
}

func (s Storage) createFile() error {
	f, err := os.Create(s.fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}