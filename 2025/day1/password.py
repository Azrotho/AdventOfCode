actual = 50
click = 0

with open("inputpass.txt") as file:
    for line in file:
        print(actual)
        direction = line[0]
        nb = int(line[1::])
        if(direction == 'L'):
            for k in range(nb):
                actual = actual - 1
                if(actual == -1):
                    actual = 99
        if(direction == 'R'):
            for k in range(nb):
                actual = actual + 1
                if(actual == 100):
                    actual = 0
        if(actual==0):
            click+=1

print(click)