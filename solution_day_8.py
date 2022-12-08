'''eigth of december'''
def part_one(grid):
    '''part one'''
    visible_count = (len(grid[0]) + len(grid) -2)*2
    transposed_grid = [[grid[i][j] for i in range(len(grid))] for j in range(len(grid[0]))]
    for i in range(1,len(grid)-1):
        for j in range(1,len(grid[0])-1):
            tree_height = grid[i][j]
            trees_left, trees_right = grid[i][:j], grid[i][j+1:]
            trees_up, trees_down = transposed_grid[j][:i], transposed_grid[j][i+1:]
            if tree_height > max(trees_left) or tree_height > max(trees_right) or tree_height > max(trees_up) or tree_height > max(trees_down):
                visible_count += 1
    return visible_count

def part_two(grid):
    '''part two'''
    ss_max = 0
    transposed_grid = [[grid[i][j] for i in range(len(grid))] for j in range(len(grid[0]))]
    for i in range(1,len(grid)-1):
        for j in range(1,len(grid[0])-1):
            tree_height = grid[i][j]
            trees_left, trees_right = grid[i][j-1::-1], grid[i][j+1:]
            trees_up, trees_down = transposed_grid[j][i-1::-1], transposed_grid[j][i+1:]

            left_score = ss_score(tree_height, trees_left)
            right_score= ss_score(tree_height, trees_right)
            up_score = ss_score(tree_height, trees_up)
            down_score = ss_score(tree_height, trees_down)

            ss_tree = left_score * right_score * up_score * down_score
            if ss_tree > ss_max:
                ss_max = ss_tree

    return ss_max

def ss_score(target, row):
    '''calculates ss_score'''
    max_score = len(row)
    for distance, height in enumerate(row, 1):
        if target <= height:
            return distance
    return max_score


with open('./data/data_day_8.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ list(line.strip()) for line in file.readlines() ]

for line in lines:
    map(int,line)

answer_one = part_one(lines)
answer_two= part_two(lines)

print(f'trees: {answer_two}')
