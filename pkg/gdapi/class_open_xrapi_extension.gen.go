// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type OpenXRAPIExtension struct {
  RefCounted
}

func (me *OpenXRAPIExtension) BaseClass() string {
  return "OpenXRAPIExtension"
}

func NewOpenXRAPIExtension() *OpenXRAPIExtension {
  str := StringNameFromStr("OpenXRAPIExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OpenXRAPIExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *OpenXRAPIExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRAPIExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRAPIExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRAPIExtension) GetInstance() int64 {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) GetSystemId() int64 {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_system_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) GetSession() int64 {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_session")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) TransformFromPose(pose unsafe.Pointer, ) Transform3D {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("transform_from_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3255299855) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pose) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&pose)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRAPIExtension) XrResult(result int64, format String, args Array, ) bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("xr_result")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3886436197) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&result) , format.AsCTypePtr(), args.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&result)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  OpenXRAPIExtensionOpenxrIsEnabled(check_run_in_editor bool, ) bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("openxr_is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2703660260) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&check_run_in_editor) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&check_run_in_editor)

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) GetInstanceProcAddr(name String, ) int64 {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_proc_addr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1597066294) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) GetErrorString(result int64, ) String {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_error_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 990163283) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&result) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&result)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRAPIExtension) GetSwapchainFormatName(swapchain_format int64, ) String {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swapchain_format_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 990163283) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&swapchain_format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&swapchain_format)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRAPIExtension) IsInitialized() bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_initialized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) IsRunning() bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_running")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) GetPlaySpace() int64 {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_play_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) GetNextFrameTime() int64 {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_frame_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OpenXRAPIExtension) CanRender() bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_render")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
