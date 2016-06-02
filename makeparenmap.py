#!/usr/bin/python3
from collections import namedtuple

Data = namedtuple('UnicodeData', 'number name category rune')

last = None
with open('UnicodeData.txt') as f:
    for line in f:
        number, name, category, rest = line.split(';', 3)
        number = int(number, 16)
        this = Data(number, name, category, chr(number))
        if category in ('Ps', 'Pi'):
            if last != None:
                print('// No matching closer: ' + str(last))
            last = this
        elif category in ('Pe', 'Pf'):
            if last == None:
                print('// No matching opener: ' + str(this))
            elif (category == 'Pe' and last.category != 'Ps' or
                  category == 'Pf' and last.category != 'Pi'):
                print('// No matching closer: ' + str(last))
                print('// No matching opener: ' + str(this))
            else:
                delta = this.number - last.number
                if delta > 1:
                    print('// delta = %d, could be false match' % delta)
                # print('// Matching ' + str(last) + ' with ' + str(this))
                print("'%s': '%s', // U+%X U+%X" % (
                    last.rune, this.rune, last.number, this.number))
            last = None
