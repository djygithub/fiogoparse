# fiogoparse
Use go to parse a valid fio report, summarizing IOPS.

Two versions, one using OS commands, one using go.

///////////////////////////////////////////////////////////////////////////////////////////////
// fiotot.go - Totals the iops values in an fio report
//
// mixed : (groupid=0, jobs=1): err= 0: pid=9103: Fri Nov 18 15:51:29 2016
// read : io=5121.4MB, bw=19953KB/s, iops=4988 , runt=262833msec
// slat (usec): min=108, max=459674, avg=188.94, stdev=1410.70
// clat (usec): min=10, max=500500, avg=303.14, stdev=2201.15
// lat (usec): min=119, max=511924, avg=492.31, stdev=2752.57
// clat percentiles (usec):
// | 1.00th=[ 17], 5.00th=[ 21], 10.00th=[ 26], 20.00th=[ 139],
// | 30.00th=[ 149], 40.00th=[ 163], 50.00th=[ 247], 60.00th=[ 270],
// | 70.00th=[ 286], 80.00th=[ 306], 90.00th=[ 398], 95.00th=[ 426],
// | 99.00th=[ 2928], 99.50th=[ 6048], 99.90th=[10176], 99.95th=[11328],
// | 99.99th=[15168]
// bw (KB /s): min= 241, max=30184, per=26.61%, avg=20199.08, stdev=11498.41
// write : io=5118.7M B, bw=19942KB/s, iops=4985 , runt=262833 msec
//
//
//
//
// This program reads stdin, and is expecting an fio result file to be piped to it as shown below
//      cat test.txt | fiotot
// The program will hang if invoked by it's lonesome
//      fiotot
//
// After compiling the program is copied to /usr/local/go/bin, which is in the PATH concatenantion
//////////////////////////////////////////////////////////////////////////////////////////////////
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
　
david@debian:~/go/src/rdstdin$ more rdstdin.go 
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
david@debian:~/go/src/rdstdin$ cat /home/david/test.txt 
mixed: (groupid=0, jobs=1): err= 0: pid=9103: Fri Nov 18 15:51:29 2016
  read : io=5121.4MB, bw=19953KB/s, iops=4988, runt=262833msec
    slat (usec): min=108, max=459674, avg=188.94, stdev=1410.70
    clat (usec): min=10, max=500500, avg=303.14, stdev=2201.15
     lat (usec): min=119, max=511924, avg=492.31, stdev=2752.57
    clat percentiles (usec):
     |  1.00th=[   17],  5.00th=[   21], 10.00th=[   26], 20.00th=[  139],
     | 30.00th=[  149], 40.00th=[  163], 50.00th=[  247], 60.00th=[  270],
     | 70.00th=[  286], 80.00th=[  306], 90.00th=[  398], 95.00th=[  426],
     | 99.00th=[ 2928], 99.50th=[ 6048], 99.90th=[10176], 99.95th=[11328],
     | 99.99th=[15168]
    bw (KB  /s): min=  241, max=30184, per=26.61%, avg=20199.08, stdev=11498.41
  write: io=5118.7MB, bw=19942KB/s, iops=4985, runt=262833msec
    slat (usec): min=2, max=492968, avg= 7.91, stdev=814.90
    clat (usec): min=2, max=509641, avg=299.62, stdev=2030.28
     lat (usec): min=6, max=509649, avg=307.65, stdev=2191.97
    clat percentiles (usec):
     |  1.00th=[   15],  5.00th=[   18], 10.00th=[   23], 20.00th=[  139],
     | 30.00th=[  147], 40.00th=[  161], 50.00th=[  245], 60.00th=[  270],
     | 70.00th=[  282], 80.00th=[  302], 90.00th=[  394], 95.00th=[  426],
     | 99.00th=[ 2896], 99.50th=[ 6048], 99.90th=[10176], 99.95th=[11584],
     | 99.99th=[15040]
    bw (KB  /s): min=  214, max=30856, per=26.61%, avg=20188.12, stdev=11502.08
    lat (usec) : 4=0.01%, 10=0.01%, 20=5.32%, 50=7.15%, 100=0.03%
    lat (usec) : 250=38.16%, 500=47.82%, 750=0.22%, 1000=0.01%
    lat (msec) : 2=0.03%, 4=0.38%, 10=0.76%, 20=0.11%, 50=0.01%
    lat (msec) : 100=0.01%, 250=0.01%, 500=0.01%, 750=0.01%
  cpu          : usr=2.15%, sys=11.07%, ctx=1311143, majf=0, minf=1444
  IO depths    : 1=0.1%, 2=0.1%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued    : total=r=1311060/w=1310380/d=0, short=r=0/w=0/d=0, drop=r=0/w=0/d=0
     latency   : target=0, window=0, percentile=100.00%, depth=4
