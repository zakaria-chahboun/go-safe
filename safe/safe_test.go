package safe

import (
	"reflect"
	"testing"
	"time"
)

// TestValue_String tests Value function with string pointers
func TestValue_String(t *testing.T) {
	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *string
		result := Value(ptr, "default")
		if result != "default" {
			t.Errorf("Expected 'default', got '%s'", result)
		}
	})

	t.Run("nil pointer without default", func(t *testing.T) {
		var ptr *string
		result := Value(ptr)
		if result != "" {
			t.Errorf("Expected empty string, got '%s'", result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		str := "Zakaria"
		ptr := &str
		result := Value(ptr)
		if result != "Zakaria" {
			t.Errorf("Expected 'Zakaria', got '%s'", result)
		}
	})

	t.Run("non-nil pointer with default", func(t *testing.T) {
		str := "Zakaria"
		ptr := &str
		result := Value(ptr, "default")
		if result != "Zakaria" {
			t.Errorf("Expected 'Zakaria', got '%s'", result)
		}
	})
}

// TestValue_Int tests Value function with int pointers
func TestValue_Int(t *testing.T) {
	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *int
		result := Value(ptr, 25)
		if result != 25 {
			t.Errorf("Expected 25, got %d", result)
		}
	})

	t.Run("nil pointer without default", func(t *testing.T) {
		var ptr *int
		result := Value(ptr)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		num := 42
		ptr := &num
		result := Value(ptr)
		if result != 42 {
			t.Errorf("Expected 42, got %d", result)
		}
	})
}

// TestValue_Struct tests Value function with struct pointers
func TestValue_Struct(t *testing.T) {
	type Info struct {
		City string
		Code int
	}

	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *Info
		defaultInfo := Info{City: "Agadir", Code: 123}
		result := Value(ptr, defaultInfo)
		if result.City != "Agadir" || result.Code != 123 {
			t.Errorf("Expected {Agadir 123}, got %+v", result)
		}
	})

	t.Run("nil pointer without default", func(t *testing.T) {
		var ptr *Info
		result := Value(ptr)
		expected := Info{}
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		info := Info{City: "Casablanca", Code: 456}
		ptr := &info
		result := Value(ptr)
		if result.City != "Casablanca" || result.Code != 456 {
			t.Errorf("Expected {Casablanca 456}, got %+v", result)
		}
	})
}

// TestValue_Slice tests Value function with slice pointers
func TestValue_Slice(t *testing.T) {
	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *[]string
		defaultSlice := []string{"a", "b", "c"}
		result := Value(ptr, defaultSlice)
		if !reflect.DeepEqual(result, defaultSlice) {
			t.Errorf("Expected %v, got %v", defaultSlice, result)
		}
	})

	t.Run("nil pointer without default", func(t *testing.T) {
		var ptr *[]string
		result := Value(ptr)
		if result != nil {
			t.Errorf("Expected nil slice, got %v", result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		slice := []string{"x", "y", "z"}
		ptr := &slice
		result := Value(ptr)
		if !reflect.DeepEqual(result, slice) {
			t.Errorf("Expected %v, got %v", slice, result)
		}
	})
}

// TestValue_Map tests Value function with map pointers
func TestValue_Map(t *testing.T) {
	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *map[string]int
		defaultMap := map[string]int{"a": 1, "b": 2}
		result := Value(ptr, defaultMap)
		if !reflect.DeepEqual(result, defaultMap) {
			t.Errorf("Expected %v, got %v", defaultMap, result)
		}
	})

	t.Run("nil pointer without default", func(t *testing.T) {
		var ptr *map[string]int
		result := Value(ptr)
		if result != nil {
			t.Errorf("Expected nil map, got %v", result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		m := map[string]int{"x": 10, "y": 20}
		ptr := &m
		result := Value(ptr)
		if !reflect.DeepEqual(result, m) {
			t.Errorf("Expected %v, got %v", m, result)
		}
	})
}

// TestValue_Time tests Value function with time.Time pointers
func TestValue_Time(t *testing.T) {
	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *time.Time
		defaultTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		result := Value(ptr, defaultTime)
		if !result.Equal(defaultTime) {
			t.Errorf("Expected %v, got %v", defaultTime, result)
		}
	})

	t.Run("nil pointer without default", func(t *testing.T) {
		var ptr *time.Time
		result := Value(ptr)
		expected := time.Time{}
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		now := time.Now()
		ptr := &now
		result := Value(ptr)
		if !result.Equal(now) {
			t.Errorf("Expected %v, got %v", now, result)
		}
	})
}

