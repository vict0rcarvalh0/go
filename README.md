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
