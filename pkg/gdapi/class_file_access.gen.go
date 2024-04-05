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

type ptrsForFileAccessList struct {
  fnOpen gdc.MethodBindPtr
  fnOpenEncrypted gdc.MethodBindPtr
  fnOpenEncryptedWithPass gdc.MethodBindPtr
  fnOpenCompressed gdc.MethodBindPtr
  fnGetOpenError gdc.MethodBindPtr
  fnGetFileAsBytes gdc.MethodBindPtr
  fnGetFileAsString gdc.MethodBindPtr
  fnFlush gdc.MethodBindPtr
  fnGetPath gdc.MethodBindPtr
  fnGetPathAbsolute gdc.MethodBindPtr
  fnIsOpen gdc.MethodBindPtr
  fnSeek gdc.MethodBindPtr
  fnSeekEnd gdc.MethodBindPtr
  fnGetPosition gdc.MethodBindPtr
  fnGetLength gdc.MethodBindPtr
  fnEofReached gdc.MethodBindPtr
  fnGet8 gdc.MethodBindPtr
  fnGet16 gdc.MethodBindPtr
  fnGet32 gdc.MethodBindPtr
  fnGet64 gdc.MethodBindPtr
  fnGetFloat gdc.MethodBindPtr
  fnGetDouble gdc.MethodBindPtr
  fnGetReal gdc.MethodBindPtr
  fnGetBuffer gdc.MethodBindPtr
  fnGetLine gdc.MethodBindPtr
  fnGetCsvLine gdc.MethodBindPtr
  fnGetAsText gdc.MethodBindPtr
  fnGetMd5 gdc.MethodBindPtr
  fnGetSha256 gdc.MethodBindPtr
  fnIsBigEndian gdc.MethodBindPtr
  fnSetBigEndian gdc.MethodBindPtr
  fnGetError gdc.MethodBindPtr
  fnGetVar gdc.MethodBindPtr
  fnStore8 gdc.MethodBindPtr
  fnStore16 gdc.MethodBindPtr
  fnStore32 gdc.MethodBindPtr
  fnStore64 gdc.MethodBindPtr
  fnStoreFloat gdc.MethodBindPtr
  fnStoreDouble gdc.MethodBindPtr
  fnStoreReal gdc.MethodBindPtr
  fnStoreBuffer gdc.MethodBindPtr
  fnStoreLine gdc.MethodBindPtr
  fnStoreCsvLine gdc.MethodBindPtr
  fnStoreString gdc.MethodBindPtr
  fnStoreVar gdc.MethodBindPtr
  fnStorePascalString gdc.MethodBindPtr
  fnGetPascalString gdc.MethodBindPtr
  fnClose gdc.MethodBindPtr
  fnFileExists gdc.MethodBindPtr
  fnGetModifiedTime gdc.MethodBindPtr
  fnGetUnixPermissions gdc.MethodBindPtr
  fnSetUnixPermissions gdc.MethodBindPtr
  fnGetHiddenAttribute gdc.MethodBindPtr
  fnSetHiddenAttribute gdc.MethodBindPtr
  fnSetReadOnlyAttribute gdc.MethodBindPtr
  fnGetReadOnlyAttribute gdc.MethodBindPtr
}

var ptrsForFileAccess ptrsForFileAccessList

