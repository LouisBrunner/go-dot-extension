package gdextension

import (
	"fmt"
	"reflect"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
	"github.com/mitchellh/mapstructure"
)

func MarshalDict(data any) (*gdapi.Dictionary, error) {
	asMap := make(map[string]interface{})

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &asMap,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(data)
	if err != nil {
		return nil, err
	}

	dict := gdapi.NewDictionary()
	for key, value := range asMap {
		va, err := variantFromReflect(reflect.ValueOf(value))
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
