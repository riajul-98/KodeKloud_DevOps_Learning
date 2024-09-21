# & = Conjunction. Only returns 1 if both bits are one
# | = Disjunction. Returns 1 if either of the bits are 1 or if both are 1
# ~ = Negation. Returns 1 if the bit is 0
# ^ = Exclusive. Returns 1 if only one of the bits are 1.

print(15 & 22)
print(15 | 22)
print(22 ~)
print(15 ^ 22)


print(22 >> 1) # Moves all bits to the right by 1
print(22 << 1) # Moves all bits to the left by 1