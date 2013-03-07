extern mod std;

use io::{file_writer, Create, WriterUtil};
use result::{chain, Ok};
use path;
use str;

fn header() -> ~str {
	let newline = ~"\r\n";
	let headerValues = &[~"%PDF-1.7", copy newline, ~"%\x93\x8c\x8b\x9e", copy newline];

	str::concat(headerValues) 
}

#[test]
fn test_output()	{
	let result = chain(file_writer(&path::Path("test.pdf"), &[Create]),	|w| { 
		Ok(w.write_str("abc")) 
	});
	//file_writer(&path::Path("test.pdf"), &[Create]).chain(|w| w.write_str("abc"); ok("done"));
	
	error!("%s", header());
}
