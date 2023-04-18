# HTTP with Nuru

You can access the internet via http protocol using the `mtandao` module.

## Importing

Import the module with:
```
tumia mtandao
```

## Methods

### peruzi()

Use this as GET method. It can either accept one positional argument which will be the URL:

```
tumia mtandao

mtandao.peruzi("http://google.com")
```

Or you can use keyword arguments to pass in parameters and headers as shown below. Note that headers and parameters must be a dictionary:

```
tumia mtandao

url = "http://mysite.com"
headers = {"Authentication": "Bearer XXXX"}

mtandao.peruzi(yuareli=url, vichwa=headers, mwili=params)
```

### tuma()

Use this as POST method. Use keyword arguments to pass in parameters and headers as shown below. Note that headers and parameters must be a dictionary:

```
tumia mtandao

url = "http://mysite.com"
headers = {"Authentication": "Bearer XXXX"}
params = {"key": "Value"}

mtandao.tuma(yuareli=url, vichwa=headers, mwili=params)
```
