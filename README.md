# Métodos Numéricos 2 - Implementação

## Visão geral

Este repositório contém implementações de métodos numéricos para cálculo de derivadas e integrais utilizando diferentes técnicas de aproximação numérica. O projeto foi desenvolvido em Go e inclui implementações de métodos de diferenças finitas para derivação numérica com diferentes ordens de precisão.

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

## Fórmulas Implementadas

### Diferenças Finitas para 1ª Derivada

**Forward (O(h)):**

```
f'(x) ≈ [f(x+h) - f(x)] / h
```

**Backward (O(h)):**

```
f'(x) ≈ [f(x) - f(x-h)] / h
```

**Central (O(h²)):**

```
f'(x) ≈ [f(x+h) - f(x-h)] / (2h)
```

## Análise de Convergência

O programa implementa análise automática de convergência através da redução iterativa do passo `h` e cálculo do erro relativo entre iterações sucessivas.

## Próximos Passos

1. Implementar derivadas de segunda e terceira ordem
2. Adicionar métodos de ordem superior (O(h³) e O(h⁴))
3. Implementar módulo de integração numérica
4. Adicionar testes unitários
5. Implementar comparação entre métodos
6. Adicionar análise de performance

## Contribuição

Para contribuir com o projeto:

1. Marque os itens implementados no checklist
2. Adicione testes para novos métodos
3. Mantenha o padrão de logging estruturado
4. Documente as fórmulas utilizadas
