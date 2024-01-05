package utils

type Daily interface {
    ParseData(string)
    Part1() (int64, error)
    Part2() (int64, error)
}
