# Day 16

## Part 1
This was a simple application of Dijkstra's algorithm, where the state includes the row, column and the direction faced. The neat trick to reduce the search space is to only consider turning if facing that direction would reveal a path to move forward in, and to do both operations at once (adding `1001` to the distance)

## Part 2

Since Dijkstra would visit states in increasing priority (distance), we can keep a separate `fromState` (4D) array, where `fromState[x][y][d]` is the list of states that would lead into the state `(x,y,d)` with the least distance. Then, in Dijkstra, empty `fromState[x][y][d]` if a new minimum is found, and append to `fromState[x][y][d]` if the distance is the same as the minimum.

Then finding all optimal paths is a simple as iterating backwards through `fromState` array.
