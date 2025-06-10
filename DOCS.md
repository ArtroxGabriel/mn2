# Relatório sobre Métodos e Estratégias Numéricas

---

## Introdução

Este documento descreve os métodos de diferenciação e integração numérica implementados e utilizados neste projeto. Ele detalha os fundamentos teóricos desses métodos e explica como eles são aplicados na prática na base de código, fornecendo uma visão abrangente das estratégias numéricas empregadas.

---

## Diferenciação Numérica

Diferenciação numérica é uma técnica utilizada para estimar a derivada de uma função em um ponto usando valores da função em pontos discretos. Isso é particularmente útil quando a função é dada como um conjunto de pontos de dados ou quando a derivada analítica é difícil ou impossível de calcular.

### Métodos de Diferenças Finitas

Métodos de diferenças finitas aproximam derivadas substituindo-as por diferenças entre valores de função em pontos próximos. A escolha de quais pontos usar determina o método específico.

#### Diferença Progressiva (Forward Difference)

O método da diferença progressiva aproxima a derivada em um ponto `x` usando o valor da função em `x` e `x+h`, onde `h` é um pequeno tamanho de passo.

**Fórmula:**

```
f'(x) ≈ (f(x+h) - f(x)) / h
```

**Explicação:**
Este método usa a inclinação da linha que conecta os pontos `(x, f(x))` e `(x+h, f(x+h))` para aproximar a tangente em `x`. É um método de primeira ordem.

##### Estratégias Implementadas

As seguintes estratégias são implementadas para o método da Diferença Progressiva e são usadas pelo `Derivador` quando a estratégia "Forward" é escolhida com a ordem correspondente:

- `firstorder.ForwardFirstOrderStrategy` (O(h))
- `secondorder.ForwardSecondOrderStrategy` (O(h^2))
- `thirdorder.ForwardThirOrderStrategy` (O(h^3))

#### Diferença Regressiva (Backward Difference)

O método da diferença regressiva aproxima a derivada em um ponto `x` usando o valor da função em `x` e `x-h`.

**Fórmula:**

```
f'(x) ≈ (f(x) - f(x-h)) / h
```

**Explicação:**
Este método usa a inclinação da linha que conecta os pontos `(x-h, f(x-h))` e `(x, f(x))` para aproximar a tangente em `x`. Também é um método de primeira ordem.

##### Estratégias Implementadas

As seguintes estratégias são implementadas para o método da Diferença Regressiva e são usadas pelo `Derivador` quando a estratégia "Backward" é escolhida com a ordem correspondente:

- `firstorder.BackwardFirstOrderStrategy` (O(h))
- `secondorder.BackwardSecondOrderStrategy` (O(h^2))
- `thirdorder.BackwardThirOrderStrategy` (O(h^3))

#### Diferença Central (Central Difference)

O método da diferença central aproxima a derivada em um ponto `x` usando os valores da função em `x-h` e `x+h`.

**Fórmula:**

```
f'(x) ≈ (f(x+h) - f(x-h)) / (2h)
```

**Explicação:**
Este método usa a inclinação da linha que conecta os pontos `(x-h, f(x-h))` e `(x+h, f(x+h))`. Geralmente, ele fornece uma aproximação mais precisa do que os métodos de diferença progressiva ou regressiva para o mesmo tamanho de passo `h`, porque considera informações simetricamente em torno de `x`. É um método de segunda ordem para aproximar a primeira derivada.

##### Estratégias Implementadas

As seguintes estratégias são implementadas para o método da Diferença Central e são usadas pelo `Derivador` quando a estratégia "Central" é escolhida com a ordem correspondente:

- `secondorder.CentralSecondOrderStrategy` (O(h^2))
- `fourthorder.CentralFourthOrderStrategy` (O(h^4))

### Ordem de Precisão

A ordem de precisão de um método de diferenças finitas indica como o erro da aproximação muda à medida que o tamanho do passo `h` é reduzido.

- **O(h) (Precisão de primeira ordem):** O erro é aproximadamente proporcional a `h`. Reduzir `h` pela metade, aproximadamente, reduz o erro pela metade. Os métodos de diferença progressiva e regressiva são tipicamente O(h).
- **O(h^2) (Precisão de segunda ordem):** O erro é aproximadamente proporcional a `h^2`. Reduzir `h` pela metade, aproximadamente, reduz o erro em um quarto. O método de diferença central para a primeira derivada é tipicamente O(h^2).
- **O(h^4) (Precisão de quarta ordem):** O erro é aproximadamente proporcional a `h^4`. Reduzir `h` pela metade, aproximadamente, reduz o erro por um fator de cerca de 16. Métodos de ordem superior podem ser derivados usando mais pontos, como o stencil de cinco pontos para a primeira derivada, que é O(h^4).

Métodos de ordem superior geralmente fornecem melhor precisão para um dado tamanho de passo `h`, mas podem exigir mais avaliações de função e podem ser mais suscetíveis a erros de arredondamento se `h` for muito pequeno.

### Uso no Código

