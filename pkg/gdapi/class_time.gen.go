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

type Time struct {
  Object
}

func (me *Time) BaseClass() string {
  return "Time"
}

func NewTime() *Time {
  str := StringNameFromStr("Time") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Time{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type TimeMonth int
const (
  TimeMonthMonthJanuary TimeMonth = 1
  TimeMonthMonthFebruary TimeMonth = 2
  TimeMonthMonthMarch TimeMonth = 3
  TimeMonthMonthApril TimeMonth = 4
  TimeMonthMonthMay TimeMonth = 5
  TimeMonthMonthJune TimeMonth = 6
  TimeMonthMonthJuly TimeMonth = 7
  TimeMonthMonthAugust TimeMonth = 8
  TimeMonthMonthSeptember TimeMonth = 9
  TimeMonthMonthOctober TimeMonth = 10
  TimeMonthMonthNovember TimeMonth = 11
  TimeMonthMonthDecember TimeMonth = 12
)

type TimeWeekday int
const (
  TimeWeekdayWeekdaySunday TimeWeekday = 0
  TimeWeekdayWeekdayMonday TimeWeekday = 1
  TimeWeekdayWeekdayTuesday TimeWeekday = 2
  TimeWeekdayWeekdayWednesday TimeWeekday = 3
  TimeWeekdayWeekdayThursday TimeWeekday = 4
  TimeWeekdayWeekdayFriday TimeWeekday = 5
  TimeWeekdayWeekdaySaturday TimeWeekday = 6
)

func (me *Time) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Time) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Time) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Time) GetDatetimeDictFromUnixTime(unix_time_val int64, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_dict_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&unix_time_val)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDateDictFromUnixTime(unix_time_val int64, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_date_dict_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&unix_time_val)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetTimeDictFromUnixTime(unix_time_val int64, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_dict_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&unix_time_val)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDatetimeStringFromUnixTime(unix_time_val int64, use_space bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_string_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311239925) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val) , gdc.ConstTypePtr(&use_space) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&unix_time_val)
  pinner.Pin(&use_space)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDateStringFromUnixTime(unix_time_val int64, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_date_string_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&unix_time_val)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetTimeStringFromUnixTime(unix_time_val int64, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_string_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&unix_time_val)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDatetimeDictFromDatetimeString(datetime String, weekday bool, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_dict_from_datetime_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3253569256) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{datetime.AsCTypePtr(), gdc.ConstTypePtr(&weekday) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&weekday)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDatetimeStringFromDatetimeDict(datetime Dictionary, use_space bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_string_from_datetime_dict")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1898123706) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{datetime.AsCTypePtr(), gdc.ConstTypePtr(&use_space) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&use_space)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetUnixTimeFromDatetimeDict(datetime Dictionary, ) int64 {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unix_time_from_datetime_dict")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3021115443) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{datetime.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Time) GetUnixTimeFromDatetimeString(datetime String, ) int64 {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unix_time_from_datetime_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{datetime.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Time) GetOffsetStringFromOffsetMinutes(offset_minutes int64, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset_string_from_offset_minutes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset_minutes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&offset_minutes)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDatetimeDictFromSystem(utc bool, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_dict_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 205769976) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&utc)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDateDictFromSystem(utc bool, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_date_dict_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 205769976) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&utc)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetTimeDictFromSystem(utc bool, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_dict_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 205769976) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&utc)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDatetimeStringFromSystem(utc bool, use_space bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_string_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1136425492) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc) , gdc.ConstTypePtr(&use_space) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&utc)
  pinner.Pin(&use_space)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetDateStringFromSystem(utc bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_date_string_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1162154673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&utc)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetTimeStringFromSystem(utc bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_string_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1162154673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&utc)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetTimeZoneFromSystem() Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_zone_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Time) GetUnixTimeFromSystem() float64 {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unix_time_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Time) GetTicksMsec() int64 {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ticks_msec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Time) GetTicksUsec() int64 {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ticks_usec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
