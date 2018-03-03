# bowling

The bowling project reads a file with the score of a bowling match and returns the final score.
The file must have the following format:

| 1 4 | 4 5 | 6 / | 5 / | X | 0 1 | 7 / | 6 / | X | 2 / 6 |

Where "/" represents a split and "X" a strike.

It will only read the first line and no format check is performed.

## Usage

To use it copy the bowling binary in the bin file and run it followed by the file path.
You can also clone the repository and build the project.

## Example

```
$ cat input.txt
| 1 4 | 4 5 | 6 / | 5 / | X | 0 1 | 7 / | 6 / | X | 2 / 6 |
$ bowling ./input.txt
133
```
