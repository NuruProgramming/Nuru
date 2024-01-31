tumia hisabati

// Mbinu za Kupanga
//orodhesha(kwanza, mwisho, umbali), huunda orodha ya nambari na umbali uliowekwa kati yao.
fanya orodhesha = unda(kwanza, mwisho, umbali){
    fanya orodha = [kwanza];
    fanya i = kwanza + umbali;
    wakati(i < mwisho){
        orodha.sukuma(i);
        i += umbali;
    }
    rudisha orodha;
}

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
    andika("vekta: ",vekta)
    fanya tokeo = [];
    kwa v ktk vekta {
        tokeo.sukuma(v * (1 - v));
    }
    
    rudisha tokeo;
}

fanya kuzidishaMatrikiVekta = unda(matriki, vekta) {
    fanya tokeo = [];
    kwa row ktk matriki {
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
    fanya jumlaYaUzani = sigmoid(kuzidishaMatrikiVekta(pembejeo, mizani));
    rudisha jumlaYaUzani;
}

fanya badiliMatriki = unda(matrix) {
    // Pata idadi ya safu mlalo na safu wima katika matrix asili
    fanya nambari_ya_safu_mlalo = matrix.idadi();
    fanya nambari_ya_safu_wima = matrix[0].idadi();

    // Unda matrix mpya na safu mlalo na safu wima zilizobadilishwa
    fanya matrikiIliyobadili = [];

    // Pita ndani ya safu wima
    kwa safu_wima_ya, safuW ktk orodhesha(0, nambari_ya_safu_wima, 1){
        // Anzisha safu mlalo mpya kwa matriki iliyopitishwa
        fanya transposed_safu_mlalo = [];

        // Pita ndani ya safu mlalo
        kwa safu_mlalo_ya, safu ktk orodhesha(0, nambari_ya_safu_mlalo, 1){
            // Sukuma kipengele kwenye safu wima ya sasa na safu mlalo hadi safu iliyopitishwa
            transposed_safu_mlalo.sukuma(matrix[safu_mlalo_ya][safu_wima_ya]);
        }

        // Sukuma safu mlalo iliyopitishwa kwenye matriki lililopitishwa
        matrikiIliyobadili.sukuma(transposed_safu_mlalo);
    }

    rudisha matrikiIliyobadili;
}

// Funza mtandao wa neva
fanya funza = unda(mizani, mafunzoPembejeo, matokeoYaMafunzo, marudioYaMafunzo) {
    fanya kurudia = 0

    andika('\nmafunzoPembejeo: ');
    andika(mafunzoPembejeo);

    andika('\nmatokeoYaMafunzo: ');
    andika(matokeoYaMafunzo);

    
    fanya orodha = orodhesha(0, marudioYaMafunzo, 1)
    // andika("orodha: ",orodha)

    kwa i ktk orodha{
        andika('\n\nkurudia: ');
        andika(i);
       // Pitisha mafunzo yaliyowekwa kupitia mtandao wa neva
        fanya pato = waza(mafunzoPembejeo, mizani);

        andika('\npato: ');
        andika(pato);
        
        // Kuhesabu kiwango cha upungufu
        fanya upungufu = [];
        kwa i, kipengee ktk matokeoYaMafunzo {
            upungufu.sukuma(kipengee - pato[i]);           
        }

        andika('\nupungufu: ');
        andika(upungufu);

        fanya sigmoidDerivative_ya_pato = sigmoidDerivative(pato)

        andika('\nsigmoidDerivative tokeo: ');
        andika(sigmoidDerivative_ya_pato);

        fanya zidishaKwaNukta_tokeo = zidishaKwaNukta(upungufu, sigmoidDerivative_ya_pato);

        andika('\nzidishaKwaNukta tokeo: ');
        andika(zidishaKwaNukta_tokeo);
        
        fanya mafunzoPembejeoYaliyobadili = badiliMatriki(mafunzoPembejeo)

        andika('\nmafunzo pembejeo yaliyobadili: ');
        andika(mafunzoPembejeoYaliyobadili);

        // Kuzidisha upungufu kwa pembejeo na upinde rangi ya kitendakazi cha sigmoid
        // Uzito mdogo wa ujasiri hurekebishwa zaidi kupitia asili ya kazi
        fanya marekebisho = kuzidishaMatrikiVekta(mafunzoPembejeo, zidishaKwaNukta_tokeo);
        andika('\nmarekebisho tokeo: ');
        andika(marekebisho);


        // Rekebisha uzani
        kwa i, j ktk mizani {
            mizani[i] = mizani[i] + marekebisho[i];           
        }

        andika('\nmizani mpya: ');
        andika(mizani);
        kurudia++
    }

    rudisha mizani;
}


andika('\nMizani ya Kuanzisha isiyo na mpangilio: ');
andika(mizani);

// Seti ya mafunzo
fanya mafunzoPembejeo = [[0, 0, 1], [1, 1, 1], [1, 0, 1], [0, 1, 1]];
fanya matokeoYaMafunzo = [0, 1, 1, 0];

// Funza mtandao wa neva
fanya mafunzoMizani = funza(mizani, mafunzoPembejeo, matokeoYaMafunzo, 10000);

andika('\nMizani baada ya mafunzo:');
andika(mafunzoMizani);

// Ingizo la mtumiaji kwa hali mpya
fanya A = 0;
fanya B = 0;
fanya C = 1;

andika('\nHali mpya: data ya pembejeo = ', A, B, C);
andika('\nData ya pato:');
andika(waza([[A, B, C]], mafunzoMizani));
