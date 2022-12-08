'''seventh of december'''
import random

def part_one(sizes):
    '''part one'''
    total_size = 0
    for size in sizes.values():
        if size <= 100000:
            total_size += size
    return total_size

def part_two(sizes, space_req = 30000000, space_disk = 70000000):
    '''part two'''
    del_dir_size = space_disk
    space_to_del =  space_req - (space_disk - sizes['/'])
    for size in sizes.values():
        if del_dir_size > size >= space_to_del:
            del_dir_size = size

    return del_dir_size

def get_dir_sizes(commands):
    '''gets sizes of directories'''
    all_directories = {}
    linked_directories = []
    for command in commands:
        if command == '$ cd ..':
            linked_directories = linked_directories[:-1]

        elif '$ cd' in command:
            current_directory = command[5:]
            if current_directory not in all_directories:
                all_directories[current_directory] = 0
            else:
                current_directory = current_directory + str(random.random())
                all_directories[current_directory] = 0
            linked_directories.append(current_directory)

        elif '$' not in command and command[:4] != 'dir ':
            ls_result = command.split()
            size = int(ls_result[0])
            for directory in linked_directories:
                all_directories[directory] += size

    return all_directories


with open('./data/data_day_7.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ line.strip() for line in file.readlines() ]

dir_sizes = get_dir_sizes(lines)

print(f'size: {part_two(dir_sizes)}')
