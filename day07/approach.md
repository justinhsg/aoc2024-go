# Day 7

Note: this was done as catch-up on Day 8

## Part 1

Immediately saw the trick to iterate through the operands in reverse order and recursively find if the target is achievable. In pseudocode:

```pseudocode
def f(target, operands) {
    if len(operands) == 1:
        return target == operands
    
    lastOperand := operands[-1]
    possible := false
    if target % lastOperand == 0 {
        possible = f(target / lastOperand, operands[:-1])
    }
    if !possible && target > lastOperand {
        possible = f(target - lastOperand, operands[:-1])
    }
}
```

## Part 2

Mistake 1:

Thought that the `||` operator concatenates the adjacent operands, so `5 * 6 || 11` would be `5 * 611`

Once that mistake was discovered, I realised the proper solution needed a `NDigits` to calculate the number of digits (rather than casting the integer to string and back) to split the number
