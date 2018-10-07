#!/usr/bin/fish

argparse 'p/problem=' -- $argv

function validate
    if test $argv[2] = $argv[3]
        echo "Problem $argv[1]: Output is correct '$argv[2] == $argv[3]'"
    else
        echo "Problem $argv[1]: Output is incorrect '$argv[2] != $argv[3]'"
    end
end

if set -q _flag_problem
    set -l answer (sed "$_flag_problem!d" answers.txt)
    set -l output (g++ $_flag_problem.cpp -o task_$_flag_problem; and ./task_$_flag_problem; and rm task_$_flag_problem)
    validate $_flag_problem $answer $output
else
    for num in (seq (math (count ls *.cpp) - 1))
        set -l answer (sed "$num!d" answers.txt)
        set -l output (g++ $num.cpp -o task_$num; and ./task_$num; and rm task_$num)
        validate $num $answer $output
    end
end
