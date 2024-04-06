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

type ptrsForTexture2DRDList struct {
	fnSetTextureRdRid gdc.MethodBindPtr
	fnGetTextureRdRid gdc.MethodBindPtr
}

var ptrsForTexture2DRD ptrsForTexture2DRDList

func initTexture2DRDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Texture2DRD")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture_rd_rid")
		defer methodName.Destroy()
		ptrsForTexture2DRD.fnSetTextureRdRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_texture_rd_rid")
		defer methodName.Destroy()
		ptrsForTexture2DRD.fnGetTextureRdRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}

}

type Texture2DRD struct {
	Texture2D
}

func (me *Texture2DRD) BaseClass() string {
	return "Texture2DRD"
}

func NewTexture2DRD() *Texture2DRD {
	str := StringNameFromStr("Texture2DRD") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Texture2DRD{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Texture2DRD) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Texture2DRD) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Texture2DRD) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Texture2DRD) SetTextureRdRid(texture_rd_rid RID) {
	cargs := []gdc.ConstTypePtr{texture_rd_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2DRD.fnSetTextureRdRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Texture2DRD) GetTextureRdRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2DRD.fnGetTextureRdRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
