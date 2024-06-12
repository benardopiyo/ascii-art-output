# ASCII-Art-Output

This project is an ASCII Art Generator written in Go. It converts input strings into ASCII art using different banner styles and write the result to a specified file.

## Features

* The output is written into a file specified by the --output=<fileName.txt> flag.
* The flag follow the exact format: --output=<fileName.txt>.
* The program displays a usage message in case of incorrect flag format
* The program supports different banner styles.

## Usage 

```bash
go run . --output=<fileName.txt> [STRING] [BANNER]
```

## Examples

### Standard

```bash
go run . --output=banner.txt "hello" standard

cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
```

### Shadow

```bash
go run . --output=banner.txt 'Hello There!' shadow

cat -e banner.txt
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/ascii-art-output.git

cd ascii-art-generator

```

2. Build the project

```bash
go build
```

3. Run the project

```bash
go run .
```

## How it Works

**main.go**

It handles the command-line arguments and calls the appropriate functions to generate ASCII art and write to a file.

**Functions**

This package has files that contains the core functionality of generating ASCII art from a string using a specified banner style.

* **CheckStatus:** Ensures the banner file exists and is correct.
* **FileName:** Determines the banner file name based on the arguments.
* **Download:** Downloads the banner file if it does not exist.
* **CreateFile:** Creates a file and writes the ASCII art to it.
* **BannerArt:** Reads the banner file and creates a map of ASCII art patterns.
* **Art:** Generates ASCII art from the input string using the map of patterns.
* **Tab** Handles instance of tab('\t') in input and replaces with string equivalent
* **Backspace** Handles instances of backspace('\b') and performs the necessary deletion

**banners/**

* Directory containing different banner style files (standard.txt, shadow.txt, thinkertoy.txt).


## Error Handling

* The program checks for various error conditions:

    * Incorrect number of arguments.
    * Invalid flag format.
    * Missing or corrupted banner files.
    * Non-printable characters in the input string.

* If any of these conditions are met, the program prints a usage message and exits.


## Authors

* [Benard Opiyo](https://learn.zone01kisumu.ke/git/beopiyo)
    
## Contributing

* Fork the repository.
* Create a new branch (git checkout -b feature-branch).
* Commit your changes (git commit -m "Add new feature").
* Push to the branch (git push origin feature-branch).
* Create a new Pull Request.
