# Day 5

Note: this was done as catch-up on Day 7

## Part 1

The input seemed relatively small, so I chose an adjacency matrix to represent the order.

I assumed that the ordering was transitive and that I needed to 'relax' the ordering (e.g. using Floyd-Warshall). e.g. if `A|B` and `B|C`, then `A,C` is still a valid ordering.

Once I got the wrong answer, I removed Floyd-Warshall and just used the ordering rules present and that passed.

Neat trick, because the numbers acutally don't mean anything until getting the final answer, the adjacency matrix was just a `map[string]map[string]bool`

## Part 2

Noticed that they just wanted the incorrectly-ordered lists to be ordered, so using a simple Sort with a custom comparator (the adjacency matrix calculated earlier) was enough to solve Part 2.