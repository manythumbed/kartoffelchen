extern mod std;

use cmp::{Eq};

use rational::Rational;
use pitch::Pitch;
use duration = rational::Rational::new;
use pitch = pitch::Pitch::new;

#[deriving_eq]
enum Value {
	Note(Rational, Pitch),
	Rest(Rational)
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
