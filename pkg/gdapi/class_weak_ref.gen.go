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

type ptrsForWeakRefList struct {
	fnGetRef gdc.MethodBindPtr
}

var ptrsForWeakRef ptrsForWeakRefList

func initWeakRefPtrs(iface gdc.Interface) {

	className := StringNameFromStr("WeakRef")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_ref")
		defer methodName.Destroy()
		ptrsForWeakRef.fnGetRef = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1214101251))
	}

}

type WeakRef struct {
	RefCounted
}

func (me *WeakRef) BaseClass() string {
	return "WeakRef"
}

func NewWeakRef() *WeakRef {
	str := StringNameFromStr("WeakRef") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &WeakRef{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *WeakRef) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *WeakRef) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *WeakRef) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *WeakRef) GetRef() Variant {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWeakRef.fnGetRef), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
