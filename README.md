# Fake Site Builder

Utility to build a site with N pages from a bunch of sample seed pages

```
Usage of ./main:
  -dumpPath string
        Directory where fabricated HTML would be stored (default "site")
  -pageCount int
        Number of pages to generate (default 100)
  -template string
        Path to seed HTML pages (default "./test-data")
```

### How to build 
```
go get 'github.com/PuerkitoBio/goquery'
go build src/main.go
```
