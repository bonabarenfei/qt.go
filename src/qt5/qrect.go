package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qrect.h
// dst-file: /src/core/qrect.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QRect::right();
extern void _ZNK5QRect5rightEv(void* qthis);
  // proto:  void QRect::moveTo(const QPoint & p);
extern void demth_ZN5QRect6moveToERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRect::moveTopLeft(const QPoint & p);
extern void demth_ZN5QRect11moveTopLeftERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRect::moveRight(int pos);
extern void demth_ZN5QRect9moveRightEi(void* qthis, int arg0);
  // proto:  void QRect::QRect(const QPoint & topleft, const QPoint & bottomright);
extern void* dector_ZN5QRectC1ERK6QPointS2_(void* arg0, void* arg1);
extern void _ZN5QRectC1ERK6QPointS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QRect QRect::translated(int dx, int dy);
extern void _ZNK5QRect10translatedEii(void* qthis, int arg0, int arg1);
  // proto:  QPoint QRect::center();
extern void _ZNK5QRect6centerEv(void* qthis);
  // proto:  void QRect::moveTopRight(const QPoint & p);
extern void demth_ZN5QRect12moveTopRightERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRect::setLeft(int pos);
extern void demth_ZN5QRect7setLeftEi(void* qthis, int arg0);
  // proto:  int QRect::left();
extern void _ZNK5QRect4leftEv(void* qthis);
  // proto:  QRect QRect::intersected(const QRect & other);
extern void demth_ZNK5QRect11intersectedERKS_(void* qthis, void* arg0);
  // proto:  bool QRect::contains(int x, int y, bool proper);
extern void demth_ZNK5QRect8containsEiib(void* qthis, int arg0, int arg1, bool arg2);
  // proto:  QPoint QRect::bottomRight();
extern void _ZNK5QRect11bottomRightEv(void* qthis);
  // proto:  bool QRect::isValid();
extern void _ZNK5QRect7isValidEv(void* qthis);
  // proto:  QSize QRect::size();
extern void _ZNK5QRect4sizeEv(void* qthis);
  // proto:  QRect QRect::united(const QRect & other);
extern void demth_ZNK5QRect6unitedERKS_(void* qthis, void* arg0);
  // proto:  void QRect::adjust(int x1, int y1, int x2, int y2);
extern void demth_ZN5QRect6adjustEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  bool QRect::isNull();
extern void _ZNK5QRect6isNullEv(void* qthis);
  // proto:  void QRect::setBottom(int pos);
extern void demth_ZN5QRect9setBottomEi(void* qthis, int arg0);
  // proto:  void QRect::setSize(const QSize & s);
extern void demth_ZN5QRect7setSizeERK5QSize(void* qthis, void* arg0);
  // proto:  int QRect::y();
extern void _ZNK5QRect1yEv(void* qthis);
  // proto:  int QRect::x();
extern void _ZNK5QRect1xEv(void* qthis);
  // proto:  QRect QRect::adjusted(int x1, int y1, int x2, int y2);
extern void _ZNK5QRect8adjustedEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QRect::QRect(const QPoint & topleft, const QSize & size);
extern void* dector_ZN5QRectC1ERK6QPointRK5QSize(void* arg0, void* arg1);
extern void _ZN5QRectC1ERK6QPointRK5QSize(void* qthis, void* arg0, void* arg1);
  // proto:  int QRect::height();
extern void _ZNK5QRect6heightEv(void* qthis);
  // proto:  void QRect::QRect(int left, int top, int width, int height);
extern void* dector_ZN5QRectC1Eiiii(int arg0, int arg1, int arg2, int arg3);
extern void _ZN5QRectC1Eiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QRect::moveBottomLeft(const QPoint & p);
extern void demth_ZN5QRect14moveBottomLeftERK6QPoint(void* qthis, void* arg0);
  // proto:  int QRect::top();
extern void _ZNK5QRect3topEv(void* qthis);
  // proto:  void QRect::moveTo(int x, int t);
extern void demth_ZN5QRect6moveToEii(void* qthis, int arg0, int arg1);
  // proto:  void QRect::getRect(int * x, int * y, int * w, int * h);
extern void demth_ZNK5QRect7getRectEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto:  bool QRect::contains(const QRect & r, bool proper);
extern void _ZNK5QRect8containsERKS_b(void* qthis, void* arg0, bool arg1);
  // proto:  QRect QRect::marginsRemoved(const QMargins & margins);
extern void _ZNK5QRect14marginsRemovedERK8QMargins(void* qthis, void* arg0);
  // proto:  void QRect::translate(int dx, int dy);
extern void demth_ZN5QRect9translateEii(void* qthis, int arg0, int arg1);
  // proto:  QPoint QRect::topLeft();
extern void _ZNK5QRect7topLeftEv(void* qthis);
  // proto:  bool QRect::contains(const QPoint & p, bool proper);
extern void _ZNK5QRect8containsERK6QPointb(void* qthis, void* arg0, bool arg1);
  // proto:  int QRect::width();
extern void _ZNK5QRect5widthEv(void* qthis);
  // proto:  void QRect::setRect(int x, int y, int w, int h);
extern void demth_ZN5QRect7setRectEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QRect::moveCenter(const QPoint & p);
extern void demth_ZN5QRect10moveCenterERK6QPoint(void* qthis, void* arg0);
  // proto:  bool QRect::intersects(const QRect & r);
extern void _ZNK5QRect10intersectsERKS_(void* qthis, void* arg0);
  // proto:  void QRect::setTopRight(const QPoint & p);
