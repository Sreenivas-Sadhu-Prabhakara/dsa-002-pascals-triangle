# Pascal's Triangle

> Striver SDE Sheet 002 · Day 1 Arrays — build the triangle row by row, each entry the sum of the two above it. Go, stdlib only, tested.

## Problem

Generate the first n rows of Pascal's triangle: row i has i+1 entries, the edges are 1, and every interior entry is the sum of the two entries diagonally above it.

## Approach

Build iteratively. Each new row starts and ends with 1; entry j is `prev[j-1] + prev[j]`. No formula tricks needed — the triangle's defining recurrence is already the cheapest way to compute all rows. (If only row n were wanted, the running-binomial trick `C(n,k+1) = C(n,k)·(n−k)/(k+1)` would do it in O(n).)

**Complexity:** O(n²) time and space — that is the size of the output itself.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
