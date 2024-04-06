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

type ptrsForImageTexture3DList struct {
	fnCreate gdc.MethodBindPtr
	fnUpdate gdc.MethodBindPtr
}

var ptrsForImageTexture3D ptrsForImageTexture3DList

func initImageTexture3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ImageTexture3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create")
		defer methodName.Destroy()
		ptrsForImageTexture3D.fnCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130379827))
	}
	{
		methodName := StringNameFromStr("update")
		defer methodName.Destroy()
		ptrsForImageTexture3D.fnUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}

}

type ImageTexture3D struct {
	Texture3D
}

func (me *ImageTexture3D) BaseClass() string {
	return "ImageTexture3D"
}

func NewImageTexture3D() *ImageTexture3D {
	str := StringNameFromStr("ImageTexture3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ImageTexture3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ImageTexture3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ImageTexture3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ImageTexture3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ImageTexture3D) Create(format ImageFormat, width int64, height int64, depth int64, use_mipmaps bool, data []Image) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&use_mipmaps), gdc.ConstTypePtr(&data)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&format)
	pinner.Pin(&width)
	pinner.Pin(&height)
	pinner.Pin(&depth)
	pinner.Pin(&use_mipmaps)
	pinner.Pin(&data)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTexture3D.fnCreate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ImageTexture3D) Update(data []Image) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTexture3D.fnUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
