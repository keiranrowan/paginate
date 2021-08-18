# Paginate

Paginate is a simple command based on Plan9's [p](https://9fans.github.io/plan9port/man/man1/p.html) command. It seeks to emulate the same behavior as the original command using Golang targeting a linux environment.

### Installation

Clone the repository and use ```go build``` to build the executable. You can use ```go install``` to install it as ```paginate``` or if you wish to keep the same name as the Plan9 command, use ```go build -o p``` and move the file to a directory in $PATH.
