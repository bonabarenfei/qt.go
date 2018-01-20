//  header block begin
// /usr/include/qt/QtGui/qdrag.h
// #include <qdrag.h>
// #include <QtGui>
package qtgui

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
type QDrag struct {
	*qtcore.QObject
}

func (this *QDrag) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQDragFromPointer(cthis unsafe.Pointer) *QDrag {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDrag{bcthis0}
}

// /usr/include/qt/QtGui/qdrag.h:59
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDrag) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:62
// index:0
// Public
// void QDrag(class QObject *)
func NewQDrag(dragSource unsafe.Pointer) *QDrag {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDragC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dragSource)
	gopp.ErrPrint(err, rv)
	gothis := NewQDragFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qdrag.h:63
// index:0
// Public virtual
// void ~QDrag()
func DeleteQDrag(*QDrag) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDragD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:65
// index:0
// Public
// void setMimeData(class QMimeData *)
func (this *QDrag) SetMimeData(data unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag11setMimeDataEP9QMimeData", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:66
// index:0
// Public
// QMimeData * mimeData()
func (this *QDrag) MimeData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag8mimeDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:68
// index:0
// Public
// void setPixmap(const class QPixmap &)
func (this *QDrag) SetPixmap(arg0 *QPixmap) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag9setPixmapERK7QPixmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:69
// index:0
// Public
// QPixmap pixmap()
func (this *QDrag) Pixmap() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag6pixmapEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:71
// index:0
// Public
// void setHotSpot(const class QPoint &)
func (this *QDrag) SetHotSpot(hotspot *qtcore.QPoint) {
	var convArg0 = hotspot.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag10setHotSpotERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:72
// index:0
// Public
// QPoint hotSpot()
func (this *QDrag) HotSpot() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag7hotSpotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:74
// index:0
// Public
// QObject * source()
func (this *QDrag) Source() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:75
// index:0
// Public
// QObject * target()
func (this *QDrag) Target() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag6targetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:81
// index:0
// Public
// void setDragCursor(const class QPixmap &, Qt::DropAction)
func (this *QDrag) SetDragCursor(cursor *QPixmap, action int) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag13setDragCursorERK7QPixmapN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:82
// index:0
// Public
// QPixmap dragCursor(Qt::DropAction)
func (this *QDrag) DragCursor(action int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag10dragCursorEN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:84
// index:0
// Public
// Qt::DropActions supportedActions()
func (this *QDrag) SupportedActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag16supportedActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:85
// index:0
// Public
// Qt::DropAction defaultAction()
func (this *QDrag) DefaultAction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag13defaultActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qdrag.h:87
// index:0
// Public static
// void cancel()
func (this *QDrag) Cancel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag6cancelEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QDrag_Cancel() {
	var nilthis *QDrag
	nilthis.Cancel()
}

// /usr/include/qt/QtGui/qdrag.h:90
// index:0
// Public
// void actionChanged(Qt::DropAction)
func (this *QDrag) ActionChanged(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag13actionChangedEN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:91
// index:0
// Public
// void targetChanged(class QObject *)
func (this *QDrag) TargetChanged(newTarget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag13targetChangedEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), newTarget)
	gopp.ErrPrint(err, rv)
}

//  body block end
