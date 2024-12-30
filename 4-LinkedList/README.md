# Estudo pessoal: Entendendo Algoritmos

## Ordenação por seleção
Arrays e lista encadeada serão abordados nesse capítulo, onde aprenderemos quando a lista encadeada é a melhor opção do que o array.
A maioria das linguagens de programação possuem nativamente os algoritmos de seleção, então raramente terá de escrever o algoritmo do zero.
O computador é composto por um conjunto de gavetas e em cada gaveta possui um endereço de um slot de memória, mas quando queremos armazenar diversos valores, podemos fazer isso de duas formas, arrays e listas. Lembrando que não existe uma forma correta para armazenar itens para cada caso, mas é importante conhecer as diferenças. 

### Arrays e listas encadeadas
Ao necessitar armazenar uma lista de elementos na memória, optaremos primeiro pelo array devido fácil compreensão, armazenando de maneira contígua (uma ao lado da outra). Uma forma de fazer uso desse mecanismo é para armazenar uma lista de afazeres. Ao optar por um array nem sempre o conteúdo será de armazenamento sequencial, pelo fato da memória estar em uso.
Quando precisamos adicionar conteúdo aos slots de memória já ocupados, é necessário buscar um novo espaço de memória que comporte todo o conteúdo e, por isso, adicionar itens a um array torna-se mais lento. Uma maneira de contornar o problema é reservando um número superior de espaços desejados a fim de evitar deslocamento. Vale ressaltar que existem desvantagens como desperdício de memória que por sua vez não ficará disponível para outros processos e possivelmente necessitará mover os itens de qualquer forma pois os "espaços reservados" serão limitados.

## Listas encadeadas
Permite manter os itens em qualquer lugar da memória, em que cada item armazena o endereço do próximo item da lista e endereços aleatórios de memória se mantém interligados, não sendo mais necessário deslocar os itens da lista para novos slots de memória quando escalado o volume dos dados. O tempo de execução para as operações em listas encadeadas dependem de fatores como a operação a ser realizada, tipo de lista encadeada, implementação da estrutura de dados na linguagem de programação escolhida e tamanho da lista. Geralmente, temos o tempo de execução médio para algumas operações em listas encadeadas conforme abaixo:

| Operação                          | Lista Encadeada: Caso médio | Lista Encadeada: Pior caso |
| --------------------------------- | --------------------------- | -------------------------- |
| Acesso aleatório ao elemento      |              O(1)           |             O(n)           | 
| Busca                             |              O(n)           |             O(n)           |
| Inserção no início                |              O(1)           |             O(1)           |
| Inserção no final                 |              O(1)           |             O(n)           |
| Remoção no ínicio                 |              O(1)           |             O(1)           |
| Remoção no final                  |              O(n)           |             O(n)           |
| Remoção de um elemento arbitrário |              O(n)           |             O(n)           |

Referente à explicação do acesso aleatório ao elemento, em listas encadeadas, para acesso aleatório, o elemento desejado é acessado diretamente no tempo médio, já no pior caso, se falamos de uma lista longa e o elemento estiver no final, será percorrida toda a lista. Para operação de busca, precisa percorrer a lista sequencialmente até encontrar o elemento buscado.
Para operação no início, consiste em somente atualizar o ponteiro inicial da lista para apontar para o novo elemento. Por outro lado, a inserção no final visa percorrer a lista até o último elemento e atualizar o ponteiro "next" do penúltimo elemento para apontar para um novo elemento e no pior caso, se a lista for muito longa, será necessário percorrer toda a lista para encontrar o penúltimo elemento.
Sobre remoção no ínicio, é necessário atualizar o ponteiro inicial da lista para apontar para o próximo elemento. A remoção no final, deve-se percorrer a lista até o penúltimo elemento para atualizar o ponteiro "next" para apontar para null. E por fim, a remoção de um elemento arbitrário, necessita percorrer a lista até encontrar o elemento a ser removido, e os ponteiros "next" e "prev" (se for uma lista duplamente encadeada) precisam ser atualizados.

### Exercícios:
2.1 Suponha que você esteja criando um aplicativo para acompanhar as suas finanças. 
</br>
  1. Compras
  </br>
  2. Cinema
  </br>
  3. Mensalidade do SFBC
 
  Todos os dias você escreve tudo o que gastou e onde. No final do mês, você revisa os seus gastos e resume o quanto gastou. Logo, você tem um monte de inserções e poucas leituras. Você deve usar um array ou uma lista para implementar este aplicativo?
  Neste caso, você está adicionando despesas na lista todos os dias e lendo todas as despesas uma vez por mês. Arrays tem leitura rápida, mas inserção lenta. Listas encadeadas têm leituras lentas e rápidas inserções. Como você inserirá mais vezes do que lerá, faz mais sentido usar uma lista encadeada. Além disso, listas encadeadas têm leitura lenta somente quando você acessa elementos aleatórios da lista. Como estará lendo todos os elementos da lista, a lista encadeada terá também uma boa velocidade de leitura. Portanto, uma lista encadeada terá também uma boa velocidade de leitura. Portanto, uma lista encadeada é uma boa solução para esse problema.

