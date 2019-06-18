package main

import (
    "bufio"
    "bytes"
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path"
    "runtime"
    "sort"
    "strconv"
    "sync"
)

type Solution struct {
    Answer string
    Path string
    Problem int
}

func SourceSolutions(ctx context.Context, root string) <- chan Solution {
    solutions := make(chan Solution)

    go func() {
        defer close(solutions)

        files, err := ioutil.ReadDir(root)
        if err != nil {
            log.Fatal(err)
        }

        for _, file := range files {
            if path.Ext(file.Name()) != ".go" || file.Name() == "test.go" {
                continue
            }

            problemNumber, _ := strconv.Atoi(file.Name()[:len(file.Name()) - len(path.Ext(file.Name()))])

            select {
            case <- ctx.Done():
                return
            case solutions <- Solution{Path: file.Name(), Problem: problemNumber}:
            }
        }
    }()

    return solutions
}

func RunSolutions(ctx context.Context, unprocessedSolutions <- chan Solution) <- chan Solution {
    solutions := make(chan Solution)

    numWorkers := runtime.GOMAXPROCS(0)
    var wg sync.WaitGroup
    wg.Add(numWorkers)

    for i := 0; i < numWorkers; i++ {
        go func() {
            for solution := range unprocessedSolutions {
                cmd := exec.Command("go", "run", solution.Path)

                answer, err := cmd.Output()
                if err != nil {
                    log.Fatal(err)
                }

                solution.Answer = string(bytes.TrimSuffix(answer, []byte{'\n'}))

                select {
                case <- ctx.Done():
                    break
                case solutions <- solution:
                }
            }

            wg.Done()
        }()
    }

    go func() {
        defer close(solutions)
        wg.Wait()
    }()

    return solutions
}

func FetchAnswers(n int) []string {
    var answers []string

    file, err := os.Open("answers.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    for i := 0; i < n; i++ {
        scanner.Scan()
        answers = append(answers, scanner.Text())
    }

    return answers
}

func main() {
    // TODO - Allow running a single problem supplied by an argument

    ctx, cancelFunc := context.WithCancel(context.Background())
    defer cancelFunc()

    var solutions []Solution

    unprocessedSolutionsStream := SourceSolutions(ctx, ".")
    processedSolutionsStream := RunSolutions(ctx, unprocessedSolutionsStream)

    for solution := range processedSolutionsStream {
        solutions = append(solutions, solution)
    }

    sort.Slice(solutions, func(i, j int) bool {
        return solutions[i].Problem < solutions[j].Problem
    })

    answers := FetchAnswers(len(solutions))

    for i := 0; i < len(solutions); i++ {
        correctness := "correct"
        if solutions[i].Answer != answers[i] {
            correctness = "incorrect"
        }

        fmt.Printf("Problem %d: Output is %s '%s == %s'\n", solutions[i].Problem, correctness, solutions[i].Answer, answers[i])
    }
}
