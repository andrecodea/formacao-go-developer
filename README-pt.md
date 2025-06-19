<div align="center">
  <h1>Formação Go Developer</h1>

  <p>Este repositório armazena os desafios que cumpri durante a Formação Go Developer da DIO.</p>

  ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
</div>

## 🎯 Sobre a Formação
Formação ministrada pela professora **Tenille Martins** na plataforma DIO, focada no desenvolvimento de habilidades práticas em Go através de desafios progressivos.

## 🚀 Desafios

### Desafio 1: Conversor Termométrico de Kelvin para Celsius | [repo](challenge1-temperature-converter.go)
- **Objetivo:** Conversão de escalas termométricas com Go
- **Desafio:** Desenvolver um programa para converter o ponto de ebulição da água de Kelvin para Celsius
- **Solução:** Criei um programa que converte qualquer temperatura de Kelvin para Celsius e identifica automaticamente os pontos de ebulição e solidificação da água com mensagens específicas

### Desafio 2: Detector de Números Divisíveis por 3 | [repo](challenge2-divisible-by-three.go)
- **Objetivo:** Detectar números divisíveis por 3 em Go
- **Desafio:** Implementar detecção usando loop `for` e operador de módulo (`%`)
- **Solução:** Desenvolvi uma função que identifica números divisíveis por 3 em um intervalo determinado utilizando estruturas de repetição e operações matemáticas

### Desafio 3: PinPan | [repo](challenge3-pinpan.go)
- **Objetivo:** Implementar um programa estilo "FizzBuzz" personalizado
- **Desafio:** Exibir "Pin" para números divisíveis por 3 e "Pan" para divisíveis por 5
- **Solução:** Criei uma função que combina as verificações dos desafios anteriores, detectando divisibilidade por 3 e 5 simultaneamente usando loop `for` e operador de módulo (`%`)

### Desafio 4: Ping Pong | [repo](challenge4-pingpong-concurrency.go)
- Objetivo: Implementar um programa pingpong usando concorrência
- Desafio: Implementar concorrência para imprimir "ping" e "pong" um após o outro, usando goroutines e channels.
- Solução: Criei dois channels que enviam "ping" e "pong", tinha uma extremidade receptora em um loop que imprime a mensagem e dorme por 1 segundo. Adicionalmente, usei sync para sincronizar ambas as goroutines e evitar que a goroutine principal termine o programa sem nenhum retorno, usando as funções WaitGroup Add(), Wait() e Done().

## 📚 Aprendizados
- Fundamentos da linguagem Go
- Estruturas de controle e repetição
- Operações matemáticas e lógicas
- Boas práticas de programação
- Resolução de problemas algorítmicos

## 🤝 Contribuições
Sinta-se à vontade para explorar os códigos e sugerir melhorias!


## 📄 Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para detalhes.

---
<div align="center">
  <b>⭐ Gostou do projeto? Deixe uma estrela no repositório!</b>
</div>
