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

type ptrsForBone2DList struct {
	fnSetRest                        gdc.MethodBindPtr
	fnGetRest                        gdc.MethodBindPtr
	fnApplyRest                      gdc.MethodBindPtr
	fnGetSkeletonRest                gdc.MethodBindPtr
	fnGetIndexInSkeleton             gdc.MethodBindPtr
	fnSetAutocalculateLengthAndAngle gdc.MethodBindPtr
	fnGetAutocalculateLengthAndAngle gdc.MethodBindPtr
	fnSetLength                      gdc.MethodBindPtr
	fnGetLength                      gdc.MethodBindPtr
	fnSetBoneAngle                   gdc.MethodBindPtr
	fnGetBoneAngle                   gdc.MethodBindPtr
}

var ptrsForBone2D ptrsForBone2DList

func initBone2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Bone2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_rest")
		defer methodName.Destroy()
		ptrsForBone2D.fnSetRest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
	}
	{
		methodName := StringNameFromStr("get_rest")
		defer methodName.Destroy()
		ptrsForBone2D.fnGetRest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("apply_rest")
		defer methodName.Destroy()
		ptrsForBone2D.fnApplyRest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_skeleton_rest")
		defer methodName.Destroy()
		ptrsForBone2D.fnGetSkeletonRest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_index_in_skeleton")
		defer methodName.Destroy()
		ptrsForBone2D.fnGetIndexInSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_autocalculate_length_and_angle")
		defer methodName.Destroy()
		ptrsForBone2D.fnSetAutocalculateLengthAndAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_autocalculate_length_and_angle")
		defer methodName.Destroy()
		ptrsForBone2D.fnGetAutocalculateLengthAndAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_length")
		defer methodName.Destroy()
		ptrsForBone2D.fnSetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_length")
		defer methodName.Destroy()
		ptrsForBone2D.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_bone_angle")
		defer methodName.Destroy()
		ptrsForBone2D.fnSetBoneAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bone_angle")
		defer methodName.Destroy()
		ptrsForBone2D.fnGetBoneAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type Bone2D struct {
	Node2D
}

func (me *Bone2D) BaseClass() string {
	return "Bone2D"
}

func NewBone2D() *Bone2D {
	str := StringNameFromStr("Bone2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Bone2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Bone2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Bone2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Bone2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Bone2D) SetRest(rest Transform2D) {
	cargs := []gdc.ConstTypePtr{rest.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnSetRest), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Bone2D) GetRest() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnGetRest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Bone2D) ApplyRest() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnApplyRest), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Bone2D) GetSkeletonRest() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnGetSkeletonRest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Bone2D) GetIndexInSkeleton() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnGetIndexInSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Bone2D) SetAutocalculateLengthAndAngle(auto_calculate bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&auto_calculate)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnSetAutocalculateLengthAndAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Bone2D) GetAutocalculateLengthAndAngle() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnGetAutocalculateLengthAndAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Bone2D) SetLength(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnSetLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Bone2D) GetLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Bone2D) SetBoneAngle(angle float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnSetBoneAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Bone2D) GetBoneAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBone2D.fnGetBoneAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
