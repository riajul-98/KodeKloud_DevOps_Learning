countries = ["USA", "Canada", "Spain"]

# Adding item to the end of a list
countries.append("UK")
print(countries)

# Inserting items at a certain index
countries.insert(2, "Singapore")
print(countries)

# Swapping values in a list
countries[0], countries[1] = countries[1], countries[0]


# Sorting a list (modifies original array)
ages = [56, 72, 24, 46]
ages.sort()
print(ages)

# Reversing items in a list
ages = [56, 72, 24, 46]
ages.reverse()
print(ages)