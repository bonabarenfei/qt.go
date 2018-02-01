package qtgui

// /usr/include/qt/QtGui/qimagereader.h
// #include <qimagereader.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QImageReader struct {
	*qtrt.CObject
}

func (this *QImageReader) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QImageReader) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQImageReaderFromPointer(cthis unsafe.Pointer) *QImageReader {
	return &QImageReader{&qtrt.CObject{cthis}}
}
func (*QImageReader) NewFromPointer(cthis unsafe.Pointer) *QImageReader {
	return NewQImageReaderFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimagereader.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageReader()
func NewQImageReader() *QImageReader {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:72
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QImageReader(QIODevice *, const QByteArray &)
func NewQImageReader_1(device *qtcore.QIODevice /*777 QIODevice **/, format *qtcore.QByteArray) *QImageReader {
	var convArg0 = device.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2EP9QIODeviceRK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:73
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImageReader(const QString &, const QByteArray &)
func NewQImageReader_2(fileName *qtcore.QString, format *qtcore.QByteArray) *QImageReader {
	var convArg0 = fileName.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderC2ERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageReaderFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qimagereader.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QImageReader()
func DeleteQImageReader(*QImageReader) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReaderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)
func (this *QImageReader) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader9setFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format()
func (this *QImageReader) Format() *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoDetectImageFormat(_Bool)
func (this *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader24setAutoDetectImageFormatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoDetectImageFormat()
func (this *QImageReader) AutoDetectImageFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader21autoDetectImageFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDecideFormatFromContent(_Bool)
func (this *QImageReader) SetDecideFormatFromContent(ignored bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader26setDecideFormatFromContentEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ignored)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool decideFormatFromContent()
func (this *QImageReader) DecideFormatFromContent() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader23decideFormatFromContentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QImageReader) SetDevice(device *qtcore.QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QImageReader) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QImageReader) SetFileName(fileName *qtcore.QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QImageReader) FileName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size()
func (this *QImageReader) Size() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QImage::Format imageFormat()
func (this *QImageReader) ImageFormat() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader11imageFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:143
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray imageFormat(const QString &)
func (this *QImageReader) ImageFormat_1(fileName *qtcore.QString) *qtcore.QByteArray /*123*/ {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11imageFormatERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QImageReader_ImageFormat_1(fileName *qtcore.QString) *qtcore.QByteArray /*123*/ {
	var nilthis *QImageReader
	rv := nilthis.ImageFormat_1(fileName)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:144
// index:2
// Public static Visibility=Default Availability=Available
// [8] QByteArray imageFormat(QIODevice *)
func (this *QImageReader) ImageFormat_2(device *qtcore.QIODevice /*777 QIODevice **/) *qtcore.QByteArray /*123*/ {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11imageFormatEP9QIODevice", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QImageReader_ImageFormat_2(device *qtcore.QIODevice /*777 QIODevice **/) *qtcore.QByteArray /*123*/ {
	var nilthis *QImageReader
	rv := nilthis.ImageFormat_2(device)
	return rv
}

// /usr/include/qt/QtGui/qimagereader.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(const QString &)
func (this *QImageReader) Text(key *qtcore.QString) *qtcore.QString /*123*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader4textERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRect &)
func (this *QImageReader) SetClipRect(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11setClipRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:99
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect clipRect()
func (this *QImageReader) ClipRect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader8clipRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaledSize(const QSize &)
func (this *QImageReader) SetScaledSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader13setScaledSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize scaledSize()
func (this *QImageReader) ScaledSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader10scaledSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuality(int)
func (this *QImageReader) SetQuality(quality int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader10setQualityEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int quality()
func (this *QImageReader) Quality() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7qualityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScaledClipRect(const QRect &)
func (this *QImageReader) SetScaledClipRect(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader17setScaledClipRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:108
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect scaledClipRect()
func (this *QImageReader) ScaledClipRect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14scaledClipRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundColor(const QColor &)
func (this *QImageReader) SetBackgroundColor(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:111
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor backgroundColor()
func (this *QImageReader) BackgroundColor() *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsAnimation()
func (this *QImageReader) SupportsAnimation() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader17supportsAnimationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] QImageIOHandler::Transformations transformation()
func (this *QImageReader) Transformation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14transformationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoTransform(_Bool)
func (this *QImageReader) SetAutoTransform(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader16setAutoTransformEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoTransform()
func (this *QImageReader) AutoTransform() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader13autoTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGamma(float)
func (this *QImageReader) SetGamma(gamma float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader8setGammaEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), gamma)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagereader.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] float gamma()
func (this *QImageReader) Gamma() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader5gammaEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray subType()
func (this *QImageReader) SubType() *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7subTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canRead()
func (this *QImageReader) CanRead() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader7canReadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:127
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage read()
func (this *QImageReader) Read() *QImage /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader4readEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:128
// index:1
// Public Visibility=Default Availability=Available
// [1] bool read(QImage *)
func (this *QImageReader) Read_1(image *QImage /*777 QImage **/) bool {
	var convArg0 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader4readEP6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:130
// index:0
// Public Visibility=Default Availability=Available
// [1] bool jumpToNextImage()
func (this *QImageReader) JumpToNextImage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader15jumpToNextImageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool jumpToImage(int)
func (this *QImageReader) JumpToImage(imageNumber int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QImageReader11jumpToImageEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), imageNumber)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimagereader.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopCount()
func (this *QImageReader) LoopCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader9loopCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] int imageCount()
func (this *QImageReader) ImageCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader10imageCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int nextImageDelay()
func (this *QImageReader) NextImageDelay() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14nextImageDelayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentImageNumber()
func (this *QImageReader) CurrentImageNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader18currentImageNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagereader.h:136
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect currentImageRect()
func (this *QImageReader) CurrentImageRect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader16currentImageRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:138
// index:0
// Public Visibility=Default Availability=Available
// [4] QImageReader::ImageReaderError error()
func (this *QImageReader) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qimagereader.h:139
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QImageReader) ErrorString() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qimagereader.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsOption(QImageIOHandler::ImageOption)
func (this *QImageReader) SupportsOption(option int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QImageReader14supportsOptionEN15QImageIOHandler11ImageOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QImageReader__ImageReaderError = int

const QImageReader__UnknownError QImageReader__ImageReaderError = 0
const QImageReader__FileNotFoundError QImageReader__ImageReaderError = 1
const QImageReader__DeviceError QImageReader__ImageReaderError = 2
const QImageReader__UnsupportedFormatError QImageReader__ImageReaderError = 3
const QImageReader__InvalidDataError QImageReader__ImageReaderError = 4

//  body block end