extern void demth_ZN5QRect11setTopRightERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRect::setCoords(int x1, int y1, int x2, int y2);
extern void demth_ZN5QRect9setCoordsEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QRect::translate(const QPoint & p);
extern void demth_ZN5QRect9translateERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRect::moveBottom(int pos);
extern void demth_ZN5QRect10moveBottomEi(void* qthis, int arg0);
  // proto:  void QRect::setBottomLeft(const QPoint & p);
extern void demth_ZN5QRect13setBottomLeftERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRect::getCoords(int * x1, int * y1, int * x2, int * y2);
extern void demth_ZNK5QRect9getCoordsEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto:  QPoint QRect::topRight();
extern void _ZNK5QRect8topRightEv(void* qthis);
  // proto:  void QRect::setBottomRight(const QPoint & p);
extern void demth_ZN5QRect14setBottomRightERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRect::setHeight(int h);
extern void demth_ZN5QRect9setHeightEi(void* qthis, int arg0);
  // proto:  bool QRect::isEmpty();
extern void _ZNK5QRect7isEmptyEv(void* qthis);
  // proto:  bool QRect::contains(int x, int y);
extern void demth_ZNK5QRect8containsEii(void* qthis, int arg0, int arg1);
  // proto:  void QRect::moveBottomRight(const QPoint & p);
extern void demth_ZN5QRect15moveBottomRightERK6QPoint(void* qthis, void* arg0);
  // proto:  QPoint QRect::bottomLeft();
extern void _ZNK5QRect10bottomLeftEv(void* qthis);
  // proto:  void QRect::setTop(int pos);
extern void demth_ZN5QRect6setTopEi(void* qthis, int arg0);
  // proto:  int QRect::bottom();
extern void _ZNK5QRect6bottomEv(void* qthis);
  // proto:  void QRect::QRect();
extern void* dector_ZN5QRectC1Ev();
extern void _ZN5QRectC1Ev(void* qthis);
  // proto:  QRect QRect::marginsAdded(const QMargins & margins);
extern void _ZNK5QRect12marginsAddedERK8QMargins(void* qthis, void* arg0);
  // proto:  QRect QRect::normalized();
extern void _ZNK5QRect10normalizedEv(void* qthis);
  // proto:  void QRect::setWidth(int w);
extern void demth_ZN5QRect8setWidthEi(void* qthis, int arg0);
  // proto:  void QRect::setY(int y);
extern void demth_ZN5QRect4setYEi(void* qthis, int arg0);
  // proto:  void QRect::moveTop(int pos);
extern void demth_ZN5QRect7moveTopEi(void* qthis, int arg0);
  // proto:  void QRect::setX(int x);
extern void demth_ZN5QRect4setXEi(void* qthis, int arg0);
  // proto:  void QRect::setRight(int pos);
extern void demth_ZN5QRect8setRightEi(void* qthis, int arg0);
  // proto:  void QRect::setTopLeft(const QPoint & p);
extern void demth_ZN5QRect10setTopLeftERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRect::moveLeft(int pos);
extern void demth_ZN5QRect8moveLeftEi(void* qthis, int arg0);
  // proto:  QRect QRect::translated(const QPoint & p);
extern void _ZNK5QRect10translatedERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRectF::QRectF();
extern void* dector_ZN6QRectFC1Ev();
extern void _ZN6QRectFC1Ev(void* qthis);
  // proto:  void QRectF::moveBottomRight(const QPointF & p);
extern void demth_ZN6QRectF15moveBottomRightERK7QPointF(void* qthis, void* arg0);
  // proto:  void QRectF::moveTo(qreal x, qreal y);
extern void demth_ZN6QRectF6moveToEdd(void* qthis, double arg0, double arg1);
  // proto:  qreal QRectF::top();
extern void _ZNK6QRectF3topEv(void* qthis);
  // proto:  QPointF QRectF::bottomLeft();
extern void _ZNK6QRectF10bottomLeftEv(void* qthis);
  // proto:  void QRectF::setHeight(qreal h);
extern void demth_ZN6QRectF9setHeightEd(void* qthis, double arg0);
  // proto:  void QRectF::setSize(const QSizeF & s);
extern void demth_ZN6QRectF7setSizeERK6QSizeF(void* qthis, void* arg0);
  // proto:  void QRectF::QRectF(const QPointF & topleft, const QPointF & bottomRight);
extern void* dector_ZN6QRectFC1ERK7QPointFS2_(void* arg0, void* arg1);
extern void _ZN6QRectFC1ERK7QPointFS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QRectF::moveTo(const QPointF & p);
extern void demth_ZN6QRectF6moveToERK7QPointF(void* qthis, void* arg0);
  // proto:  QRect QRectF::toAlignedRect();
extern void _ZNK6QRectF13toAlignedRectEv(void* qthis);
  // proto:  void QRectF::setRight(qreal pos);
extern void demth_ZN6QRectF8setRightEd(void* qthis, double arg0);
  // proto:  void QRectF::setBottomLeft(const QPointF & p);
extern void demth_ZN6QRectF13setBottomLeftERK7QPointF(void* qthis, void* arg0);
  // proto:  QPointF QRectF::topRight();
extern void _ZNK6QRectF8topRightEv(void* qthis);
  // proto:  QSizeF QRectF::size();
extern void _ZNK6QRectF4sizeEv(void* qthis);
  // proto:  void QRectF::adjust(qreal x1, qreal y1, qreal x2, qreal y2);
