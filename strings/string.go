package strings
import "fmt"


func StringsFmt1(args ...interface{}) string{
    return fmt.Sprint(args)
}


func StringsFmt2(args ...interface{}) string{
    return fmt.Sprint(args...)
}

