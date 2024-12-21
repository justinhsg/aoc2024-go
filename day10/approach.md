# Day 10

## Part 1

Simple grid-based BFS from each `0` was more than enough to get this part. Could possibly have optimised further by memoising the score on the way 'up' to `9`, but that didn't seem necessary and the backtracing would be tedious

## Part 2

Took an alternative approach, where I calculated the 'ratings' of each position recursively from 9 to 0, where the rating of each position is the sum of ratings of adjacent positions that are 1 higher.
Small optimisation in storing all positions grouped by height to speed up bottom-up DP.
