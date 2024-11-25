# Folder Tree Structure Generator

A command-line tool written in Go that generates a text-based folder tree structure visualization using depth-first search (DFS) algorithm.

## Description

This tool scans through directories and creates a visual tree-like representation of the folder structure, similar to the `tree` command in Linux.

This serve as a learning exercise for Golang, see what it can offer in terms of performance and memory usage.

## Features

- Generates ASCII tree structure of directories and files
- Customizable depth limit for directory traversal
- Option to show/hide hidden files and directories
- Built-in execution timer

### Available Flags

- `-depth`: Set maximum depth for directory traversal (default: unlimited)
- `-timer`: Show execution time
- `-hidden`: Include hidden files and directories
- `-exclude`: Comma-separated list of files/directories to exclude
- `-help`: Show help message
- `-timer`: Show execution time

```bash 
# Return the tree structure limited to 2 levels of depth, including hidden files, and excluding .git and node_modules directories
./outliner -depth 2 -hidden -exclude .git,node_modules

# Return the execution time
./outliner -timer
```




## License
Copyright 2024, Thomas Quan
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
