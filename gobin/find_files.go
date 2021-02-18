package gobin

import (
	"io/ioutil"
	"sort"
)

func GetDir(path string) ([]string, error) {
	result := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return result, err
	}

	filesList := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			dirContent, err := GetDir(path + "/" + f.Name())
			if err == nil {
				result = append(result, dirContent...)
			}
		} else {
			filesList = append(filesList, path+"/"+f.Name())
		}
	}

	sort.Slice(filesList, func(i, j int) bool {
		return filesList[i] < filesList[j]
	})

	for _, f := range filesList {
		result = append(result, f)
	}

	return result, nil
}
