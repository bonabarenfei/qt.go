//  header block begin
// /usr/include/qt/QtGui/qsurface.h
// #include <qsurface.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSurface struct {
	*qtrt.CObject
}

func (this *QSurface) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQSurfaceFromPointer(cthis unsafe.Pointer) *QSurface {
	return &QSurface{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qsurface.h:72
// index:0
// Public virtual
// void ~QSurface()
func DeleteQSurface(*QSurface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSurfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurface.h:74
// index:0
// Public
// QSurface::SurfaceClass surfaceClass()
func (this *QSurface) SurfaceClass() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface12surfaceClassEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsurface.h:76
// index:0
// Public pure virtual
// QSurfaceFormat format()
func (this *QSurface) Format() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsurface.h:77
// index:0
// Public pure virtual
// QPlatformSurface * surfaceHandle()
func (this *QSurface) SurfaceHandle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface13surfaceHandleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsurface.h:79
// index:0
// Public pure virtual
// QSurface::SurfaceType surfaceType()
func (this *QSurface) SurfaceType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface11surfaceTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsurface.h:80
// index:0
// Public
// bool supportsOpenGL()
func (this *QSurface) SupportsOpenGL() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface14supportsOpenGLEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsurface.h:82
// index:0
// Public pure virtual
// QSize size()
func (this *QSurface) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qsurface.h:85
// index:0
// Protected
// void QSurface(enum QSurface::SurfaceClass)
func NewQSurface(type_ int) *QSurface {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSurfaceC1ENS_12SurfaceClassE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQSurfaceFromPointer(cthis)
	return gothis
}

//  body block end
