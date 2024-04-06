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

type ptrsForAtlasTextureList struct {
	fnSetAtlas      gdc.MethodBindPtr
	fnGetAtlas      gdc.MethodBindPtr
	fnSetRegion     gdc.MethodBindPtr
	fnGetRegion     gdc.MethodBindPtr
	fnSetMargin     gdc.MethodBindPtr
	fnGetMargin     gdc.MethodBindPtr
	fnSetFilterClip gdc.MethodBindPtr
	fnHasFilterClip gdc.MethodBindPtr
}

var ptrsForAtlasTexture ptrsForAtlasTextureList

func initAtlasTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AtlasTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_atlas")
		defer methodName.Destroy()
		ptrsForAtlasTexture.fnSetAtlas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_atlas")
		defer methodName.Destroy()
		ptrsForAtlasTexture.fnGetAtlas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_region")
		defer methodName.Destroy()
		ptrsForAtlasTexture.fnSetRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
	}
	{
		methodName := StringNameFromStr("get_region")
		defer methodName.Destroy()
		ptrsForAtlasTexture.fnGetRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("set_margin")
		defer methodName.Destroy()
		ptrsForAtlasTexture.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
	}
	{
		methodName := StringNameFromStr("get_margin")
		defer methodName.Destroy()
		ptrsForAtlasTexture.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("set_filter_clip")
		defer methodName.Destroy()
		ptrsForAtlasTexture.fnSetFilterClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_filter_clip")
		defer methodName.Destroy()
		ptrsForAtlasTexture.fnHasFilterClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type AtlasTexture struct {
	Texture2D
}

func (me *AtlasTexture) BaseClass() string {
	return "AtlasTexture"
}

func NewAtlasTexture() *AtlasTexture {
	str := StringNameFromStr("AtlasTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AtlasTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AtlasTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AtlasTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AtlasTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AtlasTexture) SetAtlas(atlas Texture2D) {
	cargs := []gdc.ConstTypePtr{atlas.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAtlasTexture.fnSetAtlas), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AtlasTexture) GetAtlas() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAtlasTexture.fnGetAtlas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AtlasTexture) SetRegion(region Rect2) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAtlasTexture.fnSetRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AtlasTexture) GetRegion() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAtlasTexture.fnGetRegion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AtlasTexture) SetMargin(margin Rect2) {
	cargs := []gdc.ConstTypePtr{margin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAtlasTexture.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AtlasTexture) GetMargin() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAtlasTexture.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AtlasTexture) SetFilterClip(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAtlasTexture.fnSetFilterClip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AtlasTexture) HasFilterClip() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAtlasTexture.fnHasFilterClip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
