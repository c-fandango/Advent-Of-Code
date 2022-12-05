'''fifth of december'''

def solution(stacks, moves, part):
    '''solution'''
    answer = ''
    for move in moves:
        qty, frm, to = move
        frm -= 1
        to -= 1
        moving_stacks = stacks[frm][-qty:]
        stacks[frm] = stacks[frm][:-qty]
        if part == 1:
            stacks[to] += moving_stacks[::-1]
        else:
            stacks[to] += moving_stacks

    for stack in stacks:
        answer += stack[-1]
    return answer

def get_stacks():
    '''gets stacks'''
    with open('./data/data_day_5_stacks.txt', 'r', encoding = 'UTF-8') as file:
        lines = [ line for line in file.readlines() if '[' in line ]
    rows = []
    for line in lines:
        rows.append([line[i:i+3] for i in range(0, len(line),4)])
    rows.reverse()

    stacks = ['']*len(rows[0])

    for row in rows:
        for index, place in enumerate(row):
            if place.strip():
                stacks[index] += place.strip('[').strip(']')
    return stacks

def get_moves():
    '''gets moves'''
    with open('./data/data_day_5_moves.txt', 'r', encoding = 'UTF-8') as file:
        lines = [ line.strip().split(' ') for line in file.readlines() ]
    moves = []
    for line in lines:

        moves.append([ int(move) for move in line if move not in ('from', 'to', 'move') ])
    return moves

stacks_lst = get_stacks()
moves_lst = get_moves()
print(solution(stacks_lst, moves_lst, 2))
#answer_one = sum(map(part_one, lines))
#
#print(f'total points: {answer_one}')
