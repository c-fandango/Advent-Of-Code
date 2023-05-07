'''tenth of december'''
def part_one(commands):
    '''part one'''
    total = 0
    cycles = [20 + 40*i for i in range(6)]
    roll_sum = [1,1,commands[0] + 1] + [ sum(commands[:i+1]) + 1 for i in range(1, len(commands)) ]
    for cycle in cycles:
        total += cycle*roll_sum[cycle-1]

    return total

def part_two(cycles):
    '''part two'''
    roll_sum = [ 1,1,cycles[0] + 1] + [ sum(cycles[:i+1]) + 1 for i in range(1, len(cycles)) ]
    for i in range(6):
        row = ''
        for j in range(40):
            if roll_sum[i*40 + j] - 2 <  j  < roll_sum[i*40 + j] + 2:
                row += '#'
            else:
                row += '.'
        print(row)

with open('./data/data_day_10.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ line[5:].strip() for line in file.readlines() ]
clean_lines = []
for line in lines:
    clean_lines = clean_lines + [int(line),0] if line else clean_lines + [0]

answer_one = part_one(clean_lines)
part_two(clean_lines)

print(f'total points: {answer_one}')
