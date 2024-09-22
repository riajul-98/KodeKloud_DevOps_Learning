num = 100

def input_number():
    global own_num 
    own_num = 50
    result = int(input("Enter a number: ")) * own_num
    return result

print(own_num)