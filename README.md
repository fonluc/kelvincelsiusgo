### Conversor de Kelvin para Celsius

Criada a pasta temp dentro da pasta src, aberto o VsCode, criado o arquivo temp.go com o seguinte código:

```go
package main

import "fmt"

const ebulicaoK = 373

func main() {

	tempK := ebulicaoK
	tempC := (tempK - 273)

	fmt.Println("A temperatura de ebulição da água em ºK = ", tempK)
	fmt.Println("A temperatura de ebulição da água em ºC = ", tempC)
}
```

Depois foi aberto o terminal e dados os comandos: “go build temp.go” e “go run temp.go”, exibindo na tela as mensagens:

“A temperatura de ebulição da água em ºK = 373”

”A temperatura de ebulição da água em ºC = 100”
