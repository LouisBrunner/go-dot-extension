package main

import (
	"fmt"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func createFunc(out *outputFiles, cName, funcName string, funcTyp *funcType) (map[string]interface{}, *funcType) {
	var subFunc *funcType

	anyFree := false

	args := make([]map[string]interface{}, 0, len(funcTyp.params))
	for _, param := range funcTyp.params {
		name := makeGoVarName(param.name)
		def, free, use := param.typ.goToCGo(name)
		args = append(args, map[string]interface{}{
			"Name":    name,
			"GoType":  param.typ.getGoType(),
			"CGoType": param.typ.getCGoType(),
			"CType":   param.typ.getCType(),
			"GoToCGo": map[string]interface{}{
				"Def":  def,
				"Free": free,
				"Use":  use,
			},
		})
		if free != "" {
			anyFree = true
		}
	}

	retGoType := ""
	retCGoType := ""
	retCType := "void"
	cgoCast := "ret"
	if funcTyp.retType != nil {
		retGoType = funcTyp.retType.getGoType()
		retCGoType = funcTyp.retType.getCGoType()
		retCType = funcTyp.retType.getCType()
		cgoCast = funcTyp.retType.cgoToGo(cgoCast)
		subFuncPtr, isFuncPtr := out.allFuncs[retCType]
		if isFuncPtr {
			subFunc = subFuncPtr
		}
	}

	return map[string]interface{}{
		"Name":    funcName,
		"Proc":    "",
		"CName":   cName,
		"CGoType": improveTypename(cName),
		"Args":    args,
		"AnyFree": anyFree,
		"ReturnType": map[string]interface{}{
			"GoType":  retGoType,
			"CGoType": retCGoType,
			"CType":   retCType,
			"CGoToGo": cgoCast,
		},
	}, subFunc
}

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

		newFunc, subFunc := createFunc(out, cName, makeGoTypeName(procName), funcTyp)
		newFunc["Proc"] = procName
		funcs = append(funcs, newFunc)

		for subFunc != nil {
			name := improveTypename(subFunc.name)
			subNewFunc, subSubFunc := createFunc(out, subFunc.name, fmt.Sprintf("Call%s", name), subFunc)
			subNewFunc["RealName"] = name
			funcs = append(funcs, subNewFunc)
			subFunc = subSubFunc
		}
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
