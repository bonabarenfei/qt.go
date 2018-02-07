package qtmacextras

import "unsafe"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
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

//  header block end

//  body block begin
// /usr/include/qt/QtMacExtras/../../src/macextras/qmacpasteboardmime.h:56
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qRegisterDraggedTypes(const QStringList &)
func QRegisterDraggedTypes(types *qtcore.QStringList) {
	var convArg0 = types.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z21qRegisterDraggedTypesRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end