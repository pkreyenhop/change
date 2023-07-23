Change - A Command-Line Tool to Replace Text in Files

The "change" program is a simple yet powerful command-line tool written in Go that allows you to perform text replacement in files. With this tool, you can quickly and easily replace occurrences of a specific string with another string within a given file.

Usage:
Usage: change <filename> <string_to_replace> <replacement>

Features:
- Easy to Use: The "change" program is designed to be straightforward and user-friendly. Simply provide the filename, the string to be replaced, and the replacement string as command-line arguments.
- In-Place Modification: The program directly modifies the content of the specified file in-place, ensuring that the changes are applied directly to the original file.
- Powerful Text Replacement: The program uses Go's built-in `strings.ReplaceAll` function to perform a global replacement of the specified string with the replacement string. This means that all occurrences of the search string within the file will be replaced.
- Error Handling: The "change" program includes robust error handling to provide helpful feedback to the user in case of any issues, such as file reading or writing errors.

How to Use:
1. Make sure you have Go installed on your system.
2. Save the provided source code into a file named `change.go`.
3. Open your terminal or command prompt and navigate to the directory containing `change.go`.
4. Compile and run the program using the Go compiler:

   go run change.go <filename> <string_to_replace> <replacement>

   Replace <filename>, <string_to_replace>, and <replacement> with the appropriate values.

5. The program will read the contents of the specified file, perform the replacement, and write the modified contents back to the file.
6. A success message will be displayed, informing you about the number of occurrences changed.

The "change" program simplifies the process of replacing text in files, making it a valuable tool for developers, sysadmins, and anyone who needs to perform bulk text replacements efficiently.
