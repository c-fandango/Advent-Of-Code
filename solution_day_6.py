'''sixth of december'''
def solution(message, length):
    '''solution'''
    for i in range(length, len(message)):
        marker = message[i-length:i]
        if len(set(marker)) == length:
            return i
    return -1



with open('./data/data_day_6.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ line.strip() for line in file.readlines() ]

answer_one = solution(lines[0], 4)
answer_two= solution(lines[0], 14)

print(f'characters processed: {answer_two}')
