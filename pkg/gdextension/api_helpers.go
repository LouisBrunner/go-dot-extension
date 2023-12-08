package gdextension

// TODO: all of those should be generated in gdapi

/*
#include <stdlib.h>
*/
import "C"
import (
	"strings"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

func (me *extension) stringifyVariant(variant gdc.ConstVariantPtr, maxLen int) string {
	stringified, _ := me.makeString(strings.Repeat("a", maxLen))
	me.iface.VariantStringify(variant, stringified)
	name := make([]byte, maxLen)
	len := me.iface.StringToUtf8Chars(gdc.ConstStringPtr(stringified), unsafe.SliceData(name), gdc.Int(maxLen))
	return string(name[:len])
}

func (me *extension) makeString(contents string) (gdc.StringPtr, func()) {
	str := (gdc.UninitializedStringPtr)(C.malloc(gdapi.ClassSizeString))
	me.iface.StringNewWithUtf8Chars(str, contents)
	return gdc.StringPtr(str), func() {
		stringDtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeString)
		me.iface.CallPtrDestructor(stringDtr, gdc.TypePtr(str))
		C.free(unsafe.Pointer(str))
	}
}

func (me *extension) makeStringName(contents string) (gdc.StringNamePtr, func()) {
	str, clean := me.makeString(contents)
	defer clean()

	stringNameCtr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 2)
	strName := (gdc.UninitializedTypePtr)(C.malloc(gdapi.ClassSizeStringName))
	me.iface.CallPtrConstructor(stringNameCtr, strName, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(str)}))
	return gdc.StringNamePtr(strName), func() {
		me.freeStringName(gdc.StringNamePtr(strName))
	}
}

func (me *extension) freeStringName(name gdc.StringNamePtr) {
	stringNameDtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeStringName)
	me.iface.CallPtrDestructor(stringNameDtr, gdc.TypePtr(name))
	C.free(unsafe.Pointer(name))
}

func (me *extension) makeStringVariant(contents string) (gdc.VariantPtr, func()) {
	str, clean := me.makeString(contents)

	variantCtr := me.iface.GetVariantFromTypeConstructor(gdc.VariantTypeString)
	strVariant := (gdc.UninitializedVariantPtr)(C.malloc(gdapi.ClassSizeVariant))
	me.iface.CallVariantFromTypeConstructorFunc(variantCtr, strVariant, gdc.TypePtr(str))
	return gdc.VariantPtr(strVariant), func() {
		me.iface.VariantDestroy(gdc.VariantPtr(strVariant))
		C.free(unsafe.Pointer(strVariant))
		clean()
	}
}

func (me *extension) callPrint(message string) {
	printNamePtr, clean := me.makeStringName("print")
	defer clean()

	msgPtr, clean := me.makeStringVariant(message)
	defer clean()

	printPtr := me.iface.VariantGetPtrUtilityFunction(gdc.ConstStringNamePtr(printNamePtr), 2648703342)
	me.iface.CallPtrUtilityFunction(printPtr, nil, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(msgPtr)}), 1)
}
