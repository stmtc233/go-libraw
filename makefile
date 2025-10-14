print_params:
	gcc -o print_libraw_param_defaults scripts/libraw_print_param_defaults.c -I/opt/homebrew/include -L/opt/homebrew/lib -lraw
	./print_libraw_param_defaults
	rm print_libraw_param_defaults

build_test:
	g++ -o test_metadata tests/test_metadata.cpp -I/opt/homebrew/include -L/opt/homebrew/lib -lraw
