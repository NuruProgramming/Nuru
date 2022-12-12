chapa("Testing basic types...");
chapa(2 + 2);
chapa(4 * 4);
acha a = 10;
acha b = 20;

chapa(a + b);

chapa(kweli == sikweli);
chapa("Testing empty lines...")
chapa();
chapa();
chapa();
chapa("Mambo vipi");

chapa("Testing Functions... ");

acha jumlisha = fn(x, y) {x + y};

chapa(jumlisha(20,30));
chapa(jumlisha(100,1000));

acha zidisha = fn(x, y) {x * y};

chapa(zidisha(100,1000));
chapa(zidisha(200, 20));

// lists can hold any value
chapa("Testing lists...");
acha list = [1, "a", kweli, sikweli];

// a few builtins

acha list = sukuma(list, jumlisha(4,5));

chapa(list);
chapa(list[2]);
chapa(list[-100]); // prints null
chapa(idadi(list));
chapa(yamwisho(list));

// if statements
chapa("Testing if statements...");
kama (1 > 2) {
	chapa("Moja ni zaidi ya mbili");
} sivyo {
	chapa("Moja si zaidi ya mbili");
}

kama (idadi("Habari") == 6) {
	chapa("Habari yako ndugu");
}

// fibonacci example
chapa("Testing fibonacci...");

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

chapa(fibo(10));

// testing input
chapa("Testing input from user...");
acha salamu = fn() {
        acha jina = jaza("Unaitwa nani rafiki? ");
        rudisha "Mambo vipi " + jina;
    }

chapa(salamu());
