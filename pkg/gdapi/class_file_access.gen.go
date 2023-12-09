// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FileAccess struct {
  obj gdc.ObjectPtr
}

func (me *FileAccess) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FileAccess) BaseClass() string {
  return "FileAccess"
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

func  FileAccessOpen(path String, flags FileAccessModeFlags, )  {
  panic("TODO: implement")
}

func  FileAccessOpenEncrypted(path String, mode_flags FileAccessModeFlags, key PackedByteArray, )  {
  panic("TODO: implement")
}

func  FileAccessOpenEncryptedWithPass(path String, mode_flags FileAccessModeFlags, pass String, )  {
  panic("TODO: implement")
}

func  FileAccessOpenCompressed(path String, mode_flags FileAccessModeFlags, compression_mode FileAccessCompressionMode, )  {
  panic("TODO: implement")
}

func  FileAccessGetOpenError()  {
  panic("TODO: implement")
}

func  FileAccessGetFileAsBytes(path String, )  {
  panic("TODO: implement")
}

func  FileAccessGetFileAsString(path String, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) Flush()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetPath()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetPathAbsolute()  {
  panic("TODO: implement")
}

func  (me *FileAccess) IsOpen()  {
  panic("TODO: implement")
}

func  (me *FileAccess) Seek(position int, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) SeekEnd(position int, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetPosition()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetLength()  {
  panic("TODO: implement")
}

func  (me *FileAccess) EofReached()  {
  panic("TODO: implement")
}

func  (me *FileAccess) Get8()  {
  panic("TODO: implement")
}

func  (me *FileAccess) Get16()  {
  panic("TODO: implement")
}

func  (me *FileAccess) Get32()  {
  panic("TODO: implement")
}

func  (me *FileAccess) Get64()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetFloat()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetDouble()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetReal()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetBuffer(length int, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetLine()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetCsvLine(delim String, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetAsText(skip_cr bool, )  {
  panic("TODO: implement")
}

func  FileAccessGetMd5(path String, )  {
  panic("TODO: implement")
}

func  FileAccessGetSha256(path String, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) IsBigEndian()  {
  panic("TODO: implement")
}

func  (me *FileAccess) SetBigEndian(big_endian bool, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetError()  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetVar(allow_objects bool, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) Store8(value int, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) Store16(value int, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) Store32(value int, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) Store64(value int, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StoreFloat(value float32, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StoreDouble(value float32, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StoreReal(value float32, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StoreBuffer(buffer PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StoreLine(line String, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StoreCsvLine(values PackedStringArray, delim String, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StoreString(string_ String, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StoreVar(value Variant, full_objects bool, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) StorePascalString(string_ String, )  {
  panic("TODO: implement")
}

func  (me *FileAccess) GetPascalString()  {
  panic("TODO: implement")
}

func  (me *FileAccess) Close()  {
  panic("TODO: implement")
}

func  FileAccessFileExists(path String, )  {
  panic("TODO: implement")
}

func  FileAccessGetModifiedTime(file String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
