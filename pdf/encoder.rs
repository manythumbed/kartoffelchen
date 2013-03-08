extern mod std;

use io::{file_writer, Create, WriterUtil};
use result::{chain, Ok};
use path;
use str;

const newline: &str = "\r\n";

fn header() -> ~str {
	~"%PDF-1.7" + newline + ~"%\x93\x8c\x8b\x9e" + newline 
}

#[test]
fn test_output()	{
	let result = chain(file_writer(&path::Path("test.pdf"), &[Create]),	|w| { 
		w.write_str("abc");
		Ok("done") 
	});
	
	error!("%s", result.get());
	error!("%s", header());
}
