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

type ptrsForPortableCompressedTexture2DList struct {
	fnCreateFromImage               gdc.MethodBindPtr
	fnGetFormat                     gdc.MethodBindPtr
	fnGetCompressionMode            gdc.MethodBindPtr
	fnSetSizeOverride               gdc.MethodBindPtr
	fnGetSizeOverride               gdc.MethodBindPtr
	fnSetKeepCompressedBuffer       gdc.MethodBindPtr
	fnIsKeepingCompressedBuffer     gdc.MethodBindPtr
	fnSetKeepAllCompressedBuffers   gdc.MethodBindPtr
	fnIsKeepingAllCompressedBuffers gdc.MethodBindPtr
}

var ptrsForPortableCompressedTexture2D ptrsForPortableCompressedTexture2DList

func initPortableCompressedTexture2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PortableCompressedTexture2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create_from_image")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnCreateFromImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3679243433))
	}
	{
		methodName := StringNameFromStr("get_format")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3847873762))
	}
	{
		methodName := StringNameFromStr("get_compression_mode")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnGetCompressionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3265612739))
	}
	{
		methodName := StringNameFromStr("set_size_override")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnSetSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_size_override")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnGetSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_keep_compressed_buffer")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnSetKeepCompressedBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_keeping_compressed_buffer")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnIsKeepingCompressedBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_keep_all_compressed_buffers")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnSetKeepAllCompressedBuffers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_keeping_all_compressed_buffers")
		defer methodName.Destroy()
		ptrsForPortableCompressedTexture2D.fnIsKeepingAllCompressedBuffers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
}

type PortableCompressedTexture2D struct {
	Texture2D
}

func (me *PortableCompressedTexture2D) BaseClass() string {
	return "PortableCompressedTexture2D"
}

func NewPortableCompressedTexture2D() *PortableCompressedTexture2D {
	str := StringNameFromStr("PortableCompressedTexture2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PortableCompressedTexture2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type PortableCompressedTexture2DCompressionMode int

const (
	PortableCompressedTexture2DCompressionModeCompressionModeLossless       PortableCompressedTexture2DCompressionMode = 0
	PortableCompressedTexture2DCompressionModeCompressionModeLossy          PortableCompressedTexture2DCompressionMode = 1
	PortableCompressedTexture2DCompressionModeCompressionModeBasisUniversal PortableCompressedTexture2DCompressionMode = 2
	PortableCompressedTexture2DCompressionModeCompressionModeS3Tc           PortableCompressedTexture2DCompressionMode = 3
	PortableCompressedTexture2DCompressionModeCompressionModeEtc2           PortableCompressedTexture2DCompressionMode = 4
	PortableCompressedTexture2DCompressionModeCompressionModeBptc           PortableCompressedTexture2DCompressionMode = 5
)

func (me *PortableCompressedTexture2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PortableCompressedTexture2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PortableCompressedTexture2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PortableCompressedTexture2D) CreateFromImage(image Image, compression_mode PortableCompressedTexture2DCompressionMode, normal_map bool, lossy_quality float64) {
	cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&compression_mode), gdc.ConstTypePtr(&normal_map), gdc.ConstTypePtr(&lossy_quality)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnCreateFromImage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PortableCompressedTexture2D) GetFormat() ImageFormat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ImageFormat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PortableCompressedTexture2D) GetCompressionMode() PortableCompressedTexture2DCompressionMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PortableCompressedTexture2DCompressionMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnGetCompressionMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PortableCompressedTexture2D) SetSizeOverride(size Vector2) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnSetSizeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PortableCompressedTexture2D) GetSizeOverride() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnGetSizeOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PortableCompressedTexture2D) SetKeepCompressedBuffer(keep bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnSetKeepCompressedBuffer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PortableCompressedTexture2D) IsKeepingCompressedBuffer() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnIsKeepingCompressedBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func PortableCompressedTexture2DSetKeepAllCompressedBuffers(keep bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnSetKeepAllCompressedBuffers), nil, unsafe.SliceData(cargs), nil)

}

func PortableCompressedTexture2DIsKeepingAllCompressedBuffers() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPortableCompressedTexture2D.fnIsKeepingAllCompressedBuffers), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
