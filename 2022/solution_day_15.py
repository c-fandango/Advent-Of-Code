'''fourth of december'''
import re

def part_one(pairs, rng, ypos ):
    '''part one'''
    count = 0
    for xpos in range(-rng, rng):
        for pair in pairs:
            sensor, beacon = pair
            taxi_sb = abs(sensor[0] - beacon[0]) + abs(sensor[1] - beacon[1])
            taxi_me = abs(sensor[0] - xpos) + abs(sensor[1] - ypos)
            if (xpos, ypos) == beacon:
                break
            if taxi_me <= taxi_sb :
                count += 1
                break
    return count

def part_two(pairs, rng):
    '''part two'''
    for ypos in range(0, rng):
        xpos = 0
        while xpos < rng:
            for pair in pairs:
                sensor, beacon = pair
                taxi_sb = abs(sensor[0] - beacon[0]) + abs(sensor[1] - beacon[1])
                taxi_me = abs(sensor[0] - xpos) + abs(sensor[1] - ypos)
                if taxi_me <= taxi_sb :
                    xpos += taxi_sb - taxi_me
                    break
                if pair == pairs[-1]:
                    return xpos*rng + ypos
            xpos += 1


def clean_data(data):
    '''generate beacons and sensors'''
    output = []
    for line in data:
        nums =  re.findall( '[\\-]?\\d+', line)
        sensor = tuple(map(int, nums[:2]))
        beacon = tuple(map(int, nums[2:]))
        output.append([sensor, beacon])
    return output

with open('./data/data_day_15.txt', 'r', encoding = 'UTF-8') as file:
    lines = [ line.strip() for line in file.readlines() ]

positions = clean_data(lines)

answer_one = part_one(positions, 10000000, 2000000)
print(answer_one)

answer_two = part_two(positions, 4000000)
print(answer_two)
