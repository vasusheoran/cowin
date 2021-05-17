package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"text/template"
)

var lock sync.Mutex

const layout = "02-01-2006"

func ParseTemplateString(templateString string, placeholders map[string]interface{}) (string, error) {
	t, err := template.New("temp").Parse(templateString)
	if err != nil {
		return "", err
	}

	var parsedData bytes.Buffer
	if err := t.Execute(&parsedData, placeholders); err != nil {
		return "", err
	}

	value := parsedData.String()
	return value, nil
}

func MakeDirectoryIfNotExists(paths ...string) error {
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return os.Mkdir(path, os.ModeDir|0755)
		}
	}
	return nil
}

//
//
func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	r := bytes.NewReader(b)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	return json.Unmarshal(byteValue, v)
}

func GetDistrictFromFileName(fileName string) (int, error) {
	district := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return strconv.Atoi(district)
}