extern void demth_ZN6QRectF6adjustEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QRectF::moveRight(qreal pos);
extern void demth_ZN6QRectF9moveRightEd(void* qthis, double arg0);
  // proto:  qreal QRectF::y();
extern void _ZNK6QRectF1yEv(void* qthis);
  // proto:  QPointF QRectF::bottomRight();
extern void _ZNK6QRectF11bottomRightEv(void* qthis);
  // proto:  void QRectF::setBottom(qreal pos);
extern void demth_ZN6QRectF9setBottomEd(void* qthis, double arg0);
  // proto:  void QRectF::moveBottomLeft(const QPointF & p);
extern void demth_ZN6QRectF14moveBottomLeftERK7QPointF(void* qthis, void* arg0);
  // proto:  void QRectF::moveBottom(qreal pos);
extern void demth_ZN6QRectF10moveBottomEd(void* qthis, double arg0);
  // proto:  void QRectF::getRect(qreal * x, qreal * y, qreal * w, qreal * h);
extern void demth_ZNK6QRectF7getRectEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto:  qreal QRectF::x();
extern void _ZNK6QRectF1xEv(void* qthis);
  // proto:  qreal QRectF::bottom();
extern void _ZNK6QRectF6bottomEv(void* qthis);
  // proto:  bool QRectF::isNull();
extern void _ZNK6QRectF6isNullEv(void* qthis);
  // proto:  void QRectF::QRectF(const QPointF & topleft, const QSizeF & size);
extern void* dector_ZN6QRectFC1ERK7QPointFRK6QSizeF(void* arg0, void* arg1);
extern void _ZN6QRectFC1ERK7QPointFRK6QSizeF(void* qthis, void* arg0, void* arg1);
  // proto:  void QRectF::setWidth(qreal w);
extern void demth_ZN6QRectF8setWidthEd(void* qthis, double arg0);
  // proto:  qreal QRectF::height();
extern void _ZNK6QRectF6heightEv(void* qthis);
  // proto:  void QRectF::translate(const QPointF & p);
extern void demth_ZN6QRectF9translateERK7QPointF(void* qthis, void* arg0);
  // proto:  void QRectF::moveCenter(const QPointF & p);
extern void demth_ZN6QRectF10moveCenterERK7QPointF(void* qthis, void* arg0);
  // proto:  bool QRectF::contains(const QRectF & r);
extern void _ZNK6QRectF8containsERKS_(void* qthis, void* arg0);
  // proto:  QRectF QRectF::marginsRemoved(const QMarginsF & margins);
extern void _ZNK6QRectF14marginsRemovedERK9QMarginsF(void* qthis, void* arg0);
  // proto:  bool QRectF::contains(qreal x, qreal y);
extern void demth_ZNK6QRectF8containsEdd(void* qthis, double arg0, double arg1);
  // proto:  void QRectF::setX(qreal pos);
extern void demth_ZN6QRectF4setXEd(void* qthis, double arg0);
  // proto:  void QRectF::setRect(qreal x, qreal y, qreal w, qreal h);
extern void demth_ZN6QRectF7setRectEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  QPointF QRectF::center();
extern void _ZNK6QRectF6centerEv(void* qthis);
  // proto:  void QRectF::setLeft(qreal pos);
extern void demth_ZN6QRectF7setLeftEd(void* qthis, double arg0);
  // proto:  QRectF QRectF::intersected(const QRectF & other);
extern void demth_ZNK6QRectF11intersectedERKS_(void* qthis, void* arg0);
  // proto:  QPointF QRectF::topLeft();
extern void _ZNK6QRectF7topLeftEv(void* qthis);
  // proto:  qreal QRectF::left();
extern void _ZNK6QRectF4leftEv(void* qthis);
  // proto:  void QRectF::setY(qreal pos);
extern void demth_ZN6QRectF4setYEd(void* qthis, double arg0);
  // proto:  void QRectF::moveTopLeft(const QPointF & p);
extern void demth_ZN6QRectF11moveTopLeftERK7QPointF(void* qthis, void* arg0);
  // proto:  qreal QRectF::width();
extern void _ZNK6QRectF5widthEv(void* qthis);
  // proto:  void QRectF::setTop(qreal pos);
extern void demth_ZN6QRectF6setTopEd(void* qthis, double arg0);
  // proto:  bool QRectF::isValid();
extern void _ZNK6QRectF7isValidEv(void* qthis);
  // proto:  void QRectF::translate(qreal dx, qreal dy);
extern void demth_ZN6QRectF9translateEdd(void* qthis, double arg0, double arg1);
  // proto:  void QRectF::QRectF(qreal left, qreal top, qreal width, qreal height);
extern void* dector_ZN6QRectFC1Edddd(double arg0, double arg1, double arg2, double arg3);
extern void _ZN6QRectFC1Edddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  QRect QRectF::toRect();
extern void _ZNK6QRectF6toRectEv(void* qthis);
  // proto:  void QRectF::moveLeft(qreal pos);
extern void demth_ZN6QRectF8moveLeftEd(void* qthis, double arg0);
  // proto:  void QRectF::setTopLeft(const QPointF & p);
extern void demth_ZN6QRectF10setTopLeftERK7QPointF(void* qthis, void* arg0);
  // proto:  void QRectF::setBottomRight(const QPointF & p);
extern void demth_ZN6QRectF14setBottomRightERK7QPointF(void* qthis, void* arg0);
  // proto:  QRectF QRectF::marginsAdded(const QMarginsF & margins);
