# Estudo pessoal: Entendendo Algoritmos

## Programação Dinâmica
É uma técnica de resolução de problemas complexos que se baseia na divisão de problemas em subproblemas resolvidos separadamente. Devido à dificuldade em encontrar uma solução com a programação dinâmica, podemos inferir algumas características, tais como:

- Quando se busca otimizar à um limite. No problema da mochila, era necessário maximizar o valor dos itens roubados, limitados pela capacidade da mochila;
- O problema que possibilita ser repartido em subproblemas discretos que não dependam um do outro.

Pode-se listar também regras gerais sobre programação dinâmica:

- Toda solução de programação dinâmica envolve uma tabela;
- Geralmente, os valores da célula são os que necessitamos otimizar. Para o problema da mochila, os valores das células correspondem aos valores dos itens.

Trazendo à luz o problema da mochila, a solução simples envolve identificar todos os conjuntos possíveis dos itens para maximizar o valor. Sabendo que cada item adicionado ao problema dobrará o número do cálculo, o algoritmo tem tempo de execução de ```O(2^n)```, considerado de execução lenta.
Os algoritmos de aproximação garantem proximidade da solução ideal, mas não é a própria.

Por meio da programação dinâmica, a mochila maior pode ser dividido em sub-mochilas até que o problema seja resolvido. Cada algoritmo de programação dinâmica começa com uma tabela, como se fosse um "jogo da velha", em que o eixo ```X``` referem-se às capacidades das mochilas em quilos e no eixo ```Y``` são os itens que podemos escolher.

A tabela começa vazia e quando a tabela for preenchida, a resposta do problema terá sido encontrada.

- Para cada linha/item (eixo ```X```), devemos avaliar qual coluna estará o valor do item que caberá na mochila. Todos as possibilidades/capacidades disponíveis considerando somente a existência do item analisado, enquanto os demais itens ainda se encontram indisponíveis;
- A última coluna do eixo ```X``` corresponde ao melhor palpite atual do item, isto é, referente ao subproblema;
- Conforme for descendo as linhas da tabela, é importante considerar as linhas anteriores + linha atual, buscando maximizar os valores dos itens, ocupando a maior capacidade disponível possível;
- Quando tivermos dois ou mais itens que caibam na mochila, o critério de desempate será o valor máximo;
- A última linha da última coluna preenchida sempre será a combinação que maximiza o valor e capacidade ocupada pelos itens para o problema da mochila, mas pode não se aplicar aos demais casos, como acontece para o problema da maior substring comum que considera o maior número da tabela, independente de onde esteja na tabela;
- Caso houver a necessidade de inclusão de algum item, incluir e calcular como última linha da tabela;
- Com o passar dos itens, a valor da coluna sempre deve ser igual ou superior ao anterior;
- A mudança da ordem das linhas não muda o resultado;
- Não é recomendável preencher a tabela a partir das colunas ao invés das linhas, pois o resultado pode mudar;
- Caso houver algum item que não possui o peso correspondente na coluna, a tabela precisa ser modificada para atender esse item, sendo dividido proporcionalmente na tabela;
- Não é possível selecionar frações de um item quando não é possível selecionar o produto na totalidade. Nesse caso, usamos o algoritmo guloso partindo do item mais caro em diante e;
- A programação dinâmica só funciona quando os seus subproblemas são discretos, quando não são dependentes entre si. Exemplo: atrações turísticas de um mesmo lugar não será contabilizado o tempo total de deslocamento.

### Exercícios:
  9.1 Imagine que você consegue roubar outro item: um MP3 player. Ele pesa 1 quilo e vale R$ 1.000,00. Você deveria roubá-lo?
  </br>
  Sim. Então seria possível roubar o MP3, o iPhone e o violão, itens estes que totalizam R$ 4.500,00.

  9.2 Suponha que você esteja indo acampar e que sua mochila tenha capacidade para 6 quilos. Sendo assim, você pode escolher entre os itens abaixo para levar. Cada item tem um valor, e quanto mais alto este valor, mais importante o item é.

  - Água, 3 kg, 10;
  - Livro, 1 kg, 3;
  - Comida, 2 kg, 9;
  - Casaco, 2 kg, 5;
  - Câmera, 1 kg, 6.

  Qual é o conjunto de itens ideal que deve ser levado para o acampamento?
  </br>
  Você deveria levar água, comida e câmera.

  9.3 Desenhe e preencha uma tabela para calcular a maior substring comum entre blue (azul, em inglês) e clues (pistas, em inglês).
  <br/>```C L U E S```</br>
    ```B [0,0,0,0,0]```</br>
    ```L [0,1,0,0,0]```</br>
    ```U [0,0,2,0,0]```</br>
    ```E [0,0,0,3,0]```