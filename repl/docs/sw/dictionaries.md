# Kamusi Katika Nuru

Kamusi katika Nuru ni miundo ya data inayotunza jozi za funguo-thamani. Ukurasa huu unatoa maelezo kuhusu Kamusi katika Nuru, ikiwemo namna ya kutengeneza, namna ya kubadilisha, na namna ya kuzunguka ndani yake.

## Kutengeneza Kamusi

Kamusi zinawekwa kwenye mabano singasinga na hujumuisha funguo na thamani zake zikitenganishwa na nukta pacha. Mfano wa uainishwaji wa kamusi:

```go

orodha = {"jina": "Juma", "umri": 25}
```

Funguo zinawezakua tungo, namba, desimali, au buliani na thamani inaweza kua aina ya data yoyote ikiwemo tungo, namba, desimali, buliani, tupu, au kitendakazi:

```go
k = {
    "jina": "Juma",
    "umri": 25,
    kweli: "kweli",
    "salimu": unda(x) { andika("habari", x) },
    "sina thamani": tupu
}
```

## Kupata Vipengele

Unaweza kupata vipengele vya kamusi kwa kutumia funguo zake:

```go
k = {
    "jina": "Juma",
    "umri": 25,
    kweli: "kweli",
    "salimu": unda(x) { andika("habari", x) },
    "sina thamani": tupu
}

andika(k[kweli]) // kweli
andika(k["salimu"]("Juma")) // habari Juma
```

## Kuboresha Vipengele

Boresha thamani ya kipengele kwa kukipa thamani mpya kwenye funguo yake:

```go
k = {
    "jina": "Juma",
    "umri": 25,
    kweli: "kweli",
    "salimu": unda(x) { andika("habari", x) },
    "sina thamani": tupu
}

k['umri'] = 30
andika(k['umri']) // 30
```

## Kuongeza Vipengele Vipya

Ongeza jozi mpya ya funguo-thamani kwenye kamusi kwa kuipa thamani funguo ambayo haipo kwenye kamusi husika:

```go
k["lugha"] = "Kiswahili"
andika(k["lugha"]) // Kiswahili
```

## Kuunganisha Kamusi

Unganisha kamusi mbili kwa kutumia kiendeshi `+`:

```go
matunda = {"a": "apple", "b": "banana"}
mboga = {"c": "tembele", "d": "mchicha"}
vyakula = matunda + mboga
andika(vyakula) // {"a": "apple", "b": "banana", "c": "tembele", "d": "mchicha"}
```

## Angalia Kama Funguo Ipo Kwenye Kamusi

Tumia neno msingi `ktk` kuangalia kama funguo ipo kwenye kamusi:

```go

k = {
    "jina": "Juma",
    "umri": 25,
    kweli: "kweli",
    "salimu": unda(x) { andika("habari", x) },
    "sina thamani": tupu
}

"umri" ktk k // kweli
"urefu" ktk k // sikweli
```

## Kuzunguka Ndani Ya Kamusi

Zunguka ndani ya kamusi kupata funguo na thamani zake:

```go

hobby = {"a": "kulala", "b": "kucheza mpira", "c": "kuimba"}

kwa i, v ktk hobby {
    andika(i, "=>", v)
}

//Output

a => kulala
b => kucheza mpira
c => kuimba
```

Kuzunguka ndani ya kamusi na kupata thamani peke yake:

```go

hobby = {"a": "kulala", "b": "kucheza mpira", "c": "kuimba"}

kwa i, v ktk hobby {
    andika(i, "=>", v)
}

//Output

kulala
kucheza mpira
kuimba
```

Kwa ufahamu huu, unaweza ukatumia kamusi kikamilifu katika Nuru kutunza na kusimamia jozi za funguo-thamani, na kupata namna nyumbufu ya kupangilia na kupata data katika programu zako.
