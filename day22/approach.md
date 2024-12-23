# Day 22

## Part 1

One instinct I had early on (from the divison, multiplication, modulo), was that this problem would be about number theory and the Chinese Remainder Theorem.

Then, when realising that all the operands turned out to be all powers of 2, this led me to believe that the intended solution would require bitwise operations instead.

The rest of the problem was quite simple: multiplying by 64 was a left shift by 6, dividing by 32 was a right shift by 5, the modulo was just a bitwise AND by `16777216-1`

## Part 2

Since the sequence is mostly random, there probably isn't a closed form to get the answer, so I decided to use brute force:

For each item, I had a array of the price and another array of the differences between the prices. Then, for each possible sequence (all combinations of -9 to +9 across each of the 4 positions), I scanned through the difference array to find the first match, got the corresponding price, and then continued on to the next item.

This was very slow (~10 minutes), but gave the correct answer. In the mean time, I proceeded to think on how to optimise it.

My first optimisation required a change in the brute force. Instead of iterating through the difference array $19^4$ times (for each sequence), I could easily just keep a `map[Tuple4]int`, for each item, and populated it when I constructed the difference array. Then, I didn't need to repeat the linear search through the difference array for each combination of `Tuple4` seen thus far.

This got the brute force down to sub-second (on my machine), but I wanted to push it a bit further. I realised that I could actually immediately add the found price to a 'global' `map[Tuple4]int`, and just take the maximum value of that global map get the final answer. This got my code to run in sub 700ms.

The next optimisation I did was to realise that the range of differences (-9 to 9) is very constrained, meaning I could encode the value in 5 bits. Hence instead of constructing a `Tuple4` as a key, I could use an `int` type. This brought the runtime down to about 470ms.

I then took some time to (manually, using `time.Now`) profile my solution to see where the bottleneck is. Bitwise operations, as expected, took little to no time, and the bulk of the time was spent on `map` operations. Searching a bit online, there was mention of using a fixed bit-length `int` type for map keys. This was my final optimisation, which brought the runtime down further to 350ms.

I briefly entertained the idea of storing the difference array also as an `int32`, since we only needed 4 difference values at a time, and we could just use bitshifts to remove the oldest value. This had negligible improvement in runtime, and which I point I realised I probably was too deep into the weeds and should stop.
