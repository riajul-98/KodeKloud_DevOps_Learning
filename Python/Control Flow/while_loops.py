secret_number = 3
guess = int(input("Guess a number: "))
while guess != secret_number:
    guess = int(input("Guess a number: "))
else:
    print("Congratulations, you got it!")