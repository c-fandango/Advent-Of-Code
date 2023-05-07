'''third of december'''
def part_one(sack):
    '''part one'''
    alph = 'abcdefghijklmnopqrstuvwxyz'
    length = len(sack)//2
    first_comp = sack[:length]
    second_comp = sack[length:]
    intersec = [ i for i in first_comp if i in second_comp ]

    points = alph.index(intersec[0].lower()) + 1
    if intersec[0].isupper():
        points += 26

    return points

def part_two(sack):
    '''part two'''
    alph = 'abcdefghijklmnopqrstuvwxyz'
    #intersec = list(set(sack[0]) & set(sack[1]) & set(sack[2]))
    intersec = list(set(sack[0]).intersection(set(sack[1])).intersection(set(sack[2])))

    points = alph.index(intersec[0].lower()) + 1
    if intersec[0].isupper():
        points += 26

    return points

with open('./data/data_day_3.txt', 'r', encoding = 'UTF-8') as file:
    lines_one = [ line.strip() for line in file.readlines() ]

# better thing is this [input[i:i+n] for i in range(0, len(input), n)]
lines_two = []
elem = []
for index, line in enumerate(lines_one):
    elem.append(line)
    if index%3 ==2:
        lines_two.append(elem)
        elem = []


answer_one = sum(map(part_one, lines_one))
answer_two = sum(map(part_two, lines_two))

print(f'total points: {answer_two}')
