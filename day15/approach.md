# Day 15

## Part 1
A direct simulation was good enough, with a small optimisation where, instead of moving each box one step down the line, you could just remove the very first box and add another box at the end of the line.

## Part 2

This was slightly unexpected, but still doable. For the horizontal pushes, it wasn't as tricky; just keep track of the contiguous line of `[` and `]`, and shift them once the end of the line is found and if a wall is not in the way.

For vertical pushes, the following pseudocode (modified BFS) was used (using the case of pushing up):

> Initial state: `toVisit` is a list containing the positions of the box (both `[` and `]`)
>
> Loop: For each position in `toVisit` look at the position above it, if it is `[` or `]`, add that position (and the position of the other half of the box) to `toVisitNext` (deduplicating positions). If the position above is a wall `#`, then terminate early; it is not possible to push the original box. Then replace `toVisit` with `toVisitNext`.
>
> Termination: Terminate when `toVisitNext` is empty.

Then, to actually update the box positions, we can work backwards from the list of `toVisit`s
