# Métodos Numéricos 2 - Implementação

## Visão geral

Este repositório contém implementações de métodos numéricos para cálculo de derivadas e integrais utilizando diferentes técnicas de aproximação numérica. O projeto foi desenvolvido em Go e inclui implementações de métodos de diferenças finitas para derivação numérica com diferentes ordens de precisão.

As implementações focam em:
- Métodos de diferenças finitas (forward, backward, central)
- Análise de convergência e erro relativo
- Logging detalhado para análise de performance
- Estrutura modular e extensível

## Checklist de Implementação

### 📊 Derivadas Numéricas

#### Derivadas Primeiras (1ª Derivada)
- [x] **Ordem O(h)** - Diferença Forward
- [x] **Ordem O(h)** - Diferença Backward  
- [x] **Ordem O(h²)** - Diferença Central
- [ ] **Ordem O(h³)** - Diferença Central de 5 pontos
- [ ] **Ordem O(h⁴)** - Diferença Central de 7 pontos

#### Derivadas Segundas (2ª Derivada)
- [ ] **Ordem O(h)** - Diferença Forward
- [ ] **Ordem O(h)** - Diferença Backward
- [ ] **Ordem O(h²)** - Diferença Central de 3 pontos
- [ ] **Ordem O(h³)** - Diferença Central de 5 pontos
- [ ] **Ordem O(h⁴)** - Diferença Central de 7 pontos

#### Derivadas Terceiras (3ª Derivada)
- [ ] **Ordem O(h)** - Diferença Forward
- [ ] **Ordem O(h)** - Diferença Backward
- [ ] **Ordem O(h²)** - Diferença Central de 4 pontos
- [ ] **Ordem O(h³)** - Diferença Central de 6 pontos
- [ ] **Ordem O(h⁴)** - Diferença Central de 8 pontos

### 🔢 Integrais Numéricas

#### Métodos Básicos
- [ ] **Ordem O(h)** - Regra do Retângulo (Esquerda)
- [ ] **Ordem O(h)** - Regra do Retângulo (Direita)
- [ ] **Ordem O(h²)** - Regra do Ponto Médio
- [ ] **Ordem O(h²)** - Regra do Trapézio
- [ ] **Ordem O(h³)** - Regra de Simpson 1/3

#### Métodos Avançados
- [ ] **Ordem O(h⁴)** - Regra de Simpson 3/8
- [ ] **Ordem O(h⁵)** - Regra de Boole
- [ ] **Ordem O(h⁶)** - Fórmulas de Newton-Cotes de 6ª ordem

#### Métodos Adaptativos
- [ ] **Simpson Adaptativo**
- [ ] **Quadratura Gaussiana**
- [ ] **Integração de Romberg**

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
