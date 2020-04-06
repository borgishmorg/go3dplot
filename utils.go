package go3dplot

import "os"

func writeToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return file.Close()
}

func removeFile(filename string) error {
	return os.Remove(filename)
}
