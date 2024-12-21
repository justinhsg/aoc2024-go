# Day 14

## Part 1
This was easy enough to do, each robot moves independently of the other robots, so we can just take `p * v (100) mod m`, where `m` is the width / height component, and then calculate the safety rating this way

## Part 2

This was a real head-scratcher; what exactly is the christmas tree we should be looking for?

My first observation was that the search space was actually quite small (relatively); the patterns are guaranteed to repeat every 101 * 103 = 10403 seconds. So I quickly built a brute force solution that printed out all 10403 patterns and manually scanned for the Christmas Tree to get the star.

However, I was not yet satisfied, and so I decided to code a simple metric (consecutive adjacent robots), where if there are more than 10 robots in a row, that would likely be the part of the pattern of the christmas tree.

Searching online, it appears that calculating the metric in Part 1 across all 10403 states would reveal the required state having an outlier value. Still, it is quite a leap of faith to take, but makes sense in hindsight.
