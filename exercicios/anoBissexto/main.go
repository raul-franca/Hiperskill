package main

import "fmt"

func main() {

	var year int
	for {
		fmt.Println("Type a year, or 0 to exit.")
		fmt.Scanln(&year)

		if year == 0 {
			fmt.Println("Goodbye.")
			break
		}

		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			fmt.Printf("%d is a Leap year\n", year)
		} else {
			fmt.Printf("%d is a regular year\n", year)
		}

	} // Fim do for
}

//Este é um código em Go que permite ao usuário inserir um ano e determina se ele é um ano bissexto ou não. Aqui está uma descrição mais detalhada do código:
//
//O programa começa solicitando ao usuário que digite um ano ou digite 0 para sair.
//O programa lê a entrada do usuário usando o fmt.Scanln e armazena o valor na variável year.
//Se o usuário digitou 0, o programa imprime "Goodbye" e sai do loop usando a instrução break.
//Se o ano for divisível por 4 e não for divisível por 100, ou se o ano for divisível por 400, o programa imprime que o ano é bissexto usando o fmt.Printf.
//Caso contrário, o programa imprime que o ano é um ano comum.
//O programa continua em um loop infinito até que o usuário digite 0 para sair.
//Este é um exemplo simples de como a linguagem Go pode ser usada para criar um programa útil e interativo.
