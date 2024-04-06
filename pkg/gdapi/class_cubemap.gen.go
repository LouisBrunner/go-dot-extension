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

type ptrsForCubemapList struct {
	fnCreatePlaceholder gdc.MethodBindPtr
}

var ptrsForCubemap ptrsForCubemapList

func initCubemapPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Cubemap")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create_placeholder")
		defer methodName.Destroy()
		ptrsForCubemap.fnCreatePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121922552))
	}

}

type Cubemap struct {
	ImageTextureLayered
}

func (me *Cubemap) BaseClass() string {
	return "Cubemap"
}

func NewCubemap() *Cubemap {
	str := StringNameFromStr("Cubemap") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Cubemap{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Cubemap) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Cubemap) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Cubemap) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Cubemap) CreatePlaceholder() Resource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewResource()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCubemap.fnCreatePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
