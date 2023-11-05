package config

import (
	"testing"
)

func TestGet(t *testing.T) {

	// Create a ConfigObject
	c := &ConfigObject{
		values: make(map[string]map[string]interface{}),
	}

	c.Set("app", "config_dir", "../tests/stubs")

	// Call the Get method
	value := c.Get("test_config", "foo")

	// Check the returned value
	if value != "bar" {
		t.Errorf("Get returned wrong value: got %v want %v", value, "value")
	}
}

func TestSet(t *testing.T) {
	// Create a ConfigObject
	c := &ConfigObject{
		values: make(map[string]map[string]interface{}),
	}

	// Call the Set method
	c.Set("testPath", "testKey", "testValue")

	// Check the value was set correctly
	if c.values["testPath"]["testKey"] != "testValue" {
		t.Errorf("Set did not set the correct value: got %v want %v", c.values["testPath"]["testKey"], "testValue")
	}
}
