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

type ptrsForTextureLayeredList struct {
	fnXGetFormat      gdc.MethodBindPtr
	fnXGetLayeredType gdc.MethodBindPtr
	fnXGetWidth       gdc.MethodBindPtr
	fnXGetHeight      gdc.MethodBindPtr
	fnXGetLayers      gdc.MethodBindPtr
	fnXHasMipmaps     gdc.MethodBindPtr
	fnXGetLayerData   gdc.MethodBindPtr
	fnGetFormat       gdc.MethodBindPtr
	fnGetLayeredType  gdc.MethodBindPtr
	fnGetWidth        gdc.MethodBindPtr
	fnGetHeight       gdc.MethodBindPtr
	fnGetLayers       gdc.MethodBindPtr
	fnHasMipmaps      gdc.MethodBindPtr
	fnGetLayerData    gdc.MethodBindPtr
}

var ptrsForTextureLayered ptrsForTextureLayeredList

func initTextureLayeredPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextureLayered")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_format")
		defer methodName.Destroy()
		ptrsForTextureLayered.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3847873762))
	}
	{
		methodName := StringNameFromStr("get_layered_type")
		defer methodName.Destroy()
		ptrsForTextureLayered.fnGetLayeredType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 518123893))
	}
	{
		methodName := StringNameFromStr("get_width")
		defer methodName.Destroy()
		ptrsForTextureLayered.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_height")
		defer methodName.Destroy()
		ptrsForTextureLayered.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_layers")
		defer methodName.Destroy()
		ptrsForTextureLayered.fnGetLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("has_mipmaps")
		defer methodName.Destroy()
		ptrsForTextureLayered.fnHasMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_layer_data")
		defer methodName.Destroy()
		ptrsForTextureLayered.fnGetLayerData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3655284255))
	}

}

type TextureLayered struct {
	Texture
}

func (me *TextureLayered) BaseClass() string {
	return "TextureLayered"
}

func NewTextureLayered() *TextureLayered {
	str := StringNameFromStr("TextureLayered") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextureLayered{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TextureLayeredLayeredType int

const (
	TextureLayeredLayeredTypeLayeredType2DArray      TextureLayeredLayeredType = 0
	TextureLayeredLayeredTypeLayeredTypeCubemap      TextureLayeredLayeredType = 1
	TextureLayeredLayeredTypeLayeredTypeCubemapArray TextureLayeredLayeredType = 2
)

func (me *TextureLayered) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextureLayered) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextureLayered) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TextureLayered) GetFormat() ImageFormat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ImageFormat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayered.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextureLayered) GetLayeredType() TextureLayeredLayeredType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextureLayeredLayeredType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayered.fnGetLayeredType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextureLayered) GetWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayered.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureLayered) GetHeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayered.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureLayered) GetLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayered.fnGetLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureLayered) HasMipmaps() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayered.fnHasMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextureLayered) GetLayerData(layer int64) Image {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImage()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureLayered.fnGetLayerData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
