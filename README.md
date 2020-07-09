# Devaram-API

Go app that serves RESTful JSON web API for the Devaram.

Devaram means garland of poems to Lord Shiva.

Devaram denotes the first 7 volumes of the Tirumurai, the twelve-volume collection of Śaiva devotional poetry. These volumes contain the works of the three most prominent Tamil poets of the 7th and 8th centuries: Sambandar, Appar, and Sundarar.

தேவாரம் எனப்படுபவை சைவ சமயத்தின் முழுமுதற் கடவுளான சிவபெருமான் மீது, திருஞானசம்பந்தமூர்த்தி நாயனார்,திருநாவுக்கரசு நாயனார், சுந்தரமூர்த்தி நாயனார் ஆகிய மூன்று இறையடியார்களால் தமிழிற் பாடப்பட்ட பாடல்கள் ஆகும். 


| BookNum  | BookName  | Pathigams  | Verses  | Temples  |
|---|---|---:|---:|---:|
|  1 |  முதல் திருமுறை   |  136 |  1469 | 88  |
|  2 |  இரண்டாம் திருமுறை   | 122  | 1331  | 90  |
|  3 | மூன்றாம் திருமுறை  | 126  | 1358  | 85  |
|  4 | நான்காம் திருமுறை  | 113  | 1070  | 50  |
|  5 | ஐந்தாம் திருமுறை  |100   | 1015  | 76  |
|  6 | ஆறாம் திருமுறை  | 99  | 981  | 65  |
|  7 | ஏழாம் திருமுறை  |100   |1026   |84   |
|  8 | எட்டாம் திருமுறை  | 76  | 1058  | - |
|  9 | ஒன்பதாம் திருமுறை  | 19  | 301  | 14  |
| 10  | பத்தாம் திருமுறை  | 1237  | 3000  | -  |
| 11  | பதினொன்றாம் திருமுறை  | 40  | 1385  | -  |
| 12  | பன்னிரண்டாம் திருமுறை  | 71  | 4272  | -  |
|Total|| 2239|18,266|552|




## Install

```
export DEVARAM_API_DB="./devaram.db"
go run main.go
```


## Usage
To retrieve a single verse:

https://localhost:4567/BOOK/CHAPTER/VERSE


To retrieve a range of verses in a chapter:

https://localhost:4567/BOOK/CHAPTER/VERSESTART-VERSEEND


You can use the above URLs with book/chapter/verse format. It mostly just works. (Though, if it doesn't, you can let me know!)



## Copyright

Copyright Rajkumar Palani. Licensed MIT.
