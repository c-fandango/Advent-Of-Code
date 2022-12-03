'''first of december'''

def part_one(input_list):
    '''part one'''
    roll_sum = 0
    max_cal = 0
    for line in input_list:
        if line:
            roll_sum += int(line)
            continue
        if roll_sum > max_cal:
            max_cal = roll_sum
        roll_sum = 0
    print(f'max calories: {max_cal}')

def part_two(input_list):
    '''part two'''
    roll_sum = 0
    max_elves = [0,0,0]

    for line in input_list:
        if line:
            roll_sum += int(line)
            continue
        max_elves.sort()
        if roll_sum > max_elves[0]:
            max_elves[0] = roll_sum
        roll_sum = 0
    print(f'sum of max three elves calories: {sum(max_elves)}')

with open('./data/data_day_1.txt', 'r', encoding = 'UTF-8') as file:
    lines = [line.strip() for line in file.readlines()]

part_one(lines)
