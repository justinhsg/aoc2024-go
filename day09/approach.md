# Day 9

## Part 1

It is a little janky, but I came up with '2 pointer solution', starting at both ends of the input, moving either the left or right pointer when needed, and terminating when the pointers cross. I keep a 'bank' of spaces were if there are no more spaces to fill, I advance the left pointer (by 2) and add the more spaces to fill. Similarly, I keep a running count of blocks to move (from the right), and if there are no more blocks to move, I advance the right pointer (by 2) and add the next set of blocks to the count.

The checksum is calculated as the left pointer advances

## Part 2

I took a completely different approach to this part, this time keeping a list of left-most space positions, grouped by space size. Then, I could take the last file block and, find the left-most space that can fit the file block, and remove that entry from the list of space positions (and adding to a different list if there is leftover space (and sorting it))

I could have optimised this for insertion time of each space by using a heap / priority queue, but the linear search approach seems to work fast enough.
