package day01

func SolutionA(input []byte) int {
    values := 0
    first, last := -1, -1

    for i := 0; i < len(input); i ++ {
        if input[i] >= 48 && input[i] <= 57 {
            if first == -1 {
                first = int(input[i]) - 48
            } else {
                last = int(input[i]) - 48
            }
        } else if input[i] == 10 {
            if last == -1 {
                last = first
            }
            values += first * 10 + last
            first, last = -1, -1
        }
    }
    return values
}

func SolutionB(input[] byte) int {
    return 0
}
