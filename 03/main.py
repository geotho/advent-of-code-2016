print(len([x for x in [sorted([int(y) for y in line.split()]) for line in open('triangles.txt', 'r').readlines()] if x[0] + x[1] > x[2]]))

f = open('triangles.txt', 'r')

i = 0
buf = [None]*3
total = 0
for line in f:
    buf[i] = [int(y) for y in line.split()]
    i = (i + 1) % 3
    if i == 0:
        for j in range(3):
            x = sorted([x[j] for x in buf])
            total += 1 if x[0] + x[1] > x[2] else 0
print(total)

