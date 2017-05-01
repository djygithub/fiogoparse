# fiogoparse
Use go to parse a valid fio report, summarizing IOPS.

Two versions, one using OS commands, one using go.

- david@debian:~/go/src/fiotot$ cat /home/david/test.txt | fiotot
- 39187
- david@debian:~/go/src/rdstdin$ cat /home/david/test.txt | ./rdstdin
- 39187
- david@debian:~/go/src/rdstdin$ cat testfio.txt | ./rdstdin
- 8919
## References
* Programming in Go Creating Applications for the 21st Century - Mark Summerfield March 2015, Addsion-Wesley http://www.qtrac.eu/gobook.html
* The Go Programming Language - Alan A. A. Donovan Brian W. Kernighan April 2016. Addison-Wesley http://www.informit.com/store/go-programming-language-9780134190440
