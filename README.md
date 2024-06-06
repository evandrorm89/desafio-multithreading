# desafio-multithreading

Instruções:
- Rodar o programa na linha de comando colocando o cep que se deseja pesquisar nos argumentos; ex: go run main.go 01153000
- O programa pesquisa em threads separadas por duas apis - brasilapi e viacep - e retorna os dados da primeira que responder. Caso nenhuma responda em menos de um segundo, o programa sai, printando "timeout" na tela
