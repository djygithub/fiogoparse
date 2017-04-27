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
