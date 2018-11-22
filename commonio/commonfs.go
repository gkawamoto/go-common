package commonio

import (
	"io/ioutil"
	"os"
)

// FileExists returns a simple boolean to check if the file exists.
func FileExists(location string) bool {
	var err error
	_, err = os.Stat(location)
	return !os.IsNotExist(err)
}

// WriteFile writes to the specified file, keeping the file mode or assuming 0777 if the file does not exist.
func WriteFile(location string, data []byte) error {
	var mode os.FileMode
	if FileExists(location) {
		var stat os.FileInfo
		var err error
		stat, err = os.Stat(location)
		if err != nil {
			return err
		}
		mode = stat.Mode()
	} else {
		mode = 0777
	}
	return ioutil.WriteFile(location, data, mode)
}

// WriteTempFile writes the contents to a temporary file and returns the path.
func WriteTempFile(data []byte) (string, error) {
	var f *os.File
	var err error
	f, err = ioutil.TempFile("", "")
	if err != nil {
		return f.Name(), err
	}
	err = ioutil.WriteFile(f.Name(), data, 0777)
	if err != nil {
		return f.Name(), err
	}
	return f.Name(), nil
}

// ReadFile reads the specified file (specified in this package to keep company to the WriteFile method).
func ReadFile(location string) ([]byte, error) {
	return ioutil.ReadFile(location)
}
