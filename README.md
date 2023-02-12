# NFO-to-UTF8
A command line tool to convert NFO files from CP437 to UTF-8 encoding.

## The problem
NFO files containing ASCII artwork are usually encode as CP437. However, when converted to UTF-8 they are often treated as ISO-8859-1 or Shift_JIS encoded and therefore the ASCII artwork gets lost.

## The solution
NFO-to-UTF8 checks the NFO file whether it is identified as being single byte encoded e.g. as ISO-8859-X or as Shift_JIS encoded. If so it then however assumes that it is CP437 encoded and converts it to UTF-8, preserving the ASCII artwork correctly.

## Bonus
When adding the flag `-s` (or `--spaces`) the space characters are replaced with non-breaking space characters (U+00A0).
When adding the flag `-l` (or `--linebreaks`) the line break characters are replaced with the correct characters for the system (LF for Linux/Mac and CRLF for Windows).

## Usage:
```
NFO-to-UTF8 NFO [-v] [-s] [-l]
```
#### Positional Variables:
```
NFO   Path to the NFO file to be converted (Required)
```
#### Flags:
```
   --version      Displays the program version string.
-h --help         Displays help with available flag, subcommand, and positional value parameters.
-s --spaces       Convert spaces to non-breaking spaces
-l --linebreaks   Convert line breaks to correct characters for the system (LF for Linux/Mac and CRLF for Windows)
-v --verbose      Show verbose output
```