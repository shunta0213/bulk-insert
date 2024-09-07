# Bulk Insert Test

## Setup

```bash
brew install dbmate mycli
```

# Results

### seed

| record rows | milliseconds |
| :---------: | :----------: |
|     100     |      10      |
|    1000     |      20      |
|    10000    |      82      |
|   100000    |      --      |

### seed chunk

`chunk size = 1000`

| record rows | milliseconds |
| :---------: | :----------: |
|     100     |      13      |
|    1000     |      27      |
|    10000    |      91      |
|   100000    |     594      |

`chunk size = 100000`

| record rows | milliseconds |
| :---------: | :----------: |
|     100     |      10      |
|    1000     |      23      |
|    10000    |      85      |
|   100000    |     461      |
