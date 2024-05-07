# PC Gamer News Simple Scraping Like Title and Url With GoColly
## Latest News Ever Seen In pcgamer.com

### News Endpoint
| Method HTTP | URL | Input | Output | Information |
| ------ | ------ | ------ | ------ | ------ |
| GET | /:year/:month |Parameter Year and Month | List of all news objects by year and month |  displays all news titles and urls |

## Response Endpoint
```JSON
    {
        "message": "success",
        "count": 764,
        "data": [
            {
              "title": "If you're looking for a new base in Enshrouded, this three-story tavern with a 25 comfort buff is the perfect fixer-upper",
              "url": "/if-youre-looking-for-a-new-base-in-enshrouded-this-three-story-tavern-with-a-25-comfort-buff-is-the-perfect-fixer-upper",
              "baseurl": "https://www.pcgamer.com/if-youre-looking-for-a-new-base-in-enshrouded-this-three-story-tavern-with-a-25-comfort-buff-is-the-perfect-fixer-upper/"
            },
            {
              "title": "Squint at the new Silent Hill 2 remake trailer and tell me it's not just the Resident Evil 2 remake",
              "url": "/squint-at-the-new-silent-hill-2-remake-trailer-and-tell-me-its-not-just-the-resident-evil-2-remake",
              "baseurl": "https://www.pcgamer.com/squint-at-the-new-silent-hill-2-remake-trailer-and-tell-me-its-not-just-the-resident-evil-2-remake/"
            },
        ]
    }
```

disclaimer for education purposes :)