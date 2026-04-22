
# Busca de Logradouro por CEP

## Descrição
Este exercício implementa um programa que busca informações de logradouro a partir de um CEP fornecido na execução do programa.

Para facil criação da Struct ViaCep, peguei o modelo json que esta no site [ViaCep](https://viacep.com.br/)
```
{
      "cep": "01001-000",
      "logradouro": "Praça da Sé",
      "complemento": "lado ímpar",
      "unidade": "",
      "bairro": "Sé",
      "localidade": "São Paulo",
      "uf": "SP",
      "estado": "São Paulo",
      "regiao": "Sudeste",
      "ibge": "3550308",
      "gia": "1004",
      "ddd": "11",
      "siafi": "7107"
    }
```

Em seguida realizei a conversão para struct Go usando o site [transform](https://transform.tools/json-to-go)

## Funcionalidades
- Entrada de CEP via linha de comando
- Consulta de dados do logradouro
- Exibição formatada das informações

## Como executar
```bash
go run main.go <CEP>
```

## Exemplo de uso
```bash
go run main.go 01310100
```

## Requisitos
- Go 1.16 ou superior

## Estrutura do projeto
```
buscaCep/
├── main.go
└── README.md
```
