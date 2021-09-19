# gfeder

![license](https://img.shields.io/github/license/uncle-lv/feder) ![stars](https://img.shields.io/github/stars/uncle-lv/feder) ![issues](https://img.shields.io/github/issues/uncle-lv/feder) ![forks](https://img.shields.io/github/forks/uncle-lv/feder) ![go version](https://img.shields.io/github/go-mod/go-version/uncle-lv/feder?color=%23007d9c) ![database](https://img.shields.io/badge/database-sqlite3-%231296db)

**gfeder** is a simple ORM framework written in Go which can help you quickly understand the core principles of a ORM framework.

If this project is helpful for you, please star the project, thank you!!!

## Functions

- SQL Builder
- Hooks (Before/After Insert/Delete/Query/Update)
- Transactions
- Migration
- A simple logger

## Usage

**gfeder** is built step by step, and I have used git tag to manage every milestone version. You can view all important version by typing the command `git tag`, and get the version you want by typing the command `git checkout <tagname>`. You can view more details by checking commit logs.

```bash
$ git tag
step1-simple-sql-excute
step2-object-table-mapping
step3-find-and-insert
step4-chain-operation
step5-hooks
step6-transaction
step7-migrate

$ git checkout step1-simple-sql-excute
```

## Contributions

Any contribution you make are greatly appreciated.

## Lisence

[GPL-2.0](https://github.com/uncle-lv/gfeder/blob/main/LICENSE)

## Thanks

This project is based on [geektutu's](https://github.com/geektutu) [gee-orm](https://github.com/geektutu/7days-golang/tree/master/gee-orm). That's a wonderful project. I suggest you reading this [bolg](https://geektutu.com/post/geeorm.html) first.