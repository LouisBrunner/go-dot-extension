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

type ptrsForCameraTextureList struct {
	fnSetCameraFeedId gdc.MethodBindPtr
	fnGetCameraFeedId gdc.MethodBindPtr
	fnSetWhichFeed    gdc.MethodBindPtr
	fnGetWhichFeed    gdc.MethodBindPtr
	fnSetCameraActive gdc.MethodBindPtr
	fnGetCameraActive gdc.MethodBindPtr
}

var ptrsForCameraTexture ptrsForCameraTextureList

func initCameraTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("CameraTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_camera_feed_id")
		defer methodName.Destroy()
		ptrsForCameraTexture.fnSetCameraFeedId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_camera_feed_id")
		defer methodName.Destroy()
		ptrsForCameraTexture.fnGetCameraFeedId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_which_feed")
		defer methodName.Destroy()
		ptrsForCameraTexture.fnSetWhichFeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1595299230))
	}
	{
		methodName := StringNameFromStr("get_which_feed")
		defer methodName.Destroy()
		ptrsForCameraTexture.fnGetWhichFeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 91039457))
	}
	{
		methodName := StringNameFromStr("set_camera_active")
		defer methodName.Destroy()
		ptrsForCameraTexture.fnSetCameraActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_camera_active")
		defer methodName.Destroy()
		ptrsForCameraTexture.fnGetCameraActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type CameraTexture struct {
	Texture2D
}

func (me *CameraTexture) BaseClass() string {
	return "CameraTexture"
}

func NewCameraTexture() *CameraTexture {
	str := StringNameFromStr("CameraTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CameraTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CameraTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CameraTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CameraTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CameraTexture) SetCameraFeedId(feed_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feed_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraTexture.fnSetCameraFeedId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraTexture) GetCameraFeedId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraTexture.fnGetCameraFeedId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CameraTexture) SetWhichFeed(which_feed CameraServerFeedImage) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&which_feed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraTexture.fnSetWhichFeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraTexture) GetWhichFeed() CameraServerFeedImage {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CameraServerFeedImage

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraTexture.fnGetWhichFeed), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CameraTexture) SetCameraActive(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraTexture.fnSetCameraActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraTexture) GetCameraActive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraTexture.fnGetCameraActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
