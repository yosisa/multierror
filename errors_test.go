package multierror

import (
	"errors"
	"reflect"
	"testing"
)

var (
	err1 = errors.New("error 1")
	err2 = errors.New("error 2")
	err3 = errors.New("error 3")
	err4 = errors.New("error 4")
)

func TestAppend(t *testing.T) {
	var err error
	err = Append(err, err1)
	if err == nil {
		t.Error("err is nil")
	}
	err = Append(err, err2)
	if !reflect.DeepEqual(err, Errors([]error{err1, err2})) {
		t.Errorf("unexpected errors: %v", err)
	}
	if v := Append(err1, err2); !reflect.DeepEqual(err, v) {
		t.Errorf("unexpected errors: %v", v)
	}
}

func TestFlatten(t *testing.T) {
	err := Append(err1, Append(err2, err3), err4)
	if !reflect.DeepEqual(err, Errors([]error{err1, err2, err3, err4})) {
		t.Errorf("unexpected errors: %v", err)
	}
}