O tipo `Derivator` é a interface principal para realizar a diferenciação numérica. Ele é inicializado usando o construtor `NewDerivator(strategy string, order uint)`.

- O parâmetro `strategy` pode ser um de "Forward", "Backward" ou "Central".
- O parâmetro `order` especifica a ordem de precisão desejada para a estratégia escolhida.

A `DerivacaoFactory` é responsável por interpretar essas entradas e selecionar o objeto de estratégia concreto apropriado. Por exemplo, se `strategy` for "Forward" e `order` for 1, a fábrica instanciará um objeto `firstorder.ForwardFirstOrderStrategy`.

Uma vez que o `Derivator` é configurado com uma estratégia, seu método `Calculate` pode ser chamado para calcular a derivada. O método `Calculate` recebe um parâmetro `derivate`, que pode ser 1, 2 ou 3, para calcular a primeira, segunda ou terceira derivada, respectivamente.

A estratégia selecionada (por exemplo, `CentralFourthOrderStrategy`) então fornece as fórmulas e coeficientes específicos necessários para calcular a derivada solicitada (1ª, 2ª ou 3ª) de acordo com seu método de diferenças finitas subjacente e ordem de precisão.

### Análise de Convergência

A análise de convergência na diferenciação numérica examina como a aproximação se aproxima da verdadeira derivada à medida que o tamanho do passo `h` tende a zero.

Idealmente, à medida que `h` diminui, o erro de truncamento (o erro de aproximar a derivada com uma fórmula de diferença finita) também diminui. Por exemplo, para um método O(h^2), o erro de truncamento é proporcional a `h^2`. Portanto, reduzir `h` deve melhorar a precisão.

No entanto, há uma compensação. Quando `h` se torna muito pequeno, erros de arredondamento devido à precisão limitada da aritmética computacional podem se tornar significativos. Subtrair dois números muito próximos (como `f(x+h)` e `f(x)`) pode levar a uma perda de precisão, e então dividir por um `h` muito pequeno pode amplificar esse erro.

Portanto, existe um valor ótimo de `h` que equilibra o erro de truncamento e o erro de arredondamento. Escolher um `h` excessivamente pequeno pode levar a resultados menos precisos do que um `h` moderadamente pequeno. Técnicas como a extrapolação de Richardson podem ser usadas para estimar a derivada com mais precisão e também para estimar o `h` ótimo.

---

## Integração Numérica (Quadratura)

A integração numérica, também conhecida como quadratura, envolve a aproximação do valor de uma integral definida. Isso é essencial quando a antiderivada do integrando é desconhecida ou difícil de encontrar, ou quando a função é conhecida apenas em pontos de dados discretos.

### Fórmulas de Newton-Cotes

As fórmulas de Newton-Cotes são um grupo de regras de integração numérica baseadas na avaliação do integrando em pontos igualmente espaçados. A ideia geral é aproximar a função a ser integrada por um polinômio interpolador de um certo grau e, em seguida, integrar esse polinômio.

- **Ideia Geral:** Substituir a função `f(x)` sobre o intervalo `[a, b]` por um polinômio que seja fácil de integrar. O grau do polinômio e os pontos usados para interpolação determinam a regra específica.
- **Fórmulas Abertas vs. Fechadas:**
  - **Fórmulas fechadas** usam os valores da função nos pontos finais do intervalo de integração. Exemplos incluem a Regra do Trapézio e as regras de Simpson.
  - **Fórmulas abertas** usam apenas valores da função em pontos estritamente dentro do intervalo de integração. Estas são úteis para integrandos com singularidades nos pontos finais. Um exemplo é a Regra do Ponto Médio (que também é a fórmula de Newton-Cotes aberta mais simples, baseada em um polinômio de grau zero ou constante).

#### Métodos Comuns de Newton-Cotes

- **Regra do Trapézio:** Esta regra aproxima a integral ajustando um polinômio de primeiro grau (uma linha reta) entre os valores da função nos pontos finais do intervalo (ou cada subintervalo na versão composta). A área sob esta linha (um trapézio) aproxima a integral.
- **Regra de Simpson (Regra 1/3):** Esta regra usa um polinômio de segundo grau (uma parábola) para interpolar a função, usando três pontos igualmente espaçados: os dois pontos finais e o ponto médio do intervalo. Geralmente, oferece maior precisão do que a Regra do Trapézio para funções suaves.
- **Regra de Simpson (Regra 3/8):** Esta regra usa um polinômio de terceiro grau para interpolar a função, usando quatro pontos igualmente espaçados. Pode ser mais precisa do que a regra 1/3 para algumas funções, mas requer uma avaliação de função a mais.

#### Estratégias Newton-Cotes Implementadas

As seguintes estratégias Newton-Cotes são implementadas em `internal/integration/strategies.go` e são utilizadas pelo `Integrator` quando o nome da estratégia correspondente é escolhido:

- `strategies.NewtonCotesOrder1`: Implementa a **Regra do Trapézio**.
- `strategies.NewtonCotesOrder2`: Implementa a **Regra de Simpson 1/3**.
- `strategies.NewtonCotesOrder3`: Implementa a **Regra de Simpson 3/8**.
- `strategies.NewtonCotesOrder4`: Implementa uma fórmula de Newton-Cotes de quarta ordem, comumente conhecida como **Regra de Boole**.

