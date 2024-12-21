# Day 20

## Part 1

Misread the puzzle the first time, but turned out to be a happy accident. I first created a util function that generated vectors to explore that are $n$ steps away (Manhattan distance).

Then, I did 2 BFS traversals, one from the start and another from the end to get the distances from each grid to both start and end.

Then for each non-wall tile, I selected that to be the 'start' of the cheat; the 'end' of the cheats are each of the possible (non-wall) that are $n$ steps away (using the earlier util function)

Then, time saved is $d(start,end) - (d(start, cheatStart) + n + d(cheatEnd, end))$

The mistake I made was thinking that $n=2$ meant going through both 2 and 1 wall, when in fact (in a classic off-by-one fashion) we were expected to only consider 1 wall.

## Part 2

Because of my earlier mistake, my solution is easily extensible to consider $n$ of any amount.

Interestingly, my solution considers duplicate 'subsets' of cheats; for a path `...###...`, the cheats `...###.` and `.###.` are counted separately, which thankfully, matches the puzzle definition of a cheat (start and end points)