'''ninth of december'''

def solution(moves, rope_len):
    '''solution'''
    places_visited = [(0,0)]
    links = [[0,0]]*rope_len
    for move in moves:
        moving = [True]*rope_len
        target_pos = (links[0][0] + move[0], links[0][1] + move[1])
        while any(moving):
            for i in range(rope_len):
                link_pos = links[i].copy()
                if not i:
                    parent_pos = target_pos
                else:
                    parent_pos = links[i-1].copy()

                ht_dist = (parent_pos[0] - link_pos[0], parent_pos[1] - link_pos[1])

                if (abs(ht_dist[0]) > 1 or abs(ht_dist[1]) > 1):
                    moving[i] = True
                    link_pos[0] += sign(ht_dist[0])
                    link_pos[1] += sign(ht_dist[1])
                    if i == rope_len - 1:
                        places_visited.append(tuple(link_pos))
                    links[i] = link_pos.copy()
                else:
                    moving[i] = False

    return len(set(places_visited))

def get_moves(move):
    '''get moves'''
    direction, distance = move
    if direction == 'L':
        return  (-1*int(distance) - 1,0)
    if direction == 'R':
        return (int(distance) + 1,0)
    if direction == 'U':
        return (0,int(distance) + 1)
    return (0,-1*int(distance) - 1)

def sign(inpt):
    '''calc sign'''
    if inpt == 0 :
        return 0
    return inpt//abs(inpt)

with open('./data/data_day_9.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ line.strip().split(' ') for line in file.readlines() ]

all_moves = list(map(get_moves, lines))
answer_one = solution(all_moves, 2)
answer_two = solution(all_moves, 10)

print(f'total points: {answer_one}')
print(f'total points: {answer_two}')
