package main

import (
	"fmt"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func generateInterface(out *outputFiles) error {
	funcs := make([]map[string]interface{}, 0, len(out.ifaceFuncs))
	funcNames := maps.Keys(out.ifaceFuncs)
	slices.Sort(funcNames)
	for _, cName := range funcNames {
		procName := out.ifaceFuncs[cName]

		funcTyp, found := out.allFuncs[cName]
		if !found {
			return fmt.Errorf("could not find function %q", cName)
		}

		args := make([]map[string]interface{}, 0, len(funcTyp.params))
		for _, param := range funcTyp.params {
			args = append(args, map[string]interface{}{
				"Name":    makeGoVarName(param.name),
				"GoType":  param.typ.getGoType(),
				"CGoType": param.typ.getCGoType(),
				"CType":   param.typ.getCType(),
			})
		}

		retGoType := ""
		retCGoType := ""
		retCType := "void"
		if funcTyp.retType != nil {
			retGoType = funcTyp.retType.getGoType()
			retCGoType = funcTyp.retType.getCGoType()
			retCType = funcTyp.retType.getCType()
		}

		funcs = append(funcs, map[string]interface{}{
			"Name":    makeGoTypeName(procName),
			"CName":   cName,
			"Proc":    procName,
			"CGoType": improveTypename(cName),
			"Args":    args,
			"ReturnType": map[string]interface{}{
				"GoType":  retGoType,
				"CGoType": retCGoType,
				"CType":   retCType,
			},
		})
	}

	err := writeTemplate(out.files[fileInterfaceH], "interface.h", map[string]interface{}{
		"Funcs": funcs,
	})
	if err != nil {
		return err
	}
	err = writeTemplate(out.files[fileInterfaceC], "interface.c", map[string]interface{}{
		"Funcs": funcs,
	})
	if err != nil {
		return err
	}
	return writeTemplate(out.files[fileInterface], "interface.go", map[string]interface{}{
		"Funcs": funcs,
	})
}
