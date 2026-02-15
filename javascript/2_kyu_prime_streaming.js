/*

Write a generator function that yields a stream of prime numbers. The stream must be able to produce 25 million primes in seconds.

*/

class Primes {
    static* stream() {
        yield 2;

        const segmentOdds = 1 << 20;
        const basePrimes = [3];
        let nextBaseCandidate = 5;

        const preparePrimes = (limit) => {
            let last = basePrimes[basePrimes.length - 1];
            while (last < limit) {
                let n = nextBaseCandidate;
                let isPrime = true;
                const sqrtN = Math.sqrt(n);

                for (let i = 0; i < basePrimes.length; i++) {
                    const p = basePrimes[i];
                    if (p > sqrtN) break;
                    if (n % p === 0) {
                        isPrime = false;
                        break;
                    }
                }

                if (isPrime) {
                    basePrimes.push(n);
                    last = n;
                }
                nextBaseCandidate += 2;
            }
        };

        let low = 3;
        while (true) {
            const high = low + (segmentOdds << 1);
            const limit = Math.floor(Math.sqrt(high - 1));

            preparePrimes(limit);

            // Segmented sieve of Eratosthenes
            const sieve = new Uint8Array(segmentOdds);
            sieve.fill(1);

            for (let i = 0; i < basePrimes.length; i++) {
                const p = basePrimes[i];
                if (p > limit) break;

                let start = p * p;
                if (start < low) {
                    const rem = low % p;
                    start = rem === 0 ? low : low + (p - rem);
                    if ((start & 1) === 0) start += p;
                }

                const step = p << 1;
                for (let n = start; n < high; n += step) {
                    sieve[(n - low) >> 1] = 0;
                }
            }

            for (let i = 0; i < sieve.length; i++) {
                if (sieve[i]) {
                    const n = low + (i << 1);
                    if (n >= 3) yield n;
                }
            }

            low = high;
        }
    }
}
