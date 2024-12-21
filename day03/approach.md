# Day 3

This was simple string scanning + regex to get both parts right

Note: this was done as catch-up on Day 7

## Part 1

A simple `FindAllString` for the regex pattern `mul\(\d+,\d+\)` worked well enough

Had a small mistake where I didn't realise the actual input was multi-line

## Part 2

Updated the regex pattern to `mul\(\d+,\d+\)|do\(\)|don't\(\)`, and then adding an `isEnabled` toggle to keep track of whether perform the `mul()` or not.

Had another small mistake where I had assumed that each line was a new program and would start enabled.