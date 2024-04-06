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

type ptrsForPackedDataContainerRefList struct {
	fnSize gdc.MethodBindPtr
}

var ptrsForPackedDataContainerRef ptrsForPackedDataContainerRefList

func initPackedDataContainerRefPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PackedDataContainerRef")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("size")
		defer methodName.Destroy()
		ptrsForPackedDataContainerRef.fnSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type PackedDataContainerRef struct {
	RefCounted
}

func (me *PackedDataContainerRef) BaseClass() string {
	return "PackedDataContainerRef"
}

func NewPackedDataContainerRef() *PackedDataContainerRef {
	str := StringNameFromStr("PackedDataContainerRef") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PackedDataContainerRef{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PackedDataContainerRef) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PackedDataContainerRef) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PackedDataContainerRef) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PackedDataContainerRef) Size() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPackedDataContainerRef.fnSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
