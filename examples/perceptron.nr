tumia hisabati

// Kuanzisha uzani bila mpangilio
fanya mizani = [
    hisabati.random() * 2 - 1,
    hisabati.random() * 2 - 1,
    hisabati.random() * 2 - 1
];

// Undo la uanzishaji wa Sigmoid
fanya sigmoid = unda(vekta) {
    fanya tokeo = [];
    kwa v ktk vekta {
        tokeo.sukuma(1 / (1 + hisabati.exp(-1 * v)));
    }
    rudisha tokeo;
}

// Derivative ya undo la sigmoid
fanya sigmoidDerivative  = unda(vekta) {
    fanya tokeo = [];
    kwa v ktk vekta {
        tokeo.sukuma(v * (1 - v));
    }
    
    rudisha tokeo;
}

fanya kuzidishaTumboVekta = unda(tumbo, vekta) {
    fanya tokeo = [];
    kwa row ktk tumbo {
        fanya jamii = 0;
        kwa j, kipengee ktk row {
            jamii += kipengee * vekta[j];
        }
        tokeo.sukuma(jamii);        
    }
    rudisha tokeo;
}

fanya zidishaKwaNukta = unda(safu1, safu2) {
    // Angalia ikiwa safu zina urefu sawa
    kama (safu1.idadi() != safu2.idadi()) {
        andika("Safu lazima ziwe na urefu sawa kwa kuzidisha kwa busara ya kipengele.");
    }

    // Perform element-wise multiplication
    fanya tokeo = [];
    kwa i, kipengee ktk safu1 {
        tokeo.sukuma(kipengee * safu2[i]);
    }
    rudisha tokeo;
}

// Songa mbele kupitia mtandao wa neva
fanya waza = unda(pembejeo, mizani) {
    fanya jumlaYaUzani = sigmoid(kuzidishaTumboVekta(pembejeo, mizani));
    rudisha jumlaYaUzani;
}

// Funza mtandao wa neva
fanya funza = unda(mizani, mafunzoPembejeo, matokeoYaMafunzo, marudioYaMafunzo) {
    fanya kurudia = 0

    andika('mafunzoPembejeo: ');
    andika(mafunzoPembejeo);

    andika('matokeoYaMafunzo: ');
    andika(matokeoYaMafunzo);

    

    wakati (kurudia < marudioYaMafunzo) {
        andika('kurudia: ');
        andika(kurudia);
       // Pitisha mafunzo yaliyowekwa kupitia mtandao wa neva
        fanya pato = waza(mafunzoPembejeo, mizani);

        andika('pato: ');
        andika(pato);
        // Kuhesabu kiwango cha makosa
        fanya upungufu = [];
        kwa i, kipengee ktk matokeoYaMafunzo {
            upungufu.sukuma(kipengee - pato[i]);           
        }

        fanya sigmoidDerivative = sigmoidDerivative(pato)

        andika('upungufu: ');
        andika(upungufu);

        andika('sigmoidDerivative tokeo: ');
        andika(sigmoidDerivative);

        fanya zidishaKwaNukta = zidishaKwaNukta(upungufu, sigmoidDerivative);

        andika('zidishaKwaNukta tokeo: ');
        andika(zidishaKwaNukta);

        // Kuzidisha makosa kwa pembejeo na upinde rangi ya kitendakazi cha sigmoid
        // Uzito mdogo wa ujasiri hurekebishwa zaidi kupitia asili ya kazi
        fanya marekebisho = kuzidishaTumboVekta(mafunzoPembejeo, zidishaKwaNukta);
        andika('marekebisho tokeo: ');
        andika(marekebisho);


        // Rekebisha uzani
        kwa i, j ktk mizani {
            mizani[i] = mizani[i] + marekebisho[i];           
        }

        andika('mizani mpya: ');
        andika(mizani);
        kurudia++
    }
    rudisha mizani;
}



andika('Mizani ya Kuanzisha isiyo na mpangilio: ');
andika(mizani);

// Seti ya mafunzo
fanya mafunzoPembejeo = [[0, 0, 1], [1, 1, 1], [1, 0, 1], [0, 1, 1]];
fanya matokeoYaMafunzo = [0, 1, 1, 0];

// Funza mtandao wa neva
fanya mafunzoMizani = funza(mizani, mafunzoPembejeo, matokeoYaMafunzo, 10000);

andika('Mizani baada ya mafunzo:');
andika(mafunzoMizani);

// Ingizo la mtumiaji kwa hali mpya
fanya A = 1;
fanya B = 0;
fanya C = 0;

andika('Hali mpya: data ya pembejeo = ', A, B, C);
andika('Data ya pato:');
andika(waza([[A, B, C]], mafunzoMizani));