extern void _ZNK6QRectF12marginsAddedERK9QMarginsF(void* qthis, void* arg0);
  // proto:  QRectF QRectF::translated(const QPointF & p);
extern void _ZNK6QRectF10translatedERK7QPointF(void* qthis, void* arg0);
  // proto:  QRectF QRectF::normalized();
extern void _ZNK6QRectF10normalizedEv(void* qthis);
  // proto:  void QRectF::getCoords(qreal * x1, qreal * y1, qreal * x2, qreal * y2);
extern void demth_ZNK6QRectF9getCoordsEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto:  void QRectF::setTopRight(const QPointF & p);
extern void demth_ZN6QRectF11setTopRightERK7QPointF(void* qthis, void* arg0);
  // proto:  bool QRectF::contains(const QPointF & p);
extern void _ZNK6QRectF8containsERK7QPointF(void* qthis, void* arg0);
  // proto:  bool QRectF::intersects(const QRectF & r);
extern void _ZNK6QRectF10intersectsERKS_(void* qthis, void* arg0);
  // proto:  void QRectF::moveTop(qreal pos);
extern void demth_ZN6QRectF7moveTopEd(void* qthis, double arg0);
  // proto:  void QRectF::setCoords(qreal x1, qreal y1, qreal x2, qreal y2);
extern void demth_ZN6QRectF9setCoordsEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  QRectF QRectF::translated(qreal dx, qreal dy);
extern void _ZNK6QRectF10translatedEdd(void* qthis, double arg0, double arg1);
  // proto:  bool QRectF::isEmpty();
extern void _ZNK6QRectF7isEmptyEv(void* qthis);
  // proto:  void QRectF::moveTopRight(const QPointF & p);
extern void demth_ZN6QRectF12moveTopRightERK7QPointF(void* qthis, void* arg0);
  // proto:  QRectF QRectF::united(const QRectF & other);
extern void demth_ZNK6QRectF6unitedERKS_(void* qthis, void* arg0);
  // proto:  qreal QRectF::right();
extern void _ZNK6QRectF5rightEv(void* qthis);
  // proto:  void QRectF::QRectF(const QRect & rect);
extern void* dector_ZN6QRectFC1ERK5QRect(void* arg0);
extern void _ZN6QRectFC1ERK5QRect(void* qthis, void* arg0);
  // proto:  QRectF QRectF::adjusted(qreal x1, qreal y1, qreal x2, qreal y2);
extern void _ZNK6QRectF8adjustedEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QRect)=16
type QRect struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QRectF)=32
type QRectF struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QRect::right();
func (this *QRect) right(args ...interface{}) () {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect5rightEv
  default:
    qtrt.ErrorResolve("QRect", "right", args)
  }

}

  // proto:  void QRect::moveTo(const QPoint & p);
func (this *QRect) moveTo(args ...interface{}) () {
  // moveTo(const class QPoint &)
  // moveTo(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect6moveToERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN5QRect6moveToEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QRect", "moveTo", args)
  }

}

  // proto:  void QRect::moveTopLeft(const QPoint & p);
func (this *QRect) moveTopLeft(args ...interface{}) () {
  // moveTopLeft(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect11moveTopLeftERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveTopLeft", args)
  }

}

  // proto:  void QRect::moveRight(int pos);
func (this *QRect) moveRight(args ...interface{}) () {
  // moveRight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9moveRightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveRight", args)
  }

}

  // proto:  void QRect::QRect(const QPoint & topleft, const QPoint & bottomright);
func NewQRect(args ...interface{}) QRect {
  return QRect{}
}

  // proto:  QRect QRect::translated(int dx, int dy);
func (this *QRect) translated(args ...interface{}) () {
  // translated(int, int)
  // translated(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect10translatedEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZNK5QRect10translatedERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "translated", args)
  }

}

  // proto:  QPoint QRect::center();
func (this *QRect) center(args ...interface{}) () {
  // center()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6centerEv
  default:
    qtrt.ErrorResolve("QRect", "center", args)
  }

}

  // proto:  void QRect::moveTopRight(const QPoint & p);
func (this *QRect) moveTopRight(args ...interface{}) () {
  // moveTopRight(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect12moveTopRightERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveTopRight", args)
  }

}

  // proto:  void QRect::setLeft(int pos);
func (this *QRect) setLeft(args ...interface{}) () {
  // setLeft(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect7setLeftEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setLeft", args)
  }

}

  // proto:  int QRect::left();
func (this *QRect) left(args ...interface{}) () {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect4leftEv
  default:
    qtrt.ErrorResolve("QRect", "left", args)
  }

}

  // proto:  QRect QRect::intersected(const QRect & other);
func (this *QRect) intersected(args ...interface{}) () {
  // intersected(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect11intersectedERKS_
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "intersected", args)
  }

}

  // proto:  bool QRect::contains(int x, int y, bool proper);
func (this *QRect) contains(args ...interface{}) () {
  // contains(int, int, _Bool)
  // contains(const class QRect &, _Bool)
  // contains(const class QPoint &, _Bool)
  // contains(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.BoolTy(false) // "bool"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2][1] = qtrt.BoolTy(false) // "bool"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect8containsEiib
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int8_t(args[2].(int8))
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZNK5QRect8containsERKS_b
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  case 2:
    // invoke: _ZNK5QRect8containsERK6QPointb
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  case 3:
    // invoke: _ZNK5QRect8containsEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QRect", "contains", args)
  }

}

  // proto:  QPoint QRect::bottomRight();
