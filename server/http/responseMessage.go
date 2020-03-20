package server

import (
	"reflect"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

const (
	SUCCESS_CODE = 0
)

func SuccessResponseWithPayload(w rest.ResponseWriter, payload interface{}) {
	WriteErrorResponse(w, 0, "", payload)
}

func SuccessResponse(w rest.ResponseWriter) {
	WriteErrorResponse(w, 0, "", nil)
}

func ErrorResponse(w rest.ResponseWriter, errorCode int, errorMsg string) {
	WriteErrorResponse(w, errorCode, errorMsg, nil)
}

func WriteErrorResponse(w rest.ResponseWriter, errorCode int, errorMsg string, payload interface{}) {
	if errorCode != SUCCESS_CODE {
		w.WriteHeader(errorCode)
	}

	mp := make(map[string]interface{})
	mp["errorCode"] = strconv.Itoa(errorCode)
	mp["errorMessage"] = errorMsg

	if payload != nil {

		t := reflect.TypeOf(payload)
		v := reflect.ValueOf(payload)

		if v.Kind() == reflect.Struct {

			for i := 0; i < v.NumField(); i++ {
				if v.Field(i).CanInterface() {

					tagVal := t.Field(i).Tag.Get("json")
					if tagVal != "" {
						mp[tagVal] = v.Field(i).Interface()
					} else {
						mp[t.Field(i).Name] = v.Field(i).Interface()
					}
				}
			}

		}
	}

	err := w.WriteJson(mp)
	if err != nil {
		panic(err)
	}
	return

}
