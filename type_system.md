In pseudo cpp/go template syntax this file describes the `ras` type system
```
// These are actually named by a titlelized
// reflect.Kind, ie "Uint8" or "Sting" 
type Var<T> interface {
	Set(T) (error)
	Get() (T, error)
	Addr() uint
}

// unsafe
type Pointer Var<uint>

type Array<T> interface {
	Var<T>

	SetI(int) (T, error)
	GetI(int) (T, error)
}

type DyVar interface {
	Len() (int, error)
	Cap() (int, error)
}

type Slice<T> interface {
	DyVar
	Var<[]T>
	
	SetI(int) (T, error)
	GetI(int) (T, error)
}

func LoadStruct(addr uint, v interface{}) error

```

Example
```
type ExampleStruct struct {
	Name Uint8Slice
	Age Uint
}
```