# Day 25

Merry Christmas!

## Part 1

This was a simple problem, just parse the keys and locks into their 5 heights do a pairwise comparison to see if they fit. Only 500 total keys/locks, so $O(n^2)$ is still acceptable in my book, and that was enough to get the star.

I then received a helpful hint that I actually didn't need to parse the heights at all. Looking at the input, it is suspiciously (or auspiciously) small, infact only 6-by-6 grids.

This meant that I could encode each grid as a 36-bit integer, and instead of comparing the 5 heights, I could just ensure that `a ^ b == 0`, i.e. there's no overlap in whether a bit is set or not.

## Part 2