func (this *QRect) bottomRight(args ...interface{}) () {
  // bottomRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect11bottomRightEv
  default:
    qtrt.ErrorResolve("QRect", "bottomRight", args)
  }

}

  // proto:  bool QRect::isValid();
func (this *QRect) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect7isValidEv
  default:
    qtrt.ErrorResolve("QRect", "isValid", args)
  }

}

  // proto:  QSize QRect::size();
func (this *QRect) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect4sizeEv
  default:
    qtrt.ErrorResolve("QRect", "size", args)
  }

}

  // proto:  QRect QRect::united(const QRect & other);
func (this *QRect) united(args ...interface{}) () {
  // united(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6unitedERKS_
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "united", args)
  }

}

  // proto:  void QRect::adjust(int x1, int y1, int x2, int y2);
func (this *QRect) adjust(args ...interface{}) () {
  // adjust(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect6adjustEiiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRect", "adjust", args)
  }

}

  // proto:  bool QRect::isNull();
func (this *QRect) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6isNullEv
  default:
    qtrt.ErrorResolve("QRect", "isNull", args)
  }

}

  // proto:  void QRect::setBottom(int pos);
func (this *QRect) setBottom(args ...interface{}) () {
  // setBottom(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9setBottomEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setBottom", args)
  }

}

  // proto:  void QRect::setSize(const QSize & s);
func (this *QRect) setSize(args ...interface{}) () {
  // setSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect7setSizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setSize", args)
  }

}

  // proto:  int QRect::y();
func (this *QRect) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect1yEv
  default:
    qtrt.ErrorResolve("QRect", "y", args)
  }

}

  // proto:  int QRect::x();
func (this *QRect) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect1xEv
  default:
    qtrt.ErrorResolve("QRect", "x", args)
  }

}

  // proto:  QRect QRect::adjusted(int x1, int y1, int x2, int y2);
func (this *QRect) adjusted(args ...interface{}) () {
  // adjusted(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect8adjustedEiiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRect", "adjusted", args)
  }

}

  // proto:  int QRect::height();
func (this *QRect) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6heightEv
  default:
    qtrt.ErrorResolve("QRect", "height", args)
  }

}

  // proto:  void QRect::moveBottomLeft(const QPoint & p);
func (this *QRect) moveBottomLeft(args ...interface{}) () {
  // moveBottomLeft(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect14moveBottomLeftERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveBottomLeft", args)
  }

}

  // proto:  int QRect::top();
func (this *QRect) top(args ...interface{}) () {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect3topEv
  default:
    qtrt.ErrorResolve("QRect", "top", args)
  }

}

  // proto:  void QRect::getRect(int * x, int * y, int * w, int * h);
func (this *QRect) getRect(args ...interface{}) () {
  // getRect(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect7getRectEPiS0_S0_S0_
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRect", "getRect", args)
  }

}

  // proto:  QRect QRect::marginsRemoved(const QMargins & margins);
func (this *QRect) marginsRemoved(args ...interface{}) () {
  // marginsRemoved(const class QMargins &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect14marginsRemovedERK8QMargins
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "marginsRemoved", args)
  }

}

  // proto:  void QRect::translate(int dx, int dy);
func (this *QRect) translate(args ...interface{}) () {
  // translate(int, int)
  // translate(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9translateEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN5QRect9translateERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "translate", args)
  }

}

  // proto:  QPoint QRect::topLeft();
func (this *QRect) topLeft(args ...interface{}) () {
  // topLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect7topLeftEv
  default:
    qtrt.ErrorResolve("QRect", "topLeft", args)
  }

}

  // proto:  int QRect::width();
func (this *QRect) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect5widthEv
  default:
    qtrt.ErrorResolve("QRect", "width", args)
  }

}

  // proto:  void QRect::setRect(int x, int y, int w, int h);
func (this *QRect) setRect(args ...interface{}) () {
  // setRect(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect7setRectEiiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRect", "setRect", args)
  }

}

  // proto:  void QRect::moveCenter(const QPoint & p);
func (this *QRect) moveCenter(args ...interface{}) () {
  // moveCenter(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect10moveCenterERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveCenter", args)
  }

}

  // proto:  bool QRect::intersects(const QRect & r);
func (this *QRect) intersects(args ...interface{}) () {
  // intersects(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect10intersectsERKS_
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "intersects", args)
  }

}

  // proto:  void QRect::setTopRight(const QPoint & p);
func (this *QRect) setTopRight(args ...interface{}) () {
  // setTopRight(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect11setTopRightERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setTopRight", args)
  }

}

  // proto:  void QRect::setCoords(int x1, int y1, int x2, int y2);
func (this *QRect) setCoords(args ...interface{}) () {
  // setCoords(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9setCoordsEiiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRect", "setCoords", args)
  }

}

  // proto:  void QRect::moveBottom(int pos);
func (this *QRect) moveBottom(args ...interface{}) () {
  // moveBottom(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect10moveBottomEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveBottom", args)
  }

}

  // proto:  void QRect::setBottomLeft(const QPoint & p);
func (this *QRect) setBottomLeft(args ...interface{}) () {
  // setBottomLeft(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect13setBottomLeftERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setBottomLeft", args)
  }

}

  // proto:  void QRect::getCoords(int * x1, int * y1, int * x2, int * y2);