### Quadratura Gaussiana

As fórmulas de quadratura Gaussiana oferecem uma abordagem alternativa que frequentemente alcança maior precisão para o mesmo número de avaliações de função em comparação com as regras de Newton-Cotes.

- **Ideia Geral:** Em vez de fixar as abscissas (valores x) para serem igualmente espaçadas, os métodos de quadratura Gaussiana escolhem as localizações dos pontos de avaliação (nós) e os pesos de forma ótima. Esses nós são tipicamente as raízes de uma família de polinômios ortogonais. Ao escolher esses nós e pesos estrategicamente, a quadratura Gaussiana pode integrar exatamente polinômios de grau `2n-1` com apenas `n` avaliações de função.
- **Quadratura de Gauss-Legendre:** Este é um tipo comum de quadratura Gaussiana usado para integrais sobre o intervalo `[-1, 1]`. Os nós são as raízes dos polinômios de Legendre, e os pesos são escolhidos para alcançar a maior precisão possível. Integrais sobre outros intervalos `[a, b]` podem ser transformadas para `[-1, 1]` usando uma mudança linear de variável.

#### Estratégias Gauss-Legendre Implementadas

As seguintes estratégias de quadratura Gauss-Legendre são implementadas em `internal/integration/strategies.go`. Elas são usadas pelo `Integrator` e representam métodos com diferentes números de pontos otimamente escolhidos, fornecendo vários níveis de precisão:

- `strategies.GaussLegendreOrder1`
- `strategies.GaussLegendreOrder2`
- `strategies.GaussLegendreOrder3`
- `strategies.GaussLegendreOrder4`

Essas estratégias correspondem à quadratura Gauss-Legendre usando 1, 2, 3 e 4 pontos, respectivamente. Aumentar a ordem (número de pontos) geralmente leva a maior precisão para funções suaves.

### Uso no Código

O tipo `Integrator` é a interface principal para realizar a integração numérica. Ele é inicializado usando o construtor `NewIntegrator(strategyName string)`.

- O parâmetro `strategyName` é uma string que identifica unicamente a estratégia de integração desejada, por exemplo, "NewtonCotesOrder1" ou "GaussLegendreOrder2".

A `IntegrationFactory` é responsável por interpretar o `strategyName` e fornecer o objeto de estratégia concreto correspondente (por exemplo, uma instância de `strategies.NewtonCotesOrder1` ou `strategies.GaussLegendreOrder2`).

Uma vez que o `Integrator` é configurado com uma estratégia específica, seu método `Calculate(fn Func, a, b float64, n int) (float64, error)` pode ser chamado. Este método usa a estratégia selecionada para calcular a integral definida da função `fn` de `a` a `b`, potencialmente usando `n` subintervalos ou pontos, dependendo da estratégia.

### Convergência e Erro na Integração Numérica

Semelhante à diferenciação numérica, o erro na integração numérica depende do método utilizado e do número de pontos de avaliação (ou da largura dos subintervalos, `h`, em regras compostas).

- **Erro de Truncamento:** Este erro surge da aproximação da função por uma mais simples (por exemplo, um polinômio). Para regras de Newton-Cotes como a regra do Trapézio composta, o erro é tipicamente da ordem O(h^2), enquanto para a regra de Simpson 1/3 composta, é O(h^4), onde `h` é a largura dos subintervalos. A quadratura Gaussiana com `n` pontos pode integrar polinômios até o grau `2n-1` exatamente, levando a uma convergência muito rápida para funções suaves.
- **Erro de Arredondamento:** Assim como na diferenciação, realizar muitos cálculos com precisão finita pode levar ao acúmulo de erros de arredondamento. No entanto, na integração, os erros de arredondamento são geralmente menos problemáticos do que na diferenciação, porque as operações envolvidas (soma e multiplicação por pesos) são menos sensíveis a pequenos valores de `h` do que a divisão por `h` ou `h^2`.

Geralmente, aumentar o número de pontos de avaliação (ou diminuir `h` em regras compostas) melhora a precisão da aproximação, reduzindo o erro de truncamento. No entanto, para um número muito grande de pontos, o erro de arredondamento pode eventualmente começar a aumentar, embora isso seja menos preocupante do que na diferenciação. Métodos de quadratura adaptativa ajustam o tamanho do passo `h` (ou o número de pontos) em diferentes partes do domínio de integração para atingir um nível de precisão desejado de forma eficiente.

---

## Conclusão

Este documento forneceu uma visão geral dos principais métodos numéricos para diferenciação e integração empregados no projeto. Ao detalhar tanto os fundamentos teóricos quanto as implementações práticas por meio de estratégias específicas e padrões de fábrica, ele serve como um guia para o conjunto de ferramentas numéricas disponíveis. Esses métodos oferecem soluções robustas e flexíveis para as tarefas de análise numérica encontradas na aplicação.

Fontes
