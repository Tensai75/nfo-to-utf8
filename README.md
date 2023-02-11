# NFO-to-UTF8
A command line tool to convert NFO files from CP437 to UTF-8 encoding.

## The problem
NFO files containing ASCII artwork are usually encode as CP437. However, when converted to UTF-8 they are often treated as ISO-8859-1 encoded and therefore the ASCII artwork gets lost.

## The solution
NFO-to-UTF8 checks the NFO file whether it is identified as being ISO-8859-1 encoded. If so it then however assumes that it is CP437 encoded and converts it to UTF-8, preserving the ASCII artwork correctly.

## Bonus
When adding the flag -s (-spaces) the space characters are replaced with non-breaking space characters (U+00A0).

## Usage:
```
NFO-to-UTF8 [NFO] [-v] [-s]
```
#### Positional Variables:
```
NFO   Path to the NFO file to be converted (Required)
```
#### Flags:
```
   --version   Displays the program version string.
-h --help      Displays help with available flag, subcommand, and positional value parameters.
-s --spaces    Convert spaces to non-breaking spaces
```