func (this *QRect) getCoords(args ...interface{}) () {
  // getCoords(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect9getCoordsEPiS0_S0_S0_
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRect", "getCoords", args)
  }

}

  // proto:  QPoint QRect::topRight();
func (this *QRect) topRight(args ...interface{}) () {
  // topRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect8topRightEv
  default:
    qtrt.ErrorResolve("QRect", "topRight", args)
  }

}

  // proto:  void QRect::setBottomRight(const QPoint & p);
func (this *QRect) setBottomRight(args ...interface{}) () {
  // setBottomRight(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect14setBottomRightERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setBottomRight", args)
  }

}

  // proto:  void QRect::setHeight(int h);
func (this *QRect) setHeight(args ...interface{}) () {
  // setHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect9setHeightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setHeight", args)
  }

}

  // proto:  bool QRect::isEmpty();
func (this *QRect) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect7isEmptyEv
  default:
    qtrt.ErrorResolve("QRect", "isEmpty", args)
  }

}

  // proto:  void QRect::moveBottomRight(const QPoint & p);
func (this *QRect) moveBottomRight(args ...interface{}) () {
  // moveBottomRight(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect15moveBottomRightERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveBottomRight", args)
  }

}

  // proto:  QPoint QRect::bottomLeft();
func (this *QRect) bottomLeft(args ...interface{}) () {
  // bottomLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect10bottomLeftEv
  default:
    qtrt.ErrorResolve("QRect", "bottomLeft", args)
  }

}

  // proto:  void QRect::setTop(int pos);
func (this *QRect) setTop(args ...interface{}) () {
  // setTop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect6setTopEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setTop", args)
  }

}

  // proto:  int QRect::bottom();
func (this *QRect) bottom(args ...interface{}) () {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect6bottomEv
  default:
    qtrt.ErrorResolve("QRect", "bottom", args)
  }

}

  // proto:  QRect QRect::marginsAdded(const QMargins & margins);
func (this *QRect) marginsAdded(args ...interface{}) () {
  // marginsAdded(const class QMargins &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect12marginsAddedERK8QMargins
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "marginsAdded", args)
  }

}

  // proto:  QRect QRect::normalized();
func (this *QRect) normalized(args ...interface{}) () {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QRect10normalizedEv
  default:
    qtrt.ErrorResolve("QRect", "normalized", args)
  }

}

  // proto:  void QRect::setWidth(int w);
func (this *QRect) setWidth(args ...interface{}) () {
  // setWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect8setWidthEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setWidth", args)
  }

}

  // proto:  void QRect::setY(int y);
func (this *QRect) setY(args ...interface{}) () {
  // setY(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect4setYEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setY", args)
  }

}

  // proto:  void QRect::moveTop(int pos);
func (this *QRect) moveTop(args ...interface{}) () {
  // moveTop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect7moveTopEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveTop", args)
  }

}

  // proto:  void QRect::setX(int x);
func (this *QRect) setX(args ...interface{}) () {
  // setX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect4setXEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setX", args)
  }

}

  // proto:  void QRect::setRight(int pos);
func (this *QRect) setRight(args ...interface{}) () {
  // setRight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect8setRightEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setRight", args)
  }

}

  // proto:  void QRect::setTopLeft(const QPoint & p);
func (this *QRect) setTopLeft(args ...interface{}) () {
  // setTopLeft(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect10setTopLeftERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "setTopLeft", args)
  }

}

  // proto:  void QRect::moveLeft(int pos);
func (this *QRect) moveLeft(args ...interface{}) () {
  // moveLeft(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QRect8moveLeftEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRect", "moveLeft", args)
  }

}

  // proto:  void QRectF::QRectF();
func NewQRectF(args ...interface{}) QRectF {
  return QRectF{}
}

  // proto:  void QRectF::moveBottomRight(const QPointF & p);
func (this *QRectF) moveBottomRight(args ...interface{}) () {
  // moveBottomRight(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF15moveBottomRightERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveBottomRight", args)
  }

}

  // proto:  void QRectF::moveTo(qreal x, qreal y);
func (this *QRectF) moveTo(args ...interface{}) () {
  // moveTo(qreal, qreal)
  // moveTo(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF6moveToEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN6QRectF6moveToERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveTo", args)
  }

}

  // proto:  qreal QRectF::top();
func (this *QRectF) top(args ...interface{}) () {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF3topEv
  default:
    qtrt.ErrorResolve("QRectF", "top", args)
  }

}

  // proto:  QPointF QRectF::bottomLeft();
func (this *QRectF) bottomLeft(args ...interface{}) () {
  // bottomLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF10bottomLeftEv
  default:
    qtrt.ErrorResolve("QRectF", "bottomLeft", args)
  }

}

  // proto:  void QRectF::setHeight(qreal h);
func (this *QRectF) setHeight(args ...interface{}) () {
  // setHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9setHeightEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setHeight", args)
  }

}

  // proto:  void QRectF::setSize(const QSizeF & s);
func (this *QRectF) setSize(args ...interface{}) () {
  // setSize(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF7setSizeERK6QSizeF
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setSize", args)
  }

}

  // proto:  QRect QRectF::toAlignedRect();
func (this *QRectF) toAlignedRect(args ...interface{}) () {
  // toAlignedRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF13toAlignedRectEv
  default:
    qtrt.ErrorResolve("QRectF", "toAlignedRect", args)
  }

}

  // proto:  void QRectF::setRight(qreal pos);
