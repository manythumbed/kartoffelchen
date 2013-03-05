extern mod std;

use to_str::ToStr;

#[deriving_eq]
pub struct Rational	{
	num: int,
	denom: int
}

fn gcd(a: int, b: int) -> int	{
	if b == 0 {
		return a;	
	}

	return gcd(b, a % b);
}

fn lcm(a: int, b:int) -> int {
	return int::abs(a * b) / gcd(a, b);
}

fn add(a: &Rational, b: &Rational) -> Rational	{
	let l = lcm(a.denom, b.denom);
	Rational::new((a.num * (l / a.denom)) + (b.num * (l / b.denom)), l) 
}

fn mult(a: &Rational, b: &Rational) -> Rational {
	Rational::new(a.num * b.num, a.denom * b.denom)
}

const zero: &Rational = &Rational{num: 0, denom:1};

pub impl Rational {
	static fn new(num: int, denom: int) -> Rational {
		Rational{num: num, denom: denom}.reduce()
	}

	fn reduce(&self) -> Rational {
		let g = gcd(self.num, self.denom);
		Rational{ num: self.num / g, denom: self.denom / g}
	}
}

pub impl Rational : ToStr	{
	pure fn to_str() -> ~str	{
		return fmt!("%d/%d", self.num, self.denom);
	}
}

/*
impl Rational: Ord 	{
	pure fn lt(&self, other: &Rational) -> bool { false }
	pure fn le(&self, other: &Rational) -> bool { false }
	pure fn gt(&self, other: &Rational) -> bool { false }
	pure fn ge(&self, other: &Rational) -> bool { false }
}
*/

#[test]
fn test_gcd()	{
	assert gcd(54, 24) == 6;
	assert gcd(42, 56) == 14;
	assert gcd(56, 42) == 14;
	assert gcd(18, 84) == 6;
	assert gcd(2, 2) == 2;
}

#[test]
fn test_lcm() {
	assert lcm(21, 6) == 42;
}

#[test]
fn test_reduce()	{
	assert Rational{num: 3, denom: 9}.reduce() == Rational{num: 1, denom: 3};
	assert Rational{num: 6, denom: 4}.reduce() == Rational{num: 3, denom: 2};
}

#[test]
fn test_new()	{
	assert Rational::new(7, 14) == Rational{num: 1, denom: 2};
}

#[test]
fn test_add()	{
	assert add(&Rational::new(1, 2), &Rational::new(2, 3)) == Rational::new(7, 6)
}

#[test]
fn test_mult()	{
	assert mult(&Rational::new(1, 2), &Rational::new(2, 3)) == Rational::new(1, 3)
}

#[test]
fn test_zero()	{
	assert mult(zero, &Rational::new(1, 2)) == *zero;
	assert add(zero, &Rational::new(1, 2)) == Rational::new(1, 2)
}

#[test]
fn test_to_str()	{
	assert Rational::new(1, 4).to_str() == ~"1/4";
}
