//! Hii ni mwandiko wa csv.
//! Fomati hii imewekwa kawaida kwenye barua hii: https://www.rfc-editor.org/rfc/rfc4180
//! Fomati hii husomwa laini kwa laini vipengele hugauwa kutumia herufi ','
//! Kwa mfano:
//!     nambari,jina,umri
//! 	1,Nuru,2
//! Laini ya kwanza itatumika kama jina majina ya rekodi.
//!

/// Hii inasoma maneno uliyopeana na kurudisha safu ya ilivyopata.
/// Kama maneno hayapo, itarudisha tupu.
soma = unda(maneno=tupu,kigaio=',') {
	// Tunafaa kupata namna ya kurudisha makosa badala ya tupu.
	// Lakini kwa sasa, rudisha tupu.
	kama (maneno == tupu || maneno == "" || !kigawanyio_halali(kigaio)) {
		rudisha tupu
	}
	
	laini_zilizomo = laini(maneno)

	kama (laini_zilizomo.idadi() <= 0) {
		rudisha tupu
	}

	laini_csv = []

	kwa klaini ktk laini_zilizomo {
		laini_csv.sukuma(klaini)
	}

	rudisha laini_csv
}

/// Kagua kama kigawanyio ni halali ama si halali.
kigawanyio_halali = unda(kigawanyio) {
	kama (kigawanyio == '\r' || kigawanyio == '\n' || kigawanyio == '"') {
		rudisha sikweli
	}

	rudisha kweli
}

/// Gawa laini.
/// Tukipata '"', tunaendelea mbaka tupate '"' ya kufunga ndipo tuichukue kama laini.
laini = unda(maneno) {
	laini_zilizomo = []
	laini_sahii = ""
	tumepata_kindelezo = sikweli

	kwa neno ktk maneno {
		// Tunafaa kuangalia kama laini haina kitu (ukishatoa nafasi, lakini hiyo ni ya baadaye)
		kama (neno == '\n' && laini_sahii.idadi() > 0 && !tumepata_kindelezo){
			laini_zilizomo.sukuma(laini_sahii)
			laini_sahii = ""
			endelea
		}

		// Tumepata kiendelezo kwa mara ya kwanza
		kama (neno == '"' && !tumepata_kindelezo) {
			tumepata_kindelezo = kweli
			endelea
		}

		// Tumepata kiendelezo mara ya pili, iweke isiwe kweli.
		kama (neno == '"' && tumepata_kindelezo) {
			tumepata_kindelezo = sikweli
			endelea
		}

		laini_sahii += neno
	}

	// Kama iliisha kabla ya kuisha
	kama (laini_sahii.idadi() > 0) {
		laini_zilizomo.sukuma(laini_sahii)
	}

	rudisha laini_zilizomo
}

/// Gawa laini kulingana na kigawanyio kilichopeanwa.
/// Inarudisha safu na kama laini hii imeisha ama bado.
/// Laini inaweza kuwa kuliko laini moja kulingana na rfc4180 kama imeisha kabla ya `"` kufungwa.
/// Hii inamaanisha kuwa tunahitaji kuwa makini tunapo gawa laini
gawa_laini = unda(laini,kigawanyio) {
	vigawio = []
	kigao = ""
	kwa neno ktk laini {
		kama (neno == kigawanyio) {
			vigawio = kigao
			kigao = ""
			endelea
		}

		kigao += neno
	}

	rudisha vigawio
}

// hii ni ya kuangalia kama inafanya inavyo itajika
/*
mfano = "namba,jina,umri\n"
mfano += "1,yangu,32\n"
mfano += '2,"kompyuta yetu",34\n'
mfano += '3, "anza\n'
mfano += 'maliza", 54'
kwa k ktk soma(mfano) {
	andika(k)
}
*/
