package main

import (
	"fmt"
)

func handleType(out *outputFiles, typ cType) error {
	switch t := typ.(type) {
	case *enumType:
		return addGoEnumType(out, t)
	case *structType:
		return addGoStructType(out, t)
	case *funcType:
		out.allFuncs[t.ctype] = t
		return writeTemplate(out.files[fileTypes], "types_func.go", map[string]interface{}{
			"Name":  improveTypename(t.ctype),
			"CType": t.ctype,
		})
	case *baseType:
		return writeTemplate(out.files[fileTypes], "types_base.go", map[string]interface{}{
			"Name":   improveTypename(t.ctype),
			"CType":  t.ctype,
			"GoType": t.gotype,
		})
	default:
		return fmt.Errorf("unsupported type %T", typ)
	}
}

func addGoEnumType(out *outputFiles, t *enumType) error {
	members := make([]map[string]interface{}, 0, len(t.members))
	for _, v := range t.members {
		members = append(members, map[string]interface{}{
			"Name":  improveEnumValueName(v.name),
			"Value": v.value,
		})
	}
	return writeTemplate(out.files[fileTypes], "types_enum.go", map[string]interface{}{
		"Name":    improveTypename(t.name),
		"GoType":  t.gotype,
		"CType":   t.name,
		"Members": members,
	})
}

func addGoStructType(out *outputFiles, t *structType) error {
	name := improveTypename(t.name)

	callbacks := map[string]*funcType{}

	members := make([]map[string]interface{}, 0, len(t.members))
	for _, member := range t.members {
		fnName := makeGoTypeName(member.name)
		anonFn, cast := member.typ.(*funcAnonType)
		if cast {
			err := anonFn.addAnonTypes(out.files[fileTypes], name, fnName)
			if err != nil {
				return err
			}
			funcTyp := anonFn.funcType
			funcTyp.name = anonFn.name
			funcTyp.ctype = anonFn.name
			callbacks[fnName] = &funcTyp
		}
		callback, cast := member.typ.(*funcType)
		if cast {
			callbacks[fnName] = callback
		}
		def, free, use := member.typ.goToCGo(fmt.Sprintf("me.%s", fnName))
		members = append(members, map[string]interface{}{
			"Name":   fnName,
			"CName":  makeCGoName(member.name),
			"GoType": member.typ.getGoType(),
			"GoToCGo": map[string]interface{}{
				"Def":  def,
				"Free": free,
				"Use":  use,
			},
		})
	}

	if len(callbacks) > 0 {
		out.callbacks[t] = callbacks
	}

	return writeTemplate(out.files[fileTypes], "types_struct.go", map[string]interface{}{
		"Name":    name,
		"CType":   t.name,
		"Members": members,
	})
}
