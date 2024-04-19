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
``````