// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Time struct {
  obj gdc.ObjectPtr
}

func (me *Time) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Time) BaseClass() string {
  return "Time"
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

func  (me *Time) GetDatetimeDictFromUnixTime(unix_time_val int, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_dict_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDateDictFromUnixTime(unix_time_val int, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_date_dict_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetTimeDictFromUnixTime(unix_time_val int, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_dict_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDatetimeStringFromUnixTime(unix_time_val int, use_space bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_string_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311239925) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val), gdc.ConstTypePtr(&use_space), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDateStringFromUnixTime(unix_time_val int, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_date_string_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetTimeStringFromUnixTime(unix_time_val int, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_string_from_unix_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unix_time_val), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDatetimeDictFromDatetimeString(datetime String, weekday bool, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_dict_from_datetime_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3253569256) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(datetime.AsCTypePtr()), gdc.ConstTypePtr(&weekday), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDatetimeStringFromDatetimeDict(datetime Dictionary, use_space bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_string_from_datetime_dict")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1898123706) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(datetime.AsCTypePtr()), gdc.ConstTypePtr(&use_space), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetUnixTimeFromDatetimeDict(datetime Dictionary, ) int {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unix_time_from_datetime_dict")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3021115443) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(datetime.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetUnixTimeFromDatetimeString(datetime String, ) int {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unix_time_from_datetime_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(datetime.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetOffsetStringFromOffsetMinutes(offset_minutes int, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset_string_from_offset_minutes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset_minutes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDatetimeDictFromSystem(utc bool, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_dict_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 205769976) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDateDictFromSystem(utc bool, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_date_dict_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 205769976) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetTimeDictFromSystem(utc bool, ) Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_dict_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 205769976) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDatetimeStringFromSystem(utc bool, use_space bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_datetime_string_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1136425492) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc), gdc.ConstTypePtr(&use_space), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetDateStringFromSystem(utc bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_date_string_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1162154673) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetTimeStringFromSystem(utc bool, ) String {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_string_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1162154673) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&utc), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetTimeZoneFromSystem() Dictionary {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_zone_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetUnixTimeFromSystem() float32 {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unix_time_from_system")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetTicksMsec() int {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ticks_msec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Time) GetTicksUsec() int {
  classNameV := StringNameFromStr("Time")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ticks_usec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