mixed: (groupid=0, jobs=1): err= 0: pid=9104: Fri Nov 18 15:51:29 2016
  read : io=5121.4MB, bw=19820KB/s, iops=4955, runt=264593msec
    slat (usec): min=108, max=460796, avg=191.07, stdev=1532.54
    clat (usec): min=10, max=508989, avg=302.89, stdev=1962.37
     lat (usec): min=120, max=513780, avg=494.21, stdev=2653.01
    clat percentiles (usec):
     |  1.00th=[   16],  5.00th=[   21], 10.00th=[   26], 20.00th=[  141],
     | 30.00th=[  151], 40.00th=[  165], 50.00th=[  249], 60.00th=[  270],
     | 70.00th=[  286], 80.00th=[  310], 90.00th=[  398], 95.00th=[  426],
     | 99.00th=[ 2896], 99.50th=[ 6048], 99.90th=[10304], 99.95th=[11584],
     | 99.99th=[15040]
    bw (KB  /s): min=  201, max=30040, per=26.42%, avg=20055.44, stdev=11397.43
  write: io=5118.7MB, bw=19810KB/s, iops=4952, runt=264593msec
    slat (usec): min=2, max=493200, avg= 6.97, stdev=562.87
    clat (usec): min=1, max=509122, avg=303.95, stdev=2264.88
     lat (usec): min=4, max=509141, avg=311.04, stdev=2335.49
    clat percentiles (usec):
     |  1.00th=[   14],  5.00th=[   18], 10.00th=[   23], 20.00th=[  139],
     | 30.00th=[  149], 40.00th=[  163], 50.00th=[  247], 60.00th=[  270],
     | 70.00th=[  286], 80.00th=[  306], 90.00th=[  398], 95.00th=[  426],
     | 99.00th=[ 2928], 99.50th=[ 6048], 99.90th=[10176], 99.95th=[11456],
     | 99.99th=[15168]
    bw (KB  /s): min=  147, max=30784, per=26.42%, avg=20044.93, stdev=11395.41
    lat (usec) : 2=0.01%, 10=0.01%, 20=5.28%, 50=7.18%, 100=0.03%
    lat (usec) : 250=37.98%, 500=47.79%, 750=0.43%, 1000=0.01%
    lat (msec) : 2=0.03%, 4=0.38%, 10=0.76%, 20=0.11%, 50=0.01%
    lat (msec) : 100=0.01%, 250=0.01%, 500=0.01%, 750=0.01%
  cpu          : usr=2.18%, sys=11.38%, ctx=1311149, majf=0, minf=2101
  IO depths    : 1=0.1%, 2=0.1%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued    : total=r=1311060/w=1310380/d=0, short=r=0/w=0/d=0, drop=r=0/w=0/d=0
     latency   : target=0, window=0, percentile=100.00%, depth=4
