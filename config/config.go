package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ConfigObject struct {
	values map[string]map[string]interface{}
}

func (c *ConfigObject) Get(path string, key string) any {
	if c.values == nil {
		c.values = make(map[string]map[string]interface{})
	}
	// if the path is not fount on the c.values map, load it from the json file
	if c.values[path] == nil {
		// default config directory
		configDir := "config"
		if c.values["app"]["config_dir"] != nil {
			configDir = c.values["app"]["config_dir"].(string)

		}

		jsonFile, err := os.Open(fmt.Sprintf("%v/%v.json", configDir, path))

		// if we os.Open returns an error then handle it
		if err != nil {
			// TODO should I panic here?
			fmt.Println(err)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		c.values[path] = result
	}

	return c.values[path][key]
}

func (c *ConfigObject) Set(path string, key string, value any) {
	if c.values == nil {
		c.values = make(map[string]map[string]interface{})
	}

	if c.values[path] == nil {
		c.values[path] = make(map[string]interface{})
	}

	c.values[path][key] = value
}
