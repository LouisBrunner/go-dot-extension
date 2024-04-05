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

type ptrsForIPList struct {
  fnResolveHostname gdc.MethodBindPtr
  fnResolveHostnameAddresses gdc.MethodBindPtr
  fnResolveHostnameQueueItem gdc.MethodBindPtr
  fnGetResolveItemStatus gdc.MethodBindPtr
  fnGetResolveItemAddress gdc.MethodBindPtr
  fnGetResolveItemAddresses gdc.MethodBindPtr
  fnEraseResolveItem gdc.MethodBindPtr
  fnGetLocalAddresses gdc.MethodBindPtr
  fnGetLocalInterfaces gdc.MethodBindPtr
  fnClearCache gdc.MethodBindPtr
}

var ptrsForIP ptrsForIPList

func initIPPtrs(iface gdc.Interface) {

  className := StringNameFromStr("IP")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("resolve_hostname")
    defer methodName.Destroy()
    ptrsForIP.fnResolveHostname = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4283295457))
  }
  {
    methodName := StringNameFromStr("resolve_hostname_addresses")
    defer methodName.Destroy()
    ptrsForIP.fnResolveHostnameAddresses = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 773767525))
  }
  {
    methodName := StringNameFromStr("resolve_hostname_queue_item")
    defer methodName.Destroy()
    ptrsForIP.fnResolveHostnameQueueItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1749894742))
  }
  {
    methodName := StringNameFromStr("get_resolve_item_status")
    defer methodName.Destroy()
    ptrsForIP.fnGetResolveItemStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3812250196))
  }
  {
    methodName := StringNameFromStr("get_resolve_item_address")
    defer methodName.Destroy()
    ptrsForIP.fnGetResolveItemAddress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
  }
  {
    methodName := StringNameFromStr("get_resolve_item_addresses")
    defer methodName.Destroy()
    ptrsForIP.fnGetResolveItemAddresses = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
  }
  {
    methodName := StringNameFromStr("erase_resolve_item")
    defer methodName.Destroy()
    ptrsForIP.fnEraseResolveItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_local_addresses")
    defer methodName.Destroy()
    ptrsForIP.fnGetLocalAddresses = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("get_local_interfaces")
    defer methodName.Destroy()
    ptrsForIP.fnGetLocalInterfaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("clear_cache")
    defer methodName.Destroy()
    ptrsForIP.fnClearCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3005725572))
  }
}

type IP struct {
  Object
}

func (me *IP) BaseClass() string {
  return "IP"
}

func NewIP() *IP {
  str := StringNameFromStr("IP") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &IP{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  IPResolverMaxQueries = 256
  IPResolverInvalidId = -1
)

// Enums

type IPResolverStatus int
const (
  IPResolverStatusResolverStatusNone IPResolverStatus = 0
  IPResolverStatusResolverStatusWaiting IPResolverStatus = 1
  IPResolverStatusResolverStatusDone IPResolverStatus = 2
  IPResolverStatusResolverStatusError IPResolverStatus = 3
)

type IPType int
const (
  IPTypeTypeNone IPType = 0
  IPTypeTypeIpv4 IPType = 1
  IPTypeTypeIpv6 IPType = 2
  IPTypeTypeAny IPType = 3
)

func (me *IP) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *IP) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *IP) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *IP) ResolveHostname(host String, ip_type IPType, ) String {
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&ip_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&ip_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnResolveHostname), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) ResolveHostnameAddresses(host String, ip_type IPType, ) PackedStringArray {
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&ip_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()
  pinner.Pin(&ip_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnResolveHostnameAddresses), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) ResolveHostnameQueueItem(host String, ip_type IPType, ) int64 {
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&ip_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&ip_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnResolveHostnameQueueItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *IP) GetResolveItemStatus(id int64, ) IPResolverStatus {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret IPResolverStatus
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnGetResolveItemStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *IP) GetResolveItemAddress(id int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnGetResolveItemAddress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) GetResolveItemAddresses(id int64, ) Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnGetResolveItemAddresses), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) EraseResolveItem(id int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnEraseResolveItem), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *IP) GetLocalAddresses() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnGetLocalAddresses), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *IP) GetLocalInterfaces() []Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnGetLocalInterfaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *IP) ClearCache(hostname String, )  {
  cargs := []gdc.ConstTypePtr{hostname.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForIP.fnClearCache), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
