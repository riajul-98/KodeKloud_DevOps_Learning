countries = ["USA", "Canada", "Spain"]

# USA will have index 0
print(countries[0])
# Canada will have index 1
print(countries[1])
# Spain will have index 2
print(countries[2])

# Changing USA to UK
countries[0] = "UK"
print(countries)

# Checking number of items in a list
print(len(countries))

# Deleting item in a list
del countries[1]
print(countries)

# Using negative indexes
print(countries[-1])
print(countries[-2])
print(countries[-3])