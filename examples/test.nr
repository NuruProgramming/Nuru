tumia pima

unda pata_mraba(namba) {
    rudisha namba * namba;
}

unda pima_mraba() {
    fanya x = pata_mraba(5);
    pima.hakiki(x == 25, "Hesabu ya mraba sio sahihi");

    x = pata_mraba(10);
    pima.hakiki(x == 100, "Hesabu ya mraba sio sahihi");

    x = pata_mraba(100);
    pima.hakiki(x == 10000, "Hesabu ya mraba sio sahihi");

}

unda pima_mraba2() {
    fanya x = pata_mraba(5);
    pima.hakiki(x == 25, "Hesabu ya mraba sio sahihi");
}

unda pima_mraba3() {
    fanya x = pata_mraba(5);
    pima.hakiki(x == 25, "Hesabu ya mraba sio sahihi");
}

unda pima_mraba4() {
    fanya x = pata_mraba(5);
    pima.hakiki(x == 25, "Hesabu ya mraba sio sahihi");
}

unda pima_ambayo_itafeli() {
    pima.hakiki(pata_mraba(9) == 16, "Hesabu ya mraba sio sahihi");
    pima.hakiki(pata_mraba(10) == 16, "Hesabu ya mraba sio sahihi");
}


unda pima_ambayo_itafeli2() {
    pima.hakiki(pata_mraba(10) == 10 * 11, "Hesabu ya mraba sio sahihi");
}

unda pima_ambayo_itafeli3() {
    pima.hakiki(pata_mraba(10) == 10 * 11, "Hesabu ya mraba sio sahihi");
}
