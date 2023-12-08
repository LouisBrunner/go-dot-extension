package gencommon

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type ExtensionHeader struct {
	VersionMajor    int    `json:"version_major"`
	VersionMinor    int    `json:"version_minor"`
	VersionPatch    int    `json:"version_patch"`
	VersionStatus   string `json:"version_status"`
	VersionBuild    string `json:"version_build"`
	VersionFullName string `json:"version_full_name"`
}

type ExtensionClassSizeEntry struct {
	Name string `json:"name"`
	Size uint   `json:"size"`
}

type ExtensionClassSize struct {
	BuildConfiguration string                    `json:"build_configuration"`
	Sizes              []ExtensionClassSizeEntry `json:"sizes"`
}

type ExtensionClassMemberOffsetEntry struct {
	Member string `json:"member"`
	Offset uint   `json:"offset"`
	Meta   string `json:"meta"`
}

type ExtensionClassMemberOffset struct {
	Name    string                            `json:"name"`
	Members []ExtensionClassMemberOffsetEntry `json:"members"`
}

type ExtensionClassMemberOffsetConfig struct {
	BuildConfiguration string                       `json:"build_configuration"`
	Classes            []ExtensionClassMemberOffset `json:"classes"`
}

type ExtensionEnumValue struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ExtensionEnum struct {
	Name       string               `json:"name"`
	IsBitfield *bool                `json:"is_bitfield,omitempty"`
	Values     []ExtensionEnumValue `json:"values"`
}

type ExtensionNameType struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Meta    string `json:"meta,omitempty"`
	Default string `json:"default_value,omitempty"`
}

type ExtensionCommonFunction struct {
	Name       string              `json:"name"`
	ReturnType string              `json:"return_type,omitempty"`
	IsVarArg   bool                `json:"is_vararg"`
	Hash       uint                `json:"hash,omitempty"`
	Arguments  []ExtensionNameType `json:"arguments,omitempty"`
}

type ExtensionFunction struct {
	ExtensionCommonFunction
	Category string `json:"category"`
}

type ExtensionBuiltinClassOperator struct {
	Name       string `json:"name"`
	RightType  string `json:"right_type,omitempty"`
	ReturnType string `json:"return_type"`
}

type ExtensionBuiltinClassConstructor struct {
	Index     int                 `json:"index"`
	Arguments []ExtensionNameType `json:"arguments,omitempty"`
}

type ExtensionBuiltinClassMethod struct {
	ExtensionCommonFunction
	IsConst  bool `json:"is_const"`
	IsStatic bool `json:"is_static"`
}

type ExtensionClassConstant[T any] struct {
	Name  string `json:"name"`
	Type  string `json:"type,omitempty"`
	Value T      `json:"value"`
}

type ExtensionBuiltinClass struct {
	Name               string                             `json:"name"`
	IndexingReturnType string                             `json:"indexing_return_type,omitempty"`
	IsKeyed            bool                               `json:"is_keyed"`
	HasDestructor      bool                               `json:"has_destructor"`
	Enums              []ExtensionEnum                    `json:"enums,omitempty"`
	Methods            []ExtensionBuiltinClassMethod      `json:"methods,omitempty"`
	Operators          []ExtensionBuiltinClassOperator    `json:"operators"`
	Constructors       []ExtensionBuiltinClassConstructor `json:"constructors"`
	Members            []ExtensionNameType                `json:"members,omitempty"`
	Constants          []ExtensionClassConstant[string]   `json:"constants,omitempty"`
}

type ExtensionSignal struct {
	Name      string              `json:"name"`
	Arguments []ExtensionNameType `json:"arguments,omitempty"`
}

type ExtensionReturnValue struct {
	Type string `json:"type"`
	Meta string `json:"meta,omitempty"`
}

type ExtensionClassMethod struct {
	ExtensionBuiltinClassMethod
	IsVirtual   bool                  `json:"is_virtual"`
	ReturnValue *ExtensionReturnValue `json:"return_value,omitempty"`
}

type ExtensionClassProperty struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Getter string `json:"getter"`
	Setter string `json:"setter,omitempty"`
	Index  *uint  `json:"index,omitempty"`
}

type ExtensionClass struct {
	Name           string                        `json:"name"`
	IsRefCounted   bool                          `json:"is_refcounted"`
	IsInstantiable bool                          `json:"is_instantiable"`
	Inherits       string                        `json:"inherits,omitempty"`
	APIType        string                        `json:"api_type"`
	Methods        []ExtensionClassMethod        `json:"methods,omitempty"`
	Enums          []ExtensionEnum               `json:"enums,omitempty"`
	Properties     []ExtensionClassProperty      `json:"properties,omitempty"`
	Signals        []ExtensionSignal             `json:"signals,omitempty"`
	Constants      []ExtensionClassConstant[int] `json:"constants,omitempty"`
}

type ExtensionNativeFormat struct {
	Fields []ExtensionNameType `json:"fields"`
	raw    string              `json:"-"`
}

var nativeFormatRegexp = regexp.MustCompile(`^([a-zA-Z0-9_:*]+\s+\*?)([a-zA-Z0-9_]+)(\[[0-9]+\])?(?:\s*=\s*(.+))?$`)

func (f *ExtensionNativeFormat) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &f.raw)
	if err != nil {
		return err
	}
	parts := strings.Split(f.raw, ";")
	fields := make([]ExtensionNameType, len(parts))
	for i, part := range parts {
		matches := nativeFormatRegexp.FindStringSubmatch(part)
		if matches == nil {
			return fmt.Errorf("invalid native format: %s", part)
		}
		name := strings.TrimSpace(matches[2])
		typ := strings.TrimSpace(matches[1])
		defaultValue := ""
		if len(matches) > 3 {
			typ += matches[3]
		}
		if len(matches) > 4 {
			defaultValue = matches[4]
		}
		fields[i] = ExtensionNameType{
			Name:    name,
			Type:    typ,
			Default: defaultValue,
		}
	}
	f.Fields = fields
	return nil
}

func (f ExtensionNativeFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.raw)
}

type ExtensionNative struct {
	Name   string                `json:"name"`
	Format ExtensionNativeFormat `json:"format"`
}

type ExtensionAPI struct {
	Header                    ExtensionHeader                    `json:"header"`
	BuiltinClassSizes         []ExtensionClassSize               `json:"builtin_class_sizes"`
	BuiltinClassMemberOffsets []ExtensionClassMemberOffsetConfig `json:"builtin_class_member_offsets"`
	GlobalConstants           []struct{}                         `json:"global_constants"` // TODO: global_constants, ignored because empty
	GlobalEnums               []ExtensionEnum                    `json:"global_enums"`
	UtilityFunctions          []ExtensionFunction                `json:"utility_functions"`
	BuiltinClasses            []ExtensionBuiltinClass            `json:"builtin_classes"`
	Classes                   []ExtensionClass                   `json:"classes"`
	Singletons                []ExtensionNameType                `json:"singletons"`
	NativeStructures          []ExtensionNative                  `json:"native_structures"`
}
