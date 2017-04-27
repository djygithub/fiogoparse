# fiogoparse
Use go to parse a valid fio report, summarizing IOPS.

Two versions, one using OS commands, one using go.


package main
import (
     "os/exec"
     "os"
)
func main() {
    cmd := "grep iops | tr ':' ' ' | awk '{print $4}' | tr 'iops=' '    ' | awk 'BEGIN {sum=0} {sum=sum+$1} END{print sum}'"
    c1 := exec.Command("bash", "-c",cmd)
    c1.Stdout = os.Stdout
    c1.Stdin = os.Stdin
    c1.Run()
}
　
david@debian:~/go/src/fiotot$ cat /home/david/test.txt | fiotot
39187
david@debian:~/go/src/fiotot$ 
 
package main
import ( "os"
         "fmt"
         "strconv"
         "bufio"
         "strings") 
var lines string
var totiops int
func main() {
   s := bufio.NewScanner(os.Stdin)
   totiops = 0
   for s.Scan() {
      lines = s.Text()
      if strings.Contains(lines, "iops") { 
         lastBin := strings.LastIndex( lines, "p" ) + 3
         lastPin := strings.LastIndex( lines, "," )
         iops := lines[lastBin:lastPin]
         j, err := strconv.Atoi(iops)
         if err != nil {
            fmt.Println(err)
         }
         totiops = totiops + j
      }
   }
   fmt.Println(totiops)
}
　
david@debian:~/go/src/rdstdin$ cat /home/david/test.txt | ./rdstdin
39187
david@debian:~/go/src/rdstdin$ cat testfio.txt | ./rdstdin
8919
