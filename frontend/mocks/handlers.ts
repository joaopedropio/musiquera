import { http, HttpResponse } from 'msw'

type Song = {
	name: string;
	file: string;
	duration: number;
}

type Release = {
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
		let b: Artist = {
			name: 'Kream'
		}
		let c: Artist = {
			name: 'Massacration'
		}
		return HttpResponse.json([a, b, c])
	}),

	http.get('/api/release/byArtist/:artistName', () => {
		let a: Release = {
			name: 'Master of Puppets',
			artist: 'Metallica',
			releaseDate: '1986-03-03',
			songs: [
				{
					name: 'Master of Puppets',
					file: '/somewhere',
					duration: 124,
				}
			],
		}
		return HttpResponse.json([a])
	}),

	http.get('/api/release/mostRecent', () => {
		let a: Release = {
			name: 'Master of Puppets',
			artist: 'Metallica',
			releaseDate: '1986-03-03',
			songs: [],
		}
		return HttpResponse.json([a])
	}),
]
