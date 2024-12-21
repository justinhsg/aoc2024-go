# Day 17

## Part 1
After reading the prompt and also from past AoC experience, I had a sinking suspicion that either:

* the actual program given in the input would loop for quite a long time
* we would need to give a new register A value to give a certain result in Part 2

But first, I just brute-forced Part 1 and it gave the answer in a short amount of time.

## Part 2

This confirmed my suspicion, so I had to dive into the actual input and decode the instructions. The following is the pseudocode for what my input did:

```text
B = A % 8
B = B ^ 2
C = A >> B
B = B ^ C
B = B ^ 3
output B
A = A / 8
loop if A != 0
```

Two observations:

1. From the last 2 lines, the program looped by dividing A by 8 every time until A = 0
2. From the first line, there are only 8 possible values that are considered each iteration that actually matters.

With this, I could build up the required output from back to front. For a program of length `n`:

```pseudocode
A = 0
for x in (n-1..0):
    target = program[x]
    outputList = runProgram(A * 8 + x)
    if outputList[0] == target:
        A = A * 8 + x
        break
```

I noticed that the result didn't manage to get the full output; in fact it stopped working past a certain point. I realised that this was because there could be more than one possible value of `x` that could meet the target. I switch `A` from a list of possible `A`s and that got me the second star.