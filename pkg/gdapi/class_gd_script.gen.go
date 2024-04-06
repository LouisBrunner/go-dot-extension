// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForGDScriptList struct {
	fnNew gdc.MethodBindPtr
}

var ptrsForGDScript ptrsForGDScriptList

func initGDScriptPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GDScript")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("new")
		defer methodName.Destroy()
		ptrsForGDScript.fnNew = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1545262638))
	}
}

type GDScript struct {
	Script
}

func (me *GDScript) BaseClass() string {
	return "GDScript"
}

func NewGDScript() *GDScript {
	str := StringNameFromStr("GDScript") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GDScript{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GDScript) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GDScript) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GDScript) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GDScript) New(varargs ...Variant) Variant {
	cargs := make([]gdc.ConstVariantPtr, 0, 0+len(varargs))
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForGDScript.fnNew), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return *ret
	}
	return *ret
}

// Signals
