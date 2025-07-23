import { http, HttpResponse } from 'msw'

const ReleaseTypeAlbum = 'album'
//const ReleaseTypeLiveSet = 'liveSet'

type ReleaseType = string

type Track = {
	name: string;
	file: string;
	duration: number;
}

type Release = {
	name: string;
	artist: string;
	releaseDate: string;
	type: ReleaseType;
	tracks: Track[];
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
			type: ReleaseTypeAlbum,
			releaseDate: '1986-03-03',
			tracks: [
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
			type: ReleaseTypeAlbum,
			releaseDate: '1986-03-03',
			tracks: [],
		}
		return HttpResponse.json([a])
	}),
]
