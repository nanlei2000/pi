/**
 * @param {number} n
 */
function calculate(n) {
  const t = n + 10
  const b = BigInt(10) ** BigInt(t)
  let x1 = (b * 4n) / 5n
  let x2 = b / -239n
  let s = x1 + x2

  n *= 2
  for (let i = 3n; i < n; i += 2n) {
    x1 /= -25n
    x2 /= -57121n
    x = (x1 + x2) / i
    s += x
  }

  let pi = s * 4n
  pi /= 10n ** 10n
  const piStr = pi.toString()

  return `3.${piStr.slice(1)}`
}
