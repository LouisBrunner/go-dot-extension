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

type ptrsForGLTFAnimationList struct {
	fnGetOriginalName   gdc.MethodBindPtr
	fnSetOriginalName   gdc.MethodBindPtr
	fnGetLoop           gdc.MethodBindPtr
	fnSetLoop           gdc.MethodBindPtr
	fnGetAdditionalData gdc.MethodBindPtr
	fnSetAdditionalData gdc.MethodBindPtr
}

var ptrsForGLTFAnimation ptrsForGLTFAnimationList

func initGLTFAnimationPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFAnimation")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_original_name")
		defer methodName.Destroy()
		ptrsForGLTFAnimation.fnGetOriginalName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("set_original_name")
		defer methodName.Destroy()
		ptrsForGLTFAnimation.fnSetOriginalName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_loop")
		defer methodName.Destroy()
		ptrsForGLTFAnimation.fnGetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_loop")
		defer methodName.Destroy()
		ptrsForGLTFAnimation.fnSetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_additional_data")
		defer methodName.Destroy()
		ptrsForGLTFAnimation.fnGetAdditionalData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2138907829))
	}
	{
		methodName := StringNameFromStr("set_additional_data")
		defer methodName.Destroy()
		ptrsForGLTFAnimation.fnSetAdditionalData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}

}

type GLTFAnimation struct {
	Resource
}

func (me *GLTFAnimation) BaseClass() string {
	return "GLTFAnimation"
}

func NewGLTFAnimation() *GLTFAnimation {
	str := StringNameFromStr("GLTFAnimation") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFAnimation{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFAnimation) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFAnimation) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFAnimation) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GLTFAnimation) GetOriginalName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAnimation.fnGetOriginalName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFAnimation) SetOriginalName(original_name String) {
	cargs := []gdc.ConstTypePtr{original_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAnimation.fnSetOriginalName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFAnimation) GetLoop() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAnimation.fnGetLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFAnimation) SetLoop(loop bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAnimation.fnSetLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFAnimation) GetAdditionalData(extension_name StringName) Variant {
	cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAnimation.fnGetAdditionalData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFAnimation) SetAdditionalData(extension_name StringName, additional_data Variant) {
	cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr(), additional_data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAnimation.fnSetAdditionalData), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
