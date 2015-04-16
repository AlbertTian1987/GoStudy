package datafile3
import (
    "testing"
    "os"
    "path"
    "math/rand"
    "strconv"
    "time"
    "math"
    "fmt"
    "sync"
)

func GetTestDataFile() DataFile{
    dir, _ := os.Getwd()
    filePath := path.Join(dir, "data")
    dataFile, _ := NewMyDataFile(filePath, 20)
    return dataFile
}

func TestWrite(t *testing.T) {
    dataFile := GetTestDataFile()

    var wg sync.WaitGroup
    wg.Add(30)
    for i:=0;i<30;i++{
        go func(index int) {
            i:=rand.Int31()
            if i>math.MaxInt32 {
                i = math.MaxInt32
            }
            s := strconv.Itoa(int(i))
            str:=fmt.Sprintf("goruntine(%d) %s \n",index,s)

            time.Sleep(time.Duration(rand.Intn(100))*time.Nanosecond)
            wsn,err:=dataFile.Write([]byte(str))
            if err!=nil {
                t.Errorf("%s added error : %s",s,err.Error())
            }else {
                t.Logf("%s added wsn is :%d",s,wsn)
            }
            wg.Done()
        }(i)
    }
    wg.Wait()
}