mixed: (groupid=0, jobs=1): err= 0: pid=9105: Fri Nov 18 15:51:29 2016
  read : io=5121.4MB, bw=18975KB/s, iops=4743, runt=276379msec
    slat (usec): min=109, max=410217, avg=196.35, stdev=1392.59
    clat (usec): min=12, max=505943, avg=322.46, stdev=2334.62
     lat (usec): min=122, max=510400, avg=519.34, stdev=2871.30
    clat percentiles (usec):
     |  1.00th=[   18],  5.00th=[   24], 10.00th=[   33], 20.00th=[  145],
     | 30.00th=[  161], 40.00th=[  183], 50.00th=[  251], 60.00th=[  278],
     | 70.00th=[  306], 80.00th=[  346], 90.00th=[  410], 95.00th=[  470],
     | 99.00th=[ 2992], 99.50th=[ 6112], 99.90th=[10176], 99.95th=[11328],
     | 99.99th=[16064]
    bw (KB  /s): min=  208, max=28832, per=25.29%, avg=19196.93, stdev=10684.55
  write: io=5118.7MB, bw=18965KB/s, iops=4741, runt=276379msec
    slat (usec): min=2, max=492971, avg= 9.29, stdev=872.14
    clat (usec): min=2, max=508199, avg=311.74, stdev=1885.87
     lat (usec): min=7, max=508212, avg=321.18, stdev=2081.88
    clat percentiles (usec):
     |  1.00th=[   15],  5.00th=[   21], 10.00th=[   25], 20.00th=[  143],
     | 30.00th=[  159], 40.00th=[  177], 50.00th=[  249], 60.00th=[  278],
     | 70.00th=[  302], 80.00th=[  338], 90.00th=[  406], 95.00th=[  466],
     | 99.00th=[ 2896], 99.50th=[ 6048], 99.90th=[10048], 99.95th=[11200],
     | 99.99th=[14912]
    bw (KB  /s): min=  169, max=29856, per=25.29%, avg=19188.04, stdev=10704.82
    lat (usec) : 4=0.01%, 10=0.01%, 20=2.86%, 50=9.56%, 100=0.07%
    lat (usec) : 250=37.60%, 500=47.00%, 750=1.60%, 1000=0.01%
    lat (msec) : 2=0.03%, 4=0.38%, 10=0.78%, 20=0.11%, 50=0.01%
    lat (msec) : 100=0.01%, 250=0.01%, 500=0.01%, 750=0.01%
  cpu          : usr=3.02%, sys=12.95%, ctx=1311155, majf=0, minf=1443
  IO depths    : 1=0.1%, 2=0.1%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued    : total=r=1311060/w=1310380/d=0, short=r=0/w=0/d=0, drop=r=0/w=0/d=0
     latency   : target=0, window=0, percentile=100.00%, depth=4
mixed: (groupid=0, jobs=1): err= 0: pid=9106: Fri Nov 18 15:51:29 2016
  read : io=5121.4MB, bw=19652KB/s, iops=4913, runt=266853msec
    slat (usec): min=63, max=453484, avg=191.05, stdev=1403.70
    clat (usec): min=11, max=500502, avg=310.51, stdev=2276.65
     lat (usec): min=87, max=508640, avg=501.85, stdev=2795.79
    clat percentiles (usec):
     |  1.00th=[   17],  5.00th=[   21], 10.00th=[   28], 20.00th=[  141],
     | 30.00th=[  153], 40.00th=[  167], 50.00th=[  249], 60.00th=[  274],
     | 70.00th=[  290], 80.00th=[  314], 90.00th=[  402], 95.00th=[  434],
     | 99.00th=[ 2960], 99.50th=[ 6112], 99.90th=[10304], 99.95th=[11840],
     | 99.99th=[15680]
    bw (KB  /s): min=  227, max=29032, per=26.20%, avg=19887.96, stdev=11257.11
  write: io=5118.7MB, bw=19642KB/s, iops=4910, runt=266853msec
    slat (usec): min=2, max=492602, avg= 8.37, stdev=848.63
    clat (usec): min=2, max=456139, avg=301.62, stdev=1883.94
     lat (usec): min=6, max=495089, avg=310.12, stdev=2070.21
    clat percentiles (usec):
     |  1.00th=[   15],  5.00th=[   19], 10.00th=[   23], 20.00th=[  139],
     | 30.00th=[  151], 40.00th=[  165], 50.00th=[  249], 60.00th=[  270],
     | 70.00th=[  286], 80.00th=[  310], 90.00th=[  398], 95.00th=[  430],
     | 99.00th=[ 2832], 99.50th=[ 6048], 99.90th=[10304], 99.95th=[11456],
     | 99.99th=[15040]
    bw (KB  /s): min=  201, max=30384, per=26.20%, avg=19877.51, stdev=11279.16
    lat (usec) : 4=0.01%, 10=0.01%, 20=4.57%, 50=7.89%, 100=0.04%
    lat (usec) : 250=37.71%, 500=47.92%, 750=0.59%, 1000=0.01%
    lat (msec) : 2=0.03%, 4=0.37%, 10=0.76%, 20=0.12%, 50=0.01%
    lat (msec) : 100=0.01%, 250=0.01%, 500=0.01%, 750=0.01%
  cpu          : usr=2.45%, sys=11.52%, ctx=1311151, majf=0, minf=1252
  IO depths    : 1=0.1%, 2=0.1%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued    : total=r=1311060/w=1310380/d=0, short=r=0/w=0/d=0, drop=r=0/w=0/d=0
     latency   : target=0, window=0, percentile=100.00%, depth=4
　
