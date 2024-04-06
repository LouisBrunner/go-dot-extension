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

type ptrsForCompressedTexture3DList struct {
	fnLoad        gdc.MethodBindPtr
	fnGetLoadPath gdc.MethodBindPtr
}

var ptrsForCompressedTexture3D ptrsForCompressedTexture3DList

func initCompressedTexture3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CompressedTexture3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("load")
		defer methodName.Destroy()
		ptrsForCompressedTexture3D.fnLoad = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("get_load_path")
		defer methodName.Destroy()
		ptrsForCompressedTexture3D.fnGetLoadPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}

}

type CompressedTexture3D struct {
	Texture3D
}

func (me *CompressedTexture3D) BaseClass() string {
	return "CompressedTexture3D"
}

func NewCompressedTexture3D() *CompressedTexture3D {
	str := StringNameFromStr("CompressedTexture3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CompressedTexture3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CompressedTexture3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CompressedTexture3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CompressedTexture3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CompressedTexture3D) Load(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompressedTexture3D.fnLoad), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CompressedTexture3D) GetLoadPath() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompressedTexture3D.fnGetLoadPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
