package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qopenglshaderprogram.h
// dst-file: /src/gui/qopenglshaderprogram.go
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
  // proto:  void QOpenGLShader::QOpenGLShader(const QOpenGLShader & );
extern void* dector_ZN13QOpenGLShaderC1ERKS_(void* arg0);
extern void _ZN13QOpenGLShaderC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QOpenGLShader::isCompiled();
extern void _ZNK13QOpenGLShader10isCompiledEv(void* qthis);
  // proto:  const QMetaObject * QOpenGLShader::metaObject();
extern void _ZNK13QOpenGLShader10metaObjectEv(void* qthis);
  // proto:  QString QOpenGLShader::log();
extern void _ZNK13QOpenGLShader3logEv(void* qthis);
  // proto:  bool QOpenGLShader::compileSourceCode(const QString & source);
extern void _ZN13QOpenGLShader17compileSourceCodeERK7QString(void* qthis, void* arg0);
  // proto:  bool QOpenGLShader::compileSourceFile(const QString & fileName);
extern void _ZN13QOpenGLShader17compileSourceFileERK7QString(void* qthis, void* arg0);
  // proto:  QByteArray QOpenGLShader::sourceCode();
extern void _ZNK13QOpenGLShader10sourceCodeEv(void* qthis);
  // proto:  void QOpenGLShader::~QOpenGLShader();
extern void _ZN13QOpenGLShaderD0Ev(void* qthis);
  // proto:  bool QOpenGLShader::compileSourceCode(const QByteArray & source);
extern void _ZN13QOpenGLShader17compileSourceCodeERK10QByteArray(void* qthis, void* arg0);
  // proto:  GLuint QOpenGLShader::shaderId();
extern void _ZNK13QOpenGLShader8shaderIdEv(void* qthis);
  // proto:  bool QOpenGLShader::compileSourceCode(const char * source);
extern void _ZN13QOpenGLShader17compileSourceCodeEPKc(void* qthis, char* arg0);
  // proto:  bool QOpenGLShaderProgram::isLinked();