func (this *QRectF) setRight(args ...interface{}) () {
  // setRight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF8setRightEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setRight", args)
  }

}

  // proto:  void QRectF::setBottomLeft(const QPointF & p);
func (this *QRectF) setBottomLeft(args ...interface{}) () {
  // setBottomLeft(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF13setBottomLeftERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setBottomLeft", args)
  }

}

  // proto:  QPointF QRectF::topRight();
func (this *QRectF) topRight(args ...interface{}) () {
  // topRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF8topRightEv
  default:
    qtrt.ErrorResolve("QRectF", "topRight", args)
  }

}

  // proto:  QSizeF QRectF::size();
func (this *QRectF) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF4sizeEv
  default:
    qtrt.ErrorResolve("QRectF", "size", args)
  }

}

  // proto:  void QRectF::adjust(qreal x1, qreal y1, qreal x2, qreal y2);
func (this *QRectF) adjust(args ...interface{}) () {
  // adjust(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF6adjustEdddd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRectF", "adjust", args)
  }

}

  // proto:  void QRectF::moveRight(qreal pos);
func (this *QRectF) moveRight(args ...interface{}) () {
  // moveRight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9moveRightEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveRight", args)
  }

}

  // proto:  qreal QRectF::y();
func (this *QRectF) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF1yEv
  default:
    qtrt.ErrorResolve("QRectF", "y", args)
  }

}

  // proto:  QPointF QRectF::bottomRight();
func (this *QRectF) bottomRight(args ...interface{}) () {
  // bottomRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF11bottomRightEv
  default:
    qtrt.ErrorResolve("QRectF", "bottomRight", args)
  }

}

  // proto:  void QRectF::setBottom(qreal pos);
func (this *QRectF) setBottom(args ...interface{}) () {
  // setBottom(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9setBottomEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setBottom", args)
  }

}

  // proto:  void QRectF::moveBottomLeft(const QPointF & p);
func (this *QRectF) moveBottomLeft(args ...interface{}) () {
  // moveBottomLeft(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF14moveBottomLeftERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveBottomLeft", args)
  }

}

  // proto:  void QRectF::moveBottom(qreal pos);
func (this *QRectF) moveBottom(args ...interface{}) () {
  // moveBottom(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF10moveBottomEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveBottom", args)
  }

}

  // proto:  void QRectF::getRect(qreal * x, qreal * y, qreal * w, qreal * h);
func (this *QRectF) getRect(args ...interface{}) () {
  // getRect(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF7getRectEPdS0_S0_S0_
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRectF", "getRect", args)
  }

}

  // proto:  qreal QRectF::x();
func (this *QRectF) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF1xEv
  default:
    qtrt.ErrorResolve("QRectF", "x", args)
  }

}

  // proto:  qreal QRectF::bottom();
func (this *QRectF) bottom(args ...interface{}) () {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6bottomEv
  default:
    qtrt.ErrorResolve("QRectF", "bottom", args)
  }

}

  // proto:  bool QRectF::isNull();
func (this *QRectF) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6isNullEv
  default:
    qtrt.ErrorResolve("QRectF", "isNull", args)
  }

}

  // proto:  void QRectF::setWidth(qreal w);
func (this *QRectF) setWidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF8setWidthEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setWidth", args)
  }

}

  // proto:  qreal QRectF::height();
func (this *QRectF) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6heightEv
  default:
    qtrt.ErrorResolve("QRectF", "height", args)
  }

}

  // proto:  void QRectF::translate(const QPointF & p);
func (this *QRectF) translate(args ...interface{}) () {
  // translate(const class QPointF &)
  // translate(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9translateERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN6QRectF9translateEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QRectF", "translate", args)
  }

}

  // proto:  void QRectF::moveCenter(const QPointF & p);
func (this *QRectF) moveCenter(args ...interface{}) () {
  // moveCenter(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF10moveCenterERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveCenter", args)
  }

}

  // proto:  bool QRectF::contains(const QRectF & r);
func (this *QRectF) contains(args ...interface{}) () {
  // contains(const class QRectF &)
  // contains(qreal, qreal)
  // contains(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF8containsERKS_
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK6QRectF8containsEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  case 2:
    // invoke: _ZNK6QRectF8containsERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "contains", args)
  }

}

  // proto:  QRectF QRectF::marginsRemoved(const QMarginsF & margins);
func (this *QRectF) marginsRemoved(args ...interface{}) () {
  // marginsRemoved(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF14marginsRemovedERK9QMarginsF
    var arg0 = args[0].(QMarginsF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "marginsRemoved", args)
  }

}

  // proto:  void QRectF::setX(qreal pos);
func (this *QRectF) setX(args ...interface{}) () {
  // setX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF4setXEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setX", args)
  }

}

  // proto:  void QRectF::setRect(qreal x, qreal y, qreal w, qreal h);
func (this *QRectF) setRect(args ...interface{}) () {
  // setRect(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF7setRectEdddd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRectF", "setRect", args)
  }

}

  // proto:  QPointF QRectF::center();
func (this *QRectF) center(args ...interface{}) () {
  // center()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6centerEv
  default:
    qtrt.ErrorResolve("QRectF", "center", args)
  }

}

  // proto:  void QRectF::setLeft(qreal pos);
func (this *QRectF) setLeft(args ...interface{}) () {
  // setLeft(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF7setLeftEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setLeft", args)
  }

}

  // proto:  QRectF QRectF::intersected(const QRectF & other);
func (this *QRectF) intersected(args ...interface{}) () {
  // intersected(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF11intersectedERKS_
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "intersected", args)
  }

}

  // proto:  QPointF QRectF::topLeft();
