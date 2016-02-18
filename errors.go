package multierror

import (
	"bytes"
	"fmt"
)

var DefaultFormatFunc = func(es []error) string {
	b := new(bytes.Buffer)
	fmt.Fprintf(b, "%d errors: ", len(es))
	for i, err := range es {
		if i != 0 {
			b.WriteString("; ")
		}
		fmt.Fprintf(b, "[%d] %s", i+1, err)
	}
	return b.String()
}

var FormatFunc = DefaultFormatFunc

type Errors []error

func (es Errors) Error() string {
	return FormatFunc(es)
}

func Append(err error, errs ...error) error {
	if err == nil {
		return flatten(errs)
	}
	if merr, ok := err.(Errors); ok {
		return flatten(append(merr, errs...))
	}
	errs = append(errs, nil)
	copy(errs[1:], errs[0:])
	errs[0] = err
	return flatten(errs)
}

func flatten(errs Errors) Errors {
	for i := 0; i < len(errs); {
		if merr, ok := errs[i].(Errors); ok {
			errs = append(errs[:i], append(merr, errs[i+1:]...)...)
			continue
		}
		i++
	}
	return errs
}
