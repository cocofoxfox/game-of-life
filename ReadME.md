
## Directory Structure

A typical project layout might be:
- `gameoflife.go` — the main Go code.
- `README.md` — this document.
- `input/` — folder containing sample input files.

---

## Requirements

- **Go Compiler**
- A terminal or shell to build and run the program.

---

## Building & Running

1. **Build the Program**  
   From the directory containing `gameoflife.go`:
   ```bash
   go build gameoflife.go
2. **Run the Program**
   ```bash
   ./gameoflife inputs/sample1.txt > output_sample1.txt
3. **Run the Tests**
   ```bash
   go test or go test -v