# Day 24

## Part 1

At first glance I thought part 2 was going to be trying to figure out what the bit operations actually do, so I decided to just focus on part 1. Just implementing as intended, using a top-down DP to speed up the answer.

## Part 2

Good news was that this wasn't a 'reverse engineer' the circuit problem, but it's close. I immediately recalled the circuit for a full adder, and it took a while to recall the boolean equations for the outputs:

$$
\begin{align}
x_i \space \oplus \space y_i &= t \nonumber\\
t \space \oplus \space c_{i-1} &= z_i \nonumber\\
t \space \cdot \space c_{i-1} &= u \nonumber\\
x_i \space \cdot  \space y_i &= v \nonumber\\
u \space +  \space v &= c_i \nonumber\\
\end{align}
$$

My first instinct was to look at the $z_i$ outputs. If they aren't outputs of XOR operations, then those definitely need to be swapped. I then looked at the remaining XOR equations and identified outputs that should have been swapped.

This gave me 6 wires to swap, and hence there are 2 more.

I updated my code to swap those wires when parsing the input, and then printed out the equations grouped 5 ways:

1) XOR equations with `x` and `y` inputs
2) XOR equations with `z` outputs
3) AND equations with `x` and `y` inputs
4) AND equations without `x` and `y` inputs
5) OR equations

It was a messy, and required quite a lot of restarts and `fmt.Println` statements to tease out the last pair, and finally got it by iterating upwards and tracking the `carry` wires.

After getting the answer, I did a full rewrite of the code to better get out the answer. I could verify if each of the 5 equations 'types' had a valid output:

$$
\begin{align}
x_i \space \oplus \space y_i &= t \nonumber\\
t \space \oplus \space c_{i-1} &= z_i \nonumber\\
t \space \cdot \space c_{i-1} &= u \nonumber\\
x_i \space \cdot  \space y_i &= v \nonumber\\
u \space +  \space v &= c_i \nonumber\\
\end{align}
$$

1. $t$ must be an input to another XOR equation and another AND equation
2. $z_i$ must start with `z`
3. $u$ must be an input to an OR equation
4. $v$ must be an input to an OR equation
5. $c_i$ must be an input to a XOR equation and another AND equation