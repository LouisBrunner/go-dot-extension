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

type ptrsForRDPipelineDepthStencilStateList struct {
	fnSetEnableDepthTest      gdc.MethodBindPtr
	fnGetEnableDepthTest      gdc.MethodBindPtr
	fnSetEnableDepthWrite     gdc.MethodBindPtr
	fnGetEnableDepthWrite     gdc.MethodBindPtr
	fnSetDepthCompareOperator gdc.MethodBindPtr
	fnGetDepthCompareOperator gdc.MethodBindPtr
	fnSetEnableDepthRange     gdc.MethodBindPtr
	fnGetEnableDepthRange     gdc.MethodBindPtr
	fnSetDepthRangeMin        gdc.MethodBindPtr
	fnGetDepthRangeMin        gdc.MethodBindPtr
	fnSetDepthRangeMax        gdc.MethodBindPtr
	fnGetDepthRangeMax        gdc.MethodBindPtr
	fnSetEnableStencil        gdc.MethodBindPtr
	fnGetEnableStencil        gdc.MethodBindPtr
	fnSetFrontOpFail          gdc.MethodBindPtr
	fnGetFrontOpFail          gdc.MethodBindPtr
	fnSetFrontOpPass          gdc.MethodBindPtr
	fnGetFrontOpPass          gdc.MethodBindPtr
	fnSetFrontOpDepthFail     gdc.MethodBindPtr
	fnGetFrontOpDepthFail     gdc.MethodBindPtr
	fnSetFrontOpCompare       gdc.MethodBindPtr
	fnGetFrontOpCompare       gdc.MethodBindPtr
	fnSetFrontOpCompareMask   gdc.MethodBindPtr
	fnGetFrontOpCompareMask   gdc.MethodBindPtr
	fnSetFrontOpWriteMask     gdc.MethodBindPtr
	fnGetFrontOpWriteMask     gdc.MethodBindPtr
	fnSetFrontOpReference     gdc.MethodBindPtr
	fnGetFrontOpReference     gdc.MethodBindPtr
	fnSetBackOpFail           gdc.MethodBindPtr
	fnGetBackOpFail           gdc.MethodBindPtr
	fnSetBackOpPass           gdc.MethodBindPtr
	fnGetBackOpPass           gdc.MethodBindPtr
	fnSetBackOpDepthFail      gdc.MethodBindPtr
	fnGetBackOpDepthFail      gdc.MethodBindPtr
	fnSetBackOpCompare        gdc.MethodBindPtr
	fnGetBackOpCompare        gdc.MethodBindPtr
	fnSetBackOpCompareMask    gdc.MethodBindPtr
	fnGetBackOpCompareMask    gdc.MethodBindPtr
	fnSetBackOpWriteMask      gdc.MethodBindPtr
	fnGetBackOpWriteMask      gdc.MethodBindPtr
	fnSetBackOpReference      gdc.MethodBindPtr
	fnGetBackOpReference      gdc.MethodBindPtr
}

var ptrsForRDPipelineDepthStencilState ptrsForRDPipelineDepthStencilStateList

func initRDPipelineDepthStencilStatePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RDPipelineDepthStencilState")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_enable_depth_test")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetEnableDepthTest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_depth_test")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetEnableDepthTest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_enable_depth_write")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetEnableDepthWrite = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_depth_write")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetEnableDepthWrite = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_depth_compare_operator")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetDepthCompareOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2573711505))
	}
	{
		methodName := StringNameFromStr("get_depth_compare_operator")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetDepthCompareOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 269730778))
	}
	{
		methodName := StringNameFromStr("set_enable_depth_range")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetEnableDepthRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_depth_range")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetEnableDepthRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_depth_range_min")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetDepthRangeMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth_range_min")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetDepthRangeMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_depth_range_max")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetDepthRangeMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth_range_max")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetDepthRangeMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_enable_stencil")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetEnableStencil = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_stencil")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetEnableStencil = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_front_op_fail")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetFrontOpFail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2092799566))
	}
	{
		methodName := StringNameFromStr("get_front_op_fail")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetFrontOpFail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1714732389))
	}
	{
		methodName := StringNameFromStr("set_front_op_pass")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetFrontOpPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2092799566))
	}
	{
		methodName := StringNameFromStr("get_front_op_pass")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetFrontOpPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1714732389))
	}
	{
		methodName := StringNameFromStr("set_front_op_depth_fail")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetFrontOpDepthFail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2092799566))
	}
	{
		methodName := StringNameFromStr("get_front_op_depth_fail")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetFrontOpDepthFail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1714732389))
	}
	{
		methodName := StringNameFromStr("set_front_op_compare")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetFrontOpCompare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2573711505))
	}
	{
		methodName := StringNameFromStr("get_front_op_compare")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetFrontOpCompare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 269730778))
	}
	{
		methodName := StringNameFromStr("set_front_op_compare_mask")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetFrontOpCompareMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_front_op_compare_mask")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetFrontOpCompareMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_front_op_write_mask")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetFrontOpWriteMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_front_op_write_mask")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetFrontOpWriteMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_front_op_reference")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetFrontOpReference = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_front_op_reference")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetFrontOpReference = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_back_op_fail")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetBackOpFail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2092799566))
	}
	{
		methodName := StringNameFromStr("get_back_op_fail")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetBackOpFail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1714732389))
	}
	{
		methodName := StringNameFromStr("set_back_op_pass")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetBackOpPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2092799566))
	}
	{
		methodName := StringNameFromStr("get_back_op_pass")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetBackOpPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1714732389))
	}
	{
		methodName := StringNameFromStr("set_back_op_depth_fail")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetBackOpDepthFail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2092799566))
	}
	{
		methodName := StringNameFromStr("get_back_op_depth_fail")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetBackOpDepthFail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1714732389))
	}
	{
		methodName := StringNameFromStr("set_back_op_compare")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetBackOpCompare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2573711505))
	}
	{
		methodName := StringNameFromStr("get_back_op_compare")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetBackOpCompare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 269730778))
	}
	{
		methodName := StringNameFromStr("set_back_op_compare_mask")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetBackOpCompareMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_back_op_compare_mask")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetBackOpCompareMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_back_op_write_mask")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetBackOpWriteMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_back_op_write_mask")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetBackOpWriteMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_back_op_reference")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnSetBackOpReference = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_back_op_reference")
		defer methodName.Destroy()
		ptrsForRDPipelineDepthStencilState.fnGetBackOpReference = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type RDPipelineDepthStencilState struct {
	RefCounted
}

