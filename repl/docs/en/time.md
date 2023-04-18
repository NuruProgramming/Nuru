# Time in Nuru

## Importing Time

To use Time in Nuru, you first have to import the `muda` module as follows:
```so
tumia muda
```

## Time Methods

### hasahivi()
To get the current time use the the `muda.hasahivi()` method. This will return a `muda` object with the current time:
```so
tumia muda

saivi = muda.hasahivi()
```

### tangu()
Use this method to get the total time elapsed in seconds. It accepts a time object or a string in the format `HH:mm:ss dd-MM-YYYY`:

```so
tumia muda

sasa = muda.hasahivi()

muda.tangu(s) // will return the elapsed time

// alternatively:

sasa.tangu("00:00:00 01-01-1900") // will return the elapsed time in seconds since that date
```

### lala()

Use lala if you want your program to sleep. It accepts one argument which is the total time to sleep in seconds:
```so
muda.lala(10) // will sleep for ten seconds
```

### ongeza()

Use this method to add to time, better explained with an example:
```so
tumia muda

sasa = muda.hasahivi()

kesho = sasa.ongeza(siku=1)
kesho_pia = sasa.ongeza(saa=24)
mwakani = sasa.ongeza(miaka=1)
miezi_tatu_mbele = sasa.ongeza(miezi = 3)
wiki_ijayo = sasa.ongeza(siku=7)
idi = sasa.ongeza(siku=3, masaa=4, dakika=50, sekunde=3)
```
It will return a muda object with the specified time.
