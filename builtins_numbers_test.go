package xex

import (
	"math"
	"testing"
)

/////////////
//int## types
/////////////
func TestIntConversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("int")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestInt8Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("int8")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestInt16Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("int16")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestInt32Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("int32")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestInt64Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("int64")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

//////////////
//uint## types
//////////////
func TestUintConversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("uint")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestUint8Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("uint8")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestUint16Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("uint16")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestUint32Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("uint32")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestUint64Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("uint64")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

///////////////
//float## types
///////////////
func TestFloat32Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("float32")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestFloat64Conversions(t *testing.T) {
	in := 7
	fn, err := GetFunction("float64")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(int64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint8(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint16(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(uint64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float32(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if res, err := fn.Exec(float64(in)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in) {
		t.Errorf("Expected %d, got %v", in, res[0])
		return
	}
	if _, err := fn.Exec("not-a-number"); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

/////////////////
//Maths functions
/////////////////
func TestAdd(t *testing.T) {
	in1 := 7
	in2 := 99
	fn, err := GetFunction("add")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in1), int(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in1)+int(in2) {
		t.Errorf("Expected %d, got %v", int(in1)+int(in2), res[0])
		return
	}
	if res, err := fn.Exec(int8(in1), int8(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in1)+int8(in2) {
		t.Errorf("Expected %d, got %v", int8(in1)+int8(in2), res[0])
		return
	}
	if res, err := fn.Exec(int16(in1), int16(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in1)+int16(in2) {
		t.Errorf("Expected %d, got %v", int16(in1)+int16(in2), res[0])
		return
	}
	if res, err := fn.Exec(int32(in1), int32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in1)+int32(in2) {
		t.Errorf("Expected %d, got %v", int32(in1)+int32(in2), res[0])
		return
	}
	if res, err := fn.Exec(int64(in1), int64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in1)+int64(in2) {
		t.Errorf("Expected %d, got %v", int64(in1)+int64(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint(in1), uint(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in1)+uint(in2) {
		t.Errorf("Expected %d, got %v", uint(in1)+uint(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint8(in1), uint8(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in1)+uint8(in2) {
		t.Errorf("Expected %d, got %v", uint8(in1)+uint8(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint16(in1), uint16(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in1)+uint16(in2) {
		t.Errorf("Expected %d, got %v", uint16(in1)+uint16(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint32(in1), uint32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in1)+uint32(in2) {
		t.Errorf("Expected %d, got %v", uint32(in1)+uint32(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint64(in1), uint64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in1)+uint64(in2) {
		t.Errorf("Expected %d, got %v", uint64(in1)+uint64(in2), res[0])
		return
	}
	if res, err := fn.Exec(float32(in1), float32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in1)+float32(in2) {
		t.Errorf("Expected %f, got %v", float32(in1)+float32(in2), res[0])
		return
	}
	if res, err := fn.Exec(float64(in1), float64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in1)+float64(in2) {
		t.Errorf("Expected %f, got %v", float64(in1)+float64(in2), res[0])
		return
	}
	if _, err := fn.Exec(int8(in1), int64(in2)); err == nil {
		t.Error("Should have failed with different types")
		return
	}
	if _, err := fn.Exec("not-a-number", 10); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestSubtract(t *testing.T) {
	in1 := 7
	in2 := 99
	fn, err := GetFunction("subtract")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in1), int(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in1)-int(in2) {
		t.Errorf("Expected %d, got %v", int(in1)-int(in2), res[0])
		return
	}
	if res, err := fn.Exec(int8(in1), int8(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in1)-int8(in2) {
		t.Errorf("Expected %d, got %v", int8(in1)-int8(in2), res[0])
		return
	}
	if res, err := fn.Exec(int16(in1), int16(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in1)-int16(in2) {
		t.Errorf("Expected %d, got %v", int16(in1)-int16(in2), res[0])
		return
	}
	if res, err := fn.Exec(int32(in1), int32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in1)-int32(in2) {
		t.Errorf("Expected %d, got %v", int32(in1)-int32(in2), res[0])
		return
	}
	if res, err := fn.Exec(int64(in1), int64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in1)-int64(in2) {
		t.Errorf("Expected %d, got %v", int64(in1)-int64(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint(in1), uint(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in1)-uint(in2) {
		t.Errorf("Expected %d, got %v", uint(in1)-uint(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint8(in1), uint8(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in1)-uint8(in2) {
		t.Errorf("Expected %d, got %v", uint8(in1)-uint8(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint16(in1), uint16(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in1)-uint16(in2) {
		t.Errorf("Expected %d, got %v", uint16(in1)-uint16(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint32(in1), uint32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in1)-uint32(in2) {
		t.Errorf("Expected %d, got %v", uint32(in1)-uint32(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint64(in1), uint64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in1)-uint64(in2) {
		t.Errorf("Expected %d, got %v", uint64(in1)-uint64(in2), res[0])
		return
	}
	if res, err := fn.Exec(float32(in1), float32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in1)-float32(in2) {
		t.Errorf("Expected %f, got %v", float32(in1)-float32(in2), res[0])
		return
	}
	if res, err := fn.Exec(float64(in1), float64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in1)-float64(in2) {
		t.Errorf("Expected %f, got %v", float64(in1)-float64(in2), res[0])
		return
	}
	if _, err := fn.Exec(int8(in1), int64(in2)); err == nil {
		t.Error("Should have failed with different types")
		return
	}
	if _, err := fn.Exec("not-a-number", 10); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestMultiply(t *testing.T) {
	in1 := 7
	in2 := 99
	fn, err := GetFunction("multiply")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in1), int(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in1)*int(in2) {
		t.Errorf("Expected %d, got %v", int(in1)*int(in2), res[0])
		return
	}
	if res, err := fn.Exec(int8(in1), int8(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in1)*int8(in2) {
		t.Errorf("Expected %d, got %v", int8(in1)*int8(in2), res[0])
		return
	}
	if res, err := fn.Exec(int16(in1), int16(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in1)*int16(in2) {
		t.Errorf("Expected %d, got %v", int16(in1)*int16(in2), res[0])
		return
	}
	if res, err := fn.Exec(int32(in1), int32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in1)*int32(in2) {
		t.Errorf("Expected %d, got %v", int32(in1)*int32(in2), res[0])
		return
	}
	if res, err := fn.Exec(int64(in1), int64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in1)*int64(in2) {
		t.Errorf("Expected %d, got %v", int64(in1)*int64(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint(in1), uint(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in1)*uint(in2) {
		t.Errorf("Expected %d, got %v", uint(in1)*uint(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint8(in1), uint8(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in1)*uint8(in2) {
		t.Errorf("Expected %d, got %v", uint8(in1)*uint8(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint16(in1), uint16(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in1)*uint16(in2) {
		t.Errorf("Expected %d, got %v", uint16(in1)*uint16(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint32(in1), uint32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in1)*uint32(in2) {
		t.Errorf("Expected %d, got %v", uint32(in1)*uint32(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint64(in1), uint64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in1)*uint64(in2) {
		t.Errorf("Expected %d, got %v", uint64(in1)*uint64(in2), res[0])
		return
	}
	if res, err := fn.Exec(float32(in1), float32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in1)*float32(in2) {
		t.Errorf("Expected %f, got %v", float32(in1)*float32(in2), res[0])
		return
	}
	if res, err := fn.Exec(float64(in1), float64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in1)*float64(in2) {
		t.Errorf("Expected %f, got %v", float64(in1)*float64(in2), res[0])
		return
	}
	if _, err := fn.Exec(int8(in1), int64(in2)); err == nil {
		t.Error("Should have failed with different types")
		return
	}
	if _, err := fn.Exec("not-a-number", 10); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestDivide(t *testing.T) {
	in1 := 7
	in2 := 99
	fn, err := GetFunction("divide")
	if err != nil {
		t.Error(err)
		return
	}
	if res, err := fn.Exec(int(in1), int(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int(in1)/int(in2) {
		t.Errorf("Expected %d, got %v", int(in1)/int(in2), res[0])
		return
	}
	if res, err := fn.Exec(int8(in1), int8(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int8(in1)/int8(in2) {
		t.Errorf("Expected %d, got %v", int8(in1)/int8(in2), res[0])
		return
	}
	if res, err := fn.Exec(int16(in1), int16(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int16(in1)/int16(in2) {
		t.Errorf("Expected %d, got %v", int16(in1)/int16(in2), res[0])
		return
	}
	if res, err := fn.Exec(int32(in1), int32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int32(in1)/int32(in2) {
		t.Errorf("Expected %d, got %v", int32(in1)/int32(in2), res[0])
		return
	}
	if res, err := fn.Exec(int64(in1), int64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != int64(in1)/int64(in2) {
		t.Errorf("Expected %d, got %v", int64(in1)/int64(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint(in1), uint(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint(in1)/uint(in2) {
		t.Errorf("Expected %d, got %v", uint(in1)/uint(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint8(in1), uint8(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint8(in1)/uint8(in2) {
		t.Errorf("Expected %d, got %v", uint8(in1)/uint8(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint16(in1), uint16(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint16(in1)/uint16(in2) {
		t.Errorf("Expected %d, got %v", uint16(in1)/uint16(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint32(in1), uint32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint32(in1)/uint32(in2) {
		t.Errorf("Expected %d, got %v", uint32(in1)/uint32(in2), res[0])
		return
	}
	if res, err := fn.Exec(uint64(in1), uint64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != uint64(in1)/uint64(in2) {
		t.Errorf("Expected %d, got %v", uint64(in1)/uint64(in2), res[0])
		return
	}
	if res, err := fn.Exec(float32(in1), float32(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float32(in1)/float32(in2) {
		t.Errorf("Expected %f, got %v", float32(in1)/float32(in2), res[0])
		return
	}
	if res, err := fn.Exec(float64(in1), float64(in2)); err != nil {
		t.Error(err)
		return
	} else if res[0] != float64(in1)/float64(in2) {
		t.Errorf("Expected %f, got %v", float64(in1)/float64(in2), res[0])
		return
	}
	if _, err := fn.Exec(int8(in1), int64(in2)); err == nil {
		t.Error("Should have failed with different types")
		return
	}
	if _, err := fn.Exec("not-a-number", 10); err == nil {
		t.Error("String to int should have failed")
		return
	}
}

func TestPow(t *testing.T) {
	in1 := 2
	in2 := 8
	fn, err := GetFunction("pow")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := fn.Exec(float64(in1), float64(in2))
	if err != nil {
		t.Error(err)
		return
	}
	if res[0] != math.Pow(2, 8) {
		t.Errorf("Expected %f, got %v", math.Pow(float64(in1), float64(in2)), res[0])
		return
	}
}
