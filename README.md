# interlaced-spiral-cipher

let’s begin by looking at an encoding pattern.

Below is an encryption pattern :

| 0 | 4 | 8 | 1 |
| 11 | 12 | 13 | 5 |
| 7 | 15 | 14 | 9 |
| 3 | 10 | 6 | 2 |

So, if we were to encrypt the string “ABCDEFGHIJKLMNOP” and we place it in the above encryption grid, we start with index 0 in the upper left and put the letter A in there, we then find the number 1 and put the letter B in there… until we get to 15 where we put the final (16th letter) P. The pattern will now look like this :

| A | E | I | B |
| L | M | N | F |
| H | P | O | J |
| D | K | G | C |

After the grid is populated with the string to be encrypted, we just concatenate each row starting from the top. This would yield the encrypted string: AEIBLMNFHPOJDKGC

https://github.com/mevise-techblog/interlaced-spiral-cipher
https://www.codewars.com/kata/interlaced-spiral-cipher/train/go

- [x] Step 1
- [x] Step 2
- [x] Step 3
- [ ] Write Unit Tests
- [ ] Complete Decode Cypher
- [ ] Refactor
- [ ] Test at Codewars.com
- [ ] Create UI 
