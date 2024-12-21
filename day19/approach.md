# Day 19

Definitely over-thought this problem after rushing to complete day 18

## Part 1

This clearly was a DP problem, and I spent a while trying to think of the subproblem until I realised it was actually quite simple.

Took a bottom up approach where the function `possible(x)` is true iff it is possible to create the substring up to and excluding the x-th character.

Then `possible(x+len(p)) = possible(x) && s[x:x+len(p)]==p` for  all patterns `p`.

Two optimisations were done here:

1. I skipped ahead if `possible(x)` was already false.

2. I grouped all the patterns by their first characters to reduce the search space on each iteration.

## Part 2

I was hoping that this wasn't a case where I had to only use each pattern `p` once (which would transform this into a weird 0-1 Knapsack), and luckily part 2 just needed to turn `possible(x)` into `ways(x)`.

This was a simple modification to my bottom-up dp, turning (implicit) boolean ORs into addition.