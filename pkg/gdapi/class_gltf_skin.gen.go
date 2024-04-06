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

type ptrsForGLTFSkinList struct {
	fnGetSkinRoot       gdc.MethodBindPtr
	fnSetSkinRoot       gdc.MethodBindPtr
	fnGetJointsOriginal gdc.MethodBindPtr
	fnSetJointsOriginal gdc.MethodBindPtr
	fnGetInverseBinds   gdc.MethodBindPtr
	fnSetInverseBinds   gdc.MethodBindPtr
	fnGetJoints         gdc.MethodBindPtr
	fnSetJoints         gdc.MethodBindPtr
	fnGetNonJoints      gdc.MethodBindPtr
	fnSetNonJoints      gdc.MethodBindPtr
	fnGetRoots          gdc.MethodBindPtr
	fnSetRoots          gdc.MethodBindPtr
	fnGetSkeleton       gdc.MethodBindPtr
	fnSetSkeleton       gdc.MethodBindPtr
	fnGetJointIToBoneI  gdc.MethodBindPtr
	fnSetJointIToBoneI  gdc.MethodBindPtr
	fnGetJointIToName   gdc.MethodBindPtr
	fnSetJointIToName   gdc.MethodBindPtr
	fnGetGodotSkin      gdc.MethodBindPtr
	fnSetGodotSkin      gdc.MethodBindPtr
}

var ptrsForGLTFSkin ptrsForGLTFSkinList

func initGLTFSkinPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFSkin")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_skin_root")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetSkinRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_skin_root")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetSkinRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_joints_original")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetJointsOriginal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("set_joints_original")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetJointsOriginal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_inverse_binds")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetInverseBinds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("set_inverse_binds")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetInverseBinds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_joints")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetJoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("set_joints")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetJoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_non_joints")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetNonJoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("set_non_joints")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetNonJoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_roots")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetRoots = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("set_roots")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetRoots = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_skeleton")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_skeleton")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_joint_i_to_bone_i")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetJointIToBoneI = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
	}
	{
		methodName := StringNameFromStr("set_joint_i_to_bone_i")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetJointIToBoneI = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
	}
	{
		methodName := StringNameFromStr("get_joint_i_to_name")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetJointIToName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
	}
	{
		methodName := StringNameFromStr("set_joint_i_to_name")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetJointIToName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
	}
	{
		methodName := StringNameFromStr("get_godot_skin")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnGetGodotSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1032037385))
	}
	{
		methodName := StringNameFromStr("set_godot_skin")
		defer methodName.Destroy()
		ptrsForGLTFSkin.fnSetGodotSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3971435618))
	}
}

type GLTFSkin struct {
	Resource
}

func (me *GLTFSkin) BaseClass() string {
	return "GLTFSkin"
}

func NewGLTFSkin() *GLTFSkin {
	str := StringNameFromStr("GLTFSkin") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFSkin{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFSkin) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFSkin) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFSkin) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GLTFSkin) GetSkinRoot() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetSkinRoot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFSkin) SetSkinRoot(skin_root int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skin_root)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetSkinRoot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetJointsOriginal() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetJointsOriginal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkin) SetJointsOriginal(joints_original PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{joints_original.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetJointsOriginal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetInverseBinds() []Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetInverseBinds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Transform3D](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *GLTFSkin) SetInverseBinds(inverse_binds []Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inverse_binds)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetInverseBinds), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetJoints() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetJoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkin) SetJoints(joints PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{joints.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetJoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetNonJoints() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetNonJoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkin) SetNonJoints(non_joints PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{non_joints.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetNonJoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetRoots() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetRoots), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkin) SetRoots(roots PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{roots.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetRoots), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetSkeleton() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFSkin) SetSkeleton(skeleton int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skeleton)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetSkeleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetJointIToBoneI() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetJointIToBoneI), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkin) SetJointIToBoneI(joint_i_to_bone_i Dictionary) {
	cargs := []gdc.ConstTypePtr{joint_i_to_bone_i.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetJointIToBoneI), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetJointIToName() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetJointIToName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkin) SetJointIToName(joint_i_to_name Dictionary) {
	cargs := []gdc.ConstTypePtr{joint_i_to_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetJointIToName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFSkin) GetGodotSkin() Skin {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnGetGodotSkin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFSkin) SetGodotSkin(godot_skin Skin) {
	cargs := []gdc.ConstTypePtr{godot_skin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSkin.fnSetGodotSkin), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
