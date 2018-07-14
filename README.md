[![Go Report Card](https://goreportcard.com/badge/github.com/guitarpawat/go-mathapi)](https://goreportcard.com/report/github.com/guitarpawat/go-mathapi)
[![Build Status](https://travis-ci.org/guitarpawat/go-mathapi.svg?branch=master)](https://travis-ci.org/guitarpawat/go-mathapi)
[![codecov](https://codecov.io/gh/guitarpawat/go-mathapi/branch/master/graph/badge.svg)](https://codecov.io/gh/guitarpawat/go-mathapi)
[![License: LGPL v3](https://img.shields.io/badge/License-LGPL%20v3-blue.svg)](https://www.gnu.org/licenses/lgpl-3.0)

# go-mathapi
API that solves math problem.

## How to run
Browsing to the project directory and type command `go run main.go`.

The default host is localhost and port is 8000. You can specify it by using `-host` and `-port` flag.

### Example
```
go run main.go -port -8080
```
This will run the API on localhost:8080

## Usage
There are 8 basic commands:

| Symbol | Command |
|:------:|:-------:|
| +      | add     |
| -      | sub     |
| *      | mul     |
| /      | div     |
| %      | mod     |
| ^      | pow     |
| (      | par     |
| )      | end     |

### Example
Input:
```
[host]:[port]/5/mul/par/4/add/2/end
```
That means:
```
5*(4+2)
```
Output:
```
30.0000
```

## Procedure
This program converts infix to postfix expression and solves the problem by using stack datatype.

```
5*(4+2) => 4 2 + 5 * => 30.0000
```

## TODO
* Add support for constants, for example, `PI = 3.1415926535...`.
* support for trigonometry and logarithm.
* Self define function, for example, `myfunc is (x+y)^2`.