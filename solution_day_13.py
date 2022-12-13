'''thirtinth of december'''
def is_ordered(item_one, item_two):
    '''checks if ordered'''
    if isinstance(item_one, int) and isinstance(item_two, int):
        if item_one == item_two:
            return None
        return item_one < item_two

    if isinstance(item_one, int):
        item_one = [item_one]
    elif isinstance(item_two, int):
        item_two = [item_two]

    for elem_one, elem_two in zip(item_one, item_two):
        result = is_ordered(elem_one, elem_two)
        if isinstance(result, bool):
            return result

    if len(item_one) != len(item_two):
        return len(item_one) < len(item_two)

    return None

def bubble_sort(unsorted):
    '''bubble sorts'''
    swaps = True
    while swaps:
        swaps = False
        for i in range(len(unsorted) - 1):
            first, second = unsorted[i:i+2]
            if not is_ordered(first, second):
                unsorted[i], unsorted[i+1] = unsorted[i+1], unsorted[i]
                swaps = True
    return unsorted

def part_one(packets):
    '''part one'''
    total = 0
    for i in range(0, len(packets), 2):
        first, second = packets[i:i+2]
        if is_ordered(first, second):
            total += i//2 + 1

    return total


def part_two(packets):
    '''part two'''
    decoder_key = 1
    packets += [[[2]],[[6]]]

    bubble_sort(packets)

    for index, packet in enumerate(packets, 1):
        if packet in ([[2]], [[6]]):
            decoder_key *= index
    return decoder_key


with open('./data/data_day_13.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ eval(lists) for line in file.read().split('\n\n') for lists in line.split() ]

answer_one = part_one(lines)
answer_two = part_two(lines)

print(f'total points: {answer_one}')
print(f'total points: {answer_two}')
