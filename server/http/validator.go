package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func validate(jsonMap map[string]interface{}, anyStruct interface{}) (bool, error) {
	t := reflect.TypeOf(anyStruct)
	v := reflect.ValueOf(anyStruct)

	// if anyStruct is a pointer
	if v.Kind() == reflect.Ptr {
		t = reflect.TypeOf(reflect.Indirect(v).Interface())
		v = reflect.ValueOf(reflect.Indirect(v).Interface())
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanInterface() {
			jsonTagVal := t.Field(i).Tag.Get("json")
			tagVal := t.Field(i).Tag.Get("validate")
			if jsonTagVal == "" || tagVal == "" {
				continue
			}

			if tagVal == "required" {
				if field.Type().Kind() != reflect.Struct {
					continue
				}

				if jsonMap[jsonTagVal] == nil {
					return false, fmt.Errorf("Failed to find required field [%s]", jsonTagVal)
				}
				ok, err := validate(jsonMap[jsonTagVal].(map[string]interface{}), field.Interface())
				if !ok || err != nil {
					return ok, err
				}
			} else if tagVal == "notBlank" {
				if field.Type().Kind() == reflect.String {
					if jsonMap[jsonTagVal] == nil {
						return false, fmt.Errorf("Failed to find required field [%s]", jsonTagVal)
					}

					if isBlankString(jsonMap[jsonTagVal].(string)) {
						return false, fmt.Errorf("[%s] should not be blank", jsonTagVal)
					}
				} else if isIntType(field) {
					if jsonMap[jsonTagVal] == nil {
						return false, fmt.Errorf("Failed to find required field [%s]", jsonTagVal)
					}
				}
			} else if strings.HasPrefix(tagVal, "conditional") {
				//TODO:
			}
		}
	}

	return true, nil
}

// ValidateInput performs a simple validation on an input(json format) based on struct tags:
// "validate" "required" "notBlank" "conditional"
func ValidateInput(jsonString []byte, anyStruct interface{}) (bool, error) {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(jsonString, &jsonMap)
	if err != nil {
		return false, errors.New("failed to parse json string")
	}

	return validate(jsonMap, anyStruct)
}

//ValidateAndDecodeRequestPayload performs input content validation of request body (application/jason), and then decode it.
func ValidateAndDecodeRequestPayload(request http.Request, anyStructPointer interface{}) error {
	respBody, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return errors.New("failed to read request body")
	}

	ok, err := ValidateInput(respBody, anyStructPointer)
	if !ok || err != nil {
		return err
	}

	err = json.Unmarshal(respBody, anyStructPointer)
	if err != nil {
		return err
	}

	return nil
}

func isBlankString(str string) bool {
	return str == ""
}

func isIntType(field reflect.Value) bool {
	kind := field.Type().Kind()
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}
