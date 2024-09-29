# Go Boot Camp - Day 02

This repository contains the solutions for the Go Boot Camp Day 02 tasks. The focus is on creating various command-line utilities using Go to handle file system operations, including file searching, counting, command execution, and log file archiving.

## Contents

1. [Exercises](#exercises)
    - [Exercise 00: Finding Things](#exercise-00-finding-things)
    - [Exercise 01: Counting Things](#exercise-01-counting-things)
    - [Exercise 02: Running Things](#exercise-02-running-things)
    - [Exercise 03: Archiving Things](#exercise-03-archiving-things)


## Exercises

### Exercise 00: Finding Things

This exercise involves implementing a utility similar to the `find` command. The program can locate directories, regular files, and symbolic links in a specified directory, with options to filter the output.

#### Features

- Accepts a path and command-line options to find:
    - Directories (`-d`)
    - Regular files (`-f`)
    - Symbolic links (`-sl`)
- Optionally filters files by extension with `-ext` (only works when `-f` is specified).
- Resolves symlinks and handles broken symlinks gracefully.
- Skips files and directories that the current user doesn't have permission to access.

#### Example Usage

```bash
# Find all files, directories, and symlinks in /foo
./myFind /foo

# Find only files with '.go' extension
./myFind -f -ext 'go' /path/to/dir

# Find directories only
./myFind -d /path/to/dir
```

### Exercise 01: Counting Things

This exercise implements a `wc`-like utility to gather basic statistics from text files.

#### Features

- The utility supports three mutually exclusive flags:
  - `-l` for counting lines.
  - `-m` for counting characters.
  - `-w` for counting words (default if no flag is specified).
- If no flags are specified, the program defaults to word count (`-w`).
- The utility can accept multiple input files and process them concurrently using goroutines for improved performance.
- Handles UTF-8 encoded text files and supports both English and Russian. Other languages, such as Arabic, are not required to be handled for this exercise.
- Ignores punctuation and considers spaces as the only word delimiters.

#### Example Usage

```bash
# Count words in input.txt
./myWc -w input.txt

# Count lines in multiple files
./myWc -l input2.txt input3.txt

# Count characters in files
./myWc -m input4.txt input5.txt input6.txt
```

#### Output Format

The output consists of a calculated number and the filename, separated by a tab (`\t`).

Example:
```bash
777    input.txt
42     input2.txt
53     input3.txt
1337   input4.txt
2664   input5.txt
3991   input6.txt
```

### Exercise 02: Running Things

In this exercise, you will implement a utility similar to `xargs`. The program will build and execute a command using input provided via standard input (`stdin`).

#### Features

1. The utility treats all command-line parameters as a command (e.g., `wc -l`, `ls -la`).
2. It appends each line from the input (received via `stdin`) as arguments to the command and then executes it.
3. Can be combined with other utilities to create complex workflows.

#### Example Usage

```bash
# Executes 'ls -la' on /a, /b, /c
echo -e "/a\n/b\n/c" | ./myXargs ls -la

# Find '.log' files and count lines
./myFind -f -ext 'log' /path/to/logs | ./myXargs ./myWc -l
```

### Exercise 03: Archiving Things

The final tool for this day is a log rotation utility. "Log rotation" refers to the process of archiving old log files and storing them separately to prevent logs from growing indefinitely in a single file.

#### Features

- Creates a `.tar.gz` archive of each specified log file.
- The archive name is based on the original filename and the file's last modification time (mtime) represented as a UNIX timestamp.
- Supports archiving multiple files at once, optionally specifying an output directory for the archives using the `-a` flag.
- Utilizes goroutines to parallelize the archiving process, allowing for faster handling of multiple files.

#### Example Usage

```bash
# Archive a single log file
./myRotate /dataForTest/dataForMyRotate/app1.log

# Archive multiple log files into a specific directory
./myRotate -a /dataForTest/ /dataForTest/dataForMyRotate/app1.log /dataForTest/dataForMyRotate/app2.log
```
