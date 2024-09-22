usernames = {
    "lydia": "lydiahallie",
    "sarah": "sarah123",
    "max": "max_",
    "joe": "joejoe"
}

print(usernames["joe"])
print(usernames.keys())

for key in usernames.keys():
    print(key + " - " + usernames[key])


print(usernames.values())
print(usernames.items())

usernames["max"] = "max123"
usernames.update({"chloe": "chloe123"})
print(usernames)

del usernames["max"]
print(usernames)

usernames.clear()