package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qflags.h
// dst-file: /src/core/qflags.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QIncompatibleFlag::QIncompatibleFlag(int i);
extern void* C_ZN17QIncompatibleFlagC2Ei(int32_t arg0); // 1
  // proto:  void QFlag::QFlag(ushort ai);
extern void* C_ZN5QFlagC2Et(int16_t arg0); // 1
  // proto:  void QFlag::QFlag(int ai);
extern void* C_ZN5QFlagC2Ei(int32_t arg0); // 1
  // proto:  void QFlag::QFlag(short ai);
extern void* C_ZN5QFlagC2Es(int16_t arg0); // 1
  // proto:  void QFlag::QFlag(uint ai);
extern void* C_ZN5QFlagC2Ej(int32_t arg0); // 1
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QIncompatibleFlag)=4
type QIncompatibleFlag struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFlag)=4
type QFlag struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QIncompatibleFlag(int)
func GcfreeQIncompatibleFlag(this *QIncompatibleFlag) {
  qtrt.UniverseFree(this)
}
func NewQIncompatibleFlag(args ...interface{}) *QIncompatibleFlag {
  // QIncompatibleFlag(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QIncompatibleFlagC1Ei
    // invoke: void QIncompatibleFlag(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QIncompatibleFlagC2Ei(arg0)
    this := &QIncompatibleFlag{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQIncompatibleFlag)
    return this
  default:
    qtrt.ErrorResolve("QIncompatibleFlag", "QIncompatibleFlag", args)
  }

  return nil // QIncompatibleFlag{Qclsinst:qthis}
}

// QFlag(ushort)
func GcfreeQFlag(this *QFlag) {
  qtrt.UniverseFree(this)
}
func NewQFlag(args ...interface{}) *QFlag {
  // QFlag(ushort)
  // QFlag(int)
  // QFlag(short)
  // QFlag(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int16Ty(false) // "short"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QFlagC1Et
    // invoke: void QFlag(ushort)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFlagC2Et(arg0)
    this := &QFlag{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFlag)
    return this
  case 1:
    // invoke: _ZN5QFlagC1Ei
    // invoke: void QFlag(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFlagC2Ei(arg0)
    this := &QFlag{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFlag)
    return this
  case 2:
    // invoke: _ZN5QFlagC1Es
    // invoke: void QFlag(short)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFlagC2Es(arg0)
    this := &QFlag{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFlag)
    return this
  case 3:
    // invoke: _ZN5QFlagC1Ej
    // invoke: void QFlag(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QFlagC2Ej(arg0)
    this := &QFlag{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFlag)
    return this
  default:
    qtrt.ErrorResolve("QFlag", "QFlag", args)
  }

  return nil // QFlag{Qclsinst:qthis}
}

// <= body block end
