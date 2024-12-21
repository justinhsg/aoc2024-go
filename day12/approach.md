# Day 12

## Part 1

Another simple BFS / floodfill problem, with the slight modification as follows:
* area is equal to the number of positions visited
* each position adds 4 to the perimeter, but when considering adjacent positions, subtract 1 for every position

## Part 2

Had many false starts:

1. Tried to take the top left grid, and traverse around the perimeter, keeping track of the number of 'turns'. This couldn't solve the case where there is an 'internal' perimeter

2. Tried keep track of 4 lists for each cardinal direction, where each list held the positions which had a border in that direction. This list could be populated during the BFS. This missed out the edges on the borders of the grid.

With option 2, I managed to get a star, and then with some online discussion, realised I could count the corners of each shape, which is more straightforward. For each position X, consider the 2 adjacent positions and the diagonal tile between them (e.g. North, East, and NorthEast), and determine if they are inside (I) or outside the region (O).

Case 1: Adjacent positions are inside, diagonal is outside:

```text
I | O
  +--
X   I
```

Case 2: Adjacent positions are outside, diagonal is outside:

```text
O   O
--+
X | O
```

Case 3: Adjacent positions are outside, diagonal is inside. This is a special case:

```text
O | I
--+--
X | O
```

```text
III
IOI
IXO
```