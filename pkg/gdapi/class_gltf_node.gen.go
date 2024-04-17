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

type ptrsForGLTFNodeList struct {
	fnGetOriginalName   gdc.MethodBindPtr
	fnSetOriginalName   gdc.MethodBindPtr
	fnGetParent         gdc.MethodBindPtr
	fnSetParent         gdc.MethodBindPtr
	fnGetHeight         gdc.MethodBindPtr
	fnSetHeight         gdc.MethodBindPtr
	fnGetXform          gdc.MethodBindPtr
	fnSetXform          gdc.MethodBindPtr
	fnGetMesh           gdc.MethodBindPtr
	fnSetMesh           gdc.MethodBindPtr
	fnGetCamera         gdc.MethodBindPtr
	fnSetCamera         gdc.MethodBindPtr
	fnGetSkin           gdc.MethodBindPtr
	fnSetSkin           gdc.MethodBindPtr
	fnGetSkeleton       gdc.MethodBindPtr
	fnSetSkeleton       gdc.MethodBindPtr
	fnGetPosition       gdc.MethodBindPtr
	fnSetPosition       gdc.MethodBindPtr
	fnGetRotation       gdc.MethodBindPtr
	fnSetRotation       gdc.MethodBindPtr
	fnGetScale          gdc.MethodBindPtr
	fnSetScale          gdc.MethodBindPtr
	fnGetChildren       gdc.MethodBindPtr
	fnSetChildren       gdc.MethodBindPtr
	fnGetLight          gdc.MethodBindPtr
	fnSetLight          gdc.MethodBindPtr
	fnGetAdditionalData gdc.MethodBindPtr
	fnSetAdditionalData gdc.MethodBindPtr
}

var ptrsForGLTFNode ptrsForGLTFNodeList

func initGLTFNodePtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFNode")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_original_name")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetOriginalName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("set_original_name")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetOriginalName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_parent")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_parent")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_height")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_height")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_xform")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetXform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4183770049))
	}
	{
		methodName := StringNameFromStr("set_xform")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetXform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
	}
	{
		methodName := StringNameFromStr("get_mesh")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_mesh")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_camera")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetCamera = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_camera")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetCamera = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_skin")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_skin")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_skeleton")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_skeleton")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3783033775))
	}
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_rotation")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2916281908))
	}
	{
		methodName := StringNameFromStr("set_rotation")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1727505552))
	}
	{
		methodName := StringNameFromStr("get_scale")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3783033775))
	}
	{
		methodName := StringNameFromStr("set_scale")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_children")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetChildren = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("set_children")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetChildren = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
	}
	{
		methodName := StringNameFromStr("get_light")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetLight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_light")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetLight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_additional_data")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnGetAdditionalData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2138907829))
	}
	{
		methodName := StringNameFromStr("set_additional_data")
		defer methodName.Destroy()
		ptrsForGLTFNode.fnSetAdditionalData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}

}

type GLTFNode struct {
	Resource
}

func (me *GLTFNode) BaseClass() string {
	return "GLTFNode"
}

func NewGLTFNode() *GLTFNode {
	str := StringNameFromStr("GLTFNode") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFNode{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFNode) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFNode) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFNode) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GLTFNode) GetOriginalName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetOriginalName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFNode) SetOriginalName(original_name String) {
	cargs := []gdc.ConstTypePtr{original_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetOriginalName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetParent() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFNode) SetParent(parent int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetParent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetHeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFNode) SetHeight(height int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetXform() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetXform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFNode) SetXform(xform Transform3D) {
	cargs := []gdc.ConstTypePtr{xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetXform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetMesh() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFNode) SetMesh(mesh int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mesh)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetCamera() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetCamera), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFNode) SetCamera(camera int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&camera)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetCamera), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetSkin() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetSkin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFNode) SetSkin(skin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetSkin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetSkeleton() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFNode) SetSkeleton(skeleton int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skeleton)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetSkeleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetPosition() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFNode) SetPosition(position Vector3) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetRotation() Quaternion {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewQuaternion()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFNode) SetRotation(rotation Quaternion) {
	cargs := []gdc.ConstTypePtr{rotation.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetScale() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFNode) SetScale(scale Vector3) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetChildren() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetChildren), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFNode) SetChildren(children PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{children.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetChildren), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetLight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetLight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFNode) SetLight(light int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetLight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFNode) GetAdditionalData(extension_name StringName) Variant {
	cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnGetAdditionalData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFNode) SetAdditionalData(extension_name StringName, additional_data Variant) {
	cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr(), additional_data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFNode.fnSetAdditionalData), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
