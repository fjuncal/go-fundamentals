package main

import "fmt"

func PointersExample() {
	// Criando uma variável normal
	number := 10
	fmt.Println("=== Pointers in Go ===")
	fmt.Println("Valor original:", number)

	// Criando um ponteiro para armazenar o endereço da variável `number`
	var ptr *int = &number
	fmt.Println("Endereço da variável (ponteiro):", ptr) // Mostra o endereço de memória
	fmt.Println("Valor acessado pelo ponteiro:", *ptr)   // Acessa o valor armazenado no endereço

	// Modificando o valor através do ponteiro
	*ptr = 20
	fmt.Println("Novo valor da variável após alteração via ponteiro:", number) // Agora `number` vale 20

	// Criando um ponteiro não inicializado (valor padrão: nil)
	var uninitializedPointer *int
	fmt.Println("Ponteiro não inicializado:", uninitializedPointer) // nil (não aponta para nada)

	// Alocação de memória usando `new`
	allocatedPointer := new(int)                                 // Aloca memória para um inteiro e retorna um ponteiro para ele
	*allocatedPointer = 50                                       // Define o valor dentro da memória alocada
	fmt.Println("Valor do ponteiro alocado:", *allocatedPointer) // 50

	// Demonstrando a diferença entre variável normal e ponteiro
	numA := 5
	numB := numA    // Aqui estamos copiando o valor, não a referência
	numPtr := &numA // Aqui estamos copiando o endereço, então ele referencia `numA`

	fmt.Println("numA:", numA, "| numB (cópia):", numB) // Ambos começam com 5
	*numPtr = 100
	fmt.Println("numA após ponteiro modificar:", numA)  // Agora `numA` foi alterado para 100
	fmt.Println("numB (cópia) continua o mesmo:", numB) // `numB` ainda é 5, pois foi uma cópia

	// Criando uma função para modificar valores usando ponteiros
	modifyValue(&numA)
	fmt.Println("numA após função modificar via ponteiro:", numA) // Agora vale 200
}

// Função que recebe um ponteiro e modifica o valor original
func modifyValue(value *int) {
	*value = 200 // Modifica o valor no endereço apontado
}

//🔥 Explicação Detalhada
//🔹 O que é um ponteiro?
//Um ponteiro é uma variável que armazena o endereço de memória de outra variável.
//Em Go, ponteiros são declarados com *Tipo, por exemplo: var ptr *int.
//🔹 O que significa & e *?
//&variável → Obtém o endereço de memória da variável.
//*ponteiro → Acessa o valor armazenado no endereço do ponteiro.
//🔹 O que acontece ao modificar um valor via ponteiro?
//go
//Copiar
//Editar
//number := 10
//ptr := &number   // ptr agora armazena o endereço de number
//*ptr = 20        // Modifica diretamente a variável number
//fmt.Println(number) // Saída: 20
//O valor da variável number foi alterado, porque ptr aponta diretamente para ela.
//
//🔹 O que acontece ao copiar uma variável sem ponteiro?
//go
//Copiar
//Editar
//numA := 5
//numB := numA // Apenas copia o valor, sem referência
//numB = 10
//fmt.Println(numA) // Continua sendo 5
//Como numB foi apenas uma cópia, modificar numB não afeta numA.
//
//🔹 Por que new(int) é útil?
//go
//Copiar
//Editar
//ptr := new(int) // Aloca memória para um inteiro e retorna um ponteiro para ele
//*ptr = 50
//fmt.Println(*ptr) // Saída: 50
//Isso é útil quando queremos criar variáveis sem precisar declarar diretamente uma instância.
//
//🔹 Modificando valores dentro de funções
//Quando passamos uma variável para uma função sem ponteiro, ela é copiada.
//Quando passamos um ponteiro, a função pode modificar o valor original.
//go
//Copiar
//Editar
//func modifyValue(value *int) {
//	*value = 200 // Modifica o valor original na memória
//}
