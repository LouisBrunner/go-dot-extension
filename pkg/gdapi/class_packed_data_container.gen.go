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

type ptrsForPackedDataContainerList struct {
	fnPack gdc.MethodBindPtr
	fnSize gdc.MethodBindPtr
}

var ptrsForPackedDataContainer ptrsForPackedDataContainerList

func initPackedDataContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PackedDataContainer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("pack")
		defer methodName.Destroy()
		ptrsForPackedDataContainer.fnPack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 966674026))
	}
	{
		methodName := StringNameFromStr("size")
		defer methodName.Destroy()
		ptrsForPackedDataContainer.fnSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type PackedDataContainer struct {
	Resource
}

func (me *PackedDataContainer) BaseClass() string {
	return "PackedDataContainer"
}

func NewPackedDataContainer() *PackedDataContainer {
	str := StringNameFromStr("PackedDataContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PackedDataContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PackedDataContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PackedDataContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PackedDataContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PackedDataContainer) Pack(value Variant) Error {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPackedDataContainer.fnPack), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PackedDataContainer) Size() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPackedDataContainer.fnSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
