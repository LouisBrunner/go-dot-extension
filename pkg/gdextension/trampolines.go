package gdextension

/*
#include <stdlib.h>
#include <stdint.h>

// declared in C
void go_dot_gdextension_call_variant_constructor(void (*fn)(void *p_self, const void *p_args, int p_argcount), void *p_self, const void *p_args, int p_argcount);
void go_dot_gdextension_call_variant_type_constructor(void (*fn)(void *p_self, const void *inside), void *p_self, const void *inside);
void go_dot_gdextension_call_variant_destructor(void (*fn)(void *p_self), void *p_self);
void go_dot_gdextension_call_utility_function(void (*fn)(void** ret, const void *p_args, int p_argcount), void** ret, const void *p_args, int p_argcount);

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
	stringDtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeString)

	str := (gdc.UninitializedStringPtr)(C.malloc(gdapi.ClassSizeString))
	me.iface.StringNewWithUtf8Chars(str, contents)
	return str, func() {
		me.callVariantDestructor(stringDtr, gdc.VariantPtr(str))
		C.free(unsafe.Pointer(str))
	}
}

func (me *extension) makeStringName(contents string) (gdc.StringNamePtr, func()) {
	stringNameCtr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 2)
	stringNameDtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeStringName)

	str, clean := me.makeString(contents)
	defer clean()

	strName := (gdc.StringNamePtr)(C.malloc(gdapi.ClassSizeStringName))
	me.callVariantConstructor(stringNameCtr, gdc.VariantPtr(strName), [1]gdc.ConstTypePtr{gdc.ConstTypePtr(str)})
	return strName, func() {
		me.callVariantDestructor(stringNameDtr, gdc.VariantPtr(strName))
		C.free(unsafe.Pointer(strName))
	}
}

func (me *extension) makeStringVariant(contents string) (gdc.VariantPtr, func()) {
	variantCtr := me.iface.GetVariantFromTypeConstructor(gdc.VariantTypeString)

	str, clean := me.makeString(contents)

	strVariant := (gdc.VariantPtr)(C.malloc(gdapi.ClassSizeVariant))
	me.callVariantTypeConstructor(variantCtr, strVariant, unsafe.Pointer(str))
	return strVariant, func() {
		me.iface.VariantDestroy(gdc.VariantPtr(strVariant))
		C.free(unsafe.Pointer(strVariant))
		clean()
	}
}

func (me *extension) callVariantConstructor(constructor gdc.PtrConstructor, variant gdc.VariantPtr, args [1]gdc.ConstTypePtr) {
	C.go_dot_gdextension_call_variant_constructor(
		constructor,
		unsafe.Pointer(variant),
		unsafe.Pointer(&args[0]),
		C.int(len(args)),
	)
}

func (me *extension) callVariantTypeConstructor(constructor gdc.VariantFromTypeConstructorFunc, variant gdc.VariantPtr, inside unsafe.Pointer) {
	C.go_dot_gdextension_call_variant_type_constructor(
		constructor,
		unsafe.Pointer(variant),
		inside,
	)
}

func (me *extension) callVariantDestructor(destructor gdc.PtrDestructor, variant gdc.VariantPtr) {
	C.go_dot_gdextension_call_variant_destructor(
		destructor,
		unsafe.Pointer(variant),
	)
}

func (me *extension) callUtilityFunction(function gdc.PtrUtilityFunction, args []gdc.ConstVariantPtr) {
	C.go_dot_gdextension_call_utility_function(
		function,
		nil,
		unsafe.Pointer(&args[0]),
		C.int(len(args)),
	)
}

func (me *extension) callPrint(message string) error {
	printNamePtr, clean := me.makeStringName("print")
	defer clean()

	msgPtr, clean := me.makeStringVariant(message)
	defer clean()

	printPtr := me.iface.VariantGetPtrUtilityFunction(gdc.ConstStringNamePtr(printNamePtr), 2648703342)
	me.callUtilityFunction(printPtr, []gdc.ConstVariantPtr{gdc.ConstVariantPtr(msgPtr)})
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
