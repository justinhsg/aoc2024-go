# Day 11

## Part 1

Immediately noticed a cycle (2024 appeared again every other time step), but with `t=25` this didn't seem necessary. Implementing the solution as intended gave the answer easily enough.

## Part 2

Realised that while order is preserved, it does not matter for the purposes of this puzzle. So instead of keeping a slice / array for each time-step, keep a count of each unique stone number each timestep.
