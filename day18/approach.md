# Day 18

Note: this was done on day 19 as catchup

## Part 1

Another simple BFS graph traversal; I kept the falling 'bytes'/obstacles as a `map[Pair]int`, where the value is the time in which the byte would fall.

## Part 2

As a first go, I decided to just brute force the solution, which worked well enough.

Upon some further thinking, I guess I could have gone for a binary search, but I instead went for a Minimum Spanning Tree approach:

1. Generate connected components *after* all the bytes have fallen. (Each byte is it's own separate cluster)
2. Working backwards from the last fallen byte, connect it with any adjacent connected components (do not connect with bytes yet to have 'unfallen')
3. Stop when the start and end nodes now belong to the same cluster.

After a bit of derusting on the Union Find Disjoint Set, I managed to get the answer in much quicker time.