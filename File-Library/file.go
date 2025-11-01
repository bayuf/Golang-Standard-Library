package main

import "os"

func createNewFile(name, message string) error {
	// membuat file baru jika belum ada, write only, chmod 666
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	// return jika error
	if err != nil {
		return err
	}

	// close file setelah dibuat
	defer file.Close()

	// menulis isi dari file
	file.WriteString(message)

	return nil
}
func addToFile(name, message string) error {
	// membuat file baru jika belum ada, write only, chmod 666
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)

	// return jika error
	if err != nil {
		return err
	}

	// close file setelah dibuat
	defer file.Close()

	// menulis isi dari file
	file.WriteString(message)

	return nil
}

func main() {
	createNewFile("teks.txt", "Hallo ini merupakan file pertamaku.")
	addToFile("teks.txt", "\nHallo ini merupakan teks keduaku.")
}
