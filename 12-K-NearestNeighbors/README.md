# Estudo pessoal: Entendendo Algoritmos

## K-Vizinhos Mais Próximos
O algoritmo dos k-vizinhos funciona de maneira simples: escolha um objeto para classificar, olhe para os três (ou qualquer outro número considerável para comparação) vizinhos mais próximos e o resultado dos vizinhos indicará que o objeto provavelmente será tal. Para determinar os vizinhos e agrupar por similaridade, a extração de características pode ajudar, focando em características (idealmente boas e ruins) como tamanho e cor. 

O teorema de Pitágoras ajuda a determinar a distância entre dois objetos no gráfico, que pode ser calculado considerando um conjunto de números ou calcular por meio da similaridade do cosseno que compara o ângulo entre dois vetores.

O mesmo deve ser determinado após ter coordenadas para os respectivos objetos.

### Exercícios:
  10.1 No exemplo da Netflix, calculou-se a distância entre dois usuários diferentes utilizando a fórmula da distância, mas nem todos os usuários avaliam filmes da mesma maneira. Suponha que você tenha dois usuários, Yogi e Pinky, os quais têm gostos similares. No entanto, Yogi avalia qualquer filme que ele goste com 5, enquanto Pinky é mais seletivo e reserva o 5 somente para os melhores filmes. Eles têm gostos bem simlares, mas, de acordo com o algoritmo da distância, eles não são vizinhos. Como você poderia levar em conta o sistema de avaliação diferente deles?
  </br>
  Avaliando a média de cada pessoa e utilizar o valor para as avaliações. Tornar as avaliações das pessoas em uma mesmo valor pode ajudar a comparar as avaliações de mesma escala.

  10.2 Suponha que a Netflix nomeie um grupo de "influenciadores". Por exemplo, Quentin Tarantino e Wes Anderson são influenciadores na Netflix, portanto as avaliações deles contam mais do que as de um usuário comum. Como você poderia modificar o sistema de recomendações de forma que as avaliações dos influenciadores tenham um peso maior?
  </br>
  Dar o maior peso para as avaliações dos influenciadores utilizando o algoritmo dos k-vizinhos mais próximos.

  10.3 A Netflix tem milhões de usuários, e no exemplo anterior consideram-se os cinco vizinhos mais próximos ao criar o sistema de recomendações. Esse número é baixo demais? Ou talvez alto demais?
  </br>
  Baixo, pois se olhar para menos vizinhos, haverá maior chance maior de que o resultado seja tendencioso. Uma boa regra é a seguinte: se você tem N usuários, deve considerar sqrt(N) vizinhos.

Há duas etapas que devem ser levadas em consideração para chegar no valor desejado:

- Classificação: classificação de objetos em grupos e;
- Regressão: média de valores dos objetos para advinhar uma resposta.

Para se aproximar das características certas a serem comparadas, deve-se:

- Considerar características diretamente correlacionadas aos objetos que está tentando recomendar e;
- Selecionar características imparciais.

Prever a quantidade de pães a serem vendidos por dia é uma forma inteligente de aplicar o algoritmo k-vizinhos mais próximos levando em conta as características, como:

- O clima ser ruim ou ótimo na escala de 1 à 5;
- Fim de semana ou feriado classificado como 1 ou 0;
- Se haveria ou não um jogo no dia corrente classificado como 1 ou 0.

Olhando para o passado, é importante identificar os conjuntos diferentes de características que foram identificados, indicando a quantidade de pães vendidos:

A. ```(5,1,0)``` = 300 pães;
</br>
B. ```(3,1,1)``` = 225 pães;
</br>
C. ```(1,1,0)``` = 75 pães;
</br>
D. ```(4,0,1)``` = 200 pães;
</br>
E. ```(4,0,0)``` = 150 pães;
</br>
F. ```(2,0,0)``` = 50 pães.

Hoje é um fim de semana de clima bom. Assim, baseado nos dados que você observou há pouco, quantos pães venderá? Utilizaremos o algoritmo dos k-vizinhos mais próximos, em que ```K = 4```. Para isso, primeiro descubra os quatro pontos mais próximos desse ponto de ```(4,1,0)```.
Aqui temos as distâncias, onde é possível observar que os pontos ```A```, ```B```, ```D``` e ```E``` são os mais próximos, uma vez que os valores dos mesmos são: 

A. ```1```;
</br>
B. ```2```;
</br>
C. ```9```;
</br>
D. ```2```;
</br>
E. ```1```;
</br>
F. ```5```.

A média do número de pães vendidos nesses dias: ```218,75```. Esse é o número que deveria ser produzido.
