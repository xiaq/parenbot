#!/bin/bash
wget http://www.unicode.org/Public/UNIDATA/UnicodeData.txt
while IFS=';' read number name category rest
do 
    if [[ "$category" =~ Ps|Pe|Pi|Pf ]]; then 
        printf "%d U+%s \u"$number" %s %s\n" 0x$number $number $category "$name"
        # printf "%s (U+%s, %s): \u"$number"\n" "$name" "$number" "$category"
    fi
done <UnicodeData.txt
