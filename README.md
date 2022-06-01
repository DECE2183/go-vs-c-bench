# Go vs C benchmark

For benchmark was used [Rastrigin function](https://en.wikipedia.org/wiki/Rastrigin_function).

### Results

| Language | Result                  | Time        |
|----------|-------------------------|-------------|
| Go       | `3.108036677204659e+08` | `221.681ms` |
| C        | `3.108036677204659e+08` | `749.646ms` |

### Build steps

```bash
  gcc -o libs/bin/rastrigin.o -c libs/rastrigin.c
  ar rcs libs/bin/rastrigin.lib libs/bin/rastrigin.o
  go build main.go
```
