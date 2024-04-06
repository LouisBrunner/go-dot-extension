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

type ptrsForMultiMeshList struct {
	fnSetMesh                 gdc.MethodBindPtr
	fnGetMesh                 gdc.MethodBindPtr
	fnSetUseColors            gdc.MethodBindPtr
	fnIsUsingColors           gdc.MethodBindPtr
	fnSetUseCustomData        gdc.MethodBindPtr
	fnIsUsingCustomData       gdc.MethodBindPtr
	fnSetTransformFormat      gdc.MethodBindPtr
	fnGetTransformFormat      gdc.MethodBindPtr
	fnSetInstanceCount        gdc.MethodBindPtr
	fnGetInstanceCount        gdc.MethodBindPtr
	fnSetVisibleInstanceCount gdc.MethodBindPtr
	fnGetVisibleInstanceCount gdc.MethodBindPtr
	fnSetInstanceTransform    gdc.MethodBindPtr
	fnSetInstanceTransform2D  gdc.MethodBindPtr
	fnGetInstanceTransform    gdc.MethodBindPtr
	fnGetInstanceTransform2D  gdc.MethodBindPtr
	fnSetInstanceColor        gdc.MethodBindPtr
	fnGetInstanceColor        gdc.MethodBindPtr
	fnSetInstanceCustomData   gdc.MethodBindPtr
	fnGetInstanceCustomData   gdc.MethodBindPtr
	fnGetAabb                 gdc.MethodBindPtr
	fnGetBuffer               gdc.MethodBindPtr
	fnSetBuffer               gdc.MethodBindPtr
}

var ptrsForMultiMesh ptrsForMultiMeshList

func initMultiMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MultiMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mesh")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194775623))
	}
	{
		methodName := StringNameFromStr("get_mesh")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1808005922))
	}
	{
		methodName := StringNameFromStr("set_use_colors")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetUseColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_colors")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnIsUsingColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_custom_data")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetUseCustomData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_custom_data")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnIsUsingCustomData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_transform_format")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetTransformFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2404750322))
	}
	{
		methodName := StringNameFromStr("get_transform_format")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetTransformFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2444156481))
	}
	{
		methodName := StringNameFromStr("set_instance_count")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetInstanceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_instance_count")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetInstanceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_visible_instance_count")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetVisibleInstanceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_visible_instance_count")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetVisibleInstanceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_instance_transform")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetInstanceTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("set_instance_transform_2d")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetInstanceTransform2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 30160968))
	}
	{
		methodName := StringNameFromStr("get_instance_transform")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetInstanceTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("get_instance_transform_2d")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetInstanceTransform2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3836996910))
	}
	{
		methodName := StringNameFromStr("set_instance_color")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetInstanceColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_instance_color")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetInstanceColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("set_instance_custom_data")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetInstanceCustomData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_instance_custom_data")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetInstanceCustomData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("get_aabb")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
	}
	{
		methodName := StringNameFromStr("get_buffer")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnGetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675695659))
	}
	{
		methodName := StringNameFromStr("set_buffer")
		defer methodName.Destroy()
		ptrsForMultiMesh.fnSetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
	}
}

type MultiMesh struct {
	Resource
}

func (me *MultiMesh) BaseClass() string {
	return "MultiMesh"
}

func NewMultiMesh() *MultiMesh {
	str := StringNameFromStr("MultiMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MultiMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type MultiMeshTransformFormat int

const (
	MultiMeshTransformFormatTransform2D MultiMeshTransformFormat = 0
	MultiMeshTransformFormatTransform3D MultiMeshTransformFormat = 1
)

func (me *MultiMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MultiMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MultiMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MultiMesh) SetMesh(mesh Mesh) {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) GetMesh() Mesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiMesh) SetUseColors(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetUseColors), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) IsUsingColors() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnIsUsingColors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiMesh) SetUseCustomData(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetUseCustomData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) IsUsingCustomData() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnIsUsingCustomData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiMesh) SetTransformFormat(format MultiMeshTransformFormat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetTransformFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) GetTransformFormat() MultiMeshTransformFormat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret MultiMeshTransformFormat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetTransformFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *MultiMesh) SetInstanceCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetInstanceCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) GetInstanceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetInstanceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiMesh) SetVisibleInstanceCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetVisibleInstanceCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) GetVisibleInstanceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetVisibleInstanceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MultiMesh) SetInstanceTransform(instance int64, transform Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetInstanceTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) SetInstanceTransform2D(instance int64, transform Transform2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetInstanceTransform2D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) GetInstanceTransform(instance int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&instance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetInstanceTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiMesh) GetInstanceTransform2D(instance int64) Transform2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()
	pinner.Pin(&instance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetInstanceTransform2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiMesh) SetInstanceColor(instance int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetInstanceColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) GetInstanceColor(instance int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&instance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetInstanceColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiMesh) SetInstanceCustomData(instance int64, custom_data Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance), custom_data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetInstanceCustomData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMesh) GetInstanceCustomData(instance int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&instance)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetInstanceCustomData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiMesh) GetAabb() AABB {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAABB()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiMesh) GetBuffer() PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnGetBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiMesh) SetBuffer(buffer PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMesh.fnSetBuffer), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
