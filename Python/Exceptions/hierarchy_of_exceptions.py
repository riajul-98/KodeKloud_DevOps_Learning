try:
    x = int(input("Enter a number: "))
    y = 1 / x
    print(y)
except ArithmeticError:         # Can be used instead of zero division error as the zero division error falls under arithmetic errors.
    print("Calculation Failed.")
except:
    print("Something went wrong.")
print("All done!")


def calculate_user_input():
    raise ZeroDivisionError
try:
    calculate_user_input()
except ZeroDivisionError:
    print("You cannot divide by zero.")
except:
    print("Something else went wrong.")



def calculate_user_input():
    try:
        x = int(input("Enter a number: "))
        y = 1 / x
        print(y)
    except:
        print("Something else went wrong")
        raise

    return None

try:
    calculate_user_input()
except ZeroDivisionError:
    print("You cannot divide by zero.")
except:
    print("Something else went wrong")
