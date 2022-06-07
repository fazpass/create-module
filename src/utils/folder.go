package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

//GetFolderName ...
func GetFolderName(packageName string) string {
	splits := strings.Split(packageName, "/")
	return splits[len(splits)-1]
}

// create sub folder
func CreateSubFolder(projectDirectory []string, rootDirectory string) error {
	for _, p := range projectDirectory {
		err := CreateFolder(rootDirectory + "/" + p)
		if err != nil {
			return err
		}
	}
	return nil
}

//CreateFolder ...
func CreateFolder(folderName string) error {
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

//CreateFile ...
func CreateFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}
