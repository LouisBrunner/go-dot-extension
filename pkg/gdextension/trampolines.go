package gdextension

/*
#include <stdlib.h>
#include <stdint.h>

// declared in Go
void *go_dot_gdextension_class_create_instance(void *userdata);
void go_dot_gdextension_class_free_instance(void *userdata, void *instance);
*/
import "C"
import (
	"log"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

func makeTrampoline[T any](fn unsafe.Pointer) T {
	return *(*T)(unsafe.Pointer(&fn))
}

var (
	trampolineClassCreateInstance = makeTrampoline[gdc.ClassCreateInstance](C.go_dot_gdextension_class_create_instance)
	trampolineClassFreeInstance   = makeTrampoline[gdc.ClassFreeInstance](C.go_dot_gdextension_class_free_instance)
)

func (me *extension) makeString(contents string) (gdc.UninitializedStringPtr, func()) {
	str := (gdc.UninitializedStringPtr)(C.malloc(gdapi.ClassSizeString))
	me.iface.StringNewWithUtf8Chars(str, contents)
	return str, func() {
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
	me.iface.CallPtrConstructor(stringNameCtr, strName, &[]gdc.ConstTypePtr{gdc.ConstTypePtr(str)}[0])
	return gdc.StringNamePtr(strName), func() {
		stringNameDtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeStringName)
		me.iface.CallPtrDestructor(stringNameDtr, gdc.TypePtr(strName))
		C.free(unsafe.Pointer(strName))
	}
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

func (me *extension) callPrint(message string) error {
	printNamePtr, clean := me.makeStringName("print")
	defer clean()

	msgPtr, clean := me.makeStringVariant(message)
	defer clean()

	printPtr := me.iface.VariantGetPtrUtilityFunction(gdc.ConstStringNamePtr(printNamePtr), 2648703342)
	me.iface.CallPtrUtilityFunction(printPtr, nil, &[]gdc.ConstTypePtr{gdc.ConstTypePtr(msgPtr)}[0], 1)
	return nil
}

//export go_dot_gdextension_class_create_instance
func go_dot_gdextension_class_create_instance(pUserData unsafe.Pointer) unsafe.Pointer {
	class, err := restore[*classEntry](pUserData)
	if err != nil {
		log.Fatalf("could not restore class entry: %s", err.Error()) // FIXME: should not panic
	}

	parent, clean := class.ext.makeStringName(class.class.ParentClassName())
	defer clean()
	name, clean := class.ext.makeStringName(class.class.ClassName())
	defer clean()

	obj := class.ext.iface.ClassdbConstructObject(gdc.ConstStringNamePtr(parent))
	id := class.ext.iface.ObjectGetInstanceId(gdc.ConstObjectPtr(obj))
	class.ext.iface.ObjectSetInstance(obj, gdc.ConstStringNamePtr(name), gdc.ClassInstancePtr(&id))
	class.addInstance(uint64(id))
	return unsafe.Pointer(obj)
}

//export go_dot_gdextension_class_free_instance
func go_dot_gdextension_class_free_instance(pUserData unsafe.Pointer, pInstance unsafe.Pointer) {
	class, err := restore[*classEntry](pUserData)
	if err != nil {
		log.Fatalf("could not restore class entry: %s", err.Error()) // FIXME: should not panic
	}

	id := *(*uint64)(pInstance)
	class.deleteInstance(uint64(id))
}
