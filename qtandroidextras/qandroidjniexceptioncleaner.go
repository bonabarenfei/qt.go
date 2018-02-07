package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h
// #include <qandroidjnienvironment.h>
// #include <QtAndroidExtras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QAndroidJniExceptionCleaner struct {
	*qtrt.CObject
}

func (this *QAndroidJniExceptionCleaner) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidJniExceptionCleaner) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidJniExceptionCleanerFromPointer(cthis unsafe.Pointer) *QAndroidJniExceptionCleaner {
	return &QAndroidJniExceptionCleaner{&qtrt.CObject{cthis}}
}
func (*QAndroidJniExceptionCleaner) NewFromPointer(cthis unsafe.Pointer) *QAndroidJniExceptionCleaner {
	return NewQAndroidJniExceptionCleanerFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidJniExceptionCleaner(enum QAndroidJniExceptionCleaner::OutputMode)
func NewQAndroidJniExceptionCleaner(outputMode int) *QAndroidJniExceptionCleaner {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAndroidJniExceptionCleanerC2ENS_10OutputModeE", qtrt.FFI_TYPE_POINTER, outputMode)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidJniExceptionCleanerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidJniExceptionCleaner)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QAndroidJniExceptionCleaner()
func DeleteQAndroidJniExceptionCleaner(this *QAndroidJniExceptionCleaner) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAndroidJniExceptionCleanerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 4)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidjnienvironment.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clean()
func (this *QAndroidJniExceptionCleaner) Clean() {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QAndroidJniExceptionCleaner5cleanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QAndroidJniExceptionCleaner__OutputMode = int

const QAndroidJniExceptionCleaner__Silent QAndroidJniExceptionCleaner__OutputMode = 0
const QAndroidJniExceptionCleaner__Verbose QAndroidJniExceptionCleaner__OutputMode = 1

//  body block end