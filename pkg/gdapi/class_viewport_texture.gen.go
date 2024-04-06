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

type ptrsForViewportTextureList struct {
	fnSetViewportPathInScene gdc.MethodBindPtr
	fnGetViewportPathInScene gdc.MethodBindPtr
}

var ptrsForViewportTexture ptrsForViewportTextureList

func initViewportTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("ViewportTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_viewport_path_in_scene")
		defer methodName.Destroy()
		ptrsForViewportTexture.fnSetViewportPathInScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_viewport_path_in_scene")
		defer methodName.Destroy()
		ptrsForViewportTexture.fnGetViewportPathInScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
}

type ViewportTexture struct {
	Texture2D
}

func (me *ViewportTexture) BaseClass() string {
	return "ViewportTexture"
}

func NewViewportTexture() *ViewportTexture {
	str := StringNameFromStr("ViewportTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ViewportTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ViewportTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ViewportTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ViewportTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ViewportTexture) SetViewportPathInScene(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewportTexture.fnSetViewportPathInScene), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ViewportTexture) GetViewportPathInScene() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewportTexture.fnGetViewportPathInScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
