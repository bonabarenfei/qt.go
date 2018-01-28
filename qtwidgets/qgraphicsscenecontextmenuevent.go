package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
type QGraphicsSceneContextMenuEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneContextMenuEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneContextMenuEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneContextMenuEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneContextMenuEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneContextMenuEvent{bcthis0}
}
func (*QGraphicsSceneContextMenuEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneContextMenuEvent {
	return NewQGraphicsSceneContextMenuEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:174
// index:0
// Public
// void QGraphicsSceneContextMenuEvent(enum QEvent::Type)
func NewQGraphicsSceneContextMenuEvent(type_ int) *QGraphicsSceneContextMenuEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneContextMenuEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:175
// index:0
// Public virtual
// void ~QGraphicsSceneContextMenuEvent()
func DeleteQGraphicsSceneContextMenuEvent(*QGraphicsSceneContextMenuEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:177
// index:0
// Public
// QPointF pos()
func (this *QGraphicsSceneContextMenuEvent) Pos() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent3posEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:178
// index:0
// Public
// void setPos(const QPointF &)
func (this *QGraphicsSceneContextMenuEvent) SetPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:180
// index:0
// Public
// QPointF scenePos()
func (this *QGraphicsSceneContextMenuEvent) ScenePos() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent8scenePosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:181
// index:0
// Public
// void setScenePos(const QPointF &)
func (this *QGraphicsSceneContextMenuEvent) SetScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:183
// index:0
// Public
// QPoint screenPos()
func (this *QGraphicsSceneContextMenuEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:184
// index:0
// Public
// void setScreenPos(const QPoint &)
func (this *QGraphicsSceneContextMenuEvent) SetScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:186
// index:0
// Public
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneContextMenuEvent) Modifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:189
// index:0
// Public
// QGraphicsSceneContextMenuEvent::Reason reason()
func (this *QGraphicsSceneContextMenuEvent) Reason() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent6reasonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:190
// index:0
// Public
// void setReason(enum QGraphicsSceneContextMenuEvent::Reason)
func (this *QGraphicsSceneContextMenuEvent) SetReason(reason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent9setReasonENS_6ReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	gopp.ErrPrint(err, rv)
}

type QGraphicsSceneContextMenuEvent__Reason = int

const QGraphicsSceneContextMenuEvent__Mouse QGraphicsSceneContextMenuEvent__Reason = 0
const QGraphicsSceneContextMenuEvent__Keyboard QGraphicsSceneContextMenuEvent__Reason = 1
const QGraphicsSceneContextMenuEvent__Other QGraphicsSceneContextMenuEvent__Reason = 2

//  body block end
