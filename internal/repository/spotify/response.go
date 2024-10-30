package spotify

var searchResponse = `{
	"tracks": {
	  "href": "https://api.spotify.com/v1/search?query=Bruno+Mars&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.9&offset=0&limit=10",
	  "items": [
		{
		  "album": {
			"album_type": "single",
			"artists": [
			  {
				"external_urls": {
				  "spotify": "https://open.spotify.com/artist/3eVa5w3URK5duf6eyVDbu9"
				},
				"href": "https://api.spotify.com/v1/artists/3eVa5w3URK5duf6eyVDbu9",
				"id": "3eVa5w3URK5duf6eyVDbu9",
				"name": "ROSÉ",
				"type": "artist",
				"uri": "spotify:artist:3eVa5w3URK5duf6eyVDbu9"
			  },
			  {
				"external_urls": {
				  "spotify": "https://open.spotify.com/artist/0du5cEVh5yTK9QJze8zA0C"
				},
				"href": "https://api.spotify.com/v1/artists/0du5cEVh5yTK9QJze8zA0C",
				"id": "0du5cEVh5yTK9QJze8zA0C",
				"name": "Bruno Mars",
				"type": "artist",
				"uri": "spotify:artist:0du5cEVh5yTK9QJze8zA0C"
			  }
			],
			"external_urls": {
			  "spotify": "https://open.spotify.com/album/2IYQwwgxgOIn7t3iF6ufFD"
			},
			"href": "https://api.spotify.com/v1/albums/2IYQwwgxgOIn7t3iF6ufFD",
			"id": "2IYQwwgxgOIn7t3iF6ufFD",
			"images": [
			  {
				"height": 640,
				"url": "https://i.scdn.co/image/ab67616d0000b273f8c8297efc6022534f1357e1",
				"width": 640
			  },
			  {
				"height": 300,
				"url": "https://i.scdn.co/image/ab67616d00001e02f8c8297efc6022534f1357e1",
				"width": 300
			  },
			  {
				"height": 64,
				"url": "https://i.scdn.co/image/ab67616d00004851f8c8297efc6022534f1357e1",
				"width": 64
			  }
			],
			"is_playable": true,
			"name": "APT.",
			"release_date": "2024-10-18",
			"release_date_precision": "day",
			"total_tracks": 1,
			"type": "album",
			"uri": "spotify:album:2IYQwwgxgOIn7t3iF6ufFD"
		  },
		  "artists": [
			{
			  "external_urls": {
				"spotify": "https://open.spotify.com/artist/3eVa5w3URK5duf6eyVDbu9"
			  },
			  "href": "https://api.spotify.com/v1/artists/3eVa5w3URK5duf6eyVDbu9",
			  "id": "3eVa5w3URK5duf6eyVDbu9",
			  "name": "ROSÉ",
			  "type": "artist",
			  "uri": "spotify:artist:3eVa5w3URK5duf6eyVDbu9"
			},
			{
			  "external_urls": {
				"spotify": "https://open.spotify.com/artist/0du5cEVh5yTK9QJze8zA0C"
			  },
			  "href": "https://api.spotify.com/v1/artists/0du5cEVh5yTK9QJze8zA0C",
			  "id": "0du5cEVh5yTK9QJze8zA0C",
			  "name": "Bruno Mars",
			  "type": "artist",
			  "uri": "spotify:artist:0du5cEVh5yTK9QJze8zA0C"
			}
		  ],
		  "disc_number": 1,
		  "duration_ms": 169917,
		  "explicit": false,
		  "external_ids": {
			"isrc": "USAT22409172"
		  },
		  "external_urls": {
			"spotify": "https://open.spotify.com/track/5vNRhkKd0yEAg8suGBpjeY"
		  },
		  "href": "https://api.spotify.com/v1/tracks/5vNRhkKd0yEAg8suGBpjeY",
		  "id": "5vNRhkKd0yEAg8suGBpjeY",
		  "is_local": false,
		  "is_playable": true,
		  "name": "APT.",
		  "popularity": 94,
		  "preview_url": "https://p.scdn.co/mp3-preview/6b7763066dd42d94dc893f12bb135d7d5392e386?cid=cfe923b2d660439caf2b557b21f31221",
		  "track_number": 1,
		  "type": "track",
		  "uri": "spotify:track:5vNRhkKd0yEAg8suGBpjeY"
		},
		{
		  "album": {
			"album_type": "single",
			"artists": [
			  {
				"external_urls": {
				  "spotify": "https://open.spotify.com/artist/1HY2Jd0NmPuamShAr6KMms"
				},
				"href": "https://api.spotify.com/v1/artists/1HY2Jd0NmPuamShAr6KMms",
				"id": "1HY2Jd0NmPuamShAr6KMms",
				"name": "Lady Gaga",
				"type": "artist",
				"uri": "spotify:artist:1HY2Jd0NmPuamShAr6KMms"
			  },
			  {
				"external_urls": {
				  "spotify": "https://open.spotify.com/artist/0du5cEVh5yTK9QJze8zA0C"
				},
				"href": "https://api.spotify.com/v1/artists/0du5cEVh5yTK9QJze8zA0C",
				"id": "0du5cEVh5yTK9QJze8zA0C",
				"name": "Bruno Mars",
				"type": "artist",
				"uri": "spotify:artist:0du5cEVh5yTK9QJze8zA0C"
			  }
			],
			"external_urls": {
			  "spotify": "https://open.spotify.com/album/10FLjwfpbxLmW8c25Xyc2N"
			},
			"href": "https://api.spotify.com/v1/albums/10FLjwfpbxLmW8c25Xyc2N",
			"id": "10FLjwfpbxLmW8c25Xyc2N",
			"images": [
			  {
				"height": 640,
				"url": "https://i.scdn.co/image/ab67616d0000b27382ea2e9e1858aa012c57cd45",
				"width": 640
			  },
			  {
				"height": 300,
				"url": "https://i.scdn.co/image/ab67616d00001e0282ea2e9e1858aa012c57cd45",
				"width": 300
			  },
			  {
				"height": 64,
				"url": "https://i.scdn.co/image/ab67616d0000485182ea2e9e1858aa012c57cd45",
				"width": 64
			  }
			],
			"is_playable": true,
			"name": "Die With A Smile",
			"release_date": "2024-08-16",
			"release_date_precision": "day",
			"total_tracks": 1,
			"type": "album",
			"uri": "spotify:album:10FLjwfpbxLmW8c25Xyc2N"
		  },
		  "artists": [
			{
			  "external_urls": {
				"spotify": "https://open.spotify.com/artist/1HY2Jd0NmPuamShAr6KMms"
			  },
			  "href": "https://api.spotify.com/v1/artists/1HY2Jd0NmPuamShAr6KMms",
			  "id": "1HY2Jd0NmPuamShAr6KMms",
			  "name": "Lady Gaga",
			  "type": "artist",
			  "uri": "spotify:artist:1HY2Jd0NmPuamShAr6KMms"
			},
			{
			  "external_urls": {
				"spotify": "https://open.spotify.com/artist/0du5cEVh5yTK9QJze8zA0C"
			  },
			  "href": "https://api.spotify.com/v1/artists/0du5cEVh5yTK9QJze8zA0C",
			  "id": "0du5cEVh5yTK9QJze8zA0C",
			  "name": "Bruno Mars",
			  "type": "artist",
			  "uri": "spotify:artist:0du5cEVh5yTK9QJze8zA0C"
			}
		  ],
		  "disc_number": 1,
		  "duration_ms": 251667,
		  "explicit": false,
		  "external_ids": {
			"isrc": "USUM72409273"
		  },
		  "external_urls": {
			"spotify": "https://open.spotify.com/track/2plbrEY59IikOBgBGLjaoe"
		  },
		  "href": "https://api.spotify.com/v1/tracks/2plbrEY59IikOBgBGLjaoe",
		  "id": "2plbrEY59IikOBgBGLjaoe",
		  "is_local": false,
		  "is_playable": true,
		  "name": "Die With A Smile",
		  "popularity": 100,
		  "preview_url": null,
		  "track_number": 1,
		  "type": "track",
		  "uri": "spotify:track:2plbrEY59IikOBgBGLjaoe"
		}
	  ],
	  "limit": 10,
	  "next": "https://api.spotify.com/v1/search?query=Bruno+Mars&type=track&market=ID&locale=en-US%2Cen%3Bq%3D0.9&offset=10&limit=10",
	  "offset": 0,
	  "previous": null,
	  "total": 894
	}
  }`
