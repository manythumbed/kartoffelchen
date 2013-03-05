extern mod std;

use to_str::ToStr;

use rational::Rational;
use pitch::Pitch;
use duration = rational::Rational::new;
use pitch = pitch::Pitch::new;

#[deriving_eq]
enum Value {
	Note(Rational, Pitch),
	Rest(Rational)
}

impl Value : ToStr	{
	pure fn to_str() -> ~str	{
		match self {
			Note(d, p) => fmt!("%s-%s", p.to_str(), d.to_str()),
			Rest(d) => fmt!("%s", d.to_str())
		}
	}
}

#[deriving_eq]
enum Music	{
	Primitive(Value),
	Sequential(@Music, @Music),
	Parallel(@Music, @Music),
}

fn note(d: Rational, p: Pitch) -> Music	{
	Primitive(Note(d, p))
}

fn rest(d: Rational) -> Music	{
	Primitive(Rest(d))
}

impl Music : ToStr	{
	pure fn to_str() -> ~str	{
		match self {
			Primitive(v) => v.to_str(),
			Sequential(a, b) => fmt!("%s, %s", (*a).to_str(), (*b).to_str()),
			Parallel(a, b) => fmt!("[%s, %s]", (*a).to_str(), (*b).to_str())
		}
	}
}

#[test]
fn test_note_equality()	{
	assert Rest(duration(3, 2)) == Rest(duration(3, 2));
	assert Rest(duration(6, 4)) == Rest(duration(3, 2));
	assert Note(duration(3, 4), pitch(4, 0)) != Rest(duration(3, 4));
}

#[test]
fn test_music_equality()	{
	assert note(duration(1, 2), pitch(4, 0)) == note(duration(1, 2), pitch(4, 0));
	assert note(duration(1, 2), pitch(4, 0)) != rest(duration(1, 2));
}

#[test]
fn test_duration_equality()	{
	assert duration(1, 2) == duration(1, 2);
}

#[test]
fn test_value_to_str()	{
	assert Note(duration(1, 4), pitch(4, 0)).to_str() == ~"(4,0)-1/4";
	assert Rest(duration(1, 2)).to_str() == ~"1/2";
}

#[test]
fn test_music_to_str()	{
	assert Sequential(@note(duration(1,4), pitch(4, 0)), @note(duration(3, 4), pitch(5, 0))).to_str() == ~"(4,0)-1/4, (5,0)-3/4";
	assert Parallel(@note(duration(1,4), pitch(4, 0)), @note(duration(3, 4), pitch(5, 0))).to_str() == ~"[(4,0)-1/4, (5,0)-3/4]";
}
