package gdextension

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
	"github.com/mitchellh/mapstructure"
)

func MarshalDict(data any) (*gdapi.Dictionary, error) {
	dict := gdapi.NewDictionary()

	val := reflect.ValueOf(data)
	orig := val

	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s (%T)", orig.Kind(), orig.Interface())
	}

	for i := 0; i < val.NumField(); i++ {
		fieldT := val.Type().Field(i)
		key := fieldT.Name
		if fieldT.Tag.Get("json") != "" {
			key = strings.Split(fieldT.Tag.Get("json"), ",")[0]
		}
		if key == "-" {
			continue
		}
		field := val.Field(i)

		va, err := variantFromReflectWithFallback(field, func(val reflect.Value) (*gdapi.Variant, error) {
			if val.Kind() != reflect.Struct && (val.Kind() == reflect.Ptr && val.Elem().Kind() != reflect.Struct) {
				return nil, fmt.Errorf("expected struct (or ptr), got %s (%T)", val.Kind(), val.Interface())
			}
			dict, err := MarshalDict(val.Interface())
			if err != nil {
				return nil, fmt.Errorf("error marshalling dict for key %q: %w", key, err)
			}
			return dict.AsVariant(), nil
		})
		if err != nil {
			return nil, fmt.Errorf("error marshalling value for key %q: %w", key, err)
		}
		dict.Set(key, *va)
	}
	return dict, nil
}

func dictToMap(dict gdapi.Dictionary) (map[string]interface{}, error) {
	keys := dict.Keys()
	keysL := keys.Size()

	asMap := make(map[string]interface{}, keysL)

	for i := int64(0); i < keysL; i++ {
		keyVar := keys.Get(i)
		key := keyVar.String()

		value := dict.LookupWithVariant(keyVar)
		switch value.Type() {
		case gdc.VariantTypeDictionary:
			subDict, err := value.AsDictionary()
			if err != nil {
				return nil, fmt.Errorf("error fetching dict value for key %q: %w", key, err)
			}
			asMap[key], err = dictToMap(*subDict)
			if err != nil {
				return nil, fmt.Errorf("error unmarshalling dict for key %q: %w", key, err)
			}
		case gdc.VariantTypeArray:
			subArray, err := value.AsArray()
			if err != nil {
				return nil, fmt.Errorf("error fetching array value for key %q: %w", key, err)
			}
			asMap[key], err = gdapi.ConvertArrayToUnknownSlice(subArray)
			if err != nil {
				return nil, fmt.Errorf("error unmarshalling array for key %q: %w", key, err)
			}
		default:
			var err error
			asMap[key], err = gdapi.NativeFromVariant(value, nil)
			if err != nil {
				return nil, fmt.Errorf("error unmarshalling value for key %q: %w", key, err)
			}
		}
	}

	return asMap, nil
}

func UnmarshalDict(dict gdapi.Dictionary, data any) error {
	asMap, err := dictToMap(dict)
	if err != nil {
		return err
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  data,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(asMap)
}
