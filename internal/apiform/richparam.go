package apiform

import (
	"mime/multipart"
	"reflect"

	"github.com/Miuzarte/openai-go/packages/param"
)

func (e *encoder) newRichFieldTypeEncoder(t reflect.Type) encoderFunc {
	f, _ := t.FieldByName("Value")
	enc := e.newPrimitiveTypeEncoder(f.Type)
	return func(key string, value reflect.Value, writer *multipart.Writer) error {
		if opt, ok := value.Interface().(param.Optional); ok && opt.IsPresent() {
			return enc(key, value.FieldByIndex(f.Index), writer)
		} else if ok && opt.IsNull() {
			return writer.WriteField(key, "null")
		}
		return nil
	}
}