func initFileAccessPtrs(iface gdc.Interface) {

  className := StringNameFromStr("FileAccess")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("open")
    defer methodName.Destroy()
    ptrsForFileAccess.fnOpen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1247358404))
  }
  {
    methodName := StringNameFromStr("open_encrypted")
    defer methodName.Destroy()
    ptrsForFileAccess.fnOpenEncrypted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1482131466))
  }
  {
    methodName := StringNameFromStr("open_encrypted_with_pass")
    defer methodName.Destroy()
    ptrsForFileAccess.fnOpenEncryptedWithPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 790283377))
  }
  {
    methodName := StringNameFromStr("open_compressed")
    defer methodName.Destroy()
    ptrsForFileAccess.fnOpenCompressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3686439335))
  }
  {
    methodName := StringNameFromStr("get_open_error")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetOpenError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
  }
  {
    methodName := StringNameFromStr("get_file_as_bytes")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetFileAsBytes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659035735))
  }
  {
    methodName := StringNameFromStr("get_file_as_string")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetFileAsString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1703090593))
  }
  {
    methodName := StringNameFromStr("flush")
    defer methodName.Destroy()
    ptrsForFileAccess.fnFlush = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_path")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_path_absolute")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetPathAbsolute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("is_open")
    defer methodName.Destroy()
    ptrsForFileAccess.fnIsOpen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("seek")
    defer methodName.Destroy()
    ptrsForFileAccess.fnSeek = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("seek_end")
    defer methodName.Destroy()
    ptrsForFileAccess.fnSeekEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
  }
  {
    methodName := StringNameFromStr("get_position")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_length")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("eof_reached")
    defer methodName.Destroy()
    ptrsForFileAccess.fnEofReached = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_8")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGet8 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_16")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGet16 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_32")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGet32 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_64")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGet64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_float")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetFloat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_double")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetDouble = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_real")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetReal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_buffer")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4131300905))
  }
  {
    methodName := StringNameFromStr("get_line")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_csv_line")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetCsvLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2358116058))
  }
  {
    methodName := StringNameFromStr("get_as_text")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetAsText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1162154673))
  }
  {
    methodName := StringNameFromStr("get_md5")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetMd5 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1703090593))
  }
  {
    methodName := StringNameFromStr("get_sha256")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetSha256 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1703090593))
  }
  {
    methodName := StringNameFromStr("is_big_endian")
    defer methodName.Destroy()
    ptrsForFileAccess.fnIsBigEndian = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_big_endian")
    defer methodName.Destroy()
    ptrsForFileAccess.fnSetBigEndian = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_error")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3185525595))
  }
  {
    methodName := StringNameFromStr("get_var")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetVar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 189129690))
  }
  {
    methodName := StringNameFromStr("store_8")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStore8 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("store_16")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStore16 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("store_32")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStore32 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("store_64")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStore64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("store_float")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStoreFloat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("store_double")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStoreDouble = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("store_real")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStoreReal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("store_buffer")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStoreBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2971499966))
  }
  {
    methodName := StringNameFromStr("store_line")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStoreLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("store_csv_line")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStoreCsvLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2173791505))
  }
  {
    methodName := StringNameFromStr("store_string")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStoreString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("store_var")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStoreVar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 738511890))
  }
  {
    methodName := StringNameFromStr("store_pascal_string")
    defer methodName.Destroy()
    ptrsForFileAccess.fnStorePascalString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_pascal_string")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetPascalString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
  }
  {
    methodName := StringNameFromStr("close")
    defer methodName.Destroy()
    ptrsForFileAccess.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("file_exists")
    defer methodName.Destroy()
    ptrsForFileAccess.fnFileExists = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
  }
  {
    methodName := StringNameFromStr("get_modified_time")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetModifiedTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1597066294))
  }
  {
    methodName := StringNameFromStr("get_unix_permissions")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetUnixPermissions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 524341837))
  }
  {
    methodName := StringNameFromStr("set_unix_permissions")
    defer methodName.Destroy()
    ptrsForFileAccess.fnSetUnixPermissions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 846038644))
  }
  {
    methodName := StringNameFromStr("get_hidden_attribute")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetHiddenAttribute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
  }
  {
    methodName := StringNameFromStr("set_hidden_attribute")
    defer methodName.Destroy()
    ptrsForFileAccess.fnSetHiddenAttribute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2892558115))
  }
  {
    methodName := StringNameFromStr("set_read_only_attribute")
    defer methodName.Destroy()
    ptrsForFileAccess.fnSetReadOnlyAttribute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2892558115))
  }
  {
    methodName := StringNameFromStr("get_read_only_attribute")
    defer methodName.Destroy()
    ptrsForFileAccess.fnGetReadOnlyAttribute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
  }
}

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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFileAccess()
  pinner.Pin(&flags)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnOpen), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessOpenEncrypted(path String, mode_flags FileAccessModeFlags, key PackedByteArray, ) FileAccess {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&mode_flags) , key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFileAccess()
  pinner.Pin(&mode_flags)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnOpenEncrypted), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessOpenEncryptedWithPass(path String, mode_flags FileAccessModeFlags, pass String, ) FileAccess {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&mode_flags) , pass.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFileAccess()
  pinner.Pin(&mode_flags)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnOpenEncryptedWithPass), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessOpenCompressed(path String, mode_flags FileAccessModeFlags, compression_mode FileAccessCompressionMode, ) FileAccess {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&mode_flags) , gdc.ConstTypePtr(&compression_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFileAccess()
  pinner.Pin(&mode_flags)
  pinner.Pin(&compression_mode)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnOpenCompressed), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessGetOpenError() Error {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetOpenError), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessGetFileAsBytes(path String, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetFileAsBytes), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessGetFileAsString(path String, ) String {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetFileAsString), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) Flush()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnFlush), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) GetPath() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) GetPathAbsolute() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetPathAbsolute), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) IsOpen() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnIsOpen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Seek(position int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnSeek), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) SeekEnd(position int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnSeekEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) GetPosition() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetLength() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) EofReached() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnEofReached), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Get8() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGet8), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Get16() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGet16), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Get32() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGet32), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) Get64() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGet64), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetFloat() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetFloat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetDouble() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetDouble), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetReal() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetReal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) GetBuffer(length int64, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&length)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) GetLine() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) GetCsvLine(delim String, ) PackedStringArray {
  cargs := []gdc.ConstTypePtr{delim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetCsvLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) GetAsText(skip_cr bool, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skip_cr) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&skip_cr)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetAsText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessGetMd5(path String, ) String {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetMd5), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  FileAccessGetSha256(path String, ) String {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetSha256), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) IsBigEndian() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnIsBigEndian), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *FileAccess) SetBigEndian(big_endian bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&big_endian) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnSetBigEndian), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) GetError() Error {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetError), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *FileAccess) GetVar(allow_objects bool, ) Variant {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&allow_objects)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetVar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) Store8(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStore8), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) Store16(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStore16), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) Store32(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStore32), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) Store64(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStore64), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreFloat(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStoreFloat), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreDouble(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStoreDouble), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreReal(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStoreReal), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreBuffer(buffer PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStoreBuffer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreLine(line String, )  {
  cargs := []gdc.ConstTypePtr{line.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStoreLine), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreCsvLine(values PackedStringArray, delim String, )  {
  cargs := []gdc.ConstTypePtr{values.AsCTypePtr(), delim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStoreCsvLine), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreString(string_ String, )  {
  cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStoreString), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StoreVar(value Variant, full_objects bool, )  {
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(&full_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStoreVar), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) StorePascalString(string_ String, )  {
  cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnStorePascalString), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *FileAccess) GetPascalString() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetPascalString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *FileAccess) Close()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnClose), me.obj, unsafe.SliceData(cargs), nil)

}

func  FileAccessFileExists(path String, ) bool {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnFileExists), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  FileAccessGetModifiedTime(file String, ) int64 {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetModifiedTime), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  FileAccessGetUnixPermissions(file String, ) FileAccessUnixPermissionFlags {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret FileAccessUnixPermissionFlags

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetUnixPermissions), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessSetUnixPermissions(file String, permissions FileAccessUnixPermissionFlags, ) Error {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), gdc.ConstTypePtr(&permissions) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&permissions)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnSetUnixPermissions), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessGetHiddenAttribute(file String, ) bool {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetHiddenAttribute), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  FileAccessSetHiddenAttribute(file String, hidden bool, ) Error {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), gdc.ConstTypePtr(&hidden) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&hidden)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnSetHiddenAttribute), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessSetReadOnlyAttribute(file String, ro bool, ) Error {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), gdc.ConstTypePtr(&ro) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&ro)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnSetReadOnlyAttribute), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  FileAccessGetReadOnlyAttribute(file String, ) bool {
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFileAccess.fnGetReadOnlyAttribute), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
