func repeatedString(s string, n int64) int64 {
    count := int64(0)
    for _, val := range s {
        if val == 'a' {
            count++
        }
    }
     numsubstrings := n/int64(len(s))
     reminder := n%int64(len(s))
     repeated := int64(0)
     for i:=int64(0); i<reminder; i++{
         if s[i] == 'a'{
             repeated++
         }
     }
     repeated = repeated + numsubstrings*count
     return repeated
}
