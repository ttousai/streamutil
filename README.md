# Sed for Golang

1. Read input line by line.
1. If Address is specified run Operation only against lines that match
   else for each line run Operation.

Address := pattern | line number
Operation := pattern + action

## Supported Actions
1. Print (default)
1. Substitute
1. Delete
