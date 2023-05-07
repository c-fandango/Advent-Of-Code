'''eleventh of december'''
import re

def solution(monkeys, rounds, worry_factor, mod_factor):
    '''solution'''
    for _ in range(rounds):
        for monkey_no in range(len(monkeys)):
            monkey_name = 'monkey_' + str(monkey_no)
            curr_monkey = monkeys[monkey_name]
            for old in curr_monkey['items']:
                new = eval(curr_monkey['op']) // worry_factor
                throw_to = curr_monkey['false'] if new % curr_monkey['test'] else curr_monkey['true']
                new %= mod_factor
                monkeys[throw_to]['items'].append(new)
            curr_monkey['business'] += len(curr_monkey['items'])
            curr_monkey['items'] = []
    monkey_business = sorted([monkey['business'] for monkey in monkeys.values()])
    return monkey_business[-1] * monkey_business[-2]

def mult(monkeys):
    '''multiplies every element in list'''
    prod = 1
    for monkey in monkeys.values():
        prod *= monkey['test']
    return prod

def build_monkeys(all_lines):
    '''builds monkey dictionary'''
    monkeys = {}
    for line in all_lines:
        if 'Monkey' in line:
            monkey_name = 'monkey_' + re.search('\\d+', line).group()
            monkeys[monkey_name] = {'items': [], 'op': None, 'test': None, 'true': None, 'false': None, 'business': 0}
        elif 'Starting' in line:
            items = re.findall('\\d+', line)
            monkeys[monkey_name]['items'] = [ int(item) for item in items ]
        elif 'Operation' in line:
            monkeys[monkey_name]['op'] = line.split(': new = ')[1]
        elif 'Test' in line:
            monkeys[monkey_name]['test'] = int(re.search('\\d+', line).group())
        elif 'true' in line:
            monkeys[monkey_name]['true'] = 'monkey_' + re.search('\\d+', line).group()
        elif 'false' in line:
            monkeys[monkey_name]['false'] = 'monkey_' + re.search('\\d+', line).group()
    return monkeys


with open('./data/data_day_11.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ line.strip() for line in file.readlines() ]

monkeys_dict = build_monkeys(lines)

#answer_one = solution(monkeys_dict, 20, 3, 1)
#print(f'total points: {answer_one}')

modulo_factor = mult(monkeys_dict)
answer_two = solution(monkeys_dict, 10000, 1, modulo_factor)

print(f'total points: {answer_two}')
