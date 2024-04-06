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

type ptrsForSoftBody3DList struct {
	fnGetPhysicsRid                gdc.MethodBindPtr
	fnSetCollisionMask             gdc.MethodBindPtr
	fnGetCollisionMask             gdc.MethodBindPtr
	fnSetCollisionLayer            gdc.MethodBindPtr
	fnGetCollisionLayer            gdc.MethodBindPtr
	fnSetCollisionMaskValue        gdc.MethodBindPtr
	fnGetCollisionMaskValue        gdc.MethodBindPtr
	fnSetCollisionLayerValue       gdc.MethodBindPtr
	fnGetCollisionLayerValue       gdc.MethodBindPtr
	fnSetParentCollisionIgnore     gdc.MethodBindPtr
	fnGetParentCollisionIgnore     gdc.MethodBindPtr
	fnSetDisableMode               gdc.MethodBindPtr
	fnGetDisableMode               gdc.MethodBindPtr
	fnGetCollisionExceptions       gdc.MethodBindPtr
	fnAddCollisionExceptionWith    gdc.MethodBindPtr
	fnRemoveCollisionExceptionWith gdc.MethodBindPtr
	fnSetSimulationPrecision       gdc.MethodBindPtr
	fnGetSimulationPrecision       gdc.MethodBindPtr
	fnSetTotalMass                 gdc.MethodBindPtr
	fnGetTotalMass                 gdc.MethodBindPtr
	fnSetLinearStiffness           gdc.MethodBindPtr
	fnGetLinearStiffness           gdc.MethodBindPtr
	fnSetPressureCoefficient       gdc.MethodBindPtr
	fnGetPressureCoefficient       gdc.MethodBindPtr
	fnSetDampingCoefficient        gdc.MethodBindPtr
	fnGetDampingCoefficient        gdc.MethodBindPtr
	fnSetDragCoefficient           gdc.MethodBindPtr
	fnGetDragCoefficient           gdc.MethodBindPtr
	fnGetPointTransform            gdc.MethodBindPtr
	fnSetPointPinned               gdc.MethodBindPtr
	fnIsPointPinned                gdc.MethodBindPtr
	fnSetRayPickable               gdc.MethodBindPtr
	fnIsRayPickable                gdc.MethodBindPtr
}

var ptrsForSoftBody3D ptrsForSoftBody3DList

func initSoftBody3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SoftBody3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_physics_rid")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetPhysicsRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_layer")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_layer")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_mask_value")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_mask_value")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_collision_layer_value")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_layer_value")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_parent_collision_ignore")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetParentCollisionIgnore = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_parent_collision_ignore")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetParentCollisionIgnore = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_disable_mode")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetDisableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1104158384))
	}
	{
		methodName := StringNameFromStr("get_disable_mode")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetDisableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4135042476))
	}
	{
		methodName := StringNameFromStr("get_collision_exceptions")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetCollisionExceptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("add_collision_exception_with")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnAddCollisionExceptionWith = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("remove_collision_exception_with")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnRemoveCollisionExceptionWith = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("set_simulation_precision")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetSimulationPrecision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_simulation_precision")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetSimulationPrecision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_total_mass")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetTotalMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_total_mass")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetTotalMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_linear_stiffness")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetLinearStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_linear_stiffness")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetLinearStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_pressure_coefficient")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetPressureCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pressure_coefficient")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetPressureCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_damping_coefficient")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetDampingCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_damping_coefficient")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetDampingCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_drag_coefficient")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetDragCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_drag_coefficient")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetDragCoefficient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("get_point_transform")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnGetPointTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 871989493))
	}
	{
		methodName := StringNameFromStr("set_point_pinned")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetPointPinned = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814935226))
	}
	{
		methodName := StringNameFromStr("is_point_pinned")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnIsPointPinned = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_ray_pickable")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnSetRayPickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_ray_pickable")
		defer methodName.Destroy()
		ptrsForSoftBody3D.fnIsRayPickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type SoftBody3D struct {
	MeshInstance3D
}

func (me *SoftBody3D) BaseClass() string {
	return "SoftBody3D"
}

func NewSoftBody3D() *SoftBody3D {
	str := StringNameFromStr("SoftBody3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SoftBody3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type SoftBody3DDisableMode int

const (
	SoftBody3DDisableModeDisableModeRemove     SoftBody3DDisableMode = 0
	SoftBody3DDisableModeDisableModeKeepActive SoftBody3DDisableMode = 1
)

func (me *SoftBody3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SoftBody3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SoftBody3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SoftBody3D) GetPhysicsRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetPhysicsRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SoftBody3D) SetCollisionMask(collision_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetCollisionLayer(collision_layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetCollisionLayer() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetCollisionMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetCollisionMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetCollisionLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetCollisionLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetParentCollisionIgnore(parent_collision_ignore NodePath) {
	cargs := []gdc.ConstTypePtr{parent_collision_ignore.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetParentCollisionIgnore), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetParentCollisionIgnore() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetParentCollisionIgnore), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SoftBody3D) SetDisableMode(mode SoftBody3DDisableMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetDisableMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetDisableMode() SoftBody3DDisableMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret SoftBody3DDisableMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetDisableMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SoftBody3D) GetCollisionExceptions() []PhysicsBody3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetCollisionExceptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[PhysicsBody3D](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *SoftBody3D) AddCollisionExceptionWith(body Node) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnAddCollisionExceptionWith), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) RemoveCollisionExceptionWith(body Node) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnRemoveCollisionExceptionWith), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) SetSimulationPrecision(simulation_precision int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&simulation_precision)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetSimulationPrecision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetSimulationPrecision() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetSimulationPrecision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetTotalMass(mass float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetTotalMass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetTotalMass() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetTotalMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetLinearStiffness(linear_stiffness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_stiffness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetLinearStiffness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetLinearStiffness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetLinearStiffness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetPressureCoefficient(pressure_coefficient float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure_coefficient)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetPressureCoefficient), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetPressureCoefficient() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetPressureCoefficient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetDampingCoefficient(damping_coefficient float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&damping_coefficient)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetDampingCoefficient), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetDampingCoefficient() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetDampingCoefficient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetDragCoefficient(drag_coefficient float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&drag_coefficient)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetDragCoefficient), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) GetDragCoefficient() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetDragCoefficient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) GetPointTransform(point_index int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&point_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnGetPointTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SoftBody3D) SetPointPinned(point_index int64, pinned bool, attachment_path NodePath) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_index), gdc.ConstTypePtr(&pinned), attachment_path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetPointPinned), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) IsPointPinned(point_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&point_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnIsPointPinned), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SoftBody3D) SetRayPickable(ray_pickable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ray_pickable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnSetRayPickable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SoftBody3D) IsRayPickable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSoftBody3D.fnIsRayPickable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
