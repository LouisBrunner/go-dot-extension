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

type UPNPDevice struct {
  RefCounted
}

func (me *UPNPDevice) BaseClass() string {
  return "UPNPDevice"
}

func NewUPNPDevice() *UPNPDevice {
  str := StringNameFromStr("UPNPDevice") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &UPNPDevice{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type UPNPDeviceIGDStatus int
const (
  UPNPDeviceIGDStatusIgdStatusOk UPNPDeviceIGDStatus = 0
  UPNPDeviceIGDStatusIgdStatusHttpError UPNPDeviceIGDStatus = 1
  UPNPDeviceIGDStatusIgdStatusHttpEmpty UPNPDeviceIGDStatus = 2
  UPNPDeviceIGDStatusIgdStatusNoUrls UPNPDeviceIGDStatus = 3
  UPNPDeviceIGDStatusIgdStatusNoIgd UPNPDeviceIGDStatus = 4
  UPNPDeviceIGDStatusIgdStatusDisconnected UPNPDeviceIGDStatus = 5
  UPNPDeviceIGDStatusIgdStatusUnknownDevice UPNPDeviceIGDStatus = 6
  UPNPDeviceIGDStatusIgdStatusInvalidControl UPNPDeviceIGDStatus = 7
  UPNPDeviceIGDStatusIgdStatusMallocError UPNPDeviceIGDStatus = 8
  UPNPDeviceIGDStatusIgdStatusUnknownError UPNPDeviceIGDStatus = 9
)

func (me *UPNPDevice) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *UPNPDevice) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *UPNPDevice) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *UPNPDevice) IsValidGateway() bool {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_valid_gateway")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UPNPDevice) QueryExternalAddress() String {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("query_external_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UPNPDevice) AddPortMapping(port int64, port_internal int64, desc String, proto String, duration int64, ) int64 {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_port_mapping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 818314583) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , gdc.ConstTypePtr(&port_internal) , desc.AsCTypePtr(), proto.AsCTypePtr(), gdc.ConstTypePtr(&duration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&port)
  pinner.Pin(&port_internal)
  pinner.Pin(&duration)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UPNPDevice) DeletePortMapping(port int64, proto String, ) int64 {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delete_port_mapping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444187325) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , proto.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UPNPDevice) SetDescriptionUrl(url String, )  {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_description_url")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UPNPDevice) GetDescriptionUrl() String {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_description_url")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UPNPDevice) SetServiceType(type_ String, )  {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_service_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UPNPDevice) GetServiceType() String {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_service_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UPNPDevice) SetIgdControlUrl(url String, )  {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_igd_control_url")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UPNPDevice) GetIgdControlUrl() String {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_igd_control_url")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UPNPDevice) SetIgdServiceType(type_ String, )  {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_igd_service_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UPNPDevice) GetIgdServiceType() String {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_igd_service_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UPNPDevice) SetIgdOurAddr(addr String, )  {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_igd_our_addr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{addr.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UPNPDevice) GetIgdOurAddr() String {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_igd_our_addr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UPNPDevice) SetIgdStatus(status UPNPDeviceIGDStatus, )  {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_igd_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 519504122) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&status) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UPNPDevice) GetIgdStatus() UPNPDeviceIGDStatus {
  classNameV := StringNameFromStr("UPNPDevice")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_igd_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 180887011) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret UPNPDeviceIGDStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
