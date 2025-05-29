# Time in Nuru

## Importing Time

To use Time in Nuru, you first have to import the `muda` module as follows:

```so
tumia muda
```

---

## Time Methods

### `hasahivi()`

To get the current time, use `muda.hasahivi()`. It returns a `muda` object with the current time in the format `HH:mm:ss dd-MM-YYYY`.

```so
tumia muda

saivi = muda.hasahivi()
```

---

### `leo()`

To get today’s date in the format `dd-MM-YYYY`:

```so
tumia muda

leo = muda.leo()
```

---

### `tangu(time)`

Gets the total time elapsed **in seconds** from the given time to now. Accepts a `muda` object or string in `HH:mm:ss dd-MM-YYYY` format.

```so
tumia muda

muda_ulioyopita = muda.tangu("15:00:00 01-01-2024")
```

---

### `lala(sekunde)`

Pauses the program for the given number of seconds:

```so
tumia muda

muda.lala(5) // sleeps for 5 seconds
```

---

### `baada_ya(sekunde)`

Returns a `muda` object representing the time after the given number of seconds from now.

```so
tumia muda

baadaye = muda.baada_ya(60) // one minute from now
```

---

### `tofauti(muda1, muda2)`

Returns the difference between two time values in seconds.

```so
tumia muda

saa1 = muda.hasahivi()
saa2 = muda.baada_ya(30)

tofauti = muda.tofauti(saa2, saa1) // 30
```

---

### `ongeza(...)`

To add time to a `muda` object. You will be able to specify fields like `siku`, `saa`, `dakika`, `sekunde`, etc. Example:

```so
tumia muda

sasa = muda.hasahivi()
kesho = sasa.ongeza(siku=1)
mwakani = sasa.ongeza(miaka=1)
```
