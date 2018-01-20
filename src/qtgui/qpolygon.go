//  header block begin
// /usr/include/qt/QtGui/qpolygon.h
// #include <qpolygon.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 72
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
type QPolygon struct {
	*qtrt.CObject
}

func (this *QPolygon) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQPolygonFromPointer(cthis unsafe.Pointer) *QPolygon {
	return &QPolygon{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpolygon.h:59
// index:0
// Public inline
// void QPolygon()
func NewQPolygon() *QPolygon {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:61
// index:1
// Public inline
// void QPolygon(int)
func NewQPolygon_1(size int) *QPolygon {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &size)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:66
// index:2
// Public
// void QPolygon(const class QRect &, _Bool)
func NewQPolygon_2(r *qtcore.QRect, closed bool) *QPolygon {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2ERK5QRectb", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &closed)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:67
// index:3
// Public
// void QPolygon(int, const int *)
func NewQPolygon_3(nPoints int, points unsafe.Pointer) *QPolygon {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2EiPKi", ffiqt.FFI_TYPE_VOID, cthis, &nPoints, points)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:60
// index:0
// Public inline
// void ~QPolygon()
func DeleteQPolygon(*QPolygon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:74
// index:0
// Public inline
// void swap(class QPolygon &)
func (this *QPolygon) Swap(other *QPolygon) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:78
// index:0
// Public
// void translate(int, int)
func (this *QPolygon) Translate(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9translateEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:79
// index:1
// Public
// void translate(const class QPoint &)
func (this *QPolygon) Translate_1(offset *qtcore.QPoint) {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9translateERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:81
// index:0
// Public
// QPolygon translated(int, int)
func (this *QPolygon) Translated(dx int, dy int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10translatedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpolygon.h:82
// index:1
// Public inline
// QPolygon translated(const class QPoint &)
func (this *QPolygon) Translated_1(offset *qtcore.QPoint) interface{} {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10translatedERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpolygon.h:84
// index:0
// Public
// QRect boundingRect()
func (this *QPolygon) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpolygon.h:86
// index:0
// Public
// void point(int, int *, int *)
func (this *QPolygon) Point(i int, x unsafe.Pointer, y unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon5pointEiPiS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i, x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:87
// index:1
// Public
// QPoint point(int)
func (this *QPolygon) Point_1(i int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon5pointEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpolygon.h:88
// index:0
// Public
// void setPoint(int, int, int)
func (this *QPolygon) SetPoint(index int, x int, y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon8setPointEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:89
// index:1
// Public
// void setPoint(int, const class QPoint &)
func (this *QPolygon) SetPoint_1(index int, p *qtcore.QPoint) {
	var convArg1 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon8setPointEiRK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:90
// index:0
// Public
// void setPoints(int, const int *)
func (this *QPolygon) SetPoints(nPoints int, points unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9setPointsEiPKi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &nPoints, points)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:92
// index:0
// Public
// void putPoints(int, int, const int *)
func (this *QPolygon) PutPoints(index int, nPoints int, points unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiPKi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &nPoints, points)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:94
// index:1
// Public
// void putPoints(int, int, const class QPolygon &, int)
func (this *QPolygon) PutPoints_1(index int, nPoints int, from *QPolygon, fromIndex int) {
	var convArg2 = from.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiRKS_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &nPoints, convArg2, &fromIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:96
// index:0
// Public
// bool containsPoint(const class QPoint &, Qt::FillRule)
func (this *QPolygon) ContainsPoint(pt *qtcore.QPoint, fillRule int) interface{} {
	var convArg0 = pt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon13containsPointERK6QPointN2Qt8FillRuleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &fillRule)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpolygon.h:98
// index:0
// Public
// QPolygon united(const class QPolygon &)
func (this *QPolygon) United(r *QPolygon) interface{} {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon6unitedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpolygon.h:99
// index:0
// Public
// QPolygon intersected(const class QPolygon &)
func (this *QPolygon) Intersected(r *QPolygon) interface{} {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon11intersectedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpolygon.h:100
// index:0
// Public
// QPolygon subtracted(const class QPolygon &)
func (this *QPolygon) Subtracted(r *QPolygon) interface{} {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10subtractedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpolygon.h:102
// index:0
// Public
// bool intersects(const class QPolygon &)
func (this *QPolygon) Intersects(r *QPolygon) interface{} {
	var convArg0 = r.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10intersectsERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