2.2 Suponha que você esteja criando um aplicativo para anotar os pedidos dos clientes em um restaturante. Seu aplicativo precisa de uma lista de pedidos. Os garçons adicionam os pedidos a essa lista e os chefs retiram os pedidos da lista. Funciona como uma fila. Os garçons colocam os pedidos no final da fila e os chefes retiram os pedidos do começo dela para cozinhá-los. Você utilizaria um array ou uma lista encadeada para implementar essa lista (listas encadeadas são boas para inserções/eliminações e arrays são bons para acesso aleatório)?
Uma lista encadeada. Muitas inserções estão ocorrendo (garçons adicionando ordens), sendo essa uma das vantagens da lista encadeada. Você não precisa pesquisar ou ter acesso aleatório (nisso os arrays são bons), pois o chef sempre pega a primeira ordem da fila.

2.3 Vamos analisar um experimento. Imagine o Facebook guarde uma lista de usuários. Quando alguém tenta acessar o Facebook, uma busca é feita pelo nome do usuário. Se o nome da pessoa está na lista, ela pode continuar o acesso. As pessoas acessam o Facebook com muita frequência, então existem muitas buscas nessa lista. Presuma que o Facebook use a pesquisa binária para procurar um nome na lista. A pesquisa binária precisa de acesso aleatório - você precisa ser capaz de acessar o meio da lista de nome instantaneamente. Sabendo disso, você implementaria essa lista como um array ou lista encadeada?
Um array ordenado. Arrays fornecem acesso aleatório, então você pode pegar um elemento do meio do array instantaneamente. Isso não é possível com listas encadeadas. Para acessar o elemento central de uma lista encadeada, você deve iniciar com o primeiro elemento e seguir por todos os links até o elemento central.

2.4 As pessoas se inscrevem no Facebook com muita frequência também. Suponha que você decida usar um array para armazenar a lista de usuários. Quais as desvantagens de um array em relação às inserções? Em particular, imagine que você esteja usando a pesquisa binária para buscar os logins. O que acontece quando você adiciona novos usuários em um array?
Inserções em arrays são lentas. Além disso, se você estiver utilizando a pesquisa binária para procurar os nomes de usuário, o array precisará estar ordenado. Suponha que alguém chamado Adit B se registre no Facebook. O nome dele será inserido no final do array. Assim, você precisa ordenar o array cada vez que um nome for inserido.

2.5 Na verdade, o Facebook não usa nem arrays nem listas encadeadas para armazenar informações. Vamos considerar uma estrutura de dados híbrida: um array de listas encadeadas. Você tem um array com 26 slots. Cada slot aponta para uma lista encadeada. Por exemplo, o primeiro slot do array aponta para uma lista encadeada que contém todos os usuários que começam com a letra A. O segundo slot aponta para a lista encadeada que contém todos os usuários que começam com a letra B, e assim por diante.
Suponha que o Adit B se inscreva no Facebook e você queira adicioná-lo na lista. Você vai ao slot 1 do array, a seguir para a lista encadeada do slot 1, e adiciona Adit B no final. Agora, suponha que você queira procurar o Zakhir H. Você vai ao slot 26. que aponta para a lista encadeda de todos os nomes começados em Z. Então, procura a lista até encontrar o Zakhir H. 
Compare essa estrutura híbrida com arrays e listas encadeadas. É mais lento ou mais rápido fazer inserções e eliminações nesse caso? Você não precisa responder dando o tempo de execução Big (O), apenas diga se a nova estrutura de dados é mais rápida ou mais lenta do que os arrays e as listas encadeadas.
Para buscas - mais lenta do que arrays, mais rápida do que listas encadeadas. Para inserções - mais rápida do que arrays, mesmo tempo que as listas encadeadas. Portanto, é mais lenta para buscas que os arrays, porém mais rápida ou igual às listas encadeadas para tudo. Falaremos sobre outra estrutura de dados híbridos chamada tabela hash depois. Isto deve dar uma ideia sobre como é possível construir estruturas de dados mais complexas a partir das estruturas mais simples.
Então, o que o Facebook realmente utiliza? Provavelmente uma dúzia de diferentes bancos de dados com diferentes estruturas por trás deles, como tabelas hash, árvores B e outras. Os arrays e as listas encadeadas são os blocos fundamentais para estruturas de dados mais complexos.