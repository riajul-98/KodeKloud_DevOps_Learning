def input_number(num1, num2):
    return int(input("Enter a number: ")) * num1 - num2

# default value
# def input_number(num = 10):
#     return int(input("Enter a number: ")) * num1 - num2


input1 = input_number(10, 20)
input2 = input_number(num2=10, num2=20)
print(input1)
print(input2)