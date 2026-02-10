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

To add time to a `muda` object. You must specify at least one of the following fields `sekunde`, `dakika`, `masaa`, `siku`, `wiki`, `miezi`, `miaka`.

Example:

```so
tumia muda

sasa = muda.hasahivi()
kesho = sasa.ongeza(siku=1)
mwakani = sasa.ongeza(miaka=1)
```

---

### `siku(tarehe)` / `siku(mwaka, mwezi, siku)`

Returns a **tarehe** (date-only) object. One argument: string in `2006-01-02` format. Or three integer arguments: year, month, day. The date object has **panga(muundo)** to format (e.g. `d.panga("02-01-2006")`).

---

### `panga(muundo)`

Formats the `muda` object as a string using the given layout. One argument: a format string (e.g. `"02-01-2006"` for date only, `"15:04:05"` for time only). Uses Go-style layout; see [time package](https://pkg.go.dev/time#Time.Format).

```so
tumia muda
sasa = muda.hasahivi()
sasa.panga("02-01-2006")  // date part only
sasa.panga("15:04:05")   // time part only
```
