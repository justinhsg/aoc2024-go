# Day 23

## Part 1

This was a very simple problem, but I somehow got caught with weird bugs when finding if the group of 3 I found had a node that starts with a 't'.

To deduplicate the found triplets, I decided to iterate in a particular order such that : `0 <= i < j < k < n`

## Part 2

This problem rang a bell in my head; this was a very well known problem but I couldn't place my finger on it; after fruitlessly trying to find the problem / algorithm online to refresh my memory, I decided to return to brute force.

(On hindsight, I realised this was the 'clique' problem, which if I had known the name, would have immediately recalled from complexity theory lectures that this was a classic NP-complete problem)

Back to the brute force solution! The puzzle provided a hint (alphabetical sorting), which gave a heuristic on how to prune the search space, but essentially:

1. The algorithm starts with a list of all cliques of size 1.
2. Then, for each clique, consider each candidate node 'greater than' the last node in the clique.
3. check if the candidate has an edge to all nodes in the clique.If so, add it to the list of cliques of size 2.
4. Then repeat this process to generate the list of cliques of size 3, 4 and so on until the list of cliques is empty.

This produced an answer in 1s, but just like the Day 22, I really wanted to get this down to sub 500ms.

2.5 optimisations were used to aid in this:

1. Instead of using strings as keys into the map for the adjacency matrix and list, I mapped each node to an integer (naturally, its position in the alphabetically sorted list of nodes)

    1.5) Instead of storing the edges as $i \to j$ and $j \to i$, the edges were made directional, only edges $i \to j$ were considered where $i < j$
2. Instead of finding the candidate nodes $j$ from all nodes, I only considered the candidate nodes where $i \to j$ from the adjacency list.

With these 2 optimisations, the runtime dropped to <50ms.
