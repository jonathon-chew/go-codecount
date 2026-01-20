# Lines of Code Analyzer - codecount
<p align="center">
<img width="400" src="doc/images/codecount.png" alt="codecount" title="codecount" />
</p>

A simple Golang-based utility that scans a project directory and reports the number of **files** and **lines of code** for various programming languages.  

It maps file extensions to languages, counts lines across files, and outputs results in a readable, color-formatted format.

---

## 🚀 Features

- Detects and counts lines of code for popular programming languages (Python, JavaScript, Go, Rust, Java, C, C++, TypeScript, etc.).
- Supports human-readable formatting of numbers (e.g., `1,234`).
- Uses [Aphrodite](https://github.com/jonathon-chew/Aphrodite) for colorized output.
- CLI integration via [cmd](https://github.com/jonathon-chew/Omga/cmd).

---

## 🛠️ Installation

Clone the repository:

```bash
git clone https://github.com/jonathon-chew/codecount.git
cd codecount
```

Build the binary:

```bash
go build -o codecount
```

## 📦 Usage

Run the program in a project directory to analyze code:
```bash
./codecount
```

Optionally, you can pass CLI arguments that will be handled by the cmd package:
```bash
./codecount <args>
```

Version
--version -v

Help
--help -h

Ignore a folder
--ignore -i 
Followed by the exact name to exclude - this can be a partial as long as it's the last part of the folder name
eg. .git will ignore any folders called .git, it would also work

Ignore a file
--ignore-file -if
Followed by the exact name to exclude - this can be a partial as long as it's the last part of the file name
eg. README will ignore any file called README, READ would also work, this also works/applies with file extensions

## 📊 Example Output

```
Name:            No. Files:  No. Words:   No. Non Empty Lines:  No. Lines:
gitignore        1           8            6                     7
txt              1           0            0                     0
Markdown         1           479          97                    131
Golang           4           1,105        321                   387
Shell            1           205          63                    74
Z Shell          1           46           4                     5
Json             1           55           16                    16
Totals:          10          1,898        507                   620
```

## 📚 Supported Languages

The tool supports a wide range of file extensions, mapped to their respective languages:
Extension	Language
.py     Python
.js	    JavaScript
.java	Java
.go	    Golang
.rs	    Rust
.cpp	C++
.c, .cc, .cxx	C / C++
.cs	    C#
.php	PHP
.rb	    Ruby
.ts	    TypeScript
.swift	Swift
.kt	    Kotlin
.scala	Scala
.r	    R
.dart	Dart
.hs	    Haskell
.m	    Objective-C
.qml	QML
.jl	    Julia
.sh	    Shell
.pl	    Perl
.lua	Lua
.sql	SQL
.mod, .sum	Golang
.html	HTML
.css	CSS

## 📝 How It Works

Walks the project directory recursively.

For each file:
Detects its extension.

If recognized, opens the file and counts words, lines and non-empty lines.
Aggregates counts per language.
Prints results in colorized output.

## 🧩 Dependencies

Aphrodite – for colorful console output.
cmd – for argument parsing.

Install them using:

go get github.com/jonathon-chew/Aphrodite
go get github.com/jonathon-chew/codecount/cmd

### 🖌️ Attribution

The Go Gopher was originally designed by [Renee French](https://reneefrench.blogspot.com/).  
Used under the [Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/).  
