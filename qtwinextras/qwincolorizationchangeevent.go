package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h
// #include <qwinevent.h>
// #include <Qtwinextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QWinColorizationChangeEvent struct {
	*QWinEvent
}

func (this *QWinColorizationChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWinEvent.GetCthis()
	}
}
func (this *QWinColorizationChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QWinEvent = NewQWinEventFromPointer(cthis)
}
func NewQWinColorizationChangeEventFromPointer(cthis unsafe.Pointer) *QWinColorizationChangeEvent {
	bcthis0 := NewQWinEventFromPointer(cthis)
	return &QWinColorizationChangeEvent{bcthis0}
}
func (*QWinColorizationChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QWinColorizationChangeEvent {
	return NewQWinColorizationChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinColorizationChangeEvent(QRgb, _Bool)
func NewQWinColorizationChangeEvent(color uint, opaque bool) *QWinColorizationChangeEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QWinColorizationChangeEventC2Ejb", qtrt.FFI_TYPE_POINTER, color, opaque)
	gopp.ErrPrint(err, rv)
	gothis := NewQWinColorizationChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWinColorizationChangeEvent)
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinColorizationChangeEvent()
func DeleteQWinColorizationChangeEvent(this *QWinColorizationChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QWinColorizationChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QRgb color()
func (this *QWinColorizationChangeEvent) Color() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QWinColorizationChangeEvent5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinevent.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool opaqueBlend()
func (this *QWinColorizationChangeEvent) OpaqueBlend() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK27QWinColorizationChangeEvent11opaqueBlendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end