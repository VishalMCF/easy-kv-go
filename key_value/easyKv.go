package key_value

import (
	"errors"
	"fmt"
)

// ADD COMMAND
func add(key string, value interface{}) error {
	data := DATA
	if data == nil || value == nil || key == "" {
		return errors.New("Something happened to data store it might not be initilaised")
	}
	data[key] = value
	return nil
}

// UPDATE COMMAND
func update(key string, value interface{}) (interface{}, error) {
	data := DATA
	if data == nil || key == "" {
		return nil, errors.New("either key is not present or the data store is not initialised yet")
	}
	data[key] = value
	return value, nil
}

// GET COMMAND
func get(key string) (interface{}, error) {
	data := DATA
	if data == nil || key == "" {
		return nil, errors.New("either key is not present or the data store is not initialised yet")
	}
	value := data[key]
	return value, nil
}

// DELETE COMMAND
func deleteKey(key string) (string, error) {
	data := DATA
	if data == nil || key == "" {
		return "", errors.New("either key is not present or the data store is not initialised yet")
	}
	delete(data, key)
	return "DELETED !!!", nil
}

func print() {
	data := DATA
	if data != nil {
		for k, d := range DATA {
			fmt.Printf("key:%s ....%....... ", k, d)
		}
	}
}
