# Get tab applicature  

&chords[A]=A&chords[Asus4]=Asus4&chords[A6]=A6&chords[A7]=A7&chords[Amaj7]=Amaj7&chords[A7sus4]=A7sus4&chords[A/C#]=A/C#&chords[Am]=Am&chords[Am7]=Am7&chords[Am/C]=Am/C

Applicature: `https://api.ultimate-guitar.com/api/v1/tab/applicature?instrument=guitar&tuning=E%20A%20D%20G%20B%20E&chords%5BEm%5D=Em&chords%5BG7%5D=G7&chords%5BD%2FF%23%5D=D%2FF%23&chords%5BBm%5D=Bm&chords%5BG%5D=G`

Params: `instrument=guitar&tuning=E A D G B E&chords[Em]=Em&chords[G7]=G7&chords[D/F#]=D/F#&chords[Bm]=Bm&chords[G]=G`  

Response:  
```json
[{
  "chord": "Em",
  "variations": [{
    "id": "022000",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [52, 47, 43, 40, 35, 28],
    "frets": [0, 0, 0, 2, 2, 0],
    "fingers": [0, 0, 0, 2, 1, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "022003",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [55, 47, 43, 40, 35, 28],
    "frets": [3, 0, 0, 2, 2, 0],
    "fingers": [3, 0, 0, 2, 1, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "xx2000",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [52, 47, 43, 40, -1, -1],
    "frets": [0, 0, 0, 2, -1, -1],
    "fingers": [0, 0, 0, 1, 0, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "022403",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [55, 47, 47, 40, 35, 28],
    "frets": [3, 0, 4, 2, 2, 0],
    "fingers": [3, 0, 4, 2, 1, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "xx2003",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [55, 47, 43, 40, -1, -1],
    "frets": [3, 0, 0, 2, -1, -1],
    "fingers": [2, 0, 0, 1, 0, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "022453",
    "listCapos": [{
      "fret": 2,
      "startString": 3,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [55, 52, 47, 40, 35, 28],
    "frets": [3, 5, 4, 2, 2, 0],
    "fingers": [2, 4, 3, 0, 0, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "xx2403",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [55, 47, 47, 40, -1, -1],
    "frets": [3, 0, 4, 2, -1, -1],
    "fingers": [2, 0, 3, 1, 0, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "xx2453",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [55, 52, 47, 40, -1, -1],
    "frets": [3, 5, 4, 2, -1, -1],
    "fingers": [2, 4, 3, 1, 0, 0],
    "fret": 2,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "079987",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [59, 55, 52, 47, 40, 28],
    "frets": [7, 8, 9, 9, 7, 0],
    "fingers": [0, 2, 4, 3, 0, 0],
    "fret": 7,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "x79987",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [59, 55, 52, 47, 40, -1],
    "frets": [7, 8, 9, 9, 7, -1],
    "fingers": [0, 2, 4, 3, 0, 0],
    "fret": 7,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "010991212",
    "listCapos": [{
      "fret": 9,
      "startString": 2,
      "lastString": 3,
      "finger": 1
    }, {
      "fret": 12,
      "startString": 0,
      "lastString": 1,
      "finger": 4
    }],
    "noteIndex": 4,
    "notes": [64, 59, 52, 47, 43, 28],
    "frets": [12, 12, 9, 9, 10, 0],
    "fingers": [0, 0, 0, 0, 2, 0],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "0109121212",
    "listCapos": [{
      "fret": 12,
      "startString": 0,
      "lastString": 2,
      "finger": 4
    }],
    "noteIndex": 4,
    "notes": [64, 59, 55, 47, 43, 28],
    "frets": [12, 12, 12, 9, 10, 0],
    "fingers": [0, 0, 0, 1, 2, 0],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "02200x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 47, 43, 40, 35, 28],
    "frets": [-1, 0, 0, 2, 2, 0],
    "fingers": [0, 0, 0, 2, 1, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "07998x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 55, 52, 47, 40, 28],
    "frets": [-1, 8, 9, 9, 7, 0],
    "fingers": [0, 2, 4, 3, 1, 0],
    "fret": 7,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "x7998x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 55, 52, 47, 40, -1],
    "frets": [-1, 8, 9, 9, 7, -1],
    "fingers": [0, 2, 4, 3, 1, 0],
    "fret": 7,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "010998x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 55, 52, 47, 43, 28],
    "frets": [-1, 8, 9, 9, 10, 0],
    "fingers": [0, 1, 3, 2, 4, 0],
    "fret": 8,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "0109912x",
    "listCapos": [{
      "fret": 9,
      "startString": 2,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [-1, 59, 52, 47, 43, 28],
    "frets": [-1, 12, 9, 9, 10, 0],
    "fingers": [0, 4, 0, 0, 2, 0],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "01091212x",
    "listCapos": [{
      "fret": 12,
      "startString": 1,
      "lastString": 2,
      "finger": 4
    }],
    "noteIndex": 4,
    "notes": [-1, 59, 55, 47, 43, 28],
    "frets": [-1, 12, 12, 9, 10, 0],
    "fingers": [0, 0, 0, 1, 2, 0],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "0220xx",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, -1, 43, 40, 35, 28],
    "frets": [-1, -1, 0, 2, 2, 0],
    "fingers": [0, 0, 0, 2, 1, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "0254xx",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, -1, 47, 43, 35, 28],
    "frets": [-1, -1, 4, 5, 2, 0],
    "fingers": [0, 0, 3, 4, 1, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "0754xx",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, -1, 47, 43, 40, 28],
    "frets": [-1, -1, 4, 5, 7, 0],
    "fingers": [0, 0, 1, 2, 4, 0],
    "fret": 4,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "01099xx",
    "listCapos": [{
      "fret": 9,
      "startString": 2,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [-1, -1, 52, 47, 43, 28],
    "frets": [-1, -1, 9, 9, 10, 0],
    "fingers": [0, 0, 0, 0, 2, 0],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "010912xx",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, -1, 55, 47, 43, 28],
    "frets": [-1, -1, 12, 9, 10, 0],
    "fingers": [0, 0, 4, 1, 2, 0],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "121099xx",
    "listCapos": [{
      "fret": 9,
      "startString": 2,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [-1, -1, 52, 47, 43, 40],
    "frets": [-1, -1, 9, 9, 10, 12],
    "fingers": [0, 0, 0, 0, 2, 4],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "x7x987",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [59, 55, 52, -1, 40, -1],
    "frets": [7, 8, 9, -1, 7, -1],
    "fingers": [0, 2, 3, 0, 0, 0],
    "fret": 7,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "0x200x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 47, 43, 40, -1, 28],
    "frets": [-1, 0, 0, 2, -1, 0],
    "fingers": [0, 0, 0, 1, 0, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "G7",
  "variations": [{
    "id": "320001",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 47, 43, 38, 35, 31],
    "frets": [1, 0, 0, 0, 2, 3],
    "fingers": [1, 0, 0, 0, 2, 3],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "320031",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 50, 43, 38, 35, 31],
    "frets": [1, 3, 0, 0, 2, 3],
    "fingers": [1, 4, 0, 0, 2, 3],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "320401",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 47, 47, 38, 35, 31],
    "frets": [1, 0, 4, 0, 2, 3],
    "fingers": [1, 0, 4, 0, 2, 3],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "323001",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 47, 43, 41, 35, 31],
    "frets": [1, 0, 0, 3, 2, 3],
    "fingers": [1, 0, 0, 4, 2, 3],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "323003",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [55, 47, 43, 41, 35, 31],
    "frets": [3, 0, 0, 3, 2, 3],
    "fingers": [4, 0, 0, 3, 1, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "353433",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [55, 50, 47, 41, 38, 31],
    "frets": [3, 3, 4, 3, 5, 3],
    "fingers": [0, 0, 2, 0, 3, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "353463",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [55, 53, 47, 41, 38, 31],
    "frets": [3, 6, 4, 3, 5, 3],
    "fingers": [0, 4, 2, 0, 3, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "xx5463",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [55, 53, 47, 43, -1, -1],
    "frets": [3, 6, 4, 5, -1, -1],
    "fingers": [1, 4, 2, 3, 0, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "xx5467",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [59, 53, 47, 43, -1, -1],
    "frets": [7, 6, 4, 5, -1, -1],
    "fingers": [4, 3, 1, 2, 0, 0],
    "fret": 4,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "xx5767",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [59, 53, 50, 43, -1, -1],
    "frets": [7, 6, 7, 5, -1, -1],
    "fingers": [4, 2, 3, 1, 0, 0],
    "fret": 5,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x1012101210",
    "listCapos": [{
      "fret": 10,
      "startString": 0,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [62, 59, 53, 50, 43, -1],
    "frets": [10, 12, 10, 12, 10, -1],
    "fingers": [0, 4, 0, 3, 0, 0],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x1012121213",
    "listCapos": [{
      "fret": 12,
      "startString": 1,
      "lastString": 3,
      "finger": 3
    }],
    "noteIndex": 7,
    "notes": [65, 59, 55, 50, 43, -1],
    "frets": [13, 12, 12, 12, 10, -1],
    "fingers": [4, 0, 0, 0, 1, 0],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x1012101213",
    "listCapos": [{
      "fret": 10,
      "startString": 2,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [65, 59, 53, 50, 43, -1],
    "frets": [13, 12, 10, 12, 10, -1],
    "fingers": [4, 3, 0, 2, 0, 0],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "32300x",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, 47, 43, 41, 35, 31],
    "frets": [-1, 0, 0, 3, 2, 3],
    "fingers": [0, 0, 0, 3, 1, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "32303x",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, 50, 43, 41, 35, 31],
    "frets": [-1, 3, 0, 3, 2, 3],
    "fingers": [0, 4, 0, 3, 1, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "32340x",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, 47, 47, 41, 35, 31],
    "frets": [-1, 0, 4, 3, 2, 3],
    "fingers": [0, 0, 4, 3, 1, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "35343x",
    "listCapos": [{
      "fret": 3,
      "startString": 1,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [-1, 50, 47, 41, 38, 31],
    "frets": [-1, 3, 4, 3, 5, 3],
    "fingers": [0, 0, 2, 0, 3, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "35346x",
    "listCapos": [{
      "fret": 3,
      "startString": 3,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [-1, 53, 47, 41, 38, 31],
    "frets": [-1, 6, 4, 3, 5, 3],
    "fingers": [0, 4, 2, 0, 3, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x109108x",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, 55, 53, 47, 43, -1],
    "frets": [-1, 8, 10, 9, 10, -1],
    "fingers": [0, 1, 4, 2, 3, 0],
    "fret": 8,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x10121012x",
    "listCapos": [{
      "fret": 10,
      "startString": 2,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [-1, 59, 53, 50, 43, -1],
    "frets": [-1, 12, 10, 12, 10, -1],
    "fingers": [0, 4, 0, 3, 0, 0],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3230xx",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, -1, 43, 41, 35, 31],
    "frets": [-1, -1, 0, 3, 2, 3],
    "fingers": [0, 0, 0, 3, 1, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3234xx",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, -1, 47, 41, 35, 31],
    "frets": [-1, -1, 4, 3, 2, 3],
    "fingers": [0, 0, 4, 3, 1, 2],
    "fret": 2,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3534xx",
    "listCapos": [{
      "fret": 3,
      "startString": 3,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [-1, -1, 47, 41, 38, 31],
    "frets": [-1, -1, 4, 3, 5, 3],
    "fingers": [0, 0, 2, 0, 3, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x0001",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 47, 43, 38, -1, 31],
    "frets": [1, 0, 0, 0, -1, 3],
    "fingers": [1, 0, 0, 0, 0, 3],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x0401",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 47, 47, 38, -1, 31],
    "frets": [1, 0, 4, 0, -1, 3],
    "fingers": [1, 0, 4, 0, 0, 3],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x3001",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 47, 43, 41, -1, 31],
    "frets": [1, 0, 0, 3, -1, 3],
    "fingers": [1, 0, 0, 4, 0, 3],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x0431",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 50, 47, 38, -1, 31],
    "frets": [1, 3, 4, 0, -1, 3],
    "fingers": [1, 3, 4, 0, 0, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x3401",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [53, 47, 47, 41, -1, 31],
    "frets": [1, 0, 4, 3, -1, 3],
    "fingers": [1, 0, 4, 3, 0, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x3003",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [55, 47, 43, 41, -1, 31],
    "frets": [3, 0, 0, 3, -1, 3],
    "fingers": [3, 0, 0, 2, 0, 1],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x3403",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [55, 47, 47, 41, -1, 31],
    "frets": [3, 0, 4, 3, -1, 3],
    "fingers": [3, 0, 4, 2, 0, 1],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x3433",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [55, 50, 47, 41, -1, 31],
    "frets": [3, 3, 4, 3, -1, 3],
    "fingers": [0, 0, 2, 0, 0, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x3463",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [55, 53, 47, 41, -1, 31],
    "frets": [3, 6, 4, 3, -1, 3],
    "fingers": [0, 4, 2, 0, 0, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x5463",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [55, 53, 47, 43, -1, 31],
    "frets": [3, 6, 4, 5, -1, 3],
    "fingers": [0, 4, 2, 3, 0, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x10x101210",
    "listCapos": [{
      "fret": 10,
      "startString": 0,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [62, 59, 53, -1, 43, -1],
    "frets": [10, 12, 10, -1, 10, -1],
    "fingers": [0, 3, 0, 0, 0, 0],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x10x101213",
    "listCapos": [{
      "fret": 10,
      "startString": 2,
      "lastString": 4,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [65, 59, 53, -1, 43, -1],
    "frets": [13, 12, 10, -1, 10, -1],
    "fingers": [4, 3, 0, 0, 0, 0],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x10x121213",
    "listCapos": [{
      "fret": 12,
      "startString": 1,
      "lastString": 2,
      "finger": 3
    }],
    "noteIndex": 7,
    "notes": [65, 59, 55, -1, 43, -1],
    "frets": [13, 12, 12, -1, 10, -1],
    "fingers": [4, 0, 0, 0, 1, 0],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x300x",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, 47, 43, 41, -1, 31],
    "frets": [-1, 0, 0, 3, -1, 3],
    "fingers": [0, 0, 0, 2, 0, 1],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x340x",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, 47, 47, 41, -1, 31],
    "frets": [-1, 0, 4, 3, -1, 3],
    "fingers": [0, 0, 3, 2, 0, 1],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x343x",
    "listCapos": [{
      "fret": 3,
      "startString": 1,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [-1, 50, 47, 41, -1, 31],
    "frets": [-1, 3, 4, 3, -1, 3],
    "fingers": [0, 0, 2, 0, 0, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x346x",
    "listCapos": [{
      "fret": 3,
      "startString": 3,
      "lastString": 5,
      "finger": 1
    }],
    "noteIndex": 7,
    "notes": [-1, 53, 47, 41, -1, 31],
    "frets": [-1, 6, 4, 3, -1, 3],
    "fingers": [0, 4, 2, 0, 0, 0],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "3x546x",
    "listCapos": [],
    "noteIndex": 7,
    "notes": [-1, 53, 47, 43, -1, 31],
    "frets": [-1, 6, 4, 5, -1, 3],
    "fingers": [0, 4, 2, 3, 0, 1],
    "fret": 3,
    "type": "_7",
    "baseDisplayNote": ""
  }]
}]
```

Ukulele: `https://api.ultimate-guitar.com/api/v1/tab/applicature?instrument=ukulele&tuning=G%20C%20E%20A&chords%5BDm%5D=Dm&chords%5BE7%5D=E7&chords%5BE%5D=E&chords%5BF7%5D=F7&chords%5BD%5D=D&chords%5BC%5D=C&chords%5BAm%5D=Am&chords%5BF%5D=F&chords%5BC%2FE%5D=C%2FE`  

Params: `instrument=ukulele&tuning=G C E A&chords[Dm]=Dm&chords[E7]=E7&chords[E]=E&chords[F7]=F7&chords[D]=D&chords[C]=C&chords[Am]=Am&chords[F]=F&chords[C/E]=C/E`  

Response:  

```json
[{
  "chord": "Dm",
  "variations": [{
    "id": "2210",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [57, 53, 50, 57],
    "frets": [0, 1, 2, 2],
    "fingers": [0, 1, 3, 2],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "2555",
    "listCapos": [{
      "fret": 5,
      "startString": 0,
      "lastString": 2,
      "finger": 4
    }],
    "noteIndex": 2,
    "notes": [62, 57, 53, 57],
    "frets": [5, 5, 5, 2],
    "fingers": [0, 0, 0, 1],
    "fret": 2,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "7555",
    "listCapos": [{
      "fret": 5,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [62, 57, 53, 62],
    "frets": [5, 5, 5, 7],
    "fingers": [0, 0, 0, 3],
    "fret": 5,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "7558",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [65, 57, 53, 62],
    "frets": [8, 5, 5, 7],
    "fingers": [4, 0, 0, 3],
    "fret": 5,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "x555",
    "listCapos": [{
      "fret": 5,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [62, 57, 53, -1],
    "frets": [5, 5, 5, -1],
    "fingers": [0, 0, 0, 0],
    "fret": 5,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "7x58",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [65, 57, -1, 62],
    "frets": [8, 5, -1, 7],
    "fingers": [4, 1, 0, 3],
    "fret": 5,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "79108",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [65, 62, 57, 62],
    "frets": [8, 10, 9, 7],
    "fingers": [2, 4, 3, 1],
    "fret": 7,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "109108",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [65, 62, 57, 65],
    "frets": [8, 10, 9, 10],
    "fingers": [1, 4, 2, 3],
    "fret": 8,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "x9108",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [65, 62, 57, -1],
    "frets": [8, 10, 9, -1],
    "fingers": [1, 3, 2, 0],
    "fret": 8,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "10x1012",
    "listCapos": [{
      "fret": 10,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [69, 62, -1, 65],
    "frets": [12, 10, -1, 10],
    "fingers": [3, 0, 0, 0],
    "fret": 10,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "221x",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [-1, 53, 50, 57],
    "frets": [-1, 1, 2, 2],
    "fingers": [0, 1, 3, 2],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "755x",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [-1, 57, 53, 62],
    "frets": [-1, 5, 5, 7],
    "fingers": [0, 0, 0, 3],
    "fret": 5,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "10910x",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [-1, 62, 57, 65],
    "frets": [-1, 10, 9, 10],
    "fingers": [0, 3, 1, 2],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "E7",
  "variations": [{
    "id": "1202",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [59, 52, 50, 56],
    "frets": [2, 0, 2, 1],
    "fingers": [3, 0, 2, 1],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "4445",
    "listCapos": [{
      "fret": 4,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [62, 56, 52, 59],
    "frets": [5, 4, 4, 4],
    "fingers": [2, 0, 0, 0],
    "fret": 4,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "7445",
    "listCapos": [{
      "fret": 4,
      "startString": 1,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [62, 56, 52, 62],
    "frets": [5, 4, 4, 7],
    "fingers": [2, 0, 0, 4],
    "fret": 4,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "7877",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [64, 59, 56, 62],
    "frets": [7, 7, 8, 7],
    "fingers": [0, 0, 2, 0],
    "fret": 7,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "78107",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [64, 62, 56, 62],
    "frets": [7, 10, 8, 7],
    "fingers": [0, 4, 2, 0],
    "fret": 7,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "98107",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [64, 62, 56, 64],
    "frets": [7, 10, 8, 9],
    "fingers": [1, 4, 2, 3],
    "fret": 7,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x8107",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [64, 62, 56, -1],
    "frets": [7, 10, 8, -1],
    "fingers": [1, 4, 2, 0],
    "fret": 7,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "981011",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [68, 62, 56, 64],
    "frets": [11, 10, 8, 9],
    "fingers": [4, 3, 1, 2],
    "fret": 8,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "9111011",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [68, 62, 59, 64],
    "frets": [11, 10, 11, 9],
    "fingers": [4, 2, 3, 1],
    "fret": 9,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "9x1011",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [68, 62, -1, 64],
    "frets": [11, 10, -1, 9],
    "fingers": [3, 2, 0, 1],
    "fret": 9,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "13141211",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [68, 64, 62, 68],
    "frets": [11, 12, 14, 13],
    "fingers": [1, 2, 4, 3],
    "fret": 11,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x141211",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [68, 64, 62, -1],
    "frets": [11, 12, 14, -1],
    "fingers": [1, 2, 4, 0],
    "fret": 11,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "120x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 52, 50, 56],
    "frets": [-1, 0, 2, 1],
    "fingers": [0, 0, 2, 1],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "744x",
    "listCapos": [{
      "fret": 4,
      "startString": 1,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [-1, 56, 52, 62],
    "frets": [-1, 4, 4, 7],
    "fingers": [0, 0, 0, 4],
    "fret": 4,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "9810x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 62, 56, 64],
    "frets": [-1, 10, 8, 9],
    "fingers": [0, 3, 1, 2],
    "fret": 8,
    "type": "_7",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "E",
  "variations": [{
    "id": "1402",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [59, 52, 52, 56],
    "frets": [2, 0, 4, 1],
    "fingers": [2, 0, 4, 1],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "1x02",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [59, 52, -1, 56],
    "frets": [2, 0, -1, 1],
    "fingers": [2, 0, 0, 1],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "4442",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [59, 56, 52, 59],
    "frets": [2, 4, 4, 4],
    "fingers": [1, 4, 3, 2],
    "fret": 2,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "4447",
    "listCapos": [{
      "fret": 4,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [64, 56, 52, 59],
    "frets": [7, 4, 4, 4],
    "fingers": [4, 0, 0, 0],
    "fret": 4,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "4x47",
    "listCapos": [{
      "fret": 4,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [64, 56, -1, 59],
    "frets": [7, 4, -1, 4],
    "fingers": [4, 0, 0, 0],
    "fret": 4,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "9877",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [64, 59, 56, 64],
    "frets": [7, 7, 8, 9],
    "fingers": [0, 0, 2, 3],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x877",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [64, 59, 56, -1],
    "frets": [7, 7, 8, -1],
    "fingers": [0, 0, 2, 0],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "9111211",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [68, 64, 59, 64],
    "frets": [11, 12, 11, 9],
    "fingers": [3, 4, 2, 1],
    "fret": 9,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "13111211",
    "listCapos": [{
      "fret": 11,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [68, 64, 59, 68],
    "frets": [11, 12, 11, 13],
    "fingers": [0, 2, 0, 3],
    "fret": 11,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "13111214",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [71, 64, 59, 68],
    "frets": [14, 12, 11, 13],
    "fingers": [4, 2, 1, 3],
    "fret": 11,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x111211",
    "listCapos": [{
      "fret": 11,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [68, 64, 59, -1],
    "frets": [11, 12, 11, -1],
    "fingers": [0, 2, 0, 0],
    "fret": 11,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "444x",
    "listCapos": [{
      "fret": 4,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 4,
    "notes": [-1, 56, 52, 59],
    "frets": [-1, 4, 4, 4],
    "fingers": [0, 0, 0, 0],
    "fret": 4,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "987x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 59, 56, 64],
    "frets": [-1, 7, 8, 9],
    "fingers": [0, 1, 2, 3],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "131112x",
    "listCapos": [],
    "noteIndex": 4,
    "notes": [-1, 64, 59, 68],
    "frets": [-1, 12, 11, 13],
    "fingers": [0, 2, 1, 3],
    "fret": 11,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "F7",
  "variations": [{
    "id": "2310",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [57, 53, 51, 57],
    "frets": [0, 1, 3, 2],
    "fingers": [0, 1, 3, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x310",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [57, 53, 51, -1],
    "frets": [0, 1, 3, -1],
    "fingers": [0, 1, 3, 0],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "2313",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [60, 53, 51, 57],
    "frets": [3, 1, 3, 2],
    "fingers": [4, 1, 3, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "5556",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [63, 57, 53, 60],
    "frets": [6, 5, 5, 5],
    "fingers": [2, 0, 0, 0],
    "fret": 5,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "8556",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [63, 57, 53, 63],
    "frets": [6, 5, 5, 8],
    "fingers": [2, 0, 0, 4],
    "fret": 5,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "8988",
    "listCapos": [{
      "fret": 8,
      "startString": 0,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [65, 60, 57, 63],
    "frets": [8, 8, 9, 8],
    "fingers": [0, 0, 2, 0],
    "fret": 8,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "89118",
    "listCapos": [{
      "fret": 8,
      "startString": 0,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [65, 63, 57, 63],
    "frets": [8, 11, 9, 8],
    "fingers": [0, 4, 2, 0],
    "fret": 8,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "109118",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [65, 63, 57, 65],
    "frets": [8, 11, 9, 10],
    "fingers": [1, 4, 2, 3],
    "fret": 8,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "x9118",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [65, 63, 57, -1],
    "frets": [8, 11, 9, -1],
    "fingers": [1, 4, 2, 0],
    "fret": 8,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "1091112",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [69, 63, 57, 65],
    "frets": [12, 11, 9, 10],
    "fingers": [4, 3, 1, 2],
    "fret": 9,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "10121112",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [69, 63, 60, 65],
    "frets": [12, 11, 12, 10],
    "fingers": [4, 2, 3, 1],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "10x1112",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [69, 63, -1, 65],
    "frets": [12, 11, -1, 10],
    "fingers": [3, 2, 0, 1],
    "fret": 10,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "231x",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [-1, 53, 51, 57],
    "frets": [-1, 1, 3, 2],
    "fingers": [0, 1, 3, 2],
    "fret": 0,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "855x",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [-1, 57, 53, 63],
    "frets": [-1, 5, 5, 8],
    "fingers": [0, 0, 0, 4],
    "fret": 5,
    "type": "_7",
    "baseDisplayNote": ""
  }, {
    "id": "10911x",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [-1, 63, 57, 65],
    "frets": [-1, 11, 9, 10],
    "fingers": [0, 3, 1, 2],
    "fret": 9,
    "type": "_7",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "D",
  "variations": [{
    "id": "2220",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [57, 54, 50, 57],
    "frets": [0, 2, 2, 2],
    "fingers": [0, 3, 2, 1],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "2225",
    "listCapos": [{
      "fret": 2,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [62, 54, 50, 57],
    "frets": [5, 2, 2, 2],
    "fingers": [4, 0, 0, 0],
    "fret": 2,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "2x25",
    "listCapos": [{
      "fret": 2,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [62, 54, -1, 57],
    "frets": [5, 2, -1, 2],
    "fingers": [4, 0, 0, 0],
    "fret": 2,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "7655",
    "listCapos": [{
      "fret": 5,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [62, 57, 54, 62],
    "frets": [5, 5, 6, 7],
    "fingers": [0, 0, 2, 3],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x655",
    "listCapos": [{
      "fret": 5,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [62, 57, 54, -1],
    "frets": [5, 5, 6, -1],
    "fingers": [0, 0, 2, 0],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "79109",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [66, 62, 57, 62],
    "frets": [9, 10, 9, 7],
    "fingers": [3, 4, 2, 1],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "119109",
    "listCapos": [{
      "fret": 9,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [66, 62, 57, 66],
    "frets": [9, 10, 9, 11],
    "fingers": [0, 2, 0, 3],
    "fret": 9,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "1191012",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [69, 62, 57, 66],
    "frets": [12, 10, 9, 11],
    "fingers": [4, 2, 1, 3],
    "fret": 9,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x9109",
    "listCapos": [{
      "fret": 9,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [66, 62, 57, -1],
    "frets": [9, 10, 9, -1],
    "fingers": [0, 2, 0, 0],
    "fret": 9,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "11x1012",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [69, 62, -1, 66],
    "frets": [12, 10, -1, 11],
    "fingers": [3, 1, 0, 2],
    "fret": 10,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "222x",
    "listCapos": [{
      "fret": 2,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 2,
    "notes": [-1, 54, 50, 57],
    "frets": [-1, 2, 2, 2],
    "fingers": [0, 0, 0, 0],
    "fret": 2,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "765x",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [-1, 57, 54, 62],
    "frets": [-1, 5, 6, 7],
    "fingers": [0, 1, 2, 3],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "11910x",
    "listCapos": [],
    "noteIndex": 2,
    "notes": [-1, 62, 57, 66],
    "frets": [-1, 10, 9, 11],
    "fingers": [0, 2, 1, 3],
    "fret": 9,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "C",
  "variations": [{
    "id": "0003",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [60, 52, 48, 55],
    "frets": [3, 0, 0, 0],
    "fingers": [1, 0, 0, 0],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "0403",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [60, 52, 52, 55],
    "frets": [3, 0, 4, 0],
    "fingers": [1, 0, 2, 0],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "0433",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [60, 55, 52, 55],
    "frets": [3, 3, 4, 0],
    "fingers": [0, 0, 2, 0],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "5433",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [60, 55, 52, 60],
    "frets": [3, 3, 4, 5],
    "fingers": [0, 0, 2, 3],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x433",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [60, 55, 52, -1],
    "frets": [3, 3, 4, -1],
    "fingers": [0, 0, 2, 0],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "5787",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [64, 60, 55, 60],
    "frets": [7, 8, 7, 5],
    "fingers": [3, 4, 2, 1],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "0787",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [64, 60, 55, 55],
    "frets": [7, 8, 7, 0],
    "fingers": [0, 2, 0, 0],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "9787",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [64, 60, 55, 64],
    "frets": [7, 8, 7, 9],
    "fingers": [0, 2, 0, 3],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "97810",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [67, 60, 55, 64],
    "frets": [10, 8, 7, 9],
    "fingers": [4, 2, 1, 3],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x787",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [64, 60, 55, -1],
    "frets": [7, 8, 7, -1],
    "fingers": [0, 2, 0, 0],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "9x810",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [67, 60, -1, 64],
    "frets": [10, 8, -1, 9],
    "fingers": [3, 1, 0, 2],
    "fret": 8,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "0121210",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [67, 64, 60, 55],
    "frets": [10, 12, 12, 0],
    "fingers": [1, 4, 3, 0],
    "fret": 10,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "12121210",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [67, 64, 60, 67],
    "frets": [10, 12, 12, 12],
    "fingers": [1, 4, 3, 2],
    "fret": 10,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "000x",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [-1, 52, 48, 55],
    "frets": [-1, 0, 0, 0],
    "fingers": [0, 0, 0, 0],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "543x",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [-1, 55, 52, 60],
    "frets": [-1, 3, 4, 5],
    "fingers": [0, 1, 2, 3],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "978x",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [-1, 60, 55, 64],
    "frets": [-1, 8, 7, 9],
    "fingers": [0, 2, 1, 3],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "Am",
  "variations": [{
    "id": "2000",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [57, 52, 48, 57],
    "frets": [0, 0, 0, 2],
    "fingers": [0, 0, 0, 1],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "x000",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [57, 52, 48, -1],
    "frets": [0, 0, 0, -1],
    "fingers": [0, 0, 0, 0],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "2003",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [60, 52, 48, 57],
    "frets": [3, 0, 0, 2],
    "fingers": [2, 0, 0, 1],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "2403",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [60, 52, 52, 57],
    "frets": [3, 0, 4, 2],
    "fingers": [2, 0, 3, 1],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "2x03",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [60, 52, -1, 57],
    "frets": [3, 0, -1, 2],
    "fingers": [2, 0, 0, 1],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "2453",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [60, 57, 52, 57],
    "frets": [3, 5, 4, 2],
    "fingers": [2, 4, 3, 1],
    "fret": 2,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "5453",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [60, 57, 52, 60],
    "frets": [3, 5, 4, 5],
    "fingers": [1, 4, 2, 3],
    "fret": 3,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "x453",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [60, 57, 52, -1],
    "frets": [3, 5, 4, -1],
    "fingers": [1, 3, 2, 0],
    "fret": 3,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "x057",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [64, 57, 48, -1],
    "frets": [7, 5, 0, -1],
    "fingers": [3, 1, 0, 0],
    "fret": 5,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "5x57",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 9,
    "notes": [64, 57, -1, 60],
    "frets": [7, 5, -1, 5],
    "fingers": [3, 0, 0, 0],
    "fret": 5,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "9987",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [64, 60, 57, 64],
    "frets": [7, 8, 9, 9],
    "fingers": [1, 2, 4, 3],
    "fret": 7,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "9121212",
    "listCapos": [{
      "fret": 12,
      "startString": 0,
      "lastString": 2,
      "finger": 4
    }],
    "noteIndex": 9,
    "notes": [69, 64, 60, 64],
    "frets": [12, 12, 12, 9],
    "fingers": [0, 0, 0, 1],
    "fret": 9,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "200x",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [-1, 52, 48, 57],
    "frets": [-1, 0, 0, 2],
    "fingers": [0, 0, 0, 1],
    "fret": 0,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "545x",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [-1, 57, 52, 60],
    "frets": [-1, 5, 4, 5],
    "fingers": [0, 3, 1, 2],
    "fret": 4,
    "type": "MINOR",
    "baseDisplayNote": ""
  }, {
    "id": "998x",
    "listCapos": [],
    "noteIndex": 9,
    "notes": [-1, 60, 57, 64],
    "frets": [-1, 8, 9, 9],
    "fingers": [0, 1, 3, 2],
    "fret": 8,
    "type": "MINOR",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "F",
  "variations": [{
    "id": "2010",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [57, 53, 48, 57],
    "frets": [0, 1, 0, 2],
    "fingers": [0, 1, 0, 2],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x010",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [57, 53, 48, -1],
    "frets": [0, 1, 0, -1],
    "fingers": [0, 1, 0, 0],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "2013",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [60, 53, 48, 57],
    "frets": [3, 1, 0, 2],
    "fingers": [3, 1, 0, 2],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "2x13",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [60, 53, -1, 57],
    "frets": [3, 1, -1, 2],
    "fingers": [3, 1, 0, 2],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "5553",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [60, 57, 53, 60],
    "frets": [3, 5, 5, 5],
    "fingers": [1, 4, 3, 2],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x058",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [65, 57, 48, -1],
    "frets": [8, 5, 0, -1],
    "fingers": [4, 1, 0, 0],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "5558",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [65, 57, 53, 60],
    "frets": [8, 5, 5, 5],
    "fingers": [4, 0, 0, 0],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "5x58",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [65, 57, -1, 60],
    "frets": [8, 5, -1, 5],
    "fingers": [4, 0, 0, 0],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "10988",
    "listCapos": [{
      "fret": 8,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [65, 60, 57, 65],
    "frets": [8, 8, 9, 10],
    "fingers": [0, 0, 2, 3],
    "fret": 8,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x988",
    "listCapos": [{
      "fret": 8,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [65, 60, 57, -1],
    "frets": [8, 8, 9, -1],
    "fingers": [0, 0, 2, 0],
    "fret": 8,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "10121312",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [69, 65, 60, 65],
    "frets": [12, 13, 12, 10],
    "fingers": [3, 4, 2, 1],
    "fret": 10,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "201x",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [-1, 53, 48, 57],
    "frets": [-1, 1, 0, 2],
    "fingers": [0, 1, 0, 2],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "555x",
    "listCapos": [{
      "fret": 5,
      "startString": 1,
      "lastString": 3,
      "finger": 1
    }],
    "noteIndex": 5,
    "notes": [-1, 57, 53, 60],
    "frets": [-1, 5, 5, 5],
    "fingers": [0, 0, 0, 0],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "1098x",
    "listCapos": [],
    "noteIndex": 5,
    "notes": [-1, 60, 57, 65],
    "frets": [-1, 8, 9, 10],
    "fingers": [0, 1, 2, 3],
    "fret": 8,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }]
}, {
  "chord": "C\/E",
  "variations": [{
    "id": "0003",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [60, 52, 48, 55],
    "frets": [3, 0, 0, 0],
    "fingers": [1, 0, 0, 0],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "0403",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [60, 52, 52, 55],
    "frets": [3, 0, 4, 0],
    "fingers": [1, 0, 2, 0],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "0433",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [60, 55, 52, 55],
    "frets": [3, 3, 4, 0],
    "fingers": [0, 0, 2, 0],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "5433",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [60, 55, 52, 60],
    "frets": [3, 3, 4, 5],
    "fingers": [0, 0, 2, 3],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x433",
    "listCapos": [{
      "fret": 3,
      "startString": 0,
      "lastString": 1,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [60, 55, 52, -1],
    "frets": [3, 3, 4, -1],
    "fingers": [0, 0, 2, 0],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "5787",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [64, 60, 55, 60],
    "frets": [7, 8, 7, 5],
    "fingers": [3, 4, 2, 1],
    "fret": 5,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "0787",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [64, 60, 55, 55],
    "frets": [7, 8, 7, 0],
    "fingers": [0, 2, 0, 0],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "9787",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [64, 60, 55, 64],
    "frets": [7, 8, 7, 9],
    "fingers": [0, 2, 0, 3],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "97810",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [67, 60, 55, 64],
    "frets": [10, 8, 7, 9],
    "fingers": [4, 2, 1, 3],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "x787",
    "listCapos": [{
      "fret": 7,
      "startString": 0,
      "lastString": 2,
      "finger": 1
    }],
    "noteIndex": 0,
    "notes": [64, 60, 55, -1],
    "frets": [7, 8, 7, -1],
    "fingers": [0, 2, 0, 0],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "9x810",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [67, 60, -1, 64],
    "frets": [10, 8, -1, 9],
    "fingers": [3, 1, 0, 2],
    "fret": 8,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "0121210",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [67, 64, 60, 55],
    "frets": [10, 12, 12, 0],
    "fingers": [1, 4, 3, 0],
    "fret": 10,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "12121210",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [67, 64, 60, 67],
    "frets": [10, 12, 12, 12],
    "fingers": [1, 4, 3, 2],
    "fret": 10,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "000x",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [-1, 52, 48, 55],
    "frets": [-1, 0, 0, 0],
    "fingers": [0, 0, 0, 0],
    "fret": 0,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "543x",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [-1, 55, 52, 60],
    "frets": [-1, 3, 4, 5],
    "fingers": [0, 1, 2, 3],
    "fret": 3,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }, {
    "id": "978x",
    "listCapos": [],
    "noteIndex": 0,
    "notes": [-1, 60, 55, 64],
    "frets": [-1, 8, 7, 9],
    "fingers": [0, 2, 1, 3],
    "fret": 7,
    "type": "MAJOR",
    "baseDisplayNote": ""
  }]
}]
```




Piano inversions: `https://api.ultimate-guitar.com/api/v1/tab/pianoInversions?chords[]=Dm&chords[]=E7&chords[]=E&chords[]=F7&chords[]=D&chords[]=C&chords[]=Am&chords[]=F&chords[]=C%2FE`  
