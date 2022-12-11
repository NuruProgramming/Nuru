// basic types

2 + 2;
4 * 4;
acha a = 10;
acha b = 20;

a + b;

kweli == sikweli;

"Mambo vipi";

// functions

acha jumlisha = fn(x, y) {x + y};

jumlisha(20,30);
jumlisha(100,1000);

acha zidisha = fn(x, y) {x * y};

zidisha(100,1000);
zidisha(200, 20);

// lists can hold any value

acha list = [1, "a", kweli, sikweli];

// a few builtins

acha list = sukuma(list, jumlisha(4,5));

list;
list[2]
list[-100] // prints null
idadi(list)
yamwisho(list);

// if statements

kama (1 > 2) {
	"Moja ni zaidi ya mbili";
} sivyo {
	"Moja si zaidi ya mbili";
}

kama (idadi("Habari") == 6) {
	"Habari yako ndugu";
}

// fibonacci example

acha fibo = fn(x) {
	kama (x == 0) {
		rudisha 0;
	} sivyo {
		kama (x == 1) {
			rudisha 1;
		} sivyo {
			rudisha fibo(x - 1) + fibo(x - 2);
		}
	}
}

fibo(10);