// TestPointer_String tests Pointer function with string pointers
func TestPointer_String(t *testing.T) {
	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *string
		result := Pointer(ptr, "default")
		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != "default" {
			t.Errorf("Expected 'default', got '%s'", *result)
		}
	})

	t.Run("nil pointer without default", func(t *testing.T) {
		var ptr *string
		result := Pointer(ptr)
		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != "" {
			t.Errorf("Expected empty string, got '%s'", *result)
		}
	})

	t.Run("non-nil pointer reuse", func(t *testing.T) {
		str := "Zakaria"
		ptr := &str
		result := Pointer(ptr)
		if result != ptr {
			t.Error("Expected same pointer to be returned")
		}
		if *result != "Zakaria" {
			t.Errorf("Expected 'Zakaria', got '%s'", *result)
		}
	})

	t.Run("non-nil pointer with default should ignore default", func(t *testing.T) {
		str := "Zakaria"
		ptr := &str
		result := Pointer(ptr, "default")
		if result != ptr {
			t.Error("Expected same pointer to be returned")
		}
		if *result != "Zakaria" {
			t.Errorf("Expected 'Zakaria', got '%s'", *result)
		}
	})
}

// TestPointer_Int tests Pointer function with int pointers
func TestPointer_Int(t *testing.T) {
	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *int
		result := Pointer(ptr, 25)
		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != 25 {
			t.Errorf("Expected 25, got %d", *result)
		}
	})

	t.Run("nil pointer without default", func(t *testing.T) {
		var ptr *int
		result := Pointer(ptr)
		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != 0 {
			t.Errorf("Expected 0, got %d", *result)
		}
	})

	t.Run("non-nil pointer reuse", func(t *testing.T) {
		num := 42
		ptr := &num
		result := Pointer(ptr)
		if result != ptr {
			t.Error("Expected same pointer to be returned")
		}
		if *result != 42 {
			t.Errorf("Expected 42, got %d", *result)
		}
	})
}

// TestPointer_Identity tests pointer identity behavior
func TestPointer_Identity(t *testing.T) {
	t.Run("same non-nil pointer returns identical pointers", func(t *testing.T) {
		str := "test"
		ptr := &str

		result1 := Pointer(ptr)
		result2 := Pointer(ptr)

		if result1 != result2 {
			t.Error("Expected identical pointers for same non-nil input")
		}
		if result1 != ptr || result2 != ptr {
			t.Error("Expected returned pointers to be the same as input pointer")
		}
	})

	t.Run("nil pointer returns different pointers", func(t *testing.T) {
		var ptr *int

		result1 := Pointer(ptr)
		result2 := Pointer(ptr)

		if result1 == result2 {
			t.Error("Expected different pointers for nil input")
		}
		if *result1 != *result2 {
			t.Error("Expected same values in different pointers")
		}
	})
}

// TestPointer_Struct tests Pointer function with struct pointers
func TestPointer_Struct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	t.Run("nil pointer with default", func(t *testing.T) {
		var ptr *User
		defaultUser := User{Name: "Guest", Age: 25}
		result := Pointer(ptr, defaultUser)
		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if result.Name != "Guest" || result.Age != 25 {
			t.Errorf("Expected {Guest 25}, got %+v", *result)
		}
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		user := User{Name: "Zakaria", Age: 30}
		ptr := &user
		result := Pointer(ptr)
		if result != ptr {
			t.Error("Expected same pointer to be returned")
		}
	})
}

// TestValue_MultipleDefaults tests behavior with multiple default values
func TestValue_MultipleDefaults(t *testing.T) {
	t.Run("multiple defaults uses first one", func(t *testing.T) {
		var ptr *string
		result := Value(ptr, "first", "second", "third")
		if result != "first" {
			t.Errorf("Expected 'first', got '%s'", result)
		}
	})
}

// TestPointer_MultipleDefaults tests behavior with multiple default values
func TestPointer_MultipleDefaults(t *testing.T) {
	t.Run("multiple defaults uses first one", func(t *testing.T) {
		var ptr *int
		result := Pointer(ptr, 10, 20, 30)
		if result == nil {
			t.Error("Expected non-nil pointer")
		}
		if *result != 10 {
			t.Errorf("Expected 10, got %d", *result)
		}
	})
}

// BenchmarkValue tests performance of Value function
func BenchmarkValue(b *testing.B) {
	str := "benchmark"
	ptr := &str

	b.Run("non-nil pointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Value(ptr)
		}
	})

	b.Run("nil pointer with default", func(b *testing.B) {
		var nilPtr *string
		for i := 0; i < b.N; i++ {
			_ = Value(nilPtr, "default")
		}
	})
}

// BenchmarkPointer tests performance of Pointer function
func BenchmarkPointer(b *testing.B) {
	str := "benchmark"
	ptr := &str

	b.Run("non-nil pointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Pointer(ptr)
		}
	})

	b.Run("nil pointer with default", func(b *testing.B) {
		var nilPtr *string
		for i := 0; i < b.N; i++ {
			_ = Pointer(nilPtr, "default")
		}
	})
}
