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

type ptrsForCompositorEffectList struct {
	fnXRenderCallback          gdc.MethodBindPtr
	fnSetEnabled               gdc.MethodBindPtr
	fnGetEnabled               gdc.MethodBindPtr
	fnSetEffectCallbackType    gdc.MethodBindPtr
	fnGetEffectCallbackType    gdc.MethodBindPtr
	fnSetAccessResolvedColor   gdc.MethodBindPtr
	fnGetAccessResolvedColor   gdc.MethodBindPtr
	fnSetAccessResolvedDepth   gdc.MethodBindPtr
	fnGetAccessResolvedDepth   gdc.MethodBindPtr
	fnSetNeedsMotionVectors    gdc.MethodBindPtr
	fnGetNeedsMotionVectors    gdc.MethodBindPtr
	fnSetNeedsNormalRoughness  gdc.MethodBindPtr
	fnGetNeedsNormalRoughness  gdc.MethodBindPtr
	fnSetNeedsSeparateSpecular gdc.MethodBindPtr
	fnGetNeedsSeparateSpecular gdc.MethodBindPtr
}

var ptrsForCompositorEffect ptrsForCompositorEffectList

func initCompositorEffectPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CompositorEffect")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_enabled")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enabled")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnGetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_effect_callback_type")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnSetEffectCallbackType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1390728419))
	}
	{
		methodName := StringNameFromStr("get_effect_callback_type")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnGetEffectCallbackType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1221912590))
	}
	{
		methodName := StringNameFromStr("set_access_resolved_color")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnSetAccessResolvedColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_access_resolved_color")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnGetAccessResolvedColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_access_resolved_depth")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnSetAccessResolvedDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_access_resolved_depth")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnGetAccessResolvedDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_needs_motion_vectors")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnSetNeedsMotionVectors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_needs_motion_vectors")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnGetNeedsMotionVectors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_needs_normal_roughness")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnSetNeedsNormalRoughness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_needs_normal_roughness")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnGetNeedsNormalRoughness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_needs_separate_specular")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnSetNeedsSeparateSpecular = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_needs_separate_specular")
		defer methodName.Destroy()
		ptrsForCompositorEffect.fnGetNeedsSeparateSpecular = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type CompositorEffect struct {
	Resource
}

func (me *CompositorEffect) BaseClass() string {
	return "CompositorEffect"
}

func NewCompositorEffect() *CompositorEffect {
	str := StringNameFromStr("CompositorEffect") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CompositorEffect{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CompositorEffectEffectCallbackType int

const (
	CompositorEffectEffectCallbackTypeEffectCallbackTypePreOpaque       CompositorEffectEffectCallbackType = 0
	CompositorEffectEffectCallbackTypeEffectCallbackTypePostOpaque      CompositorEffectEffectCallbackType = 1
	CompositorEffectEffectCallbackTypeEffectCallbackTypePostSky         CompositorEffectEffectCallbackType = 2
	CompositorEffectEffectCallbackTypeEffectCallbackTypePreTransparent  CompositorEffectEffectCallbackType = 3
	CompositorEffectEffectCallbackTypeEffectCallbackTypePostTransparent CompositorEffectEffectCallbackType = 4
	CompositorEffectEffectCallbackTypeEffectCallbackTypeMax             CompositorEffectEffectCallbackType = 5
)

func (me *CompositorEffect) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CompositorEffect) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CompositorEffect) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CompositorEffect) SetEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CompositorEffect) GetEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnGetEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CompositorEffect) SetEffectCallbackType(effect_callback_type CompositorEffectEffectCallbackType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&effect_callback_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnSetEffectCallbackType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CompositorEffect) GetEffectCallbackType() CompositorEffectEffectCallbackType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CompositorEffectEffectCallbackType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnGetEffectCallbackType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CompositorEffect) SetAccessResolvedColor(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnSetAccessResolvedColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CompositorEffect) GetAccessResolvedColor() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnGetAccessResolvedColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CompositorEffect) SetAccessResolvedDepth(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnSetAccessResolvedDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CompositorEffect) GetAccessResolvedDepth() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnGetAccessResolvedDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CompositorEffect) SetNeedsMotionVectors(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnSetNeedsMotionVectors), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CompositorEffect) GetNeedsMotionVectors() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnGetNeedsMotionVectors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CompositorEffect) SetNeedsNormalRoughness(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnSetNeedsNormalRoughness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CompositorEffect) GetNeedsNormalRoughness() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnGetNeedsNormalRoughness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CompositorEffect) SetNeedsSeparateSpecular(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnSetNeedsSeparateSpecular), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CompositorEffect) GetNeedsSeparateSpecular() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositorEffect.fnGetNeedsSeparateSpecular), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
