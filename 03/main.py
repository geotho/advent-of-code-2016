print(len([x for x in [sorted([int(y) for y in line.split()]) for line in open('triangles.txt', 'r').readlines()] if x[0] + x[1] > x[2]]))
