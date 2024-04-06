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

type ptrsForTextureLayeredRDList struct {
	fnSetTextureRdRid gdc.MethodBindPtr
	fnGetTextureRdRid gdc.MethodBindPtr
}

var ptrsForTextureLayeredRD ptrsForTextureLayeredRDList

func initTextureLayeredRDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextureLayeredRD")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture_rd_rid")
		defer methodName.Destroy()
		ptrsForTextureLayeredRD.fnSetTextureRdRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_texture_rd_rid")
		defer methodName.Destroy()
		ptrsForTextureLayeredRD.fnGetTextureRdRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}

}

type TextureLayeredRD struct {
	TextureLayered
}

func (me *TextureLayeredRD) BaseClass() string {
	return "TextureLayeredRD"
}

func NewTextureLayeredRD() *TextureLayeredRD {
	str := StringNameFromStr("TextureLayeredRD") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextureLayeredRD{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TextureLayeredRD) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextureLayeredRD) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextureLayeredRD) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TextureLayeredRD) SetTextureRdRid(texture_rd_rid RID) {
	cargs := []gdc.ConstTypePtr{texture_rd_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayeredRD.fnSetTextureRdRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextureLayeredRD) GetTextureRdRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayeredRD.fnGetTextureRdRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
