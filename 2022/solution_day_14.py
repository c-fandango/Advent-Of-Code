'''fourteenth of december'''
def solution(cave, sand_limit_index):
    '''solution'''
    while True:
        place_sand(cave)
        if sum(cave[sand_limit_index]):
            break

    return sum( point for row in cave for point in row if point == 1 )


def place_sand(cave):
    '''falls sand'''
    sandx, sandy = 500, 0
    while True :
        if grid[sandy+1][sandx]:
            if grid[sandy+1][sandx -1]:
                if grid[sandy+1][sandx+1]:
                    cave[sandy][sandx] = 1
                    return cave
                sandx += 1
            else:
                sandx -= 1
        sandy += 1


def sign(inpt):
    '''calc sign'''
    return inpt//abs(inpt)


def add_rock(cave, rock):
    '''generate grid of cave'''
    for vein in rock:
        for i in range(len(vein) - 1):
            curr, nxt = vein[i], vein[i+1]
            if curr[0] == nxt[0]:
                sgn = sign(nxt[1] - curr[1])
                for j in range(curr[1], nxt[1] + sgn, sgn):
                    cave[j][curr[0]] = -1
            else:
                sgn = sign(nxt[0] - curr[0])
                for j in range(curr[0], nxt[0] + sgn, sgn):
                    cave[curr[1]][j] = -1

    cave += [[ -1 for i in range(1000) ]] #only makes sense for second part but also works for first part
    return cave


with open('./data/data_day_14.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ [ eval(splt) for splt in line.strip().split(' -> ') ] for line in file.readlines() ]

ceiling = max( move[1] for line in lines for move in line )
grid = [ [0 for j in range(1000)] for i in range(ceiling + 2) ]

add_rock(grid, lines)
answer_one = solution(grid, -2) - 1
answer_two = solution(grid, 0)

print(f'total points: {answer_one}')
print(f'total points: {answer_two}')