extern void _ZNK20QOpenGLShaderProgram8isLinkedEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QVector3D & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector3D(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QPoint & point);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QPoint(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcffff(void* qthis, char* arg0, float arg1, float arg2, float arg3, float arg4);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QVector3D & value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector3D(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QVector3D * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector3Di(void* qthis, char* arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QTransform & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK10QTransform(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeBuffer(int location, GLenum type, int offset, int tupleSize, int stride);
extern void _ZN20QOpenGLShaderProgram18setAttributeBufferEijiii(void* qthis, int arg0, unsigned int arg1, int arg2, int arg3, int arg4);
  // proto: static bool QOpenGLShaderProgram::hasOpenGLShaderPrograms(QOpenGLContext * context);
extern void _ZN20QOpenGLShaderProgram23hasOpenGLShaderProgramsEP14QOpenGLContext(void* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const GLfloat [2][2] value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA2_Kf(void* qthis, char* arg0, float** arg1);
  // proto:  void QOpenGLShaderProgram::setPatchVertexCount(int count);
extern void _ZN20QOpenGLShaderProgram19setPatchVertexCountEi(void* qthis, int arg0);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const GLfloat * values, int tupleSize, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPKfii(void* qthis, char* arg0, float* arg1, int arg2, int arg3);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const QVector3D * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector3Di(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const GLfloat [3][3] value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA3_Kf(void* qthis, char* arg0, float** arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QVector2D & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector2D(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::bindAttributeLocation(const QByteArray & name, int location);
extern void _ZN20QOpenGLShaderProgram21bindAttributeLocationERK10QByteArrayi(void* qthis, void* arg0, int arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const GLfloat * values, int count, int tupleSize);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKfii(void* qthis, char* arg0, float* arg1, int arg2, int arg3);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, GLfloat x, GLfloat y);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEiff(void* qthis, int arg0, float arg1, float arg2);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QVector2D & value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector2D(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const QColor & value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEiRK6QColor(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, GLfloat x, GLfloat y, GLfloat z);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEifff(void* qthis, int arg0, float arg1, float arg2, float arg3);
  // proto:  bool QOpenGLShaderProgram::bind();
extern void _ZN20QOpenGLShaderProgram4bindEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLfloat value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEif(void* qthis, int arg0, float arg1);
  // proto:  void QOpenGLShaderProgram::enableAttributeArray(int location);
extern void _ZN20QOpenGLShaderProgram20enableAttributeArrayEi(void* qthis, int arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLint value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEii(void* qthis, int arg0, int arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLuint value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEij(void* qthis, int arg0, unsigned int arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QSize & size);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK5QSize(void* qthis, int arg0, void* arg1);
  // proto:  bool QOpenGLShaderProgram::addShader(QOpenGLShader * shader);
extern void _ZN20QOpenGLShaderProgram9addShaderEP13QOpenGLShader(void* qthis, void* arg0);
  // proto:  int QOpenGLShaderProgram::attributeLocation(const QString & name);
extern void _ZNK20QOpenGLShaderProgram17attributeLocationERK7QString(void* qthis, void* arg0);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, GLenum type, const void * values, int tupleSize, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEijPKvii(void* qthis, int arg0, unsigned int arg1, void* arg2, int arg3, int arg4);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QPoint & point);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QPoint(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const QMatrix4x4 * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK10QMatrix4x4i(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, const QVector2D * values, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector2Di(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLfloat x, GLfloat y);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiff(void* qthis, int arg0, float arg1, float arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QMatrix4x4 * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK10QMatrix4x4i(void* qthis, char* arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QColor & value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK6QColor(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcffff(void* qthis, char* arg0, float arg1, float arg2, float arg3, float arg4);
  // proto:  GLuint QOpenGLShaderProgram::programId();
extern void _ZNK20QOpenGLShaderProgram9programIdEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const GLuint * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKji(void* qthis, int arg0, unsigned int* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::disableAttributeArray(int location);
extern void _ZN20QOpenGLShaderProgram21disableAttributeArrayEi(void* qthis, int arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const GLint * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKii(void* qthis, char* arg0, int* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const GLuint * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKji(void* qthis, char* arg0, unsigned int* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const GLfloat * values, int columns, int rows);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcPKfii(void* qthis, char* arg0, float* arg1, int arg2, int arg3);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const QVector4D * values, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector4Di(void* qthis, char* arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const QVector2D & value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector2D(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLfloat x, GLfloat y, GLfloat z);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEifff(void* qthis, int arg0, float arg1, float arg2, float arg3);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEiffff(void* qthis, int arg0, float arg1, float arg2, float arg3, float arg4);
  // proto:  void QOpenGLShaderProgram::bindAttributeLocation(const char * name, int location);
extern void _ZN20QOpenGLShaderProgram21bindAttributeLocationEPKci(void* qthis, char* arg0, int arg1);
  // proto:  int QOpenGLShaderProgram::attributeLocation(const char * name);
extern void _ZNK20QOpenGLShaderProgram17attributeLocationEPKc(void* qthis, char* arg0);
  // proto:  int QOpenGLShaderProgram::uniformLocation(const QString & name);
extern void _ZNK20QOpenGLShaderProgram15uniformLocationERK7QString(void* qthis, void* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QPointF & point);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK7QPointF(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QVector4D & value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector4D(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const QVector4D * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector4Di(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  int QOpenGLShaderProgram::uniformLocation(const char * name);
extern void _ZNK20QOpenGLShaderProgram15uniformLocationEPKc(void* qthis, char* arg0);
  // proto:  void QOpenGLShaderProgram::~QOpenGLShaderProgram();
extern void _ZN20QOpenGLShaderProgramD0Ev(void* qthis);
  // proto:  QVector<float> QOpenGLShaderProgram::defaultInnerTessellationLevels();
extern void _ZNK20QOpenGLShaderProgram30defaultInnerTessellationLevelsEv(void* qthis);
  // proto:  bool QOpenGLShaderProgram::link();
extern void _ZN20QOpenGLShaderProgram4linkEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QVector4D * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector4Di(void* qthis, char* arg0, void* arg1, int arg2);
  // proto:  QList<QOpenGLShader *> QOpenGLShaderProgram::shaders();
extern void _ZNK20QOpenGLShaderProgram7shadersEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, GLfloat value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcf(void* qthis, char* arg0, float arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const QVector3D & value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector3D(void* qthis, int arg0, void* arg1);
  // proto:  QVector<float> QOpenGLShaderProgram::defaultOuterTessellationLevels();
extern void _ZNK20QOpenGLShaderProgram30defaultOuterTessellationLevelsEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QColor & color);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QColor(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QVector2D * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector2Di(void* qthis, char* arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const GLfloat [4][4] value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA4_Kf(void* qthis, char* arg0, float** arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, GLfloat x, GLfloat y);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcff(void* qthis, char* arg0, float arg1, float arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const GLint * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKii(void* qthis, int arg0, int* arg1, int arg2);
  // proto:  QString QOpenGLShaderProgram::log();
extern void _ZNK20QOpenGLShaderProgram3logEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QSize & size);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK5QSize(void* qthis, char* arg0, void* arg1);
  // proto:  int QOpenGLShaderProgram::patchVertexCount();
extern void _ZNK20QOpenGLShaderProgram16patchVertexCountEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const GLfloat [2][2] value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiPA2_Kf(void* qthis, int arg0, float** arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QMatrix4x4 & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QMatrix4x4(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const QVector4D & value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector4D(void* qthis, int arg0, void* arg1);
  // proto:  bool QOpenGLShaderProgram::create();
extern void _ZN20QOpenGLShaderProgram6createEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QSizeF & size);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QSizeF(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::removeShader(QOpenGLShader * shader);
extern void _ZN20QOpenGLShaderProgram12removeShaderEP13QOpenGLShader(void* qthis, void* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QVector4D & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector4D(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLfloat x, GLfloat y);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcff(void* qthis, char* arg0, float arg1, float arg2);
  // proto:  void QOpenGLShaderProgram::setAttributeBuffer(const char * name, GLenum type, int offset, int tupleSize, int stride);
extern void _ZN20QOpenGLShaderProgram18setAttributeBufferEPKcjiii(void* qthis, char* arg0, unsigned int arg1, int arg2, int arg3, int arg4);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, const GLfloat * values, int columns, int rows);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEiPKfii(void* qthis, int arg0, float* arg1, int arg2, int arg3);
  // proto:  void QOpenGLShaderProgram::disableAttributeArray(const char * name);
extern void _ZN20QOpenGLShaderProgram21disableAttributeArrayEPKc(void* qthis, char* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QPointF & point);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK7QPointF(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const GLfloat [4][4] value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiPA4_Kf(void* qthis, int arg0, float** arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(int location, GLfloat value);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEif(void* qthis, int arg0, float arg1);
  // proto:  void QOpenGLShaderProgram::bindAttributeLocation(const QString & name, int location);
extern void _ZN20QOpenGLShaderProgram21bindAttributeLocationERK7QStringi(void* qthis, void* arg0, int arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, const QVector3D * values, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector3Di(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QColor & color);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QColor(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QTransform & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QTransform(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QVector3D & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector3D(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const QVector2D * values, int count);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector2Di(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  int QOpenGLShaderProgram::attributeLocation(const QByteArray & name);
extern void _ZNK20QOpenGLShaderProgram17attributeLocationERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLfloat x, GLfloat y, GLfloat z);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcfff(void* qthis, char* arg0, float arg1, float arg2, float arg3);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, const QVector4D * values, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector4Di(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, GLfloat x, GLfloat y, GLfloat z);
extern void _ZN20QOpenGLShaderProgram17setAttributeValueEPKcfff(void* qthis, char* arg0, float arg1, float arg2, float arg3);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const GLfloat [3][3] value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiPA3_Kf(void* qthis, int arg0, float** arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, GLenum type, const void * values, int tupleSize, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcjPKvii(void* qthis, char* arg0, unsigned int arg1, void* arg2, int arg3, int arg4);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLuint value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcj(void* qthis, char* arg0, unsigned int arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLint value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKci(void* qthis, char* arg0, int arg1);
  // proto:  void QOpenGLShaderProgram::QOpenGLShaderProgram(QObject * parent);
extern void* dector_ZN20QOpenGLShaderProgramC1EP7QObject(void* arg0);
extern void _ZN20QOpenGLShaderProgramC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, GLfloat value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcf(void* qthis, char* arg0, float arg1);
  // proto:  void QOpenGLShaderProgram::enableAttributeArray(const char * name);
extern void _ZN20QOpenGLShaderProgram20enableAttributeArrayEPKc(void* qthis, char* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QMatrix4x4 & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK10QMatrix4x4(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, const QVector4D & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector4D(void* qthis, int arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(int location, const GLfloat * values, int tupleSize, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEiPKfii(void* qthis, int arg0, float* arg1, int arg2, int arg3);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QSizeF & size);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QSizeF(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::QOpenGLShaderProgram(const QOpenGLShaderProgram & );
extern void* dector_ZN20QOpenGLShaderProgramC1ERKS_(void* arg0);
extern void _ZN20QOpenGLShaderProgramC1ERKS_(void* qthis, void* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(int location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEiffff(void* qthis, int arg0, float arg1, float arg2, float arg3, float arg4);
  // proto:  void QOpenGLShaderProgram::removeAllShaders();
extern void _ZN20QOpenGLShaderProgram16removeAllShadersEv(void* qthis);
  // proto:  int QOpenGLShaderProgram::maxGeometryOutputVertices();
extern void _ZNK20QOpenGLShaderProgram25maxGeometryOutputVerticesEv(void* qthis);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const QVector3D * values, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector3Di(void* qthis, char* arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::release();
extern void _ZN20QOpenGLShaderProgram7releaseEv(void* qthis);
  // proto:  const QMetaObject * QOpenGLShaderProgram::metaObject();
extern void _ZNK20QOpenGLShaderProgram10metaObjectEv(void* qthis);
  // proto:  int QOpenGLShaderProgram::uniformLocation(const QByteArray & name);
extern void _ZNK20QOpenGLShaderProgram15uniformLocationERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QVector2D & value);
extern void _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector2D(void* qthis, char* arg0, void* arg1);
  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const QVector2D * values, int stride);
extern void _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector2Di(void* qthis, char* arg0, void* arg1, int arg2);
  // proto:  void QOpenGLShaderProgram::setUniformValueArray(int location, const GLfloat * values, int count, int tupleSize);
extern void _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKfii(void* qthis, int arg0, float* arg1, int arg2, int arg3);
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

// class sizeof(QOpenGLShader)=1
type QOpenGLShader struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QOpenGLShaderProgram)=1
type QOpenGLShaderProgram struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QOpenGLShader::QOpenGLShader(const QOpenGLShader & );
func NewQOpenGLShader(args ...interface{}) QOpenGLShader {
  return QOpenGLShader{}
}

  // proto:  bool QOpenGLShader::isCompiled();
func (this *QOpenGLShader) isCompiled(args ...interface{}) () {
  // isCompiled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10isCompiledEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "isCompiled", args)
  }

}

  // proto:  const QMetaObject * QOpenGLShader::metaObject();
func (this *QOpenGLShader) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "metaObject", args)
  }

}

  // proto:  QString QOpenGLShader::log();
func (this *QOpenGLShader) log(args ...interface{}) () {
  // log()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader3logEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "log", args)
  }

}

  // proto:  bool QOpenGLShader::compileSourceCode(const QString & source);
func (this *QOpenGLShader) compileSourceCode(args ...interface{}) () {
  // compileSourceCode(const class QString &)
  // compileSourceCode(const class QByteArray &)
  // compileSourceCode(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeERK10QByteArray
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  case 2:
    // invoke: _ZN13QOpenGLShader17compileSourceCodeEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShader", "compileSourceCode", args)
  }

}

  // proto:  bool QOpenGLShader::compileSourceFile(const QString & fileName);
func (this *QOpenGLShader) compileSourceFile(args ...interface{}) () {
  // compileSourceFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLShader17compileSourceFileERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShader", "compileSourceFile", args)
  }

}

  // proto:  QByteArray QOpenGLShader::sourceCode();
func (this *QOpenGLShader) sourceCode(args ...interface{}) () {
  // sourceCode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader10sourceCodeEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "sourceCode", args)
  }

}

  // proto:  void QOpenGLShader::~QOpenGLShader();
func (this *QOpenGLShader) FreeQOpenGLShader(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLShader", "~QOpenGLShader", args)
  }

}

  // proto:  GLuint QOpenGLShader::shaderId();
func (this *QOpenGLShader) shaderId(args ...interface{}) () {
  // shaderId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLShader8shaderIdEv
  default:
    qtrt.ErrorResolve("QOpenGLShader", "shaderId", args)
  }

}

  // proto:  bool QOpenGLShaderProgram::isLinked();
func (this *QOpenGLShaderProgram) isLinked(args ...interface{}) () {
  // isLinked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram8isLinkedEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "isLinked", args)
  }

}

  // proto:  void QOpenGLShaderProgram::setUniformValue(const char * name, const QVector3D & value);
func (this *QOpenGLShaderProgram) setUniformValue(args ...interface{}) () {
  // setUniformValue(const char *, const class QVector3D &)
  // setUniformValue(int, const class QPoint &)
  // setUniformValue(const char *, GLfloat, GLfloat, GLfloat, GLfloat)
  // setUniformValue(int, const class QTransform &)
  // setUniformValue(int, const QMatrix2x3 &)
  // setUniformValue(const char *, const GLfloat (*)[2])
  // setUniformValue(const char *, const QMatrix3x3 &)
  // setUniformValue(int, const QMatrix3x2 &)
  // setUniformValue(const char *, const QMatrix4x3 &)
  // setUniformValue(const char *, const GLfloat (*)[3])
  // setUniformValue(int, const class QVector2D &)
  // setUniformValue(int, GLfloat)
  // setUniformValue(const char *, const QMatrix2x4 &)
  // setUniformValue(int, GLint)
  // setUniformValue(int, GLuint)
  // setUniformValue(int, const class QSize &)
  // setUniformValue(int, const QMatrix2x2 &)
  // setUniformValue(const char *, const class QPoint &)
  // setUniformValue(int, GLfloat, GLfloat)
  // setUniformValue(int, GLfloat, GLfloat, GLfloat)
  // setUniformValue(const char *, const class QPointF &)
  // setUniformValue(const char *, const QMatrix3x2 &)
  // setUniformValue(const char *, const QMatrix2x3 &)
  // setUniformValue(int, const class QColor &)
  // setUniformValue(const char *, const GLfloat (*)[4])
  // setUniformValue(const char *, const QMatrix2x2 &)
  // setUniformValue(const char *, const class QSize &)
  // setUniformValue(const char *, const QMatrix4x2 &)
  // setUniformValue(int, const GLfloat (*)[2])
  // setUniformValue(const char *, const class QMatrix4x4 &)
  // setUniformValue(int, const class QSizeF &)
  // setUniformValue(const char *, const QMatrix3x4 &)
  // setUniformValue(const char *, const class QVector4D &)
  // setUniformValue(const char *, GLfloat, GLfloat)
  // setUniformValue(int, const QMatrix3x4 &)
  // setUniformValue(int, const class QPointF &)
  // setUniformValue(int, const QMatrix4x2 &)
  // setUniformValue(int, const GLfloat (*)[4])
  // setUniformValue(int, const QMatrix3x3 &)
  // setUniformValue(const char *, const class QColor &)
  // setUniformValue(const char *, const class QTransform &)
  // setUniformValue(int, const class QVector3D &)
  // setUniformValue(const char *, GLfloat, GLfloat, GLfloat)
  // setUniformValue(int, const GLfloat (*)[3])
  // setUniformValue(const char *, GLuint)
  // setUniformValue(const char *, GLint)
  // setUniformValue(int, const QMatrix4x3 &)
  // setUniformValue(const char *, GLfloat)
  // setUniformValue(int, const class QMatrix4x4 &)
  // setUniformValue(int, const class QVector4D &)
  // setUniformValue(const char *, const class QSizeF &)
  // setUniformValue(int, GLfloat, GLfloat, GLfloat, GLfloat)
  // setUniformValue(int, const QMatrix2x4 &)
  // setUniformValue(const char *, const class QVector2D &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[2][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[2][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[2][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  // vtys[4][1] = reflect.TypeOf(QGenericMatrix<2, 3, float>{}) // "const QMatrix2x3 &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"
  vtys[5][1] = qtrt.VoidpTy() // "const GLfloat [2][2]"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[6][1] = reflect.TypeOf(QGenericMatrix<3, 3, float>{}) // "const QMatrix3x3 &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  // vtys[7][1] = reflect.TypeOf(QGenericMatrix<3, 2, float>{}) // "const QMatrix3x2 &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[8][1] = reflect.TypeOf(QGenericMatrix<4, 3, float>{}) // "const QMatrix4x3 &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = qtrt.VoidpTy() // "const GLfloat [3][3]"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int32Ty(false) // "int"
  vtys[10][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "int"
  vtys[11][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[12][1] = reflect.TypeOf(QGenericMatrix<2, 4, float>{}) // "const QMatrix2x4 &"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.Int32Ty(false) // "int"
  vtys[13][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.Int32Ty(false) // "int"
  vtys[14][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.Int32Ty(false) // "int"
  vtys[15][1] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  // vtys[16][1] = reflect.TypeOf(QGenericMatrix<2, 2, float>{}) // "const QMatrix2x2 &"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.ByteTy(true) // "const char *"
  vtys[17][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = qtrt.Int32Ty(false) // "int"
  vtys[18][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[18][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = qtrt.Int32Ty(false) // "int"
  vtys[19][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[19][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[19][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[20] = make(map[int32]reflect.Type)
  vtys[20][0] = qtrt.ByteTy(true) // "const char *"
  vtys[20][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[21] = make(map[int32]reflect.Type)
  vtys[21][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[21][1] = reflect.TypeOf(QGenericMatrix<3, 2, float>{}) // "const QMatrix3x2 &"
  vtys[22] = make(map[int32]reflect.Type)
  vtys[22][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[22][1] = reflect.TypeOf(QGenericMatrix<2, 3, float>{}) // "const QMatrix2x3 &"
  vtys[23] = make(map[int32]reflect.Type)
  vtys[23][0] = qtrt.Int32Ty(false) // "int"
  vtys[23][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[24] = make(map[int32]reflect.Type)
  vtys[24][0] = qtrt.ByteTy(true) // "const char *"
  vtys[24][1] = qtrt.VoidpTy() // "const GLfloat [4][4]"
  vtys[25] = make(map[int32]reflect.Type)
  vtys[25][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[25][1] = reflect.TypeOf(QGenericMatrix<2, 2, float>{}) // "const QMatrix2x2 &"
  vtys[26] = make(map[int32]reflect.Type)
  vtys[26][0] = qtrt.ByteTy(true) // "const char *"
  vtys[26][1] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[27] = make(map[int32]reflect.Type)
  vtys[27][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[27][1] = reflect.TypeOf(QGenericMatrix<4, 2, float>{}) // "const QMatrix4x2 &"
  vtys[28] = make(map[int32]reflect.Type)
  vtys[28][0] = qtrt.Int32Ty(false) // "int"
  vtys[28][1] = qtrt.VoidpTy() // "const GLfloat [2][2]"
  vtys[29] = make(map[int32]reflect.Type)
  vtys[29][0] = qtrt.ByteTy(true) // "const char *"
  vtys[29][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[30] = make(map[int32]reflect.Type)
  vtys[30][0] = qtrt.Int32Ty(false) // "int"
  vtys[30][1] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[31] = make(map[int32]reflect.Type)
  vtys[31][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[31][1] = reflect.TypeOf(QGenericMatrix<3, 4, float>{}) // "const QMatrix3x4 &"
  vtys[32] = make(map[int32]reflect.Type)
  vtys[32][0] = qtrt.ByteTy(true) // "const char *"
  vtys[32][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[33] = make(map[int32]reflect.Type)
  vtys[33][0] = qtrt.ByteTy(true) // "const char *"
  vtys[33][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[33][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[34] = make(map[int32]reflect.Type)
  vtys[34][0] = qtrt.Int32Ty(false) // "int"
  // vtys[34][1] = reflect.TypeOf(QGenericMatrix<3, 4, float>{}) // "const QMatrix3x4 &"
  vtys[35] = make(map[int32]reflect.Type)
  vtys[35][0] = qtrt.Int32Ty(false) // "int"
  vtys[35][1] = reflect.TypeOf(QPointF{}) // "const QPointF &"
  vtys[36] = make(map[int32]reflect.Type)
  vtys[36][0] = qtrt.Int32Ty(false) // "int"
  // vtys[36][1] = reflect.TypeOf(QGenericMatrix<4, 2, float>{}) // "const QMatrix4x2 &"
  vtys[37] = make(map[int32]reflect.Type)
  vtys[37][0] = qtrt.Int32Ty(false) // "int"
  vtys[37][1] = qtrt.VoidpTy() // "const GLfloat [4][4]"
  vtys[38] = make(map[int32]reflect.Type)
  vtys[38][0] = qtrt.Int32Ty(false) // "int"
  // vtys[38][1] = reflect.TypeOf(QGenericMatrix<3, 3, float>{}) // "const QMatrix3x3 &"
  vtys[39] = make(map[int32]reflect.Type)
  vtys[39][0] = qtrt.ByteTy(true) // "const char *"
  vtys[39][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[40] = make(map[int32]reflect.Type)
  vtys[40][0] = qtrt.ByteTy(true) // "const char *"
  vtys[40][1] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[41] = make(map[int32]reflect.Type)
  vtys[41][0] = qtrt.Int32Ty(false) // "int"
  vtys[41][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[42] = make(map[int32]reflect.Type)
  vtys[42][0] = qtrt.ByteTy(true) // "const char *"
  vtys[42][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[42][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[42][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[43] = make(map[int32]reflect.Type)
  vtys[43][0] = qtrt.Int32Ty(false) // "int"
  vtys[43][1] = qtrt.VoidpTy() // "const GLfloat [3][3]"
  vtys[44] = make(map[int32]reflect.Type)
  vtys[44][0] = qtrt.ByteTy(true) // "const char *"
  vtys[44][1] = qtrt.Int32Ty(false) // "GLuint"
  vtys[45] = make(map[int32]reflect.Type)
  vtys[45][0] = qtrt.ByteTy(true) // "const char *"
  vtys[45][1] = qtrt.Int32Ty(false) // "GLint"
  vtys[46] = make(map[int32]reflect.Type)
  vtys[46][0] = qtrt.Int32Ty(false) // "int"
  // vtys[46][1] = reflect.TypeOf(QGenericMatrix<4, 3, float>{}) // "const QMatrix4x3 &"
  vtys[47] = make(map[int32]reflect.Type)
  vtys[47][0] = qtrt.ByteTy(true) // "const char *"
  vtys[47][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[48] = make(map[int32]reflect.Type)
  vtys[48][0] = qtrt.Int32Ty(false) // "int"
  vtys[48][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 &"
  vtys[49] = make(map[int32]reflect.Type)
  vtys[49][0] = qtrt.Int32Ty(false) // "int"
  vtys[49][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[50] = make(map[int32]reflect.Type)
  vtys[50][0] = qtrt.ByteTy(true) // "const char *"
  vtys[50][1] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"
  vtys[51] = make(map[int32]reflect.Type)
  vtys[51][0] = qtrt.Int32Ty(false) // "int"
  vtys[51][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[51][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[51][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[51][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[52] = make(map[int32]reflect.Type)
  vtys[52][0] = qtrt.Int32Ty(false) // "int"
  // vtys[52][1] = reflect.TypeOf(QGenericMatrix<2, 4, float>{}) // "const QMatrix2x4 &"
  vtys[53] = make(map[int32]reflect.Type)
  vtys[53][0] = qtrt.ByteTy(true) // "const char *"
  vtys[53][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector3D
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QPoint
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcffff
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK10QTransform
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTransform).qclsinst
    if false {fmt.Println(arg1)}
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi2ELi3EfE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 3, float>).qclsinst
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA2_Kf
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi3ELi3EfE
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 3, float>).qclsinst
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi3ELi2EfE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 2, float>).qclsinst
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi4ELi3EfE
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<4, 3, float>).qclsinst
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA3_Kf
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector2D
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEif
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi2ELi4EfE
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 4, float>).qclsinst
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 14:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEij
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 15:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK5QSize
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSize).qclsinst
    if false {fmt.Println(arg1)}
  case 16:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi2ELi2EfE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 2, float>).qclsinst
  case 17:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QPoint
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
  case 18:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiff
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
  case 19:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEifff
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
  case 20:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK7QPointF
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
  case 21:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi3ELi2EfE
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 2, float>).qclsinst
  case 22:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi2ELi3EfE
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 3, float>).qclsinst
  case 23:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QColor
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
  case 24:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcPA4_Kf
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 25:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi2ELi2EfE
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 2, float>).qclsinst
  case 26:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK5QSize
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSize).qclsinst
    if false {fmt.Println(arg1)}
  case 27:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi4ELi2EfE
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<4, 2, float>).qclsinst
  case 28:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA2_Kf
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 29:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QMatrix4x4
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg1)}
  case 30:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK6QSizeF
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSizeF).qclsinst
    if false {fmt.Println(arg1)}
  case 31:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK14QGenericMatrixILi3ELi4EfE
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 4, float>).qclsinst
  case 32:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector4D
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
  case 33:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcff
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
  case 34:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi3ELi4EfE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 4, float>).qclsinst
  case 35:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK7QPointF
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPointF).qclsinst
    if false {fmt.Println(arg1)}
  case 36:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi4ELi2EfE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<4, 2, float>).qclsinst
  case 37:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA4_Kf
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 38:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi3ELi3EfE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 3, float>).qclsinst
  case 39:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QColor
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
  case 40:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK10QTransform
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTransform).qclsinst
    if false {fmt.Println(arg1)}
  case 41:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector3D
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
  case 42:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcfff
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
  case 43:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiPA3_Kf
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 44:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcj
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 45:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKci
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 46:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi4ELi3EfE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<4, 3, float>).qclsinst
  case 47:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcf
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
  case 48:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK10QMatrix4x4
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg1)}
  case 49:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK9QVector4D
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
  case 50:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK6QSizeF
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSizeF).qclsinst
    if false {fmt.Println(arg1)}
  case 51:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiffff
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
  case 52:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEiRK14QGenericMatrixILi2ELi4EfE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 4, float>).qclsinst
  case 53:
    // invoke: _ZN20QOpenGLShaderProgram15setUniformValueEPKcRK9QVector2D
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setUniformValue", args)
  }

}

  // proto:  void QOpenGLShaderProgram::setAttributeValue(const char * name, const QVector3D & value);
func (this *QOpenGLShaderProgram) setAttributeValue(args ...interface{}) () {
  // setAttributeValue(const char *, const class QVector3D &)
  // setAttributeValue(int, GLfloat, GLfloat)
  // setAttributeValue(const char *, const class QVector2D &)
  // setAttributeValue(int, const class QColor &)
  // setAttributeValue(int, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(const char *, const class QColor &)
  // setAttributeValue(const char *, GLfloat, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(const char *, const GLfloat *, int, int)
  // setAttributeValue(int, const class QVector2D &)
  // setAttributeValue(int, GLfloat, GLfloat, GLfloat, GLfloat)
  // setAttributeValue(const char *, const class QVector4D &)
  // setAttributeValue(const char *, GLfloat)
  // setAttributeValue(int, const class QVector3D &)
  // setAttributeValue(const char *, GLfloat, GLfloat)
  // setAttributeValue(int, const class QVector4D &)
  // setAttributeValue(int, const GLfloat *, int, int)
  // setAttributeValue(int, GLfloat)
  // setAttributeValue(const char *, GLfloat, GLfloat, GLfloat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[1][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[4][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[4][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"
  vtys[5][1] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[6][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.ByteTy(true) // "const char *"
  vtys[7][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int32Ty(false) // "int"
  vtys[8][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.Int32Ty(false) // "int"
  vtys[9][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[9][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[9][3] = qtrt.FloatTy(false) // "GLfloat"
  vtys[9][4] = qtrt.FloatTy(false) // "GLfloat"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(true) // "const char *"
  vtys[10][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.ByteTy(true) // "const char *"
  vtys[11][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.Int32Ty(false) // "int"
  vtys[12][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D &"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.ByteTy(true) // "const char *"
  vtys[13][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[13][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.Int32Ty(false) // "int"
  vtys[14][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D &"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.Int32Ty(false) // "int"
  vtys[15][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[15][2] = qtrt.Int32Ty(false) // "int"
  vtys[15][3] = qtrt.Int32Ty(false) // "int"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  vtys[16][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.ByteTy(true) // "const char *"
  vtys[17][1] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17][2] = qtrt.FloatTy(false) // "GLfloat"
  vtys[17][3] = qtrt.FloatTy(false) // "GLfloat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector3D
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiff
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector2D
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK6QColor
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEifff
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK6QColor
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcffff
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcPKfii
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector2D
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiffff
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
    var arg4 = C.float(args[4].(float32))
    if false {fmt.Println(arg4)}
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcRK9QVector4D
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcf
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector3D
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcff
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
  case 14:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiRK9QVector4D
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
  case 15:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEiPKfii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  case 16:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEif
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
  case 17:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeValueEPKcfff
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.float(args[1].(float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.float(args[2].(float32))
    if false {fmt.Println(arg2)}
    var arg3 = C.float(args[3].(float32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeValue", args)
  }

}

  // proto:  void QOpenGLShaderProgram::setUniformValueArray(const char * name, const QVector3D * values, int count);
func (this *QOpenGLShaderProgram) setUniformValueArray(args ...interface{}) () {
  // setUniformValueArray(const char *, const class QVector3D *, int)
  // setUniformValueArray(const char *, const QMatrix3x3 *, int)
  // setUniformValueArray(int, const QMatrix2x3 *, int)
  // setUniformValueArray(int, const class QVector3D *, int)
  // setUniformValueArray(const char *, const GLfloat *, int, int)
  // setUniformValueArray(int, const QMatrix2x4 *, int)
  // setUniformValueArray(const char *, const QMatrix3x4 *, int)
  // setUniformValueArray(int, const class QMatrix4x4 *, int)
  // setUniformValueArray(int, const QMatrix3x4 *, int)
  // setUniformValueArray(const char *, const class QMatrix4x4 *, int)
  // setUniformValueArray(const char *, const QMatrix2x3 *, int)
  // setUniformValueArray(int, const GLuint *, int)
  // setUniformValueArray(const char *, const QMatrix4x2 *, int)
  // setUniformValueArray(const char *, const GLint *, int)
  // setUniformValueArray(const char *, const GLuint *, int)
  // setUniformValueArray(int, const QMatrix3x2 *, int)
  // setUniformValueArray(int, const class QVector4D *, int)
  // setUniformValueArray(const char *, const class QVector4D *, int)
  // setUniformValueArray(const char *, const QMatrix2x4 *, int)
  // setUniformValueArray(int, const QMatrix4x2 *, int)
  // setUniformValueArray(const char *, const class QVector2D *, int)
  // setUniformValueArray(int, const GLint *, int)
  // setUniformValueArray(const char *, const QMatrix4x3 *, int)
  // setUniformValueArray(const char *, const QMatrix3x2 *, int)
  // setUniformValueArray(const char *, const QMatrix2x2 *, int)
  // setUniformValueArray(int, const QMatrix2x2 *, int)
  // setUniformValueArray(int, const class QVector2D *, int)
  // setUniformValueArray(int, const QMatrix4x3 *, int)
  // setUniformValueArray(int, const QMatrix3x3 *, int)
  // setUniformValueArray(int, const GLfloat *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[1][1] = reflect.TypeOf(QGenericMatrix<3, 3, float>{}) // "const QMatrix3x3 *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  // vtys[2][1] = reflect.TypeOf(QGenericMatrix<2, 3, float>{}) // "const QMatrix2x3 *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(true) // "const char *"
  vtys[4][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[4][3] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  // vtys[5][1] = reflect.TypeOf(QGenericMatrix<2, 4, float>{}) // "const QMatrix2x4 *"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[6][1] = reflect.TypeOf(QGenericMatrix<3, 4, float>{}) // "const QMatrix3x4 *"
  vtys[6][2] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int32Ty(false) // "int"
  // vtys[8][1] = reflect.TypeOf(QGenericMatrix<3, 4, float>{}) // "const QMatrix3x4 *"
  vtys[8][2] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = reflect.TypeOf(QMatrix4x4{}) // "const QMatrix4x4 *"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[10][1] = reflect.TypeOf(QGenericMatrix<2, 3, float>{}) // "const QMatrix2x3 *"
  vtys[10][2] = qtrt.Int32Ty(false) // "int"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "int"
  vtys[11][1] = qtrt.Int32Ty(true) // "const GLuint *"
  vtys[11][2] = qtrt.Int32Ty(false) // "int"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[12][1] = reflect.TypeOf(QGenericMatrix<4, 2, float>{}) // "const QMatrix4x2 *"
  vtys[12][2] = qtrt.Int32Ty(false) // "int"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.ByteTy(true) // "const char *"
  vtys[13][1] = qtrt.Int32Ty(true) // "const GLint *"
  vtys[13][2] = qtrt.Int32Ty(false) // "int"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.ByteTy(true) // "const char *"
  vtys[14][1] = qtrt.Int32Ty(true) // "const GLuint *"
  vtys[14][2] = qtrt.Int32Ty(false) // "int"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.Int32Ty(false) // "int"
  // vtys[15][1] = reflect.TypeOf(QGenericMatrix<3, 2, float>{}) // "const QMatrix3x2 *"
  vtys[15][2] = qtrt.Int32Ty(false) // "int"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.Int32Ty(false) // "int"
  vtys[16][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[16][2] = qtrt.Int32Ty(false) // "int"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.ByteTy(true) // "const char *"
  vtys[17][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[17][2] = qtrt.Int32Ty(false) // "int"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[18][1] = reflect.TypeOf(QGenericMatrix<2, 4, float>{}) // "const QMatrix2x4 *"
  vtys[18][2] = qtrt.Int32Ty(false) // "int"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = qtrt.Int32Ty(false) // "int"
  // vtys[19][1] = reflect.TypeOf(QGenericMatrix<4, 2, float>{}) // "const QMatrix4x2 *"
  vtys[19][2] = qtrt.Int32Ty(false) // "int"
  vtys[20] = make(map[int32]reflect.Type)
  vtys[20][0] = qtrt.ByteTy(true) // "const char *"
  vtys[20][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[20][2] = qtrt.Int32Ty(false) // "int"
  vtys[21] = make(map[int32]reflect.Type)
  vtys[21][0] = qtrt.Int32Ty(false) // "int"
  vtys[21][1] = qtrt.Int32Ty(true) // "const GLint *"
  vtys[21][2] = qtrt.Int32Ty(false) // "int"
  vtys[22] = make(map[int32]reflect.Type)
  vtys[22][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[22][1] = reflect.TypeOf(QGenericMatrix<4, 3, float>{}) // "const QMatrix4x3 *"
  vtys[22][2] = qtrt.Int32Ty(false) // "int"
  vtys[23] = make(map[int32]reflect.Type)
  vtys[23][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[23][1] = reflect.TypeOf(QGenericMatrix<3, 2, float>{}) // "const QMatrix3x2 *"
  vtys[23][2] = qtrt.Int32Ty(false) // "int"
  vtys[24] = make(map[int32]reflect.Type)
  vtys[24][0] = qtrt.ByteTy(true) // "const char *"
  // vtys[24][1] = reflect.TypeOf(QGenericMatrix<2, 2, float>{}) // "const QMatrix2x2 *"
  vtys[24][2] = qtrt.Int32Ty(false) // "int"
  vtys[25] = make(map[int32]reflect.Type)
  vtys[25][0] = qtrt.Int32Ty(false) // "int"
  // vtys[25][1] = reflect.TypeOf(QGenericMatrix<2, 2, float>{}) // "const QMatrix2x2 *"
  vtys[25][2] = qtrt.Int32Ty(false) // "int"
  vtys[26] = make(map[int32]reflect.Type)
  vtys[26][0] = qtrt.Int32Ty(false) // "int"
  vtys[26][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[26][2] = qtrt.Int32Ty(false) // "int"
  vtys[27] = make(map[int32]reflect.Type)
  vtys[27][0] = qtrt.Int32Ty(false) // "int"
  // vtys[27][1] = reflect.TypeOf(QGenericMatrix<4, 3, float>{}) // "const QMatrix4x3 *"
  vtys[27][2] = qtrt.Int32Ty(false) // "int"
  vtys[28] = make(map[int32]reflect.Type)
  vtys[28][0] = qtrt.Int32Ty(false) // "int"
  // vtys[28][1] = reflect.TypeOf(QGenericMatrix<3, 3, float>{}) // "const QMatrix3x3 *"
  vtys[28][2] = qtrt.Int32Ty(false) // "int"
  vtys[29] = make(map[int32]reflect.Type)
  vtys[29][0] = qtrt.Int32Ty(false) // "int"
  vtys[29][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[29][2] = qtrt.Int32Ty(false) // "int"
  vtys[29][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector3Di
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi3ELi3EfEi
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 3, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi2ELi3EfEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 3, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector3Di
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKfii
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi2ELi4EfEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 4, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi3ELi4EfEi
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 4, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK10QMatrix4x4i
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi3ELi4EfEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 4, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK10QMatrix4x4i
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMatrix4x4).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 10:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi2ELi3EfEi
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 3, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 11:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKji
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 12:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi4ELi2EfEi
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<4, 2, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 13:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKii
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 14:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPKji
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 15:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi3ELi2EfEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 2, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 16:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector4Di
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 17:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector4Di
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 18:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi2ELi4EfEi
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 4, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 19:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi4ELi2EfEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<4, 2, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 20:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK9QVector2Di
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 21:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 22:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi4ELi3EfEi
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<4, 3, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 23:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi3ELi2EfEi
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 2, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 24:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEPKcPK14QGenericMatrixILi2ELi2EfEi
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 2, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 25:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi2ELi2EfEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<2, 2, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 26:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK9QVector2Di
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 27:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi4ELi3EfEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<4, 3, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 28:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPK14QGenericMatrixILi3ELi3EfEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    // var arg1 = args[1].(QGenericMatrix<3, 3, float>).qclsinst
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 29:
    // invoke: _ZN20QOpenGLShaderProgram20setUniformValueArrayEiPKfii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setUniformValueArray", args)
  }

}

  // proto:  void QOpenGLShaderProgram::setAttributeBuffer(int location, GLenum type, int offset, int tupleSize, int stride);
func (this *QOpenGLShaderProgram) setAttributeBuffer(args ...interface{}) () {
  // setAttributeBuffer(int, GLenum, int, int, int)
  // setAttributeBuffer(const char *, GLenum, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram18setAttributeBufferEijiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram18setAttributeBufferEPKcjiii
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeBuffer", args)
  }

}

  // proto: static bool QOpenGLShaderProgram::hasOpenGLShaderPrograms(QOpenGLContext * context);
func (this *QOpenGLShaderProgram) hasOpenGLShaderPrograms_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "hasOpenGLShaderPrograms", args)
  }

}

  // proto:  void QOpenGLShaderProgram::setPatchVertexCount(int count);
func (this *QOpenGLShaderProgram) setPatchVertexCount(args ...interface{}) () {
  // setPatchVertexCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram19setPatchVertexCountEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setPatchVertexCount", args)
  }

}

  // proto:  void QOpenGLShaderProgram::setAttributeArray(const char * name, const GLfloat * values, int tupleSize, int stride);
func (this *QOpenGLShaderProgram) setAttributeArray(args ...interface{}) () {
  // setAttributeArray(const char *, const GLfloat *, int, int)
  // setAttributeArray(int, GLenum, const void *, int, int)
  // setAttributeArray(int, const class QVector2D *, int)
  // setAttributeArray(const char *, const class QVector4D *, int)
  // setAttributeArray(int, const class QVector3D *, int)
  // setAttributeArray(int, const class QVector4D *, int)
  // setAttributeArray(const char *, GLenum, const void *, int, int)
  // setAttributeArray(int, const GLfloat *, int, int)
  // setAttributeArray(const char *, const class QVector3D *, int)
  // setAttributeArray(const char *, const class QVector2D *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[1][2] = qtrt.VoidpTy() // "const void *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = reflect.TypeOf(QVector4D{}) // "const QVector4D *"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6][1] = qtrt.Int32Ty(false) // "GLenum"
  vtys[6][2] = qtrt.VoidpTy() // "const void *"
  vtys[6][3] = qtrt.Int32Ty(false) // "int"
  vtys[6][4] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.FloatTy(true) // "const GLfloat *"
  vtys[7][2] = qtrt.Int32Ty(false) // "int"
  vtys[7][3] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.ByteTy(true) // "const char *"
  vtys[8][1] = reflect.TypeOf(QVector3D{}) // "const QVector3D *"
  vtys[8][2] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = reflect.TypeOf(QVector2D{}) // "const QVector2D *"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPKfii
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEijPKvii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector2Di
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 3:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector4Di
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 4:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector3Di
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 5:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPK9QVector4Di
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector4D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 6:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcjPKvii
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(unsafe.Pointer)
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
  case 7:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEiPKfii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.float)(args[1].(*float32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  case 8:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector3Di
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector3D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  case 9:
    // invoke: _ZN20QOpenGLShaderProgram17setAttributeArrayEPKcPK9QVector2Di
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVector2D).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "setAttributeArray", args)
  }

}

  // proto:  void QOpenGLShaderProgram::bindAttributeLocation(const QByteArray & name, int location);
func (this *QOpenGLShaderProgram) bindAttributeLocation(args ...interface{}) () {
  // bindAttributeLocation(const class QByteArray &, int)
  // bindAttributeLocation(const char *, int)
  // bindAttributeLocation(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationERK10QByteArrayi
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationEPKci
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 2:
    // invoke: _ZN20QOpenGLShaderProgram21bindAttributeLocationERK7QStringi
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "bindAttributeLocation", args)
  }

}

  // proto:  bool QOpenGLShaderProgram::bind();
func (this *QOpenGLShaderProgram) bind(args ...interface{}) () {
  // bind()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram4bindEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "bind", args)
  }

}

  // proto:  void QOpenGLShaderProgram::enableAttributeArray(int location);
func (this *QOpenGLShaderProgram) enableAttributeArray(args ...interface{}) () {
  // enableAttributeArray(int)
  // enableAttributeArray(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram20enableAttributeArrayEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram20enableAttributeArrayEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "enableAttributeArray", args)
  }

}

  // proto:  bool QOpenGLShaderProgram::addShader(QOpenGLShader * shader);
func (this *QOpenGLShaderProgram) addShader(args ...interface{}) () {
  // addShader(class QOpenGLShader *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLShader{}) // "QOpenGLShader *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram9addShaderEP13QOpenGLShader
    var arg0 = args[0].(QOpenGLShader).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "addShader", args)
  }

}

  // proto:  int QOpenGLShaderProgram::attributeLocation(const QString & name);
func (this *QOpenGLShaderProgram) attributeLocation(args ...interface{}) () {
  // attributeLocation(const class QString &)
  // attributeLocation(const char *)
  // attributeLocation(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  case 2:
    // invoke: _ZNK20QOpenGLShaderProgram17attributeLocationERK10QByteArray
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "attributeLocation", args)
  }

}

  // proto:  GLuint QOpenGLShaderProgram::programId();
func (this *QOpenGLShaderProgram) programId(args ...interface{}) () {
  // programId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram9programIdEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "programId", args)
  }

}

  // proto:  void QOpenGLShaderProgram::disableAttributeArray(int location);
func (this *QOpenGLShaderProgram) disableAttributeArray(args ...interface{}) () {
  // disableAttributeArray(int)
  // disableAttributeArray(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram21disableAttributeArrayEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN20QOpenGLShaderProgram21disableAttributeArrayEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "disableAttributeArray", args)
  }

}

  // proto:  int QOpenGLShaderProgram::uniformLocation(const QString & name);
func (this *QOpenGLShaderProgram) uniformLocation(args ...interface{}) () {
  // uniformLocation(const class QString &)
  // uniformLocation(const char *)
  // uniformLocation(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationEPKc
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
  case 2:
    // invoke: _ZNK20QOpenGLShaderProgram15uniformLocationERK10QByteArray
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "uniformLocation", args)
  }

}

  // proto:  void QOpenGLShaderProgram::~QOpenGLShaderProgram();
func (this *QOpenGLShaderProgram) FreeQOpenGLShaderProgram(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "~QOpenGLShaderProgram", args)
  }

}

  // proto:  QVector<float> QOpenGLShaderProgram::defaultInnerTessellationLevels();
func (this *QOpenGLShaderProgram) defaultInnerTessellationLevels(args ...interface{}) () {
  // defaultInnerTessellationLevels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram30defaultInnerTessellationLevelsEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "defaultInnerTessellationLevels", args)
  }

}

  // proto:  bool QOpenGLShaderProgram::link();
func (this *QOpenGLShaderProgram) link(args ...interface{}) () {
  // link()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram4linkEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "link", args)
  }

}

  // proto:  QList<QOpenGLShader *> QOpenGLShaderProgram::shaders();
func (this *QOpenGLShaderProgram) shaders(args ...interface{}) () {
  // shaders()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram7shadersEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "shaders", args)
  }

}

  // proto:  QVector<float> QOpenGLShaderProgram::defaultOuterTessellationLevels();
func (this *QOpenGLShaderProgram) defaultOuterTessellationLevels(args ...interface{}) () {
  // defaultOuterTessellationLevels()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram30defaultOuterTessellationLevelsEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "defaultOuterTessellationLevels", args)
  }

}

  // proto:  QString QOpenGLShaderProgram::log();
func (this *QOpenGLShaderProgram) log(args ...interface{}) () {
  // log()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram3logEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "log", args)
  }

}

  // proto:  int QOpenGLShaderProgram::patchVertexCount();
func (this *QOpenGLShaderProgram) patchVertexCount(args ...interface{}) () {
  // patchVertexCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram16patchVertexCountEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "patchVertexCount", args)
  }

}

  // proto:  bool QOpenGLShaderProgram::create();
func (this *QOpenGLShaderProgram) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram6createEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "create", args)
  }

}

  // proto:  void QOpenGLShaderProgram::removeShader(QOpenGLShader * shader);
func (this *QOpenGLShaderProgram) removeShader(args ...interface{}) () {
  // removeShader(class QOpenGLShader *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QOpenGLShader{}) // "QOpenGLShader *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram12removeShaderEP13QOpenGLShader
    var arg0 = args[0].(QOpenGLShader).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "removeShader", args)
  }

}

  // proto:  void QOpenGLShaderProgram::QOpenGLShaderProgram(QObject * parent);
func NewQOpenGLShaderProgram(args ...interface{}) QOpenGLShaderProgram {
  return QOpenGLShaderProgram{}
}

  // proto:  void QOpenGLShaderProgram::removeAllShaders();
func (this *QOpenGLShaderProgram) removeAllShaders(args ...interface{}) () {
  // removeAllShaders()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram16removeAllShadersEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "removeAllShaders", args)
  }

}

  // proto:  int QOpenGLShaderProgram::maxGeometryOutputVertices();
func (this *QOpenGLShaderProgram) maxGeometryOutputVertices(args ...interface{}) () {
  // maxGeometryOutputVertices()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram25maxGeometryOutputVerticesEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "maxGeometryOutputVertices", args)
  }

}

  // proto:  void QOpenGLShaderProgram::release();
func (this *QOpenGLShaderProgram) release(args ...interface{}) () {
  // release()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QOpenGLShaderProgram7releaseEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "release", args)
  }

}

  // proto:  const QMetaObject * QOpenGLShaderProgram::metaObject();
func (this *QOpenGLShaderProgram) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QOpenGLShaderProgram10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLShaderProgram", "metaObject", args)
  }

}

// <= body block end

