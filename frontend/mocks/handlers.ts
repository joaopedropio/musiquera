import { http, HttpResponse } from 'msw'

type Song = {
	name: string;
	file: string;
}

type Album = {
	name: string;
	artist: string;
	releaseDate: string;
	songs: Song[];
}

type Artist = {
	name: string;
}


export const handlers = [
	http.get('/api/artist/', () => {
		let a: Artist = {
			name: 'Metallica'
		}
		return HttpResponse.json([a])
	}),

	http.get('/api/album/byArtist/:artistName', () => {
		let a: Album = {
			name: 'Master of Puppest',
			artist: 'Metallica',
			releaseDate: '1986-03-03',
			songs: [
				{
					name: 'Master of Puppets',
					file: '/somewhere',
				}
			],
		}
		return HttpResponse.json([a])
	}),

	http.get('/api/album/mostRecent', () => {
		let a: Album = {
			name: 'Master of Puppest',
			artist: 'Metallica',
			releaseDate: '1986-03-03',
			songs: [],
		}
		return HttpResponse.json([a])
	}),
]
