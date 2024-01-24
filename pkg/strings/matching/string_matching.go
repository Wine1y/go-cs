package stringmatching

type StringMatchingFunction func(str, substr string) (index int, found bool)
