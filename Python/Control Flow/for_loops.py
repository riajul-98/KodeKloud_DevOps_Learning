for i in range(10):
    print("i is: ", i)


print("\n" * 3)
for i in range(5):
    if i == 2:
        break
    print("i is: ", i)

print("\n" * 3)
for i in range(5):
    if i == 2:
        continue
    print("i is: ", i)