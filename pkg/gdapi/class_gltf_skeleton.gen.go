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

type ptrsForGLTFSkeletonList struct {
	fnGetJoints              gdc.MethodBindPtr
	fnSetJoints              gdc.MethodBindPtr
	fnGetRoots               gdc.MethodBindPtr
	fnSetRoots               gdc.MethodBindPtr
	fnGetGodotSkeleton       gdc.MethodBindPtr
	fnGetUniqueNames         gdc.MethodBindPtr
	fnSetUniqueNames         gdc.MethodBindPtr
	fnGetGodotBoneNode       gdc.MethodBindPtr
	fnSetGodotBoneNode       gdc.MethodBindPtr
	fnGetBoneAttachmentCount gdc.MethodBindPtr
	fnGetBoneAttachment      gdc.MethodBindPtr
}

var ptrsForGLTFSkeleton ptrsForGLTFSkeletonList

func initGLTFSkeletonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFSkeleton")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_joints")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnGetJoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("set_joints")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnSetJoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_roots")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnGetRoots = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("set_roots")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnSetRoots = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_godot_skeleton")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnGetGodotSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1814733083))
	}
	{
		methodName := StringNameFromStr("get_unique_names")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnGetUniqueNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("set_unique_names")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnSetUniqueNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_godot_bone_node")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnGetGodotBoneNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
	}
	{
		methodName := StringNameFromStr("set_godot_bone_node")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnSetGodotBoneNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
	}
	{
		methodName := StringNameFromStr("get_bone_attachment_count")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnGetBoneAttachmentCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_bone_attachment")
		defer methodName.Destroy()
		ptrsForGLTFSkeleton.fnGetBoneAttachment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 945440495))
	}
}

type GLTFSkeleton struct {
	Resource
}

func (me *GLTFSkeleton) BaseClass() string {
	return "GLTFSkeleton"
}

func NewGLTFSkeleton() *GLTFSkeleton {
	str := StringNameFromStr("GLTFSkeleton") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFSkeleton{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFSkeleton) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFSkeleton) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFSkeleton) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GLTFSkeleton) GetJoints() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnGetJoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkeleton) SetJoints(joints PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{joints.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnSetJoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkeleton) GetRoots() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnGetRoots), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkeleton) SetRoots(roots PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{roots.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnSetRoots), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkeleton) GetGodotSkeleton() Skeleton3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkeleton3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnGetGodotSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkeleton) GetUniqueNames() []String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnGetUniqueNames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[String](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *GLTFSkeleton) SetUniqueNames(unique_names []String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unique_names)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnSetUniqueNames), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkeleton) GetGodotBoneNode() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnGetGodotBoneNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkeleton) SetGodotBoneNode(godot_bone_node Dictionary) {
	cargs := []gdc.ConstTypePtr{godot_bone_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnSetGodotBoneNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkeleton) GetBoneAttachmentCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnGetBoneAttachmentCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFSkeleton) GetBoneAttachment(idx int64) BoneAttachment3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBoneAttachment3D()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkeleton.fnGetBoneAttachment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
