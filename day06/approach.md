# Day 6

Note: this was done as catch-up on Day 8

## Part 1

Simple graph traversal problem via BFS. Obstacles were a simple `map[IntPair]bool`

## Part 2

Observation:

The only positions that have a possibility of getting a loop are those on the path found in Part 1, so just restart BFS but with an extra entry in the obstacles map / set.

Optimisation 1:

Instead of storing the entire path in the `visited` array, only store the points (and direction) where you 'collide' with the obstacle.

Optimisation 2:
We can immediately find these collision points by searching for next immediate obstruction.

Now, use two `map[int][]int` (though `[][]int` probably works in hindsight), one to store the obstacles by row and another by column. So for an obstacle at row $5$, column $13$, `obstaclesByRow[5]` would have element $13$ and `obstaclesByColumn[13]` would have element $5$.

Then, you can just do a linear search from either direction to get the collsion point. i.e. if you are on row $5$, moving from east to west, iterate `obstaclesByRow[5]` in reverse order until you find an element less than $5$. If no such element found, then you will exit the grid (there is no loop). If an element $x$ exists, then the collision would be at row $5$, column $x+1$.

This optimisation sped things up quite considerably, as now we don't iterate one step at a time, and instead 'fast-forward' to the next collision point.
