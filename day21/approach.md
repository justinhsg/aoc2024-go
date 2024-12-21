# Day 21

Funny prompt / preamble!

## Part 1

This was an interesting DP problem.

If bot `x` needs to key in `<<^A` to move the `x-1` bot, and I needed what shortest key sequence for bot `x+1` to move `x` bot, note that `x` always starts at `A`. Hence we just need to consider the fastest way to move from `A` to `<`, then `<` to `<`, then `<` to `^`, then `^` back to `A`. These can all be done independently.

So, I can create the recursive function `minMoves(sequence, layer)`, to denote the minimum number of inputs required to move the bot according to the sequence at that layer.

The tricky part is memoising all the possible key sequences to move the bot from one position to another. After a few false starts, I did a bit more thinking and realised that all movement will consist of horizontal movement in one direction, vertical movement in one direction, or both. Also, we can assume that we should repeat the same direction as much as possible: `>^>A` will require the `n+1` bot to move a lot more to input than `>>^A`.

The approach I finally went with was to designate grid coordinates for each key, and use the vector difference to generate the horizontal and vertical movements. This produced a bug where paths would 'cross' the gap, which is illegal:

```text
Moving from 4 -> 0: vv>A

789   789
456   456
123   v23
 0A   v>A
```

To overcome this, I created a simple verifier, where I would re-traverse the grid with the given path using a simplified adjacently list, and if the traversal was illegal I would exclude that path.

After some debugging (more on the path finding than the DP), I managed to get the first star

## Part 2

This was a simple extension: increase the initial number of layers from 2 to 25 (in my code this is increasing the layer id from 1 to 24), and adding a simple `map[string]map[int]int` cache to store already computed answers was enough to get the answer in (very) reasonable time