func (this *QRectF) topLeft(args ...interface{}) () {
  // topLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF7topLeftEv
  default:
    qtrt.ErrorResolve("QRectF", "topLeft", args)
  }

}

  // proto:  qreal QRectF::left();
func (this *QRectF) left(args ...interface{}) () {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF4leftEv
  default:
    qtrt.ErrorResolve("QRectF", "left", args)
  }

}

  // proto:  void QRectF::setY(qreal pos);
func (this *QRectF) setY(args ...interface{}) () {
  // setY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF4setYEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setY", args)
  }

}

  // proto:  void QRectF::moveTopLeft(const QPointF & p);
func (this *QRectF) moveTopLeft(args ...interface{}) () {
  // moveTopLeft(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF11moveTopLeftERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveTopLeft", args)
  }

}

  // proto:  qreal QRectF::width();
func (this *QRectF) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF5widthEv
  default:
    qtrt.ErrorResolve("QRectF", "width", args)
  }

}

  // proto:  void QRectF::setTop(qreal pos);
func (this *QRectF) setTop(args ...interface{}) () {
  // setTop(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF6setTopEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setTop", args)
  }

}

  // proto:  bool QRectF::isValid();
func (this *QRectF) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF7isValidEv
  default:
    qtrt.ErrorResolve("QRectF", "isValid", args)
  }

}

  // proto:  QRect QRectF::toRect();
func (this *QRectF) toRect(args ...interface{}) () {
  // toRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6toRectEv
  default:
    qtrt.ErrorResolve("QRectF", "toRect", args)
  }

}

  // proto:  void QRectF::moveLeft(qreal pos);
func (this *QRectF) moveLeft(args ...interface{}) () {
  // moveLeft(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF8moveLeftEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveLeft", args)
  }

}

  // proto:  void QRectF::setTopLeft(const QPointF & p);
func (this *QRectF) setTopLeft(args ...interface{}) () {
  // setTopLeft(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF10setTopLeftERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setTopLeft", args)
  }

}

  // proto:  void QRectF::setBottomRight(const QPointF & p);
func (this *QRectF) setBottomRight(args ...interface{}) () {
  // setBottomRight(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF14setBottomRightERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setBottomRight", args)
  }

}

  // proto:  QRectF QRectF::marginsAdded(const QMarginsF & margins);
func (this *QRectF) marginsAdded(args ...interface{}) () {
  // marginsAdded(const class QMarginsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMarginsF{}) // "const QMarginsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF12marginsAddedERK9QMarginsF
    var arg0 = args[0].(QMarginsF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "marginsAdded", args)
  }

}

  // proto:  QRectF QRectF::translated(const QPointF & p);
func (this *QRectF) translated(args ...interface{}) () {
  // translated(const class QPointF &)
  // translated(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[1][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF10translatedERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK6QRectF10translatedEdd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QRectF", "translated", args)
  }

}

  // proto:  QRectF QRectF::normalized();
func (this *QRectF) normalized(args ...interface{}) () {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF10normalizedEv
  default:
    qtrt.ErrorResolve("QRectF", "normalized", args)
  }

}

  // proto:  void QRectF::getCoords(qreal * x1, qreal * y1, qreal * x2, qreal * y2);
func (this *QRectF) getCoords(args ...interface{}) () {
  // getCoords(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF9getCoordsEPdS0_S0_S0_
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRectF", "getCoords", args)
  }

}

  // proto:  void QRectF::setTopRight(const QPointF & p);
func (this *QRectF) setTopRight(args ...interface{}) () {
  // setTopRight(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF11setTopRightERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "setTopRight", args)
  }

}

  // proto:  bool QRectF::intersects(const QRectF & r);
func (this *QRectF) intersects(args ...interface{}) () {
  // intersects(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF10intersectsERKS_
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "intersects", args)
  }

}

  // proto:  void QRectF::moveTop(qreal pos);
func (this *QRectF) moveTop(args ...interface{}) () {
  // moveTop(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF7moveTopEd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveTop", args)
  }

}

  // proto:  void QRectF::setCoords(qreal x1, qreal y1, qreal x2, qreal y2);
func (this *QRectF) setCoords(args ...interface{}) () {
  // setCoords(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF9setCoordsEdddd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRectF", "setCoords", args)
  }

}

  // proto:  bool QRectF::isEmpty();
func (this *QRectF) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF7isEmptyEv
  default:
    qtrt.ErrorResolve("QRectF", "isEmpty", args)
  }

}

  // proto:  void QRectF::moveTopRight(const QPointF & p);
func (this *QRectF) moveTopRight(args ...interface{}) () {
  // moveTopRight(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QRectF12moveTopRightERK7QPointF
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "moveTopRight", args)
  }

}

  // proto:  QRectF QRectF::united(const QRectF & other);
func (this *QRectF) united(args ...interface{}) () {
  // united(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF6unitedERKS_
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QRectF", "united", args)
  }

}

  // proto:  qreal QRectF::right();
func (this *QRectF) right(args ...interface{}) () {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF5rightEv
  default:
    qtrt.ErrorResolve("QRectF", "right", args)
  }

}

  // proto:  QRectF QRectF::adjusted(qreal x1, qreal y1, qreal x2, qreal y2);
func (this *QRectF) adjusted(args ...interface{}) () {
  // adjusted(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QRectF8adjustedEdddd
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QRectF", "adjusted", args)
  }

}

// <= body block end

