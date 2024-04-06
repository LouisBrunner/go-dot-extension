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

type ptrsForOptimizedTranslationList struct {
	fnGenerate gdc.MethodBindPtr
}

var ptrsForOptimizedTranslation ptrsForOptimizedTranslationList

func initOptimizedTranslationPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OptimizedTranslation")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("generate")
		defer methodName.Destroy()
		ptrsForOptimizedTranslation.fnGenerate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1466479800))
	}

}

type OptimizedTranslation struct {
	Translation
}

func (me *OptimizedTranslation) BaseClass() string {
	return "OptimizedTranslation"
}

func NewOptimizedTranslation() *OptimizedTranslation {
	str := StringNameFromStr("OptimizedTranslation") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OptimizedTranslation{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OptimizedTranslation) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OptimizedTranslation) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OptimizedTranslation) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OptimizedTranslation) Generate(from Translation) {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptimizedTranslation.fnGenerate), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
