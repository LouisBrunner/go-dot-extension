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

type ptrsForXRInterfaceExtensionList struct {
	fnXGetName                     gdc.MethodBindPtr
	fnXGetCapabilities             gdc.MethodBindPtr
	fnXIsInitialized               gdc.MethodBindPtr
	fnXInitialize                  gdc.MethodBindPtr
	fnXUninitialize                gdc.MethodBindPtr
	fnXGetSystemInfo               gdc.MethodBindPtr
	fnXSupportsPlayAreaMode        gdc.MethodBindPtr
	fnXGetPlayAreaMode             gdc.MethodBindPtr
	fnXSetPlayAreaMode             gdc.MethodBindPtr
	fnXGetPlayArea                 gdc.MethodBindPtr
	fnXGetRenderTargetSize         gdc.MethodBindPtr
	fnXGetViewCount                gdc.MethodBindPtr
	fnXGetCameraTransform          gdc.MethodBindPtr
	fnXGetTransformForView         gdc.MethodBindPtr
	fnXGetProjectionForView        gdc.MethodBindPtr
	fnXGetVrsTexture               gdc.MethodBindPtr
	fnXProcess                     gdc.MethodBindPtr
	fnXPreRender                   gdc.MethodBindPtr
	fnXPreDrawViewport             gdc.MethodBindPtr
	fnXPostDrawViewport            gdc.MethodBindPtr
	fnXEndFrame                    gdc.MethodBindPtr
	fnXGetSuggestedTrackerNames    gdc.MethodBindPtr
	fnXGetSuggestedPoseNames       gdc.MethodBindPtr
	fnXGetTrackingStatus           gdc.MethodBindPtr
	fnXTriggerHapticPulse          gdc.MethodBindPtr
	fnXGetAnchorDetectionIsEnabled gdc.MethodBindPtr
	fnXSetAnchorDetectionIsEnabled gdc.MethodBindPtr
	fnXGetCameraFeedId             gdc.MethodBindPtr
	fnXGetColorTexture             gdc.MethodBindPtr
	fnXGetDepthTexture             gdc.MethodBindPtr
	fnXGetVelocityTexture          gdc.MethodBindPtr
	fnGetColorTexture              gdc.MethodBindPtr
	fnGetDepthTexture              gdc.MethodBindPtr
	fnGetVelocityTexture           gdc.MethodBindPtr
	fnAddBlit                      gdc.MethodBindPtr
	fnGetRenderTargetTexture       gdc.MethodBindPtr
}

var ptrsForXRInterfaceExtension ptrsForXRInterfaceExtensionList

func initXRInterfaceExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRInterfaceExtension")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_color_texture")
		defer methodName.Destroy()
		ptrsForXRInterfaceExtension.fnGetColorTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("get_depth_texture")
		defer methodName.Destroy()
		ptrsForXRInterfaceExtension.fnGetDepthTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("get_velocity_texture")
		defer methodName.Destroy()
		ptrsForXRInterfaceExtension.fnGetVelocityTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("add_blit")
		defer methodName.Destroy()
		ptrsForXRInterfaceExtension.fnAddBlit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 258596971))
	}
	{
		methodName := StringNameFromStr("get_render_target_texture")
		defer methodName.Destroy()
		ptrsForXRInterfaceExtension.fnGetRenderTargetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 41030802))
	}

}

type XRInterfaceExtension struct {
	XRInterface
}

func (me *XRInterfaceExtension) BaseClass() string {
	return "XRInterfaceExtension"
}

func NewXRInterfaceExtension() *XRInterfaceExtension {
	str := StringNameFromStr("XRInterfaceExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRInterfaceExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *XRInterfaceExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRInterfaceExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRInterfaceExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRInterfaceExtension) GetColorTexture() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterfaceExtension.fnGetColorTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterfaceExtension) GetDepthTexture() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterfaceExtension.fnGetDepthTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterfaceExtension) GetVelocityTexture() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterfaceExtension.fnGetVelocityTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRInterfaceExtension) AddBlit(render_target RID, src_rect Rect2, dst_rect Rect2i, use_layer bool, layer int64, apply_lens_distortion bool, eye_center Vector2, k1 float64, k2 float64, upscale float64, aspect_ratio float64) {
	cargs := []gdc.ConstTypePtr{render_target.AsCTypePtr(), src_rect.AsCTypePtr(), dst_rect.AsCTypePtr(), gdc.ConstTypePtr(&use_layer), gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&apply_lens_distortion), eye_center.AsCTypePtr(), gdc.ConstTypePtr(&k1), gdc.ConstTypePtr(&k2), gdc.ConstTypePtr(&upscale), gdc.ConstTypePtr(&aspect_ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterfaceExtension.fnAddBlit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRInterfaceExtension) GetRenderTargetTexture(render_target RID) RID {
	cargs := []gdc.ConstTypePtr{render_target.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRInterfaceExtension.fnGetRenderTargetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
