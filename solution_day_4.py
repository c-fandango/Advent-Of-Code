'''fourth of december'''
def part_one(pair):
    '''part one'''
    pair = [[int(i) for i in elv.split('-')] for elv in pair]
    elv_one = set(range(pair[0][0], pair[0][1] +1))
    elv_two = set(range(pair[1][0], pair[1][1] +1))

    intersec = elv_one.intersection(elv_two)
    if intersec in (elv_one, elv_two):
        return 1

    return 0

def part_two(pair):
    '''part two'''
    pair = [[int(i) for i in elv.split('-')] for elv in pair]
    elv_one = set(range(pair[0][0], pair[0][1] +1))
    elv_two = set(range(pair[1][0], pair[1][1] +1))

    intersec = elv_one.intersection(elv_two)
    if intersec:
        return 1

    return 0

with open('./data/data_day_4.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ line.strip().split(',') for line in file.readlines() ]

#answer_one = sum(map(part_one, lines))
answer_two = sum(map(part_two, lines))

print(f'total points: {answer_two}')
