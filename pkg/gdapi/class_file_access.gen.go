// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type FileAccess struct {
  RefCounted
}

func (me *FileAccess) BaseClass() string {
  return "FileAccess"
}

func NewFileAccess() *FileAccess {
  str := StringNameFromStr("FileAccess") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &FileAccess{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type FileAccessModeFlags int
const (
  FileAccessModeFlagsRead FileAccessModeFlags = 1
  FileAccessModeFlagsWrite FileAccessModeFlags = 2
  FileAccessModeFlagsReadWrite FileAccessModeFlags = 3
  FileAccessModeFlagsWriteRead FileAccessModeFlags = 7
)

type FileAccessCompressionMode int
const (
  FileAccessCompressionModeCompressionFastlz FileAccessCompressionMode = 0
  FileAccessCompressionModeCompressionDeflate FileAccessCompressionMode = 1
  FileAccessCompressionModeCompressionZstd FileAccessCompressionMode = 2
  FileAccessCompressionModeCompressionGzip FileAccessCompressionMode = 3
  FileAccessCompressionModeCompressionBrotli FileAccessCompressionMode = 4
)

type FileAccessUnixPermissionFlags int
const (
  FileAccessUnixPermissionFlagsUnixReadOwner FileAccessUnixPermissionFlags = 256
  FileAccessUnixPermissionFlagsUnixWriteOwner FileAccessUnixPermissionFlags = 128
  FileAccessUnixPermissionFlagsUnixExecuteOwner FileAccessUnixPermissionFlags = 64
  FileAccessUnixPermissionFlagsUnixReadGroup FileAccessUnixPermissionFlags = 32
  FileAccessUnixPermissionFlagsUnixWriteGroup FileAccessUnixPermissionFlags = 16
  FileAccessUnixPermissionFlagsUnixExecuteGroup FileAccessUnixPermissionFlags = 8
  FileAccessUnixPermissionFlagsUnixReadOther FileAccessUnixPermissionFlags = 4
  FileAccessUnixPermissionFlagsUnixWriteOther FileAccessUnixPermissionFlags = 2
  FileAccessUnixPermissionFlagsUnixExecuteOther FileAccessUnixPermissionFlags = 1
  FileAccessUnixPermissionFlagsUnixSetUserId FileAccessUnixPermissionFlags = 2048
  FileAccessUnixPermissionFlagsUnixSetGroupId FileAccessUnixPermissionFlags = 1024
  FileAccessUnixPermissionFlagsUnixRestrictedDelete FileAccessUnixPermissionFlags = 512
)

func (me *FileAccess) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FileAccess) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FileAccess) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  FileAccessOpen(path String, flags FileAccessModeFlags, ) FileAccess {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1247358404) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  ret := NewFileAccess()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessOpenEncrypted(path String, mode_flags FileAccessModeFlags, key PackedByteArray, ) FileAccess {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_encrypted")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1482131466) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&mode_flags), gdc.ConstTypePtr(key.AsCTypePtr()), }
  ret := NewFileAccess()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessOpenEncryptedWithPass(path String, mode_flags FileAccessModeFlags, pass String, ) FileAccess {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_encrypted_with_pass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 790283377) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&mode_flags), gdc.ConstTypePtr(pass.AsCTypePtr()), }
  ret := NewFileAccess()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessOpenCompressed(path String, mode_flags FileAccessModeFlags, compression_mode FileAccessCompressionMode, ) FileAccess {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_compressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3686439335) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&mode_flags), gdc.ConstTypePtr(&compression_mode), }
  ret := NewFileAccess()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessGetOpenError() Error {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_open_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessGetFileAsBytes(path String, ) PackedByteArray {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_as_bytes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659035735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessGetFileAsString(path String, ) String {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file_as_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) Flush()  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("flush")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) GetPath() String {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) GetPathAbsolute() String {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_absolute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) IsOpen() bool {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_open")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Seek(position int64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("seek")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) SeekEnd(position int64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("seek_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) GetPosition() int64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetLength() int64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) EofReached() bool {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("eof_reached")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Get8() int64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_8")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Get16() int64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_16")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Get32() int64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_32")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Get64() int64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetFloat() float64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_float")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetDouble() float64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_double")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetReal() float64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_real")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetBuffer(length int64, ) PackedByteArray {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4131300905) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) GetLine() String {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) GetCsvLine(delim String, ) PackedStringArray {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_csv_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2358116058) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(delim.AsCTypePtr()), }
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) GetAsText(skip_cr bool, ) String {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_as_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1162154673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skip_cr), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessGetMd5(path String, ) String {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_md5")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessGetSha256(path String, ) String {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sha256")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) IsBigEndian() bool {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_big_endian")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) SetBigEndian(big_endian bool, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_big_endian")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&big_endian), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) GetError() Error {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3185525595) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FileAccess) GetVar(allow_objects bool, ) Variant {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_var")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 189129690) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_objects), }
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) Store8(value int64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_8")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) Store16(value int64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_16")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) Store32(value int64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_32")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) Store64(value int64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreFloat(value float64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_float")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreDouble(value float64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_double")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreReal(value float64, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_real")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreBuffer(buffer PackedByteArray, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2971499966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreLine(line String, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(line.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreCsvLine(values PackedStringArray, delim String, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_csv_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2173791505) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(values.AsCTypePtr()), gdc.ConstTypePtr(delim.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreString(string_ String, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreVar(value Variant, full_objects bool, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_var")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 738511890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(value.AsCTypePtr()), gdc.ConstTypePtr(&full_objects), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StorePascalString(string_ String, )  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("store_pascal_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) GetPascalString() String {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pascal_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) Close()  {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  FileAccessFileExists(path String, ) bool {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("file_exists")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  FileAccessGetModifiedTime(file String, ) int64 {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modified_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1597066294) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  FileAccessGetUnixPermissions(file String, ) FileAccessUnixPermissionFlags {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unix_permissions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 524341837) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  var ret FileAccessUnixPermissionFlags

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessSetUnixPermissions(file String, permissions FileAccessUnixPermissionFlags, ) Error {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unix_permissions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 846038644) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), gdc.ConstTypePtr(&permissions), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessGetHiddenAttribute(file String, ) bool {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hidden_attribute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  FileAccessSetHiddenAttribute(file String, hidden bool, ) Error {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hidden_attribute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2892558115) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), gdc.ConstTypePtr(&hidden), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessSetReadOnlyAttribute(file String, ro bool, ) Error {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_read_only_attribute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2892558115) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), gdc.ConstTypePtr(&ro), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessGetReadOnlyAttribute(file String, ) bool {
  classNameV := StringNameFromStr("FileAccess")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_read_only_attribute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
