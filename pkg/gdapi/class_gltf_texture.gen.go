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

type ptrsForGLTFTextureList struct {
	fnGetSrcImage gdc.MethodBindPtr
	fnSetSrcImage gdc.MethodBindPtr
	fnGetSampler  gdc.MethodBindPtr
	fnSetSampler  gdc.MethodBindPtr
}

var ptrsForGLTFTexture ptrsForGLTFTextureList

func initGLTFTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_src_image")
		defer methodName.Destroy()
		ptrsForGLTFTexture.fnGetSrcImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_src_image")
		defer methodName.Destroy()
		ptrsForGLTFTexture.fnSetSrcImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_sampler")
		defer methodName.Destroy()
		ptrsForGLTFTexture.fnGetSampler = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_sampler")
		defer methodName.Destroy()
		ptrsForGLTFTexture.fnSetSampler = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}

}

type GLTFTexture struct {
	Resource
}

func (me *GLTFTexture) BaseClass() string {
	return "GLTFTexture"
}

func NewGLTFTexture() *GLTFTexture {
	str := StringNameFromStr("GLTFTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GLTFTexture) GetSrcImage() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTexture.fnGetSrcImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFTexture) SetSrcImage(src_image int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&src_image)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTexture.fnSetSrcImage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFTexture) GetSampler() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTexture.fnGetSampler), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFTexture) SetSampler(sampler int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sampler)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTexture.fnSetSampler), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
