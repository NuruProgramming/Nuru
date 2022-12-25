andika("Testing basic types...");
andika(2 + 2);
andika(4 * 4);
fanya a = 10;
fanya b = 20;

andika(a + b);

andika(kweli == sikweli);
andika("Testing empty lines...")
andika();
andika();
andika();
andika("Mambo vipi");

andika("Testing Functions... ");

fanya jumlisha = unda(x, y) {x + y};

andika(jumlisha(20,30));
andika(jumlisha(100,1000));

fanya zidisha = unda(x, y) {x * y};

andika(zidisha(100,1000));
andika(zidisha(200, 20));

// lists can hold any value
andika("Testing lists...");
fanya list = [1, "a", kweli, sikweli];

// a few builtins

fanya list = sukuma(list, jumlisha(4,5));

andika(list);
andika(list[2]);
andika(list[-100]); // prints null
andika(idadi(list));
andika(yamwisho(list));
andika([1,2,3] + [4,5,6]);
list[0] = 1000
andika(list[0])

// if statements
andika("Testing if statements...");
kama (1 > 2) {
	andika("Moja ni zaidi ya mbili");
} sivyo {
	andika("Moja si zaidi ya mbili");
}

kama (idadi("Habari") == 6) {
	andika("Habari yako ndugu");
}

// fibonacci example
andika("Testing fibonacci...");

fanya fibo = unda(x) {
	kama (x == 0) {
		rudisha 0;
	} au kama (x == 1) {
			rudisha 1;
	} sivyo {
			rudisha fibo(x - 1) + fibo(x - 2);
	}
}


andika(fibo(10));

// testing input
andika("Testing input from user...");
fanya salamu = unda() {
        fanya jina = jaza("Unaitwa nani rafiki? ");
        rudisha "Mambo vipi " + jina;
    }

andika(salamu());

/*
Multiline comment
*/

// test dictionaries

andika("Testing dictionaries...")

fanya watu = [{"jina": "Mojo", "kabila": "Mnyakusa"}, {"jina": "Avi", "kabila": "Mwarabu wa dubai"}]

andika(watu, watu[0], watu[0]["jina"], watu[0]["kabila"])

watu[0]["jina"] = "MJ";
andika(watu[0]["jina"]);

andika({"a":1} + {"b": 2})
// testing while loop

andika("Testing while loop...");

fanya i = 10;

wakati (i > 0) {
	andika(i);
	i = i - 1;
}
