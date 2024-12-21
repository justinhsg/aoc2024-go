# Day 13

# Part 1

At first thought this was a knapsack problem, but then realised this was in fact a simple linear system, with 2 equations and 2 unknowns. Given an input:

```text
Button A: X+(Ax), Y+(Ay)
Button B: X+(Bx), Y+(By)
Prize: X=(Cx), Y=(Cy)
```

$$
a \cdot A_x + b \cdot B_x = C_x \\
a \cdot A_y + b \cdot B_y = C_y \\

\begin{pmatrix}
A_x & B_x \\
A_y & B_y
\end{pmatrix}
\begin{pmatrix}
a \\ b
\end{pmatrix}
=
\begin{pmatrix}
C_x \\ C_y
\end{pmatrix}
$$

Solving this by Cramer's Rule was simple enough, I decided to ignore the case where the equations were linearly dependent, which would require a bit more complex handling, and luckily none of the inputs had this case.

## Part 2

Probably the easiest modification by far; just adding `10000000000000` to the `C` values and resolving the equations gave the required answer.