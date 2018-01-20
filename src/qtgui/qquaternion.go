//  header block begin
// /usr/include/qt/QtGui/qquaternion.h
// #include <qquaternion.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
type QQuaternion struct {
	*qtrt.CObject
}

func (this *QQuaternion) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQQuaternionFromPointer(cthis unsafe.Pointer) *QQuaternion {
	return &QQuaternion{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qquaternion.h:59
// index:0
// Public
// void QQuaternion()
func NewQQuaternion() *QQuaternion {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:60
// index:1
// Public inline
// void QQuaternion(Qt::Initialization)
func NewQQuaternion_1(arg0 int) *QQuaternion {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, &arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:61
// index:2
// Public
// void QQuaternion(float, float, float, float)
func NewQQuaternion_2(scalar float32, xpos float32, ypos float32, zpos float32) *QQuaternion {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2Effff", ffiqt.FFI_TYPE_VOID, cthis, &scalar, &xpos, &ypos, &zpos)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:63
// index:3
// Public
// void QQuaternion(float, const class QVector3D &)
func NewQQuaternion_3(scalar float32, vector *QVector3D) *QQuaternion {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2EfRK9QVector3D", ffiqt.FFI_TYPE_VOID, cthis, &scalar, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:66
// index:4
// Public
// void QQuaternion(const class QVector4D &)
func NewQQuaternion_4(vector *QVector4D) *QQuaternion {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternionC2ERK9QVector4D", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuaternionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qquaternion.h:69
// index:0
// Public
// bool isNull()
func (this *QQuaternion) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:70
// index:0
// Public
// bool isIdentity()
func (this *QQuaternion) IsIdentity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion10isIdentityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:73
// index:0
// Public
// QVector3D vector()
func (this *QQuaternion) Vector() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion6vectorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:74
// index:0
// Public
// void setVector(const class QVector3D &)
func (this *QQuaternion) SetVector(vector *QVector3D) {
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion9setVectorERK9QVector3D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:76
// index:1
// Public
// void setVector(float, float, float)
func (this *QQuaternion) SetVector_1(x float32, y float32, z float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion9setVectorEfff", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x, &y, &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:78
// index:0
// Public
// float x()
func (this *QQuaternion) X() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:79
// index:0
// Public
// float y()
func (this *QQuaternion) Y() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:80
// index:0
// Public
// float z()
func (this *QQuaternion) Z() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion1zEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:81
// index:0
// Public
// float scalar()
func (this *QQuaternion) Scalar() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion6scalarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:83
// index:0
// Public
// void setX(float)
func (this *QQuaternion) SetX(x float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion4setXEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:84
// index:0
// Public
// void setY(float)
func (this *QQuaternion) SetY(y float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion4setYEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:85
// index:0
// Public
// void setZ(float)
func (this *QQuaternion) SetZ(z float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion4setZEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:86
// index:0
// Public
// void setScalar(float)
func (this *QQuaternion) SetScalar(scalar float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion9setScalarEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &scalar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:88
// index:0
// Public static inline
// float dotProduct(const class QQuaternion &, const class QQuaternion &)
func (this *QQuaternion) DotProduct(q1 *QQuaternion, q2 *QQuaternion) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion10dotProductERKS_S1_", ffiqt.FFI_TYPE_POINTER, q1, q2)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_DotProduct(q1 *QQuaternion, q2 *QQuaternion) {
	var nilthis *QQuaternion
	nilthis.DotProduct(q1, q2)
}

// /usr/include/qt/QtGui/qquaternion.h:90
// index:0
// Public
// float length()
func (this *QQuaternion) Length() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:91
// index:0
// Public
// float lengthSquared()
func (this *QQuaternion) LengthSquared() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion13lengthSquaredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:93
// index:0
// Public
// QQuaternion normalized()
func (this *QQuaternion) Normalized() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion10normalizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:94
// index:0
// Public
// void normalize()
func (this *QQuaternion) Normalize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion9normalizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:96
// index:0
// Public inline
// QQuaternion inverted()
func (this *QQuaternion) Inverted() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion8invertedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:98
// index:0
// Public
// QQuaternion conjugated()
func (this *QQuaternion) Conjugated() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion10conjugatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:100
// index:0
// Public
// QQuaternion conjugate()
func (this *QQuaternion) Conjugate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion9conjugateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:103
// index:0
// Public
// QVector3D rotatedVector(const class QVector3D &)
func (this *QQuaternion) RotatedVector(vector *QVector3D) interface{} {
	var convArg0 = vector.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion13rotatedVectorERK9QVector3D", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:124
// index:0
// Public
// QVector4D toVector4D()
func (this *QQuaternion) ToVector4D() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion10toVector4DEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:130
// index:0
// Public inline
// void getAxisAndAngle(class QVector3D *, float *)
func (this *QQuaternion) GetAxisAndAngle(axis unsafe.Pointer, angle unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion15getAxisAndAngleEP9QVector3DPf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), axis, angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:133
// index:1
// Public
// void getAxisAndAngle(float *, float *, float *, float *)
func (this *QQuaternion) GetAxisAndAngle_1(x unsafe.Pointer, y unsafe.Pointer, z unsafe.Pointer, angle unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion15getAxisAndAngleEPfS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, z, angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:131
// index:0
// Public static
// QQuaternion fromAxisAndAngle(const class QVector3D &, float)
func (this *QQuaternion) FromAxisAndAngle(axis *QVector3D, angle float32) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion16fromAxisAndAngleERK9QVector3Df", ffiqt.FFI_TYPE_POINTER, axis, angle)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_FromAxisAndAngle(axis *QVector3D, angle float32) {
	var nilthis *QQuaternion
	nilthis.FromAxisAndAngle(axis, angle)
}

// /usr/include/qt/QtGui/qquaternion.h:134
// index:1
// Public static
// QQuaternion fromAxisAndAngle(float, float, float, float)
func (this *QQuaternion) FromAxisAndAngle_1(x float32, y float32, z float32, angle float32) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion16fromAxisAndAngleEffff", ffiqt.FFI_TYPE_POINTER, x, y, z, angle)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_FromAxisAndAngle_1(x float32, y float32, z float32, angle float32) {
	var nilthis *QQuaternion
	nilthis.FromAxisAndAngle_1(x, y, z, angle)
}

// /usr/include/qt/QtGui/qquaternion.h:138
// index:0
// Public inline
// QVector3D toEulerAngles()
func (this *QQuaternion) ToEulerAngles() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion13toEulerAnglesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:139
// index:0
// Public static inline
// QQuaternion fromEulerAngles(const class QVector3D &)
func (this *QQuaternion) FromEulerAngles(eulerAngles *QVector3D) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion15fromEulerAnglesERK9QVector3D", ffiqt.FFI_TYPE_POINTER, eulerAngles)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_FromEulerAngles(eulerAngles *QVector3D) {
	var nilthis *QQuaternion
	nilthis.FromEulerAngles(eulerAngles)
}

// /usr/include/qt/QtGui/qquaternion.h:142
// index:1
// Public static
// QQuaternion fromEulerAngles(float, float, float)
func (this *QQuaternion) FromEulerAngles_1(pitch float32, yaw float32, roll float32) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion15fromEulerAnglesEfff", ffiqt.FFI_TYPE_POINTER, pitch, yaw, roll)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_FromEulerAngles_1(pitch float32, yaw float32, roll float32) {
	var nilthis *QQuaternion
	nilthis.FromEulerAngles_1(pitch, yaw, roll)
}

// /usr/include/qt/QtGui/qquaternion.h:141
// index:0
// Public
// void getEulerAngles(float *, float *, float *)
func (this *QQuaternion) GetEulerAngles(pitch unsafe.Pointer, yaw unsafe.Pointer, roll unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion14getEulerAnglesEPfS0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pitch, yaw, roll)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:144
// index:0
// Public
// QMatrix3x3 toRotationMatrix()
func (this *QQuaternion) ToRotationMatrix() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion16toRotationMatrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qquaternion.h:148
// index:0
// Public
// void getAxes(class QVector3D *, class QVector3D *, class QVector3D *)
func (this *QQuaternion) GetAxes(xAxis unsafe.Pointer, yAxis unsafe.Pointer, zAxis unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QQuaternion7getAxesEP9QVector3DS1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), xAxis, yAxis, zAxis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qquaternion.h:149
// index:0
// Public static
// QQuaternion fromAxes(const class QVector3D &, const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) FromAxes(xAxis *QVector3D, yAxis *QVector3D, zAxis *QVector3D) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion8fromAxesERK9QVector3DS2_S2_", ffiqt.FFI_TYPE_POINTER, xAxis, yAxis, zAxis)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_FromAxes(xAxis *QVector3D, yAxis *QVector3D, zAxis *QVector3D) {
	var nilthis *QQuaternion
	nilthis.FromAxes(xAxis, yAxis, zAxis)
}

// /usr/include/qt/QtGui/qquaternion.h:151
// index:0
// Public static
// QQuaternion fromDirection(const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) FromDirection(direction *QVector3D, up *QVector3D) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion13fromDirectionERK9QVector3DS2_", ffiqt.FFI_TYPE_POINTER, direction, up)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_FromDirection(direction *QVector3D, up *QVector3D) {
	var nilthis *QQuaternion
	nilthis.FromDirection(direction, up)
}

// /usr/include/qt/QtGui/qquaternion.h:153
// index:0
// Public static
// QQuaternion rotationTo(const class QVector3D &, const class QVector3D &)
func (this *QQuaternion) RotationTo(from *QVector3D, to *QVector3D) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion10rotationToERK9QVector3DS2_", ffiqt.FFI_TYPE_POINTER, from, to)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_RotationTo(from *QVector3D, to *QVector3D) {
	var nilthis *QQuaternion
	nilthis.RotationTo(from, to)
}

// /usr/include/qt/QtGui/qquaternion.h:156
// index:0
// Public static
// QQuaternion slerp(const class QQuaternion &, const class QQuaternion &, float)
func (this *QQuaternion) Slerp(q1 *QQuaternion, q2 *QQuaternion, t float32) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion5slerpERKS_S1_f", ffiqt.FFI_TYPE_POINTER, q1, q2, t)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_Slerp(q1 *QQuaternion, q2 *QQuaternion, t float32) {
	var nilthis *QQuaternion
	nilthis.Slerp(q1, q2, t)
}

// /usr/include/qt/QtGui/qquaternion.h:158
// index:0
// Public static
// QQuaternion nlerp(const class QQuaternion &, const class QQuaternion &, float)
func (this *QQuaternion) Nlerp(q1 *QQuaternion, q2 *QQuaternion, t float32) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QQuaternion5nlerpERKS_S1_f", ffiqt.FFI_TYPE_POINTER, q1, q2, t)
	gopp.ErrPrint(err, rv)
	return rv
}
func QQuaternion_Nlerp(q1 *QQuaternion, q2 *QQuaternion, t float32) {
	var nilthis *QQuaternion
	nilthis.Nlerp(q1, q2, t)
}

//  body block end
