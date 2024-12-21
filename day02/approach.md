# Day 2

This was surprisingly difficult, probably due to misreading Part 2 and overcomplicating things.

Note: this was done as catch-up on Day 7

## Part 1

Started with a simple check on the first 2 values to get the 'direction', and then iterate through to see if any increase/decrease is 'unsafe'.

## Part 2

Mistake 1: I thought that it would tolerate one 'unsafe' delta, so I just ignored the first 'unsafe' increase/decrease, accounting for the case where the first delta was the 'unsafe' one.

After realising that the problem meant tolerating one 'unsafe' level, so I had to rethink my solution.

The naive solution would be to try removing a level each time and checking if the list now passed validation, but that would be $O(n^2)$

I then reworked my part 1 solution to use a difference array (storing the differences between each level), and found the first 'difference' that broke the rules, and then 'merged' the adjacent differences. e.g.:

$$
\text{array} = [1, \space 3, \space 9, \space 5] \\

\text{diffArray} = [3-1, \space 9-3, \space 5-9]\\
=[2, \space 6, \space-4]
\\
\text{Suppose 6 is the outlier; merge  diffArray:} \\
\\
\text{newDiffArray} = [2, \space 6 + (-4)]\\
= [2, \space 2]\\
$$

Mistake 2: This assumes that you merge the outlier element and the element after (i.e. $9$ is the outlier in the above example); there could be the case e.g. if `array` = $[1, 4, 2, 3]$,  `diffArray` = $[3,-2,1]$ which would attempt to merge $-2$  and $1$ from `diffArray`, when in fact it should merge $-2$ and $1$.

This meant that if an outlier was found I now have to check 2 cases; `mergeBefore` and `mergeAfter`, and also include the edge-cases for the first and last elements as usual

On hindsight, perhaps the naive approach would be easier to implement; plus the size of each array is relatively small (max of 8).
