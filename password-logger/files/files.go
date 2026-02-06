package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile("test.txt")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println("error:", err)
	}

	_, err = file.Write(content)

	defer file.Close()

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("success writing")
}
