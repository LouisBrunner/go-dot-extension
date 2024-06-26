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

type ptrsForGLTFCameraList struct {
	fnFromNode       gdc.MethodBindPtr
	fnToNode         gdc.MethodBindPtr
	fnFromDictionary gdc.MethodBindPtr
	fnToDictionary   gdc.MethodBindPtr
	fnGetPerspective gdc.MethodBindPtr
	fnSetPerspective gdc.MethodBindPtr
	fnGetFov         gdc.MethodBindPtr
	fnSetFov         gdc.MethodBindPtr
	fnGetSizeMag     gdc.MethodBindPtr
	fnSetSizeMag     gdc.MethodBindPtr
	fnGetDepthFar    gdc.MethodBindPtr
	fnSetDepthFar    gdc.MethodBindPtr
	fnGetDepthNear   gdc.MethodBindPtr
	fnSetDepthNear   gdc.MethodBindPtr
}

var ptrsForGLTFCamera ptrsForGLTFCameraList

func initGLTFCameraPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFCamera")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("from_node")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnFromNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 237784))
	}
	{
		methodName := StringNameFromStr("to_node")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnToNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2285090890))
	}
	{
		methodName := StringNameFromStr("from_dictionary")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnFromDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2495512509))
	}
	{
		methodName := StringNameFromStr("to_dictionary")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnToDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_perspective")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnGetPerspective = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_perspective")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnSetPerspective = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_fov")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnGetFov = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fov")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnSetFov = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_size_mag")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnGetSizeMag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_size_mag")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnSetSizeMag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth_far")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnGetDepthFar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_depth_far")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnSetDepthFar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth_near")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnGetDepthNear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_depth_near")
		defer methodName.Destroy()
		ptrsForGLTFCamera.fnSetDepthNear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}

}

type GLTFCamera struct {
	Resource
}

func (me *GLTFCamera) BaseClass() string {
	return "GLTFCamera"
}

func NewGLTFCamera() *GLTFCamera {
	str := StringNameFromStr("GLTFCamera") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFCamera{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFCamera) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFCamera) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFCamera) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func GLTFCameraFromNode(camera_node Camera3D) GLTFCamera {
	cargs := []gdc.ConstTypePtr{camera_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFCamera()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnFromNode), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFCamera) ToNode() Camera3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCamera3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnToNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func GLTFCameraFromDictionary(dictionary Dictionary) GLTFCamera {
	cargs := []gdc.ConstTypePtr{dictionary.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFCamera()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnFromDictionary), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFCamera) ToDictionary() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnToDictionary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFCamera) GetPerspective() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnGetPerspective), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFCamera) SetPerspective(perspective bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&perspective)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnSetPerspective), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFCamera) GetFov() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnGetFov), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFCamera) SetFov(fov float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fov)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnSetFov), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFCamera) GetSizeMag() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnGetSizeMag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFCamera) SetSizeMag(size_mag float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_mag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnSetSizeMag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFCamera) GetDepthFar() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnGetDepthFar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFCamera) SetDepthFar(zdepth_far float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zdepth_far)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnSetDepthFar), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFCamera) GetDepthNear() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnGetDepthNear), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFCamera) SetDepthNear(zdepth_near float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zdepth_near)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFCamera.fnSetDepthNear), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
