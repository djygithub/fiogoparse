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
// After compiling the program is copied to /usr/local/go/bin, which is in the PATH concatenation
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
