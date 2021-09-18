const { log } = console

function callatz(n) {
    if (n % 2 == 0) {
        return n / 2;
    } else {
        return (n * 3) + 1
    }
}

function callatzSeries(n) {
    let i = 0;
    log(`n = ${n}`)
    for (; n != 1; n = callatz(n), i++) {
    }
    log(`iterations: ${i}`)
    log('-----')
    return n
}

function main() {
    const begin = new Date()
    let arr = []
    for (let i = 1; i < 10000; i++) {
        n = callatzSeries(i)
        arr.push(n)
    }
    const end = new Date()
    log(`Array: `, arr)
    log(`Execution time: ${end.getTime() - begin.getTime()}ms`)
}

main()