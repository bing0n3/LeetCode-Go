package backtracking 

func generatePossibleNextMoves(s string) []string {
    // backtracking
    if len(s) == 0 {
        return []string{}
    }
    
    res := []string{}

    news := []rune(s)
    for i := 0; i < len(news) - 1; i++ {
        if news[i] == '+' && news[i+1] == '+' {
            news[i] = '-'
            news[i+1] = '-'
            res = append(res, string(news))
            news[i] = '+'
            news[i+1] = '+'
        }
    }
    
    return res
}
