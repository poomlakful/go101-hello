#### #Learning source:
- Awesome-go
- http://go-database-sql.org/
- tour.golang // this online IDE for golang
- https://gobyexample.com/
- https://mholt.github.io/json-to-go/ - use for pars json to go struct

#### #Go install

#### #Go folder structure
bin/
src/

#### #Go modulea

#### #Go keyword

Default value is not nil some variable type

    - 0 int
    - False bool
    - “” (the empty string) string

##### While is for
    for i < 10  {
        i++
    }

##### slide:
https://tour.golang.org/moretypes/18

##### Custom Type:
    - default datatype
    - struct
    - function

##### method:

##### function support:
    - closure-fuction keep state
    - anonymous-function like lambda function in java

##### defer:
    defer some_function()
    will call some_function after end current fucntion

##### interface:

##### web freamwork:
    - https://echo.labstack.com/
    - https://gin-gonic.com/

##### goroutine:
	go ${call_function}

    sync.WaitGroup
    wg.add()
    wg.wait()

    chanel

##### recover:
    This'll show really error
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
    }


#### #Go command
```
$go version
$go env
$go mod init
$go run ${filename_and_ext}
$go test -v
$go build -o main.go
$GOOS=windows go build -o server.exe main.go
```