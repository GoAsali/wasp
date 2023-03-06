package make

import (
	"fmt"
	"log"
	"os"
)

type Make struct {
	folder     string
	ModuleType string
}

func CreateNewMakeService(folder string, moduleType string) *Make {
	return &Make{
		folder:     folder,
		ModuleType: moduleType,
	}
}

func (make Make) checkDirectoryExists() error {
	if _, err := os.Stat(make.folder); !os.IsNotExist(err) {
		return nil
	}
	err := os.Mkdir(make.folder, 0755)
	if err == nil {
		return nil
	}
	if os.IsPermission(err) {
		return fmt.Errorf("you need permission to current folder to create %s", make.ModuleType)
	}
	return fmt.Errorf("Could not create " + make.ModuleType)
}

func (make Make) getFilePath(fileName string) (string, error) {
	filePath := fmt.Sprintf("%s/%s.go", make.folder, fileName)
	if _, err := os.Stat(filePath); err != nil {
		return filePath, nil
	}
	return filePath, fmt.Errorf("%s %s already exists\n", make.ModuleType, fileName)
}

func (make Make) createFile(name string, content string) (string, error) {
	err := make.checkDirectoryExists()
	if err != nil {
		return "", err
	}

	filePath, fileError := make.getFilePath(name)

	if fileError != nil {
		return "", fileError
	}

	file, e := os.Create(filePath)
	if e != nil {
		log.Fatal(e)
		return "", nil
	}

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}
	err = file.Close()
	if err != nil {
		return "", fmt.Errorf("error during writing template")
	}

	return fmt.Sprintf("%s created successfully\n", name), nil
}
