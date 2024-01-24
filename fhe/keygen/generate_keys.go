package keygen

import (
	"math/big"

	"github.com/a-t-0/golang-pallier-fhe/fhe/helper"
)

func getKeys() (*big.Int, *big.Int, *big.Int, *big.Int) {
	p, q := pickTwoLargePrimes()
	n := computePrimeProduct(p, q)
	g := getRandG(n)
	lambda := ComputeLambda(p, q)
	mu := getModularMultiplicativeInverse(g, lambda, n)
	return n, g, lambda, mu
}

// pickTwoLargePrimes returns two large prime numbers randomly and
// independently such that gcd(pq,(p−1)(q−1)) is 1.
func pickTwoLargePrimes() (*big.Int, *big.Int) {
	p := big.NewInt(13)
	q := big.NewInt(17)
	return p, q
}

// computePrimeProduct returns the multiplication of two prime integers.
func computePrimeProduct(p *big.Int, q *big.Int) *big.Int {
	// TODO: assert p is prime.
	// TODO: assert q is prime.
	// var n int = p * q
	var n big.Int
	n.Mul(p, q)
	return &n
}

// ComputeL computes the L function as named by Pascal Pallier in the FHE
// presentation. L(x) = [(x-1)/n], THIS IS NOT DIVIDE. This is: how often does
// the bottom fit in the top, like [2/3]=0, [5/3=1], [7/3=2]. A more
// mathematical description is:  the largest integer value v ≥ 0 to satisfy the
// relation a >= v*b. The Python equivalent is 11 // 4 (=2).
func ComputeL(x *big.Int, n *big.Int) *big.Int {
	// Ensure x is non-negative
	if x.Sign() < 0 {
		panic("x must be non-negative")
	}

	// Ensure n is positive
	if n.Sign() <= 0 {
		panic("n must be positive")
	}

	// Compute the result using big.Int's Div method
	// TODO: Important: verify you should also subtract 1 at the bottom.
	result := new(big.Int).Sub(x, big.NewInt(1)).Div(new(big.Int).Sub(x, big.NewInt(1)), n)

	// Convert the result to an int
	return result
}

// ComputeLambda returns the least common multiple of primes p and q (minus
// one) like: lcm(p−1,q−1).
func ComputeLambda(p *big.Int, q *big.Int) *big.Int {
	lambda := helper.Lcm(p.Sub(p, big.NewInt(1)), q.Sub(q, big.NewInt(1)))
	return lambda
}

// getRandG returns a random integer g in {Z_{n^2}}^* between 1 and and n^2.
//   - {Z_{n^2}}​: This denotes the set of integers modulo n^2. It represents the
//     integers from 0 to n^2 − 1 when considering the residue classes.
//   - The residue classes are: if we consider the modulus n=5, then the
//     residue class of 2 modulo 5 is represented as {[2]}_5​ and includes
//     integers like 2, 7, 12, −3, etc., because they all leave the same
//     remainder (2) when divided by 5.
//   - The asterisk (∗) typically indicates that we are excluding 0 from the
//     set. Therefore, {Z_{n^2}}^* is the set of integers modulo n2n2
//     excluding 0.
//   - Multiplicative group: This means that the elements in the set form a
//     group under multiplication. In this context, it implies that for any two
//     elements a and b in {Z_{n^2}}^*, their product a*b is also in
//     {Z_{n^2}}^*.
func getRandG(n *big.Int) *big.Int {
	// TODO: assert g is integer.
	// TODO: assert g is larger than 1.
	// TODO: assert g < n^2.
	// var g int = 4886
	g := big.NewInt(4886)
	return g
}

// getModularMultiplicativeInverse returns an integer x of a such that the
// product a*x is congruent to 1 with respect to the modulus m.[1] In the
// standard notation of modular arithmetic this congruence is written as
// a x ≡ 1 ( mod m ). In this function, a = g^lambda mod n^2.
// Question: in case mu does not exist, what are the security implications if g
// is modified instead of starting again at pickTwoLargePrimes?
func getModularMultiplicativeInverse(g *big.Int, lambda *big.Int, n *big.Int) *big.Int {
	// TODO: separate into separate functions and test each step.
	// TODO: determine whether cast to float64 is required to compute
	// exponents of int.
	// var exponent int = int(math.Pow(float64(g), float64(lambda)))
	exponent := new(big.Int).Exp(g, lambda, nil)

	// var a int = exponent % (n * n)
	a := new(big.Int).Mod(exponent, new(big.Int).Mul(n, n))

	L := ComputeL(a, n)

	// var inverseL int = int(math.Pow(float64(L), float64(-1)))
	inverseL := new(big.Int).Exp(L, big.NewInt(-1), nil)

	// var mu int = inverseL % n
	mu := new(big.Int).Mod(inverseL, n)
	return mu
}