func (me *RDPipelineDepthStencilState) BaseClass() string {
	return "RDPipelineDepthStencilState"
}

func NewRDPipelineDepthStencilState() *RDPipelineDepthStencilState {
	str := StringNameFromStr("RDPipelineDepthStencilState") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RDPipelineDepthStencilState{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RDPipelineDepthStencilState) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RDPipelineDepthStencilState) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RDPipelineDepthStencilState) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RDPipelineDepthStencilState) SetEnableDepthTest(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetEnableDepthTest), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetEnableDepthTest() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetEnableDepthTest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetEnableDepthWrite(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetEnableDepthWrite), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetEnableDepthWrite() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetEnableDepthWrite), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetDepthCompareOperator(p_member RenderingDeviceCompareOperator) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetDepthCompareOperator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetDepthCompareOperator() RenderingDeviceCompareOperator {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceCompareOperator

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetDepthCompareOperator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetEnableDepthRange(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetEnableDepthRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetEnableDepthRange() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetEnableDepthRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetDepthRangeMin(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetDepthRangeMin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetDepthRangeMin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetDepthRangeMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetDepthRangeMax(p_member float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetDepthRangeMax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetDepthRangeMax() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetDepthRangeMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetEnableStencil(p_member bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetEnableStencil), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetEnableStencil() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetEnableStencil), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetFrontOpFail(p_member RenderingDeviceStencilOperation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetFrontOpFail), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetFrontOpFail() RenderingDeviceStencilOperation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceStencilOperation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetFrontOpFail), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetFrontOpPass(p_member RenderingDeviceStencilOperation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetFrontOpPass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetFrontOpPass() RenderingDeviceStencilOperation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceStencilOperation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetFrontOpPass), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetFrontOpDepthFail(p_member RenderingDeviceStencilOperation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetFrontOpDepthFail), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetFrontOpDepthFail() RenderingDeviceStencilOperation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceStencilOperation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetFrontOpDepthFail), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetFrontOpCompare(p_member RenderingDeviceCompareOperator) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetFrontOpCompare), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetFrontOpCompare() RenderingDeviceCompareOperator {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceCompareOperator

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetFrontOpCompare), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetFrontOpCompareMask(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetFrontOpCompareMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetFrontOpCompareMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetFrontOpCompareMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetFrontOpWriteMask(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetFrontOpWriteMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetFrontOpWriteMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetFrontOpWriteMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetFrontOpReference(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetFrontOpReference), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetFrontOpReference() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetFrontOpReference), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetBackOpFail(p_member RenderingDeviceStencilOperation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetBackOpFail), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetBackOpFail() RenderingDeviceStencilOperation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceStencilOperation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetBackOpFail), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetBackOpPass(p_member RenderingDeviceStencilOperation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetBackOpPass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetBackOpPass() RenderingDeviceStencilOperation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceStencilOperation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetBackOpPass), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetBackOpDepthFail(p_member RenderingDeviceStencilOperation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetBackOpDepthFail), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetBackOpDepthFail() RenderingDeviceStencilOperation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceStencilOperation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetBackOpDepthFail), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetBackOpCompare(p_member RenderingDeviceCompareOperator) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetBackOpCompare), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetBackOpCompare() RenderingDeviceCompareOperator {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceCompareOperator

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetBackOpCompare), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RDPipelineDepthStencilState) SetBackOpCompareMask(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetBackOpCompareMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetBackOpCompareMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetBackOpCompareMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetBackOpWriteMask(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetBackOpWriteMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetBackOpWriteMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetBackOpWriteMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RDPipelineDepthStencilState) SetBackOpReference(p_member int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnSetBackOpReference), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RDPipelineDepthStencilState) GetBackOpReference() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineDepthStencilState.fnGetBackOpReference), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
