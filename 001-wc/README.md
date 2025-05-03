
# wc CLI Tool

A Go implementation of the classic Unix `wc` (word count) command using the Cobra CLI framework.

This tool counts **lines**, **words**, **bytes**, and **non-space characters** from a file or standard input.

---

## 🚀 Installation

```bash
go build -o wc
```

This will create an executable named `wc`.

---

## 🧑‍💻 Usage

```bash
./wc [flags] [file]
```

- If no file is provided, input is read from `stdin`.

---

## 📄 Examples

### ✅ Count all (lines, words, bytes)
```bash
./wc sample.txt
# Output: 10 34 198
```

### ✅ Count only lines
```bash
./wc -l sample.txt
# Output: 10
```

### ✅ Count only words
```bash
./wc -w sample.txt
# Output: 34
```

### ✅ Count only bytes
```bash
./wc -c sample.txt
# Output: 198
```

### ✅ Count only non-space characters
```bash
./wc -m sample.txt
# Output: 164
```

### ✅ Read from stdin
```bash
echo "hello world" | ./wc
# Output: 1 2 12
```

---

## 🏁 Supported Flags

| Flag        | Short | Description                          |
|-------------|-------|--------------------------------------|
| `--lines`   | `-l`  | Print the number of lines            |
| `--words`   | `-w`  | Print the number of words            |
| `--bytes`   | `-c`  | Print the number of bytes            |
| `--chars`   | `-m`  | Print the number of non-space chars  |

---

## 🛠 Built With

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)

---

## 📜 License

MIT
