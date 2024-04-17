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

type ptrsForGLTFPhysicsShapeList struct {
	fnFromNode        gdc.MethodBindPtr
	fnToNode          gdc.MethodBindPtr
	fnFromResource    gdc.MethodBindPtr
	fnToResource      gdc.MethodBindPtr
	fnFromDictionary  gdc.MethodBindPtr
	fnToDictionary    gdc.MethodBindPtr
	fnGetShapeType    gdc.MethodBindPtr
	fnSetShapeType    gdc.MethodBindPtr
	fnGetSize         gdc.MethodBindPtr
	fnSetSize         gdc.MethodBindPtr
	fnGetRadius       gdc.MethodBindPtr
	fnSetRadius       gdc.MethodBindPtr
	fnGetHeight       gdc.MethodBindPtr
	fnSetHeight       gdc.MethodBindPtr
	fnGetIsTrigger    gdc.MethodBindPtr
	fnSetIsTrigger    gdc.MethodBindPtr
	fnGetMeshIndex    gdc.MethodBindPtr
	fnSetMeshIndex    gdc.MethodBindPtr
	fnGetImporterMesh gdc.MethodBindPtr
	fnSetImporterMesh gdc.MethodBindPtr
}

var ptrsForGLTFPhysicsShape ptrsForGLTFPhysicsShapeList

func initGLTFPhysicsShapePtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFPhysicsShape")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("from_node")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnFromNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3613751275))
	}
	{
		methodName := StringNameFromStr("to_node")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnToNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 563689933))
	}
	{
		methodName := StringNameFromStr("from_resource")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnFromResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3845569786))
	}
	{
		methodName := StringNameFromStr("to_resource")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnToResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1913542110))
	}
	{
		methodName := StringNameFromStr("from_dictionary")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnFromDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2390691823))
	}
	{
		methodName := StringNameFromStr("to_dictionary")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnToDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_shape_type")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnGetShapeType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_shape_type")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnSetShapeType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_height")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_height")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_is_trigger")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnGetIsTrigger = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_is_trigger")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnSetIsTrigger = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_mesh_index")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnGetMeshIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_mesh_index")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnSetMeshIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_importer_mesh")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnGetImporterMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3161779525))
	}
	{
		methodName := StringNameFromStr("set_importer_mesh")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsShape.fnSetImporterMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2255166972))
	}

}

type GLTFPhysicsShape struct {
	Resource
}

func (me *GLTFPhysicsShape) BaseClass() string {
	return "GLTFPhysicsShape"
}

func NewGLTFPhysicsShape() *GLTFPhysicsShape {
	str := StringNameFromStr("GLTFPhysicsShape") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFPhysicsShape{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFPhysicsShape) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFPhysicsShape) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFPhysicsShape) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func GLTFPhysicsShapeFromNode(shape_node CollisionShape3D) GLTFPhysicsShape {
	cargs := []gdc.ConstTypePtr{shape_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFPhysicsShape()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnFromNode), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsShape) ToNode(cache_shapes bool) CollisionShape3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_shapes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCollisionShape3D()
	pinner.Pin(&cache_shapes)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnToNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func GLTFPhysicsShapeFromResource(shape_resource Shape3D) GLTFPhysicsShape {
	cargs := []gdc.ConstTypePtr{shape_resource.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFPhysicsShape()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnFromResource), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsShape) ToResource(cache_shapes bool) Shape3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_shapes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewShape3D()
	pinner.Pin(&cache_shapes)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnToResource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func GLTFPhysicsShapeFromDictionary(dictionary Dictionary) GLTFPhysicsShape {
	cargs := []gdc.ConstTypePtr{dictionary.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFPhysicsShape()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnFromDictionary), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsShape) ToDictionary() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnToDictionary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsShape) GetShapeType() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnGetShapeType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsShape) SetShapeType(shape_type String) {
	cargs := []gdc.ConstTypePtr{shape_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnSetShapeType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsShape) GetSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsShape) SetSize(size Vector3) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsShape) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFPhysicsShape) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsShape) GetHeight() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFPhysicsShape) SetHeight(height float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsShape) GetIsTrigger() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnGetIsTrigger), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFPhysicsShape) SetIsTrigger(is_trigger bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&is_trigger)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnSetIsTrigger), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsShape) GetMeshIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnGetMeshIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFPhysicsShape) SetMeshIndex(mesh_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mesh_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnSetMeshIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsShape) GetImporterMesh() ImporterMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImporterMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnGetImporterMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsShape) SetImporterMesh(importer_mesh ImporterMesh) {
	cargs := []gdc.ConstTypePtr{importer_mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsShape.fnSetImporterMesh), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
