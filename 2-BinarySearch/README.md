# Estudo pessoal: Entendendo Algoritmos

## Pesquisa ou Busca Binária
Consiste em repartir uma lista ordenada até chegar no conteúdo buscado, uma solução que possibilita ser realizada com o menos número de tentativas possíveis. A estratégia da busca binária é chutar um número intermediário e eliminar a metade dos números a cada vez que a estrutura for repartida.
A pesquisa simples é quando você precisa percorrer uma lista inteira para identificar o conteúdo buscado, minando a performance do programa. De maneira geral, uma lista de n números, a pesquisa simples necessitará verificar em n etapas, enquanto a pesquisa binária necessita de ```log² n```. 

### Exercícios:
1.1 Suponha que você tenha uma lista com 128 nomes e esteja fazendo uma pesquisa binária. Qual seria o número máximo de etapas que você levaria para encontrar o nome desejado?
Cálculo dos 128 nomes: ```128 nomes -> 64 -> 32 -> 16 -> 8 -> 4 -> 2 -> 1```,  totalizando 7 etapas até encontrar o número desejado.

1.2 Suponha que você duplique o tamanho da lista. Qual seria o número máximo de etapas agora?
Cálculo dos 256 nomes: ```256 nomes -> 128 -> 64 -> 32 -> 16 -> 8 -> 4 -> 2 -> 1```, totalizando 8 etapas até encontrar o número desejado.

## Tempo de execução
Quando o número máximo de tentativas é igual ao tamanho da lista, é chamado de tempo linear. A pesquisa binária é executada de acordo com o tempo logarítimo.
</br>

### Exercícios:
Forneça o tempo de execução para cada um dos casos a seguir em termos da notação Big O.

1.3 Você tem um nome e deseja encontrar o número de telefone para esse nome em um agenda telefônica.
</br>
```O(log n)```

1.4 Você tem um número de telefone e deseja encontrar o dono dele em uma agenda telefônica (Dica: Deve procurar pela agenda inteira).
</br>
```O(n)```

1.5 Você quer ler o número de cada pessoa na agenda telefônica.
</br>
```O(n)```

1.6 Você quer ler os números apenas dos nomes que começam com A. (Isso é complicado! Esse algoritmo envolve conceitos que serão abordados mais profundamente no Capítulo 4: Quicksort. Leia e responda - você ficará surpreso).
</br>
```O(n)```. Você pode pensar: "Só estou fazendo isso para 1 dentre 26 caracteres, portanto o tempo de execução deve ser ```O(n/26)```". Uma regra simples é a de ignorar números que são somados, subtraídos, multiplicados ou divididos. Nenhum desses são tempos de execução Big O: ```O(n + 26)```, ```O(n - 26)```, ```O(n * 26)```, ```O(n / 26)```. Eles são todos o mesmo que ```O(n)```! Por quê? Se você está com dúvidas, vá para "Notação Big O revisada", no Capítulo 4, e leia a parte sobre constantes na notação Big O (uma constante é apenas um número; 26 era a constante dessa questão).