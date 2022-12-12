'''twelth of december'''
def part_one(grid, start, end):
    '''part one'''
    grid = set_start(grid, start)
    nodes = djikstra(grid, test_func = not_too_high)
    end_node = [ node['dist'] for node in nodes if node['pos'] == end ]
    return end_node[0]

def part_two(grid, start):
    '''part two'''
    grid = set_start(grid, start)
    nodes = djikstra(grid, test_func = not_too_low)
    low_nodes = [ node['dist'] for node in nodes if not node['alt'] ]
    return min(low_nodes)

def set_start(grid, start):
    '''sets stert of djikstra'''
    for i, row in enumerate(grid):
        for j, node in enumerate(row):
            if node['pos'] == start:
                grid[i][j]['dist'] = 0
    return grid

def not_too_high( inpt_one, inpt_two):
    '''checks if step is too high'''
    return inpt_one >= inpt_two - 1

def not_too_low( inpt_one, inpt_two):
    '''checks if step is too low'''
    return inpt_one <= inpt_two + 1

def djikstra(grid, test_func):
    '''implementation of dkiksta algorithm'''
    ibound = len(grid)
    jbound = len(grid[0])
    unvisited = [ grid[i][j] for i in range(ibound) for j in range(jbound) ]
    visited = []
    while unvisited:
        curr_node = min(unvisited, key = lambda x: x['dist'])
        ipos = curr_node['pos'][0]
        jpos = curr_node['pos'][1]

        if ipos != 0:
            upp = grid[ipos-1][jpos]
            if upp not in visited and test_func(curr_node['alt'], upp['alt']):
                upp['dist'] = curr_node['dist'] + 1 if curr_node['dist'] + 1 < upp['dist'] else upp['dist']

        if ipos != ibound -1:
            down = grid[ipos+1][jpos]
            if down not in visited and test_func(curr_node['alt'], down['alt']):
                down['dist'] = curr_node['dist'] + 1 if curr_node['dist'] + 1 < down['dist'] else down['dist']

        if jpos != 0:
            left = grid[ipos][jpos-1]
            if left not in visited and test_func(curr_node['alt'], left['alt']):
                left['dist'] = curr_node['dist'] + 1 if curr_node['dist'] + 1 < left['dist'] else left['dist']

        if jpos != jbound -1:
            right = grid[ipos][jpos+1]
            if right not in visited and test_func(curr_node['alt'], right['alt']):
                right['dist'] = curr_node['dist'] + 1 if curr_node['dist'] + 1 < right['dist'] else right['dist']

        unvisited.remove(curr_node)
        visited.append(curr_node)
    return visited

def get_grid(grid):
    '''generate grid'''
    alph = 'abcdefghijklmnopqrstuvwxyz'
    for i, row in enumerate(lines):
        for j, letter in enumerate(row):
            grid[i][j] = { 'pos': (i, j), 'alt': alph.find(letter), 'dist': float('inf')}
            if letter == 'S':
                start_pos = (i, j)
                grid[i][j] = {'pos': (i, j), 'alt': 0, 'dist': float('inf')}
            elif letter == 'E':
                end_pos = (i, j)
                grid[i][j] = {'pos': end_pos, 'alt': 25, 'dist': float('inf')}

    return (grid, start_pos, end_pos)

with open('./data/data_day_12.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ list(line.strip()) for line in file.readlines() ]

data_grid, beg, fin  = get_grid(lines)

#distance_one = part_one(data_grid, beg, fin)
#print(f'min dist: {distance_one}')

distance_two= part_two(data_grid, fin)
print(f'min dist: {distance_two}')
