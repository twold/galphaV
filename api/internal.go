// galphaV api package contains primary application interface.
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/twold/galphaV/client"
)

func dayOfWeek(input string) (*string, error) {
	// add day of week
	date, err := time.Parse("2006-01-02", input)
	if err != nil {
		return nil, err
	}
	v := date.Weekday().String()
	return &v, nil
}

func get(f client.Builder) (*Response, error) {
	resp, err := f.Do("GET")
	if err != nil {
		return nil, err
	}
	return &Response{r: resp}, nil
}

func marshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func read(r *Response) ([]byte, error) {
	b, err := ioutil.ReadAll(r.r.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func readFile(path, fileName string) ([]byte, error) {
	b, err := ioutil.ReadFile(filepath.Join(path, "input", fileName))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshal(byt []byte, obj interface{}) error {
	return json.Unmarshal(byt, obj)
}

func writeFile(date, path, symbol string, obj interface{}) error {
	err := os.Mkdir(filepath.Join(path, "output", symbol), 0777)
	if err != nil {
		if check := strings.Contains(err.Error(), "file exists"); check == false {
			return err
		}
	}

	b, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		return err
	}

	name := filepath.Join(path, "output", symbol, fmt.Sprintf("%v.json", date))
	if _, err := os.Stat(name); os.IsNotExist(err) == false {
		return nil
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}
