package et
/*
下面是样本测试：
//Output: hello world必须紧跟在最后一个}之前，跟前面的//之间不能有空格
可以分多行，除第一行外，其他行不需要以Output:开头
*/
func TypeCategoryof(input interface{}) string{
    switch input.(type) {
        case bool:
        return "boolean"
        case int,uint,int8,uint8,int16,uint16,int32,uint32,int64,uint64:
        return "integer"
        case float32,float64:
        return "float"
        case complex64,complex128:
        return "complex"
        case string:
        return "string"
    }
    return "unknown"
}