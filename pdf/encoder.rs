extern mod std;

use str;

fn header() -> ~str {
	let newline = ~"\r\n";
	let headerValues = &[~"%PDF-1.7", copy newline, ~"%\x93\x8c\x8b\x9e", copy newline];

	str::concat(headerValues) 
}

#[test]
fn test_output()	{
	error!("%s", header());
}
