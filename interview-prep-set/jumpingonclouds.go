func jumpingOnClouds(c []int32) int32 {
    // Write your code here
    count := int32(0)
    for i:=0; i < len(c)-1; {
     if c[i] == 0 {
            if i+2 < len(c) && c[i+2] == 0 {
                count++
                i = i + 2
            } else if c[i+1] == 0 {
                count++
                i++
            }
        }
    }
    return count 

}
