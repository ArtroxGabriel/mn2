# Métodos Numéricos 2 - Implementação

## Visão geral

Este repositório contém implementações de métodos numéricos para cálculo de derivadas e integrais utilizando diferentes técnicas de aproximação numérica, desenvolvido em Go. O foco principal é na aplicação de métodos de diferenças finitas para derivação e diversas regras para integração.

Para uma discussão detalhada sobre os fundamentos teóricos dos métodos numéricos implementados, incluindo as fórmulas de diferenças finitas, ordens de precisão e análise de convergência, consulte o documento [DOCS.md](./DOCS.md).

## Checklist de Implementação

### 📊 Derivadas Numéricas

#### Derivadas Primeiras (1ª Derivada

- [x] **O(h)** - Diferença Forward
- [x] **O(h²)** - Diferença Forward
- [] **O(h³)** - Diferença Forward
- [x] **O(h)** - Diferença Backward
- [x] **O(h²)** - Diferença Backward
- [] **O(h³)** - Diferença Backward
- [x] **Ordem O(h²)** - Diferença Central
- [ ] **Ordem O(h⁴)** - Diferença Central

#### Derivadas Segundas (2ª Derivada)

- [x] **O(h)** - Diferença Forward
- [x] **O(h²)** - Diferença Forward
- [] **O(h³)** - Diferença Forward
- [x] **O(h)** - Diferença Backward
- [x] **O(h²)** - Diferença Backward
- [] **O(h³)** - Diferença Backward
- [x] **Ordem O(h²)** - Diferença Central
- [ ] **Ordem O(h⁴)** - Diferença Central

#### Derivadas Terceiras (3ª Derivada)

- [x] **O(h)** - Diferença Forward
- [x] **O(h²)** - Diferença Forward
- [] **O(h³)** - Diferença Forward
- [x] **O(h)** - Diferença Backward
- [x] **O(h²)** - Diferença Backward
- [] **O(h³)** - Diferença Backward
- [x] **Ordem O(h²)** - Diferença Central
- [ ] **Ordem O(h⁴)** - Diferença Central

### 🔢 Integrais Numéricas

#### Métodos Newton-Cotes

##### Abertos

- [ ] **Ordem O(h)** - Regra do Trapézio
- [ ] **Ordem O($h^5$)** - Regra de Simpson 1/3
- [ ] **Ordem O($h^5$)** - Regra de Simpson 3/8
- [ ] **Ordem O($h^7$)** - Regra de Boole

##### Fechados

- [ ] **Ordem O($h^3$)** - Regra do Retângulo
- [ ] **Ordem O($h^3$)** - Regra do Ponto Médio
- [ ] **Ordem O($h^5$)** - Regra de Milne

#### Métodos Gauss-Legendre

- [ ] **Ordem O(h⁴)** - Gauss-Legendre
- [ ] **Ordem O(h⁵)** - Gauss-Hermite
- [ ] **Ordem O(h⁶)** - Gauss-Lagerre
- [ ] **Ordem O(h)** - Gauss-Chebyshev

## Estrutura do Projeto

```
implementacao/
├── derivacao_numerica/
│   ├── main.go          # Programa principal com testes
│   └── derivadas.go     # Implementações das derivadas
├── integracao_numerica/ # (A ser implementado)
│   ├── main.go
│   └── integrais.go
└── README.md           # Este arquivo
```

## Como Executar

### Derivação Numérica

```bash
cd derivacao_numerica
go run *.go
```

### Configuração de Logs

O projeto utiliza `slog` para logging estruturado. Os níveis disponíveis são:

- `DEBUG`: Logs detalhados de cada cálculo
- `INFO`: Informações principais do processo
- `WARN`: Avisos sobre convergência
- `ERROR`: Erros críticos

## Documentação Teórica

Os conceitos teóricos detalhados, incluindo as derivações das fórmulas, análise de erro e discussões sobre a ordem de precisão para os métodos de diferenciação e integração numérica implementados neste projeto, estão documentados em [DOCS.md](./DOCS.md). Este documento serve como referência para entender os fundamentos por trás dos algoritmos.

## Análise de Convergência

O programa implementa análise de convergência para os métodos numéricos, geralmente através da observação do comportamento do erro conforme o passo `h` é ajustado. Detalhes sobre a teoria da convergência e a interpretação dos resultados podem ser encontrados em [DOCS.md](./DOCS.md).

## Próximos Passos

1. Implementar derivadas de segunda e terceira ordem.
2. Adicionar métodos de ordem superior (O(h³) e O(h⁴)) para maior precisão.
3. Implementar o módulo de integração numérica com os métodos listados no checklist.
4. Adicionar testes unitários robustos para todas as funções implementadas.
5. Implementar funcionalidades para comparar o desempenho e a precisão de diferentes métodos.
6. Adicionar análise de performance para otimizar os cálculos.

## Contribuição

Para contribuir com o projeto:

1. Marque os itens implementados no checklist.
2. Adicione testes unitários para quaisquer novos métodos ou modificações.
3. Siga o padrão de logging estruturado com `slog`.
4. Consulte [DOCS.md](./DOCS.md) para a base teórica e, se necessário, atualize-o ao introduzir novos conceitos.
