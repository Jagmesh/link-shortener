package request

import (
	"encoding/json"
	"fmt"
	"io"
	"link-shortener/pkg/logger"
	"link-shortener/pkg/validation"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

var log = logger.GetWithScopes("REQUEST")

func GetBody[T any](r *http.Request) (T, error) {
	body, err := jsonDecode[T](r.Body)
	if err != nil {
		return *new(T), err
	}

	if err := validation.IsValidStruct[T](&body); err != nil {
		return *new(T), err
	}

	return body, nil
}

func GetQueryParams[T any](r *http.Request) (*T, error) {
	q := r.URL.Query()

	var payload T

	err := mapFieldsToStruct(&payload, q)
	if err != nil {
		log.Errorln("GetQueryParams err:", err)
		return nil, err
	}
	if err := validation.IsValidStruct[T](&payload); err != nil {
		log.Errorln("GetQueryParams err:", err)
		return nil, err
	}

	return &payload, nil
}

func mapFieldsToStruct(obj any, queryValMap url.Values) error {
	for key, queryValues := range queryValMap {
		if len(queryValues) == 0 {
			return nil
		}
		if len(queryValues) > 1 {
			return fmt.Errorf("more than 1 (%d) value in key '%s'", len(queryValues), key)
		}

		if err := setField(obj, key, queryValues[0]); err != nil {
			return err
		}
	}
	return nil
}

func setField(obj any, key string, val string) error {
	structElemVal := reflect.ValueOf(obj).Elem().FieldByNameFunc(func(s string) bool {
		return strings.EqualFold(s, key)
	})
	queryValue := reflect.ValueOf(val)

	if !structElemVal.IsValid() {
		return fmt.Errorf("key '%s' not exists", key)
	}
	if !structElemVal.CanSet() {
		return fmt.Errorf("cannot set value of key '%s'", key)
	}
	if !structElemVal.Type().ConvertibleTo(queryValue.Type()) {
		return fmt.Errorf("value of type '%T' is not assignable to type '%T'", queryValue, key)
	}

	structElemVal.Set(queryValue.Convert(structElemVal.Type()))
	return nil
}

func jsonDecode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