david@debian:~/go/src/rdstdin$ cat testfio.txt 
random-read: (groupid=0, jobs=1): err= 0: pid=3456: Sat Mar 19 13:26:49 2016
  read : io=524704KB, bw=17854KB/s, iops=4463, runt= 29389msec
    clat (usec): min=30, max=29950, avg=211.08, stdev=295.63
     lat (usec): min=30, max=29950, avg=211.22, stdev=295.63
    clat percentiles (usec):
     |  1.00th=[  149],  5.00th=[  161], 10.00th=[  167], 20.00th=[  177],
     | 30.00th=[  181], 40.00th=[  185], 50.00th=[  189], 60.00th=[  193],
     | 70.00th=[  197], 80.00th=[  203], 90.00th=[  217], 95.00th=[  278],
     | 99.00th=[  612], 99.50th=[  980], 99.90th=[ 1992], 99.95th=[ 4512],
     | 99.99th=[15680]
    bw (KB  /s): min= 7984, max=23488, per=99.67%, avg=17794.62, stdev=3960.49
  write: io=523872KB, bw=17825KB/s, iops=4456, runt= 29389msec
    clat (usec): min=1, max=2784, avg= 7.67, stdev=13.93
     lat (usec): min=1, max=2784, avg= 7.86, stdev=13.96
    clat percentiles (usec):
     |  1.00th=[    2],  5.00th=[    2], 10.00th=[    2], 20.00th=[    3],
     | 30.00th=[    3], 40.00th=[    5], 50.00th=[    7], 60.00th=[    7],
     | 70.00th=[    8], 80.00th=[    9], 90.00th=[   13], 95.00th=[   25],
     | 99.00th=[   32], 99.50th=[   37], 99.90th=[   54], 99.95th=[   66],
     | 99.99th=[  115]
    bw (KB  /s): min= 8200, max=22952, per=99.67%, avg=17766.76, stdev=3977.87
    lat (usec) : 2=0.02%, 4=15.35%, 10=25.14%, 20=5.44%, 50=3.96%
    lat (usec) : 100=0.09%, 250=47.08%, 500=1.89%, 750=0.60%, 1000=0.21%
    lat (msec) : 2=0.17%, 4=0.02%, 10=0.02%, 20=0.01%, 50=0.01%
  cpu          : usr=0.04%, sys=31.29%, ctx=131848, majf=0, minf=6
  IO depths    : 1=100.0%, 2=0.0%, 4=0.0%, 8=0.0%, 16=0.0%, 32=0.0%, >=64=0.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     issued    : total=r=131176/w=130968/d=0, short=r=0/w=0/d=0
     latency   : target=0, window=0, percentile=100.00%, depth=1
　
david@debian:~/go/src/rdstdin$ 
