//  header block begin
// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QMetaClassInfo struct {
	*qtrt.CObject
}

func (this *QMetaClassInfo) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQMetaClassInfoFromPointer(cthis unsafe.Pointer) *QMetaClassInfo {
	return &QMetaClassInfo{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qmetaobject.h:303
// index:0
// Public inline
// void QMetaClassInfo()
func NewQMetaClassInfo() *QMetaClassInfo {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMetaClassInfoC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMetaClassInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:304
// index:0
// Public
// const char * name()
func (this *QMetaClassInfo) Name() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMetaClassInfo4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:305
// index:0
// Public
// const char * value()
func (this *QMetaClassInfo) Value() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMetaClassInfo5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:306
// index:0
// Public inline
// const QMetaObject * enclosingMetaObject()
func (this *QMetaClassInfo) EnclosingMetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMetaClassInfo19enclosingMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
