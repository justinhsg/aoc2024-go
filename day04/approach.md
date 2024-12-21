# Day 4

Note: this was done as catch-up on Day 7

## Part 1

Finding all `X`s in the grid, and searching in all 8 directions for `XMAS` while respecting boundaries was enough to get the star.

## Part 2

Observation: since `A` is in the centre, we can just skip the first and last rows and columns of the search, and search for `A`, while searching for adjacent `M`s and `S`s

Had one mistake where I assumed that a cross (+) also counted as an `X`, but once that was cleared up, it was relatively simple, if not tedious to get the second star.
