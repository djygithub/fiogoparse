# fiogoparse
Use go to parse a valid fio report, summarizing IOPS.

Two versions, one using OS commands, one using go.

# david@debian:~/go/src/fiotot$ cat /home/david/test.txt | fiotot
39187
 
 

# david@debian:~/go/src/rdstdin$ cat /home/david/test.txt | ./rdstdin
39187
# david@debian:~/go/src/rdstdin$ cat testfio.txt | ./rdstdin
8919
