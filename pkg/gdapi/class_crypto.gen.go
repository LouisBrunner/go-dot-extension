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

type ptrsForCryptoList struct {
  fnGenerateRandomBytes gdc.MethodBindPtr
  fnGenerateRsa gdc.MethodBindPtr
  fnGenerateSelfSignedCertificate gdc.MethodBindPtr
  fnSign gdc.MethodBindPtr
  fnVerify gdc.MethodBindPtr
  fnEncrypt gdc.MethodBindPtr
  fnDecrypt gdc.MethodBindPtr
  fnHmacDigest gdc.MethodBindPtr
  fnConstantTimeCompare gdc.MethodBindPtr
}

var ptrsForCrypto ptrsForCryptoList

func initCryptoPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Crypto")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("generate_random_bytes")
    defer methodName.Destroy()
    ptrsForCrypto.fnGenerateRandomBytes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 47165747))
  }
  {
    methodName := StringNameFromStr("generate_rsa")
    defer methodName.Destroy()
    ptrsForCrypto.fnGenerateRsa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1237515462))
  }
  {
    methodName := StringNameFromStr("generate_self_signed_certificate")
    defer methodName.Destroy()
    ptrsForCrypto.fnGenerateSelfSignedCertificate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 492266173))
  }
  {
    methodName := StringNameFromStr("sign")
    defer methodName.Destroy()
    ptrsForCrypto.fnSign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1673662703))
  }
  {
    methodName := StringNameFromStr("verify")
    defer methodName.Destroy()
    ptrsForCrypto.fnVerify = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2805902225))
  }
  {
    methodName := StringNameFromStr("encrypt")
    defer methodName.Destroy()
    ptrsForCrypto.fnEncrypt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2361793670))
  }
  {
    methodName := StringNameFromStr("decrypt")
    defer methodName.Destroy()
    ptrsForCrypto.fnDecrypt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2361793670))
  }
  {
    methodName := StringNameFromStr("hmac_digest")
    defer methodName.Destroy()
    ptrsForCrypto.fnHmacDigest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2368951203))
  }
  {
    methodName := StringNameFromStr("constant_time_compare")
    defer methodName.Destroy()
    ptrsForCrypto.fnConstantTimeCompare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1024142237))
  }
}

type Crypto struct {
  RefCounted
}

func (me *Crypto) BaseClass() string {
  return "Crypto"
}

func NewCrypto() *Crypto {
  str := StringNameFromStr("Crypto") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Crypto{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Crypto) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Crypto) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Crypto) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Crypto) GenerateRandomBytes(size int64, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&size)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnGenerateRandomBytes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Crypto) GenerateRsa(size int64, ) CryptoKey {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCryptoKey()
  pinner.Pin(&size)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnGenerateRsa), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Crypto) GenerateSelfSignedCertificate(key CryptoKey, issuer_name String, not_before String, not_after String, ) X509Certificate {
  cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), issuer_name.AsCTypePtr(), not_before.AsCTypePtr(), not_after.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewX509Certificate()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnGenerateSelfSignedCertificate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Crypto) Sign(hash_type HashingContextHashType, hash PackedByteArray, key CryptoKey, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hash_type) , hash.AsCTypePtr(), key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&hash_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnSign), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Crypto) Verify(hash_type HashingContextHashType, hash PackedByteArray, signature PackedByteArray, key CryptoKey, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hash_type) , hash.AsCTypePtr(), signature.AsCTypePtr(), key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&hash_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnVerify), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Crypto) Encrypt(key CryptoKey, plaintext PackedByteArray, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), plaintext.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnEncrypt), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Crypto) Decrypt(key CryptoKey, ciphertext PackedByteArray, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), ciphertext.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnDecrypt), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Crypto) HmacDigest(hash_type HashingContextHashType, key PackedByteArray, msg PackedByteArray, ) PackedByteArray {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hash_type) , key.AsCTypePtr(), msg.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&hash_type)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnHmacDigest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Crypto) ConstantTimeCompare(trusted PackedByteArray, received PackedByteArray, ) bool {
  cargs := []gdc.ConstTypePtr{trusted.AsCTypePtr(), received.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCrypto.fnConstantTimeCompare), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
