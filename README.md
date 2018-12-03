# Advent of Code

Monorepo for all AoC solutions

* [2018](./2018)
* [2017](./2017)
* [2016](./2016)

## Set up new year

```sh
YEAR=2018
mkdir -p ${YEAR}
sed s/YEAR/${YEAR}/g main.tpl > ${YEAR}/main.go
for DAY in {01..25}; do
    mkdir -p ${YEAR}/${DAY}
    sed s/DAY/day${DAY}/g solve.tpl > ${YEAR}/${DAY}/solve.go
done
git add -N ${YEAR}
```

## Create runtime

```sh
YEAR=2018
go build -o adventofcode-${YEAR} ./${YEAR}
```

## Solve puzzle

```sh
YEAR=2018
DAY=01
./adventofcode-${YEAR} day${DAY} -input ${YEAR}/${DAY}/input.txt
```
