## Explore Tabs  

#### Params:  
- `decade[]` - array of decades (2020, 2010, 1990, etc)
- `capo[]` - array of int/bool. 0 for no, 1 for yes
- `tuning[]` - array of ints, predefined (dumb!)
- `genres[]` - array of ints, predefined
- `difficulty[]` - array of ints, predefined
- `order` - string
- `date` - int, not sure what the purpose is? set to 0 on the "top 100" for "today" and "all time", otherwise ommitted.  

##### Order enum:  
- `hitstotal_desc`  
- `hitsdailygroup_desc`  
- `songname_asc`  
- `artistname_asc`  
- `rating_desc`  
- `hits`  
- `hits_daily`  
- Possibly also `typesort_asc`, `date_desc` although I haven't seen it being used yet  

#### Request examples:  

- Most popular all time: `https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&order=hitstotal_desc&genres[]=4`  
- Today: `https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&order=hitsdailygroup_desc&genres[]=4`  
- Recently added: `https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&order=date_desc&genres[]=4`  
- Song title A-Z: `https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&order=songname_asc&genres[]=4`  
- Artist A-Z: `https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&order=artistname_asc&genres[]=4`  
- Rating: `https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&order=rating_desc&genres[]=4`  
- Top all time: `https://api.ultimate-guitar.com/api/v1/tab/explore?date=0&genre=0&level=0&order=hits&page=1&type=800&official=0`  
- Top 100 tabs today: `https://api.ultimate-guitar.com/api/v1/tab/explore?date=0&genre=0&level=0&order=hits_daily&page=1&type=800&official=0`  


Example requests:  

```
https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&subgenres[]=24&subgenres[]=46&decade[]=2020&capo[]=1&tuning[]=1&tuning[]=12&tuning[]=2&tuning[]=10&tuning[]=6&tuning[]=11&tuning[]=9&tuning[]=4&tuning[]=3&genres[]=4

https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&decade[]=2020&capo[]=0&tuning[]=1&tuning[]=12&tuning[]=2&tuning[]=10&tuning[]=6&tuning[]=11&tuning[]=9&tuning[]=4&tuning[]=3&genres[]=4&subgenres[]=24&subgenres[]=46

https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&genres[]=666

https://api.ultimate-guitar.com/api/v1/tab/explore?explore=1&group=1&difficulty[]=4&difficulty[]=3&difficulty[]=2&difficulty[]=1
```