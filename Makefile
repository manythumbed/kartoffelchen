test:
	rustc --test -o all-tests kartoffelchen.rc
	./all-tests

clean:
	rm all-tests
	rm -rf *.dSYM
		 
.PHONY: clean
