'''second of december'''

def part_one(rps_round):
    '''part one'''
    for key, val in enumerate(rps_round):
        if val in ('X','A'):
            rps_round[key] = 1
        elif val in ('Y','B'):
            rps_round[key] = 2
        elif val in ('Z','C'):
            rps_round[key] = 3

    points = rps_round[1]
    if rps_round  == [1,3]:
        pass
    elif rps_round[0] < rps_round[1] or rps_round == [3,1]:
        points += 6
    elif rps_round[0] == rps_round[1]:
        points += 3

    return points

def part_two(rps_round):
    '''part two'''
    mapping = {'A': {'X': 3, 'Y': 4, 'Z': 8} , 'B': {'X': 1, 'Y': 5, 'Z': 9}, 'C': {'X': 2, 'Y': 6, 'Z': 7}}

    points = mapping[rps_round[0]][rps_round[1]]

    return points


with open('./data/data_day_2.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ part_two(line.strip().split(' ')) for line in file.readlines() ]


print(f'total points: {sum(lines)}')
