# When python comes across errors, it expects something to take care of it. If there is nothing to take care of it, an error message
# will be displayed.

try:
    name = "Lydia"
    print("My name is " + naem)
except:
    print("Something went wrong")

print("All done!")



try:
    x = int(input("Enter a number: "))
    y = 1 / x
    print(y)
except ZeroDivisionError:
    print("You cannot divide by zero.")
except:
    print("Something went wrong.")
print("All done!")