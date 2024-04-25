# Go
Repository to register my Golang studies

### Installation(Linux)
1. Download Go: <a href="https://go.dev/dl/">https://go.dev/dl/</a>
2. Run the following commands:
    - `sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz`

    - `export PATH=$PATH:/usr/local/go/bin`

    - `go version`


### Folder structure

- **cmd/**: This folder contains the executable applications. Each subfolder within cmd/ represents a specific application and usually contains a main.go file that serves as the application's entry point.

- **pkg/**: This folder contains reusable code that can be imported by other projects. Each subfolder within pkg/ represents a separate module or package. The files within these subfolders contain the relevant code for that package.

- **internal/**: This folder contains code that can only be imported by other packages in the same module. This means that code inside internal/ cannot be imported by packages outside the module's root directory. It's useful for encapsulating code that shouldn't be accessible outside of your project.

``````
meu_projeto/
│
├── cmd/
│   └── my_app/
│       └── main.go
│
├── pkg/
│   ├── module1/
│   │   ├── file1.go
│   │   └── file2.go
│   │
│   └── module2/
│       ├── file1.go
│       └── file2.go
│
└── internal/
    ├── private_module1/
    │   ├── file1.go
    │   └── file2.go
    │
    └── private_module2/
        ├── file1.go
        └── file2.go
``````

### go.mod
Usado para definir as dependências do projeto, bem como o próprio módulo do projeto. O Go usa um sistema de controle de versão chamado "modules" para gerenciar dependências. Ao importar um pacote em seu código, o Go procura no módulo go.mod para determinar de onde deve baixar o pacote e em qual versão.

- `go mod init example_mod`
    - Inicia um novo módulo Go em um diretório. Ele cria um arquivo go.mod inicial com o nome do módulo especificado.
- `go mod tidy`
    - Verifica quais pacotes seu programa realmente depende e adiciona ou remove dependências em seu arquivo go.mod conforme necessário.
- `go mod diretorio_exemplo`
    - Copia todas as dependências do seu módulo para um diretório localizado no seu projeto.

## Syntax
#### Packages
The entire Go program is made up of packages, at the beginning of each file it is necessary to declare the `package file_name` and import the modules as follows:
``` go
import (
    "fmt"
    "math/rand"
)
```

#### Types
- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32, represents a Unicode code point
- float32 float64
- complex64 complex128

#### Type Conversion
There are two ways to convert types:
1. **Explicit**
```go
var i int = 10
var f float64 = float64(i)
var u uint = uint(f)
```
2. **Implicit**
```go
i := 42
f := float64(i)
u := uint(f)
```

#### Functions
In Go, functions are values too, thay can be passed around just like other values.
```go
func MyFunction(var1 type, var2 type) return_type {
    return x + y
}
```

```go
func MyFunction(var1 type, var2 type) (return_type, return_type) {
    return x + y
}
```

#### Variables
The var statement declares a list of variables; as in function argument lists, the type is last.

```go
var a, b, c bool

func main() {
	var i int
	fmt.Println(i, a, b, c)
}
```

#### Constants
Are declared like variables, but with the const keyword.
```go
const MyConst
```

#### For
```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

Go can also loop forever by using the following syntax:
```go
for {
}
```

#### While(in Go is a for)
```go
sum := 1
for sum < 1000 {
    sum += sum
}
```

#### Conditional(if/else)
```go
var x int
if x < 10; x > 0 {
    return x
} else {
    fmt.Printf("x does not satisfy condition :(")
}
```

#### Switch
```go
switch time.Saturday {
case today + 0:
    fmt.Println("Today.")
case today + 1:
    fmt.Println("Tomorrow.")
case today + 2:
    fmt.Println("In two days.")
default:
    fmt.Println("Too far away.")
}
```

#### Defer
Defer statements is used to ensure that a function is executed when the current function is about to return. This is useful for freeing up resources such as files or network connections, or for ensuring that certain actions are performed regardless of how the current function ends (whether by an error or by success).
```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

#### Pointers
Holds the memory address of a value.

1. The type *T is a pointer to a T value. Its zero value is nil.
```go
var p *int
```

2. The & operator generates a pointer to its operand.
```go
i := 42
p = &i
```

The * operator denotes the pointer's underlying value.
```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```
This is known as "dereferencing" or "indirecting".

#### Struct
It's a composite data structure that allows you to group related heterogeneous data under a single type. Structures are defined by fields, which are variables with names and types, grouped under a single structure name.
```go
type StructName struct {
    X type1
    Y type2
}

v := StructName{1, 2} // assigns values to the struct fields
v.X = 4 // assigns a value to one field

// pointers
p := &v
p.X = 1e9
```

#### Arrays
```go
// first declare and then assign
var a [2]string
a[0] = "Hello"
a[1] = "World"

// declaring and assigning
primes := [6]int{2, 3, 5, 7, 11, 13}
```

#### Slices
- A slice is a dynamically-sized, flexible view into the elements of an array.
- Slice does not store any data, it just describes a section of an underlying array.
- Changing the elements of a slice modifies the corresponding elements of its underlying array.
- The length and capacity of a slice s can be obtained using the expressions `len(s)` and `cap(s)`.
- The zero value of a slice is `nil` -> Capacity and Length of 0.

```go
[]T // -> Slice of elements with type T
primes := [6]int{2, 3, 5, 7, 11, 13} // -> length 6, type int, values
var s []int = primes[1:4] // -> output: 3, 5, 7

// Slice the slice to give it zero length.
s = s[:0]
printSlice(s)

// Extend its length.
s = s[:4]
printSlice(s)

// Drop its first two values.
s = s[2:]
printSlice(s)

// Creating Slice with make(built-in function that allocates a zeroed array and returns a slice that refers to that array)
a := make([]int, 5)  // len(a)=5

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

// Slice of Slices
board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
}

// Append works on nil slices.
s = append(s, 0)
printSlice(s)

// The slice grows as needed.
s = append(s, 1)
printSlice(s)

// We can add more than one element at a time.
s = append(s, 2, 3, 4)
printSlice(s)

// The range form of the for loop iterates over a slice or map.
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// You can skip the index or value by assigning to _.
for i, _ := range pow
for _, value := range pow

// If you only want the index, you can omit the second variable.
for i := range pow
```

#### Maps
- Maps is useful to map keys to values
- Zero value of a map is nil(has no keys and cannot be added)
- The make function returns a map of the given type, initialized and ready for use

```go
    var m map[type]StructName 
    m = make(map[type]StructName)
    m["My Key 1"] = StructName{
		40.68433, -74.39967,
	}

    // Map literals 
    var m = map[type]StructName{
        "My Key 2": StructName{
            40.68433, -74.39967,
        },
        "My Key 3": StructName {
            40.68433, -74.39967,
        },
    }

    // Map literals continued -> Top-level type is just a type name
    var m = map[type]StructName{
	"My Key 2": {40.68433, -74.39967},
	"My Key 3": {37.42202, -122.08408},
    }

    // Inserting or updating an element in map m
    m := make(map[string]int)
    m[key] = elem

    // Retrieve an element
    elem = m[key]

    // Delete an element
    delete(m, key)

    // Test that a key is present with a two-value assignment:
    elem, ok = m[key] // if key is in m ok is true, if not ok is false
```

#### Function closures

#### Methods
- Go doesn't have classes! But methods can be defined on types
- Method is a function with a special receiver(appears in its own arg list between func word and method name) argument
- The receiver can be a pointer. There are two reasons to use a pointer receiver:
    - The first is so that the method can modify the value that its receiver points to.
    - The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
- In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

```go
func (v StructName) MethodName() return_type {
    return "blablabla"
}
```

#### Interfaces
- Interfaces are implemented implicitly by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
- Under the hood, interface values can be thought of as a tuple of a value and a concrete type.
    - Calling a method on an interface value executes the method of the same name on its underlying type.

```go
type InterfaceName interface {
	MethodName() return_type
}

var a InterfaceName
f := MyFloat(-math.Sqrt2)
s := StructName{1, 2}

a = f  // a MyFloat implements InterfaceName
a = &v // a *StructName implements InterfaceName

type MyFloat float64

func (f MyFloat) MethodName() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Empty Interface
var i interface{}

```

#### Special Interfaces
**Stringers**
- A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
```go
type Stringer interface {
    String() string
}
```

**Errors**
- The error type is a built-in interface similar to fmt.Stringer.
- Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
- A nil error denotes success; a non-nil error denotes failure.
```go
type error interface {
    Error() string
}
```

**Readers**
- The io package specifies the io.Reader interface, which represents the read end of a stream of data.
- The io.Reader interface has a Read method
    - Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.
```go
func (T) Read(b []byte) (n int, err error)
```

**Images**
- Package image defines the Image interface
```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

#### Generics
- Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets, before the function's arguments.
- Comparable is a useful constraint that makes it possible to use the == and != operators on values of the type
```go
func Index[T comparable](s []T, x T) int // s is a slice of any type T that fulfills the built-in constraint comparable. x is also a value of the same type.

// Generic Types
type List[T any] struct {
	next *List[T]
	val  T
}
```
