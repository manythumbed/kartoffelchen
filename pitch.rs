extern mod std;

#[deriving_eq]
pub struct Pitch	{
	octave: int,
	index: int
}

pub impl Pitch	{
	fn absolute(&self) -> int { (self.octave * 12) + self.index }
	fn transpose(&self, transposition: int) -> Pitch {
		let t = self.absolute() + transposition;
		if t < 0 && t % 12 != 0	{
			return Pitch{octave: (t / 12) - 1, index: (12 + (t % 12)) % 12}
		}
		return Pitch{octave: t / 12, index: t % 12}
	}
	static fn new(octave: int, index: int) -> Pitch	{ 
		Pitch{octave: octave, index: index % 12}
	}
}

#[test]
fn test_transpose()	{
	assert Pitch{octave: 4, index: 0}.transpose(0) == Pitch{octave: 4, index: 0};
	assert Pitch{octave: 4, index: 0}.transpose(1) == Pitch{octave: 4, index: 1};
	assert Pitch{octave: 4, index: 0}.transpose(12) == Pitch{octave: 5, index: 0};
	assert Pitch{octave: 4, index: 0}.transpose(-12) == Pitch{octave: 3, index: 0};
	assert Pitch{octave: 0, index: 0}.transpose(-12) == Pitch{octave: -1, index: 0};
	assert Pitch{octave: 0, index: 0}.transpose(-13) == Pitch{octave: -2, index: 11};
	assert Pitch{octave: -2, index: 11}.transpose(-13) == Pitch{octave: -3, index: 10};